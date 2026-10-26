package http

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"time"

	"study/api/internal/domain"
	"study/api/internal/events"
	"study/api/internal/store"
	sc "study/api/pkg/studyconfig"
)

const (
	sessionCookie = "study_session"
	clientHeader  = "X-Study-Client"
	maxEventBody  = 64 << 10
)

// codeAlphabet has no look-alike characters (0/o, 1/l/i).
const codeAlphabet = "abcdefghjkmnpqrstuvwxyz23456789"

// Run link codes. DECISION (2026-10-25): the architecture's six characters were
// Claude's choice, not a requirement from Oliver; the default is now twelve
// (about 59 bits) and the length is configurable with RUN_CODE_LENGTH.
const (
	DefaultRunCodeLength = 12
	MinRunCodeLength     = 6
	MaxRunCodeLength     = 32
)

// runCodeLength returns the configured run code length within its bounds.
func (s *Server) runCodeLength() int {
	n := s.o.RunCodeLength
	switch {
	case n == 0:
		return DefaultRunCodeLength
	case n < MinRunCodeLength:
		return MinRunCodeLength
	case n > MaxRunCodeLength:
		return MaxRunCodeLength
	}
	return n
}

func randomCode(n int) string {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	for i := range b {
		b[i] = codeAlphabet[int(b[i])%len(codeAlphabet)]
	}
	return string(b)
}

func randomToken() string {
	b := make([]byte, 32) // 256 bits (section 6)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	return base64.RawURLEncoding.EncodeToString(b)
}

func hashToken(t string) string {
	h := sha256.Sum256([]byte(t))
	return hex.EncodeToString(h[:])
}

// DECISION: the session cookie holds "<sessionId>.<token>", so a reload can be
// matched to its session without a token-hash index; only the token's
// SHA-256 is stored.
func (s *Server) setSessionCookie(w http.ResponseWriter, id, token string) {
	http.SetCookie(w, &http.Cookie{Name: sessionCookie, Value: id + "." + token, Path: "/api/v1",
		HttpOnly: true, Secure: !s.o.InsecureCookies, SameSite: http.SameSiteStrictMode, MaxAge: 7 * 24 * 3600})
}

func readSessionCookie(r *http.Request) (id, token string, ok bool) {
	c, err := r.Cookie(sessionCookie)
	if err != nil {
		return "", "", false
	}
	id, token, ok = strings.Cut(c.Value, ".")
	return id, token, ok && id != "" && token != ""
}

// dropUnauthenticated answers a participant call without a valid token: 401,
// counted per IP, nothing else touched. Past the hourly limit it is 429.
func (s *Server) dropUnauthenticated(w http.ResponseWriter, r *http.Request, code, msg string) {
	ip := ClientIP(r, s.o.TrustedProxyHops)
	if s.authFail.Hit(ip) > AuthFailuresPerHour {
		writeErr(w, domain.ErrRateLimited, "too many unauthenticated requests")
		return
	}
	writeErr(w, code, msg)
}

type sessionHandler func(w http.ResponseWriter, r *http.Request, sess *domain.Session)

// participantAuth enforces section 6 on /api/v1/sessions/{id}/...: a valid
// token cookie (401 otherwise), the X-Study-Client header (403), and a token
// bound to the session id in the path (403 for another session's path).
func (s *Server) participantAuth(next sessionHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, token, ok := readSessionCookie(r)
		if !ok {
			s.dropUnauthenticated(w, r, domain.ErrNoSession, "no session token")
			return
		}
		if r.Header.Get(clientHeader) != "1" {
			writeErrStatus(w, http.StatusForbidden, ErrForbidden, "missing X-Study-Client header")
			return
		}
		ctx, cancel := reqCtx(r)
		defer cancel()
		sess, err := s.o.Store.GetSession(ctx, id)
		if err != nil || subtle.ConstantTimeCompare([]byte(sess.TokenHash), []byte(hashToken(token))) != 1 {
			if err != nil && !errors.Is(err, store.ErrNotFound) {
				s.o.Log.Error("session lookup", "err", err)
				writeErr(w, domain.ErrUnavailable, "store unavailable")
				return
			}
			s.dropUnauthenticated(w, r, domain.ErrBadToken, "invalid session token")
			return
		}
		if r.PathValue("id") != id {
			writeErrStatus(w, http.StatusForbidden, ErrForbidden, "token belongs to another session")
			return
		}
		next(w, r, sess)
	}
}

// participantFallback answers any other path under /api/v1/: without a valid
// token it is dropped like every unauthenticated call.
func (s *Server) participantFallback(w http.ResponseWriter, r *http.Request) {
	id, _, ok := readSessionCookie(r)
	if !ok {
		s.dropUnauthenticated(w, r, domain.ErrNoSession, "no session token")
		return
	}
	r.SetPathValue("id", id)
	s.participantAuth(func(w http.ResponseWriter, r *http.Request, _ *domain.Session) {
		writeErr(w, domain.ErrNotFound, "no such endpoint")
	})(w, r)
}

