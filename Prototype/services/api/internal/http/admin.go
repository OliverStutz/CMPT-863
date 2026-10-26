package http

import (
	"context"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"golang.org/x/crypto/bcrypt"

	"study/api/internal/analysis"
	"study/api/internal/configs"
	"study/api/internal/domain"
	"study/api/internal/export"
	"study/api/internal/store"
)

const (
	adminCookie  = "admin"
	gateCookie   = "admin_gate"
	adminHeader  = "X-Study-Admin"
	adminTTL     = 12 * time.Hour      // password login; asked again after this
	gateTTL      = 30 * 24 * time.Hour // access-key gate; renewed by every page load behind it
	maxConfigLen = 4 << 20
)

var (
	fallbackKeyOnce sync.Once
	fallbackKey     []byte
)

// signingKey returns admin-session-key, or a random per-process key when the
// secret is not configured (dev; cookies then end with the process).
func (s *Server) signingKey(ctx context.Context) []byte {
	if k := s.o.SessionKey.Get(ctx); len(k) > 0 {
		return k
	}
	fallbackKeyOnce.Do(func() {
		fallbackKey = make([]byte, 32)
		_, _ = rand.Read(fallbackKey)
		s.o.Log.Warn("ADMIN_SESSION_KEY not configured; using a random per-process key")
	})
	return fallbackKey
}

func (s *Server) sign(ctx context.Context, payload string) string {
	m := hmac.New(sha256.New, s.signingKey(ctx))
	m.Write([]byte(payload))
	return base64.RawURLEncoding.EncodeToString([]byte(payload)) + "." + base64.RawURLEncoding.EncodeToString(m.Sum(nil))
}

// verify returns the payload of a signed value, or "" when invalid.
func (s *Server) verify(ctx context.Context, v string) string {
	p64, sig64, ok := strings.Cut(v, ".")
	if !ok {
		return ""
	}
	p, err1 := base64.RawURLEncoding.DecodeString(p64)
	sig, err2 := base64.RawURLEncoding.DecodeString(sig64)
	if err1 != nil || err2 != nil {
		return ""
	}
	m := hmac.New(sha256.New, s.signingKey(ctx))
	m.Write(p)
	if !hmac.Equal(sig, m.Sum(nil)) {
		return ""
	}
	return string(p)
}

// unexpired checks a payload "<kind>|...|<unix expiry>" (expiry last).
func (s *Server) unexpired(payload, kind string) []string {
	parts := strings.Split(payload, "|")
	if len(parts) < 2 || parts[0] != kind {
		return nil
	}
	exp, err := strconv.ParseInt(parts[len(parts)-1], 10, 64)
	if err != nil || s.now().Unix() >= exp {
		return nil
	}
	return parts
}

func keyFingerprint(k []byte) string {
	h := sha256.Sum256(k)
	return hex.EncodeToString(h[:8])
}

// adminGate hides the admin area (SPA and API) behind the access key: 404
// unless the request carries the key as ?k= on an /admin/ page request (which
// sets the signed admin_gate cookie and redirects without the query) or a
// valid admin_gate cookie. The cookie is bound to the current key, so a
// rotated key invalidates it.
//
// The gate cookie lives gateTTL (30 days) and every page load behind the gate
// renews it, so a browser that opened the key link once keeps rendering the
// page while it is in use; the password login (adminTTL, 12 hours) is what
// expires first and is asked again by the admin app. Before, both lived 12
// hours and expired together, so an open admin tab met the gate's 404 with no
// way back but the key link. The gate cookie is SameSite=Lax, not Strict, so
// the redirect after a key link opened from another site or app still carries
// it: it grants nothing but rendering the page; data and writes need the Strict
// admin cookie plus the CSRF header.
func (s *Server) adminGate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		key := s.o.AccessKey.Get(ctx)
		if len(key) == 0 {
			http.NotFound(w, r)
			return
		}
		fp := keyFingerprint(key)
		isPage := r.URL.Path == "/admin" || strings.HasPrefix(r.URL.Path, "/admin/")
		if c, err := r.Cookie(gateCookie); err == nil {
			if parts := s.unexpired(s.verify(ctx, c.Value), "gate"); len(parts) == 3 && parts[1] == fp {
				if isPage && !strings.HasPrefix(r.URL.Path, "/admin/assets/") {
					s.setGateCookie(ctx, w, fp)
				}
				next.ServeHTTP(w, r)
				return
			}
		}
		k := r.URL.Query().Get("k")
		ip := ClientIP(r, s.o.TrustedProxyHops)
		if k != "" && isPage && s.adminGateFail.Count(ip) < GateFailuresPerHour {
			if subtle.ConstantTimeCompare([]byte(k), key) == 1 {
				s.setGateCookie(ctx, w, fp)
				http.Redirect(w, r, "/admin/", http.StatusFound)
				return
			}
			s.adminGateFail.Hit(ip)
		}
		http.NotFound(w, r)
	})
}

// setGateCookie issues a fresh admin_gate cookie bound to the key fingerprint.
func (s *Server) setGateCookie(ctx context.Context, w http.ResponseWriter, fp string) {
	exp := s.now().Add(gateTTL).Unix()
	http.SetCookie(w, &http.Cookie{Name: gateCookie, Value: s.sign(ctx, "gate|"+fp+"|"+strconv.FormatInt(exp, 10)),
		Path: "/", HttpOnly: true, Secure: !s.o.InsecureCookies, SameSite: http.SameSiteLaxMode,
		MaxAge: int(gateTTL.Seconds())})
}

func (s *Server) adminAPI() http.Handler {
	m := http.NewServeMux()
	m.HandleFunc("POST /api/admin/login", s.adminLogin)
	m.HandleFunc("POST /api/admin/logout", s.adminLogout)
	m.HandleFunc("GET /api/admin/runs", s.requireAdmin(s.listRuns))
	m.HandleFunc("POST /api/admin/runs", s.requireAdmin(s.createRun))
	m.HandleFunc("PATCH /api/admin/runs/{id}", s.requireAdmin(s.patchRun))
	m.HandleFunc("DELETE /api/admin/runs/{id}", s.requireAdmin(s.deleteRun))
	m.HandleFunc("GET /api/admin/runs/{id}/analysis", s.requireAdmin(s.runAnalysis))
	m.HandleFunc("GET /api/admin/participants", s.requireAdmin(s.listParticipants))
	m.HandleFunc("GET /api/admin/participants/{id}", s.requireAdmin(s.participantDetail))
	m.HandleFunc("DELETE /api/admin/sessions/{id}", s.requireAdmin(s.deleteSession))
	m.HandleFunc("GET /api/admin/export", s.requireAdmin(s.export))
	m.HandleFunc("GET /api/admin/configs", s.requireAdmin(s.listConfigs))
	m.HandleFunc("POST /api/admin/configs", s.requireAdmin(s.uploadConfig))
	m.HandleFunc("GET /api/admin/configs/{version}", s.requireAdmin(s.getConfig))
	m.HandleFunc("/api/admin/", s.requireAdmin(func(w http.ResponseWriter, r *http.Request) {
		writeErr(w, domain.ErrNotFound, "no such endpoint")
	}))
	return m
}

// requireAdmin checks the admin session cookie and the CSRF header (8.2).
func (s *Server) requireAdmin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie(adminCookie)
		if err != nil || len(s.unexpired(s.verify(r.Context(), c.Value), "admin")) != 3 {
			writeErr(w, domain.ErrAdminLoginRequired, "admin login required")
			return
		}
		if r.Header.Get(adminHeader) != "1" {
			writeErrStatus(w, http.StatusForbidden, ErrForbidden, "missing X-Study-Admin header")
			return
		}
		next(w, r)
	}
}