// CreateResponse is the body of POST /api/v1/runs/{code}/sessions.
type CreateResponse struct {
	SessionID         string       `json:"sessionId"`
	ConfigVersion     string       `json:"configVersion"`
	ScenePackURL      string       `json:"scenePackUrl"`
	AssetBaseURL      string       `json:"assetBaseUrl"`
	ExpectedClientSeq int64        `json:"expectedClientSeq"`
	Next              *domain.Step `json:"next"`
}

type createRequest struct {
	Screen struct {
		W int `json:"w"`
		H int `json:"h"`
	} `json:"screen"`
	TZ string `json:"tz"`
	// Test is honoured only on a run with TestOverrides (admin); otherwise it
	// is ignored, so a participant adding ?group= to a real link changes nothing.
	Test *struct {
		Group       string `json:"group"`
		DenialSlide string `json:"denialSlide"`
	} `json:"test,omitempty"`
}

// AssetBaseURL is the prefix for image asset keys; the consent PDF is at
// /assets/consent/<pdfAssetKey>.
// DECISION: content images are served from images/ (the public prefix), so
// assetBaseUrl is "/assets/images/".
const AssetBaseURL = "/assets/images/"

func (s *Server) createSession(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get(clientHeader) != "1" {
		writeErrStatus(w, http.StatusForbidden, ErrForbidden, "missing X-Study-Client header")
		return
	}
	ctx, cancel := reqCtx(r)
	defer cancel()
	code := r.PathValue("code")
	run, err := s.o.Store.GetRunByCode(ctx, code)
	if errors.Is(err, store.ErrNotFound) {
		writeErr(w, domain.ErrUnknownRun, "unknown run")
		return
	}
	if err != nil {
		writeErr(w, domain.ErrUnavailable, "store unavailable")
		return
	}
	if run.Status != domain.RunOpen {
		writeErr(w, domain.ErrRunClosed, "this study link is closed")
		return
	}
	ct, err := s.o.Registry.Content(run.ConfigVersion)
	if err != nil {
		s.o.Log.Error("run config not loaded", "run", run.ID, "version", run.ConfigVersion)
		writeErr(w, domain.ErrUnavailable, "study content unavailable")
		return
	}
	now := s.now()

	// Resume: a valid cookie for an unfinished session of this run (8.1).
	if id, token, ok := readSessionCookie(r); ok {
		if sess, err := s.o.Store.GetSession(ctx, id); err == nil &&
			subtle.ConstantTimeCompare([]byte(sess.TokenHash), []byte(hashToken(token))) == 1 &&
			sess.RunID == run.ID && !sess.Final() {
			var out *domain.Session
			err := s.o.Store.RunTx(ctx, id, func(tx domain.Tx) error {
				if err := domain.Reload(tx, ct, now); err != nil {
					return err
				}
				out = tx.Session()
				return nil
			})
			if err != nil {
				writeErr(w, domain.ErrUnavailable, "store unavailable")
				return
			}
			writeJSON(w, http.StatusOK, s.createBody(out, ct, now, true))
			return
		}
	}

	ip := ClientIP(r, s.o.TrustedProxyHops)
	n := s.sessionCreate.Hit(ip)
	if n > SessionsHardLimit {
		writeErr(w, domain.ErrRateLimited, "too many sessions from this address")
		return
	}
	if s.xffLogged.Add(1) <= logRawXFFFirstNCount {
		s.o.Log.Info("session creation client address", "xForwardedFor", r.Header.Values("X-Forwarded-For"),
			"remoteAddr", r.RemoteAddr, "ip", ip, "trustedProxyHops", s.o.TrustedProxyHops)
	}
	var body createRequest
	_ = json.NewDecoder(io.LimitReader(r.Body, 4096)).Decode(&body)
	forcedGroup, forcedDenial := "", ""
	if body.Test != nil && run.TestOverrides {
		g := strings.ToLower(strings.TrimSpace(body.Test.Group))
		if g != "" {
			if g != string(sc.GroupChecker) && g != string(sc.GroupNeutral) {
				writeErr(w, domain.ErrBadPayload, "test.group must be checker or neutral")
				return
			}
			forcedGroup = g
		}
		if d := strings.TrimSpace(body.Test.DenialSlide); d != "" {
			if forcedGroup == "" {
				writeErr(w, domain.ErrBadPayload, "test.denialSlide needs test.group")
				return
			}
			ok := false
			for _, id := range ct.FlaggedSlides() {
				ok = ok || id == d
			}
			if !ok {
				writeErr(w, domain.ErrBadPayload, "test.denialSlide is not a flagged slide of this config")
				return
			}
			forcedDenial = d
		}
	}

	token := randomToken()
	ua := r.UserAgent()
	if len(ua) > 1024 {
		ua = ua[:1024]
	}
	first := ct.Plan.Steps[0]
	sess := &domain.Session{
		ID: randomCode(12), RunID: run.ID, RunCode: run.Code, Source: run.Source, ConfigVersion: run.ConfigVersion,
		Status: domain.StatusOpened, CurrentState: first.Key(), ProgressIndex: 1,
		Deadlines: domain.Deadlines{StepStartedAt: &now},
		OpenedAt:  now, LastEventAt: now, UpdatedAt: now,
		TokenHash: hashToken(token), IP: ip, UserAgent: ua, UAFamily: UAFamily(ua),
		Flags: []string{}, Screen: domain.Screen{W: body.Screen.W, H: body.Screen.H}, TZ: truncate(body.TZ, 64),
		ActionChoices: map[string]string{}, ForcedGroup: forcedGroup, ForcedDenialSlide: forcedDenial,
	}
	for _, st := range domain.PreAllocationSteps {
		sess.Plan = append(sess.Plan, st.Key())
	}
	if n > SessionsFlagAfter {
		sess.AddFlag(domain.FlagRateLimited)
	}
	// Duplicate detection (1.21): same IP and user agent within 60 minutes.
	since := now.Add(-time.Hour)
	others, err := s.o.Store.ListSessions(ctx, store.SessionFilter{IP: ip, OpenedSince: &since})
	if err != nil {
		writeErr(w, domain.ErrUnavailable, "store unavailable")
		return
	}
	var dups []string
	for _, o := range others {
		if o.UserAgent == ua {
			dups = append(dups, o.ID)
		}
	}
	if len(dups) > 0 {
		sess.AddFlag(domain.FlagPossibleDuplicate)
	}
	if err := s.o.Store.CreateSession(ctx, sess); err != nil {
		writeErr(w, domain.ErrUnavailable, "store unavailable")
		return
	}
	for _, id := range dups {
		_ = s.o.Store.RunTx(ctx, id, func(tx domain.Tx) error {
			x := tx.Session()
			if !x.HasFlag(domain.FlagPossibleDuplicate) {
				x.AddFlag(domain.FlagPossibleDuplicate)
				x.UpdatedAt = now
			}
			return nil
		})
	}
	s.setSessionCookie(w, sess.ID, token)
	writeJSON(w, http.StatusCreated, s.createBody(sess, ct, now, false))
}