// adminLogin checks the shared password. Only wrong passwords count against
// the per-IP limit (AdminLoginAttempts in 15 minutes); a correct login never
// counts and clears the failures of its address.
// DECISION (2026-10-25): the first version counted every attempt, so test runs
// from the same network locked the admin out before the first try.
func (s *Server) adminLogin(w http.ResponseWriter, r *http.Request) {
	ip := ClientIP(r, s.o.TrustedProxyHops)
	if s.loginLimit.Count(ip) >= AdminLoginAttempts {
		writeErr(w, domain.ErrRateLimited, "too many wrong passwords from this address; wait 15 minutes")
		return
	}
	var body struct {
		Password string `json:"password"`
	}
	if err := json.NewDecoder(io.LimitReader(r.Body, 4096)).Decode(&body); err != nil {
		writeErr(w, domain.ErrBadRequest, "body must be { password }")
		return
	}
	hash := s.o.PasswordHash.Get(r.Context())
	if len(hash) == 0 || bcrypt.CompareHashAndPassword(hash, []byte(body.Password)) != nil {
		s.loginLimit.Hit(ip)
		writeErr(w, domain.ErrBadPassword, "wrong password")
		return
	}
	s.loginLimit.Reset(ip)
	exp := s.now().Add(adminTTL).Unix()
	http.SetCookie(w, &http.Cookie{Name: adminCookie, Value: s.sign(r.Context(), "admin|"+randomToken()+"|"+strconv.FormatInt(exp, 10)),
		Path: "/api/admin", HttpOnly: true, Secure: !s.o.InsecureCookies, SameSite: http.SameSiteStrictMode,
		MaxAge: int(adminTTL.Seconds())})
	writeJSON(w, http.StatusOK, map[string]any{"ok": true, "expiresAt": time.Unix(exp, 0).UTC()})
}

func (s *Server) adminLogout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{Name: adminCookie, Value: "", Path: "/api/admin", HttpOnly: true,
		Secure: !s.o.InsecureCookies, SameSite: http.SameSiteStrictMode, MaxAge: -1})
	w.Header().Set("Cache-Control", "no-store")
	w.WriteHeader(http.StatusNoContent)
}

// ---------- runs ----------

type runView struct {
	*domain.Run
	Counts          map[string]int `json:"counts"`
	FinishedByGroup map[string]int `json:"finishedByGroup"`
}

func (s *Server) listRuns(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := reqCtx(r)
	defer cancel()
	runs, err := s.o.Store.ListRuns(ctx)
	if err != nil {
		writeErr(w, domain.ErrUnavailable, "store unavailable")
		return
	}
	sessions, err := s.o.Store.ListSessions(ctx, store.SessionFilter{})
	if err != nil {
		writeErr(w, domain.ErrUnavailable, "store unavailable")
		return
	}
	out := make([]runView, 0, len(runs))
	for _, run := range runs {
		v := runView{Run: run, Counts: map[string]int{"opened": 0, "started": 0, "ongoing": 0, "finished": 0, "abandoned": 0, "blocked": 0},
			FinishedByGroup: map[string]int{"checker": 0, "neutral": 0}}
		for _, x := range sessions {
			if x.RunID != run.ID {
				continue
			}
			v.Counts[strings.ToLower(x.Status)]++
			if x.Status == domain.StatusFinished && x.Group != "" {
				v.FinishedByGroup[x.Group]++
			}
		}
		out = append(out, v)
	}
	writeJSON(w, http.StatusOK, out)
}

func (s *Server) baseURL(r *http.Request) string {
	if s.o.PublicBaseURL != "" {
		return strings.TrimSuffix(s.o.PublicBaseURL, "/")
	}
	scheme := "https"
	if r.TLS == nil && s.o.InsecureCookies {
		scheme = "http"
	}
	return scheme + "://" + r.Host
}

func (s *Server) createRun(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Name          string `json:"name"`
		Source        string `json:"source"`
		ConfigVersion string `json:"configVersion"`
		TestOverrides bool   `json:"testOverrides"`
	}
	if err := json.NewDecoder(io.LimitReader(r.Body, 8192)).Decode(&body); err != nil {
		writeErr(w, domain.ErrBadRequest, "body must be { name, source, configVersion, testOverrides? }")
		return
	}
	body.Name, body.Source = strings.TrimSpace(body.Name), strings.TrimSpace(body.Source)
	if body.Name == "" || body.Source == "" {
		writeErr(w, domain.ErrBadPayload, "name and source are required")
		return
	}
	if _, ok := s.o.Registry.Get(body.ConfigVersion); !ok {
		writeErr(w, domain.ErrBadPayload, "configVersion is not a loaded config")
		return
	}
	ctx, cancel := reqCtx(r)
	defer cancel()
	run := &domain.Run{ID: randomCode(16), Name: body.Name, Source: body.Source, ConfigVersion: body.ConfigVersion,
		Status: domain.RunOpen, CreatedAt: s.now(), CreatedFromIP: ClientIP(r, s.o.TrustedProxyHops), TestOverrides: body.TestOverrides}
	var err error
	for i := 0; i < 8; i++ {
		run.Code = randomCode(s.runCodeLength())
		if err = s.o.Store.CreateRun(ctx, run); !errors.Is(err, store.ErrConflict) {
			break
		}
	}
	if err != nil {
		writeErr(w, domain.ErrUnavailable, "could not create run")
		return
	}
	writeJSON(w, http.StatusCreated, map[string]string{"id": run.ID, "code": run.Code, "url": s.baseURL(r) + "/s/" + run.Code})
}

func (s *Server) runHasSessions(ctx context.Context, id string) (bool, error) {
	ss, err := s.o.Store.ListSessions(ctx, store.SessionFilter{RunID: id})
	return len(ss) > 0, err
}

func (s *Server) patchRun(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Name          *string `json:"name"`
		Source        *string `json:"source"`
		Status        *string `json:"status"`
		ConfigVersion *string `json:"configVersion"`
		TestOverrides *bool   `json:"testOverrides"`
	}
	if err := json.NewDecoder(io.LimitReader(r.Body, 8192)).Decode(&body); err != nil {
		writeErr(w, domain.ErrBadRequest, "invalid body")
		return
	}
	ctx, cancel := reqCtx(r)
	defer cancel()
	run, err := s.o.Store.GetRun(ctx, r.PathValue("id"))
	if errors.Is(err, store.ErrNotFound) {
		writeErr(w, domain.ErrNotFound, "no such run")
		return
	}
	if err != nil {
		writeErr(w, domain.ErrUnavailable, "store unavailable")
		return
	}
	if body.Name != nil {
		run.Name = strings.TrimSpace(*body.Name)
	}
	if body.Source != nil {
		run.Source = strings.TrimSpace(*body.Source)
	}
	if body.Status != nil {
		if *body.Status != domain.RunOpen && *body.Status != domain.RunClosed {
			writeErr(w, domain.ErrBadPayload, "status must be open or closed")
			return
		}
		run.Status = *body.Status
	}
	if body.TestOverrides != nil {
		run.TestOverrides = *body.TestOverrides
	}
	if body.ConfigVersion != nil && *body.ConfigVersion != run.ConfigVersion {
		has, err := s.runHasSessions(ctx, run.ID)
		if err != nil {
			writeErr(w, domain.ErrUnavailable, "store unavailable")
			return
		}
		if has {
			writeErr(w, domain.ErrConflict, "configVersion cannot change once the run has sessions (2.16)")
			return
		}
		if _, ok := s.o.Registry.Get(*body.ConfigVersion); !ok {
			writeErr(w, domain.ErrBadPayload, "configVersion is not a loaded config")
			return
		}
		run.ConfigVersion = *body.ConfigVersion
	}
	if err := s.o.Store.UpdateRun(ctx, run); err != nil {
		writeErr(w, domain.ErrUnavailable, "store unavailable")
		return
	}
	writeJSON(w, http.StatusOK, run)
}

func (s *Server) deleteRun(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := reqCtx(r)
	defer cancel()
	switch err := s.o.Store.DeleteRun(ctx, r.PathValue("id")); {
	case errors.Is(err, store.ErrNotFound):
		writeErr(w, domain.ErrNotFound, "no such run")
	case errors.Is(err, store.ErrConflict):
		writeErr(w, domain.ErrConflict, "the run has sessions")
	case err != nil:
		writeErr(w, domain.ErrUnavailable, "store unavailable")
	default:
		w.WriteHeader(http.StatusNoContent)
	}
}

// ---------- participants ----------

type participantRow struct {
	SessionID   string                   `json:"sessionId"`
	RunCode     string                   `json:"runCode"`
	Source      string                   `json:"source"`
	Group       *string                  `json:"group"`
	OpenedAt    time.Time                `json:"openedAt"`
	ConsentAt   *time.Time               `json:"consentAt"`
	FinishedAt  *time.Time               `json:"finishedAt"`
	Progress    domain.Progress          `json:"progress"`
	Status      string                   `json:"status"`
	LastStep    string                   `json:"lastStep"`
	LastEventAt time.Time                `json:"lastEventAt"`
	AudioCheck  *domain.AudioCheckResult `json:"audioCheck"`
	ReloadCount int                      `json:"reloadCount"`
	Flags       []string                 `json:"flags"`
}