func truncate(s string, n int) string {
	if len(s) > n {
		return s[:n]
	}
	return s
}

func (s *Server) createBody(sess *domain.Session, ct *domain.Content, now time.Time, resumed bool) *CreateResponse {
	return &CreateResponse{SessionID: sess.ID, ConfigVersion: sess.ConfigVersion,
		ScenePackURL: "/assets/scenes/" + sess.ConfigVersion + ".json", AssetBaseURL: AssetBaseURL,
		ExpectedClientSeq: sess.LastClientSeq + 1, Next: ct.Describe(sess, now, resumed)}
}

func (s *Server) postEvent(w http.ResponseWriter, r *http.Request, sess *domain.Session) {
	ct, err := s.o.Registry.Content(sess.ConfigVersion)
	if err != nil {
		writeErr(w, domain.ErrUnavailable, "study content unavailable")
		return
	}
	if s.sessionEvents.Hit(sess.ID) > EventsPerMinute {
		writeErr(w, domain.ErrRateLimited, "too many events per minute")
		return
	}
	var req events.Request
	dec := json.NewDecoder(io.LimitReader(r.Body, maxEventBody))
	if err := dec.Decode(&req); err != nil || req.Type == "" {
		exp := sess.LastClientSeq + 1
		writeJSON(w, http.StatusBadRequest, &domain.ErrorBody{Error: &domain.APIError{Code: domain.ErrBadRequest,
			Message: "body must be { clientSeq, type, clientTs, payload }", ExpectedClientSeq: &exp,
			Next: ct.Describe(sess, s.now(), false)}})
		return
	}
	ctx, cancel := reqCtx(r)
	defer cancel()
	var res *domain.Result
	err = s.o.Store.RunTx(ctx, sess.ID, func(tx domain.Tx) error {
		var err error
		res, err = domain.HandleEvent(tx, ct, s.now(), &req)
		return err
	})
	if err != nil {
		s.o.Log.Error("event transaction", "session", sess.ID, "err", err)
		writeErr(w, domain.ErrUnavailable, "please retry")
		return
	}
	writeJSON(w, res.Status, res.Body)
}

func (s *Server) getNext(w http.ResponseWriter, r *http.Request, sess *domain.Session) {
	ct, err := s.o.Registry.Content(sess.ConfigVersion)
	if err != nil {
		writeErr(w, domain.ErrUnavailable, "study content unavailable")
		return
	}
	ctx, cancel := reqCtx(r)
	defer cancel()
	var res *domain.Result
	err = s.o.Store.RunTx(ctx, sess.ID, func(tx domain.Tx) error {
		var err error
		res, err = domain.Next(tx, ct, s.now())
		return err
	})
	if err != nil {
		writeErr(w, domain.ErrUnavailable, "please retry")
		return
	}
	writeJSON(w, res.Status, res.Body)
}