func (s *Server) listParticipants(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	f := store.SessionFilter{RunID: q.Get("runId")}
	if st := q.Get("status"); st != "" {
		f.Statuses = []string{st}
	}
	if us := q.Get("updatedSince"); us != "" {
		t, err := time.Parse(time.RFC3339Nano, us)
		if err != nil {
			writeErr(w, domain.ErrBadRequest, "updatedSince must be RFC 3339")
			return
		}
		f.UpdatedSince = &t
	}
	limit := 500
	if l, err := strconv.Atoi(q.Get("limit")); err == nil && l > 0 && l <= 1000 {
		limit = l
	}
	offset, _ := strconv.Atoi(q.Get("cursor"))
	group, flagged := q.Get("group"), q.Get("flagged") == "true"
	serverTime := s.now()
	ctx, cancel := reqCtx(r)
	defer cancel()
	sessions, err := s.o.Store.ListSessions(ctx, f)
	if err != nil {
		writeErr(w, domain.ErrUnavailable, "store unavailable")
		return
	}
	var rows []participantRow
	for _, x := range sessions {
		if group != "" && x.Group != group {
			continue
		}
		if flagged && len(x.Flags) == 0 {
			continue
		}
		var g *string
		if x.Group != "" {
			gg := x.Group
			g = &gg
		}
		flags := x.Flags
		if flags == nil {
			flags = []string{}
		}
		rows = append(rows, participantRow{SessionID: x.ID, RunCode: x.RunCode, Source: x.Source, Group: g,
			OpenedAt: x.OpenedAt, ConsentAt: x.ConsentAt, FinishedAt: x.FinishedAt, Progress: domain.ProgressOf(x),
			Status: x.Status, LastStep: x.CurrentState, LastEventAt: x.LastEventAt, AudioCheck: x.AudioCheck,
			ReloadCount: x.ReloadCount, Flags: flags})
	}
	cursor := ""
	if offset > len(rows) {
		offset = len(rows)
	}
	rows = rows[offset:]
	if len(rows) > limit {
		rows = rows[:limit]
		cursor = strconv.Itoa(offset + limit)
	}
	if rows == nil {
		rows = []participantRow{}
	}
	writeJSON(w, http.StatusOK, map[string]any{"items": rows, "cursor": cursor, "serverTime": serverTime})
}

func (s *Server) participantDetail(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := reqCtx(r)
	defer cancel()
	id := r.PathValue("id")
	sess, err := s.o.Store.GetSession(ctx, id)
	if errors.Is(err, store.ErrNotFound) {
		writeErr(w, domain.ErrNotFound, "no such session")
		return
	}
	if err != nil {
		writeErr(w, domain.ErrUnavailable, "store unavailable")
		return
	}
	evs, err1 := s.o.Store.Events(ctx, id)
	rts, err2 := s.o.Store.Ratings(ctx, id)
	ans, err3 := s.o.Store.Answers(ctx, id)
	if err := errors.Join(err1, err2, err3); err != nil {
		writeErr(w, domain.ErrUnavailable, "store unavailable")
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"session": sess, "progress": domain.ProgressOf(sess),
		"events": evs, "ratings": rts, "answers": ans})
}

// maxDeleteReason is the longest reason a participant deletion takes.
const maxDeleteReason = 500

// deleteSession removes one participant's data on request (plan 06, section
// 6): the session with its events, ratings and answers; a tombstone without
// personal data remains. The run's counters are untouched (decision 1.6).
func (s *Server) deleteSession(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var body struct {
		Confirm string `json:"confirm"`
		Reason  string `json:"reason"`
	}
	if err := json.NewDecoder(io.LimitReader(r.Body, 8192)).Decode(&body); err != nil {
		writeErr(w, domain.ErrBadRequest, "body must be { confirm, reason }")
		return
	}
	body.Reason = strings.TrimSpace(body.Reason)
	if body.Confirm != id {
		writeErr(w, domain.ErrBadRequest, "confirm must repeat the session id")
		return
	}
	// DECISION: the limit counts characters (runes), not bytes.
	if body.Reason == "" || len([]rune(body.Reason)) > maxDeleteReason {
		writeErr(w, domain.ErrBadRequest, "reason must be 1 to 500 characters")
		return
	}
	ctx, cancel := reqCtx(r)
	defer cancel()
	sess, err := s.o.Store.GetSession(ctx, id)
	if errors.Is(err, store.ErrNotFound) {
		writeErr(w, domain.ErrNotFound, "no such session")
		return
	}
	if err != nil {
		writeErr(w, domain.ErrUnavailable, "store unavailable")
		return
	}
	t := &domain.Deleted{SessionID: id, RunID: sess.RunID, Group: sess.Group, DeletedAt: s.now(), Reason: body.Reason}
	switch err := s.o.Store.DeleteSession(ctx, id, t); {
	case errors.Is(err, store.ErrNotFound):
		writeErr(w, domain.ErrNotFound, "no such session")
		return
	case err != nil:
		s.o.Log.Error("session delete failed", "session", id, "err", err)
		writeErr(w, domain.ErrUnavailable, "store unavailable")
		return
	}
	s.o.Log.Info("participant data deleted", "session", id, "run", sess.RunID,
		"admin", ClientIP(r, s.o.TrustedProxyHops))
	w.WriteHeader(http.StatusNoContent)
}

// runAnalysis is the run dashboard document (plan 06, section 6).
func (s *Server) runAnalysis(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	opts := analysis.Options{Include: q.Get("include"), Group: q.Get("group")}
	switch opts.Include {
	case "":
		opts.Include = analysis.IncludeFinished
	case analysis.IncludeFinished, analysis.IncludeConsented:
	default:
		writeErr(w, domain.ErrBadRequest, "include must be finished or consented")
		return
	}
	switch opts.Group {
	case "":
		opts.Group = analysis.GroupBoth
	case analysis.GroupBoth, analysis.GroupChecker, analysis.GroupNeutral:
	default:
		writeErr(w, domain.ErrBadRequest, "group must be both, checker or neutral")
		return
	}
	// DECISION: exclude= present but empty excludes no flag (forced_allocation
	// still is); exclude absent means the defaults.
	if q.Has("exclude") {
		opts.Exclude = []string{}
		for _, f := range strings.Split(q.Get("exclude"), ",") {
			if f = strings.TrimSpace(f); f != "" {
				opts.Exclude = append(opts.Exclude, f)
			}
		}
	}
	ctx, cancel := reqCtx(r)
	defer cancel()
	run, err := s.o.Store.GetRun(ctx, r.PathValue("id"))
	if errors.Is(err, store.ErrNotFound) {
		writeErr(w, domain.ErrNotFound, "no such run")
		return
	}
	if err != nil {
		writeErr(w, domain.ErrUnavailable, "store unavailable")
		return
	}
	ct, err := s.o.Registry.Content(run.ConfigVersion)
	if err != nil {
		writeErr(w, domain.ErrUnavailable, "the run's config version is not loaded")
		return
	}
	sessions, err := s.o.Store.ListSessions(ctx, store.SessionFilter{RunID: run.ID})
	if err != nil {
		writeErr(w, domain.ErrUnavailable, "store unavailable")
		return
	}
	// DECISION: ratings and answers are read only for consented sessions;
	// a session without a group has none.
	ratings := map[string][]*domain.Rating{}
	answers := map[string][]*domain.AnswerDoc{}
	for _, x := range sessions {
		if x.Group == "" {
			continue
		}
		rts, err1 := s.o.Store.Ratings(ctx, x.ID)
		ans, err2 := s.o.Store.Answers(ctx, x.ID)
		if err := errors.Join(err1, err2); err != nil {
			if errors.Is(err, store.ErrNotFound) {
				continue // deleted while reading
			}
			writeErr(w, domain.ErrUnavailable, "store unavailable")
			return
		}
		ratings[x.ID], answers[x.ID] = rts, ans
	}
	writeJSON(w, http.StatusOK, analysis.Compute(run, ct, sessions, ratings, answers, opts, s.now()))
}

func (s *Server) export(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	format := q.Get("format")
	if format != "csv" && format != "json" {
		writeErr(w, domain.ErrBadRequest, "format must be csv or json")
		return
	}
	opt := export.Options{RunID: q.Get("runId"), IncludeInvalid: q.Get("includeInvalid") == "true", IncludeIP: q.Get("includeIp") == "true"}
	ctx, cancel := reqCtx(r)
	defer cancel()
	data, err := export.Collect(ctx, s.o.Store, s.o.Registry, opt)
	if err != nil {
		writeErr(w, domain.ErrUnavailable, "export failed")
		return
	}
	stamp := s.now().Format("20060102-150405")
	w.Header().Set("Cache-Control", "no-store")
	if format == "csv" {
		w.Header().Set("Content-Type", "text/csv; charset=utf-8")
		w.Header().Set("Content-Disposition", `attachment; filename="study-export-`+stamp+`.csv"`)
		_ = export.WriteCSV(w, data, opt)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Content-Disposition", `attachment; filename="study-export-`+stamp+`.json"`)
	_ = export.WriteJSON(w, data, opt)
}

// ---------- configs ----------

func (s *Server) lockedVersions(ctx context.Context) (map[string]bool, error) {
	runs, err := s.o.Store.ListRuns(ctx)
	if err != nil {
		return nil, err
	}
	sessions, err := s.o.Store.ListSessions(ctx, store.SessionFilter{})
	if err != nil {
		return nil, err
	}
	withSessions := map[string]bool{}
	for _, x := range sessions {
		withSessions[x.RunID] = true
	}
	out := map[string]bool{}
	for _, run := range runs {
		if withSessions[run.ID] {
			out[run.ConfigVersion] = true
		}
	}
	return out, nil
}

func (s *Server) listConfigs(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := reqCtx(r)
	defer cancel()
	locked, err := s.lockedVersions(ctx)
	if err != nil {
		writeErr(w, domain.ErrUnavailable, "store unavailable")
		return
	}
	out := []map[string]any{}
	for _, e := range s.o.Registry.List() {
		rec := *e.Record
		if locked[rec.Version] {
			rec.Status = domain.ConfigLocked
		}
		lint := rec.LintReport
		if lint == nil {
			lint = []string{}
		}
		out = append(out, map[string]any{"version": rec.Version, "sha256": rec.SHA256, "status": rec.Status,
			"lintReport": lint, "uploadedAt": rec.UploadedAt, "gcsPath": rec.GCSPath, "scenePackPath": rec.ScenePackPath,
			"manifestPath": rec.ManifestPath, "slides": len(e.Content.Config.Slides()),
			"flaggedSlides": e.Content.FlaggedSlides(), "progressTotal": e.Content.Plan.Total,
			"estimatedAudio": len(e.Content.Estimated) > 0})
	}
	sort.Slice(out, func(i, j int) bool { return out[i]["version"].(string) < out[j]["version"].(string) })
	writeJSON(w, http.StatusOK, out)
}

func (s *Server) uploadConfig(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(io.LimitReader(r.Body, maxConfigLen))
	if err != nil {
		writeErr(w, domain.ErrBadRequest, "could not read body")
		return
	}
	ctx, cancel := reqCtx(r)
	defer cancel()
	locked, err := s.lockedVersions(ctx)
	if err != nil {
		writeErr(w, domain.ErrUnavailable, "store unavailable")
		return
	}
	e, err := s.o.Registry.Upload(ctx, data, func(v string) bool { return locked[v] })
	var ue *configs.UploadError
	switch {
	case errors.As(err, &ue):
		writeJSON(w, http.StatusUnprocessableEntity, &domain.ErrorBody{Error: &domain.APIError{Code: ue.Code,
			Message: strings.Join(ue.Errors, "; ")}})
		return
	case errors.Is(err, configs.ErrLocked):
		writeErr(w, domain.ErrConflict, "this config version is locked by a run with sessions (2.16)")
		return
	case err != nil:
		writeErr(w, domain.ErrUnavailable, "could not store the config")
		return
	}
	writeJSON(w, http.StatusCreated, map[string]any{"version": e.Record.Version, "status": e.Record.Status, "lintReport": e.Record.LintReport})
}

func (s *Server) getConfig(w http.ResponseWriter, r *http.Request) {
	e, ok := s.o.Registry.Get(r.PathValue("version"))
	if !ok {
		writeErr(w, domain.ErrNotFound, "no such config version")
		return
	}
	writeJSON(w, http.StatusOK, e.Content.Config)
}
