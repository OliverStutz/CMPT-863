package http

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
	"time"

	"golang.org/x/crypto/bcrypt"

	"study/api/internal/assets"
	"study/api/internal/configs"
	"study/api/internal/domain"
	"study/api/internal/secrets"
	"study/api/internal/store/memory"
	"study/api/internal/testutil"
	sc "study/api/pkg/studyconfig"
)

const (
	realContent = "../../../../content"
	testPW      = "correct horse battery"
	testGateKey = "gate-key-0123456789"
)

var quiet = slog.New(slog.NewTextHandler(io.Discard, nil))

type env struct {
	t   *testing.T
	srv *Server
	st  *memory.Store
	reg *configs.Registry
	now time.Time
	dir string // content dir the registry loads from
}

// newEnv builds a server over a content dir. With fast=true the dir is a temp
// dir holding the generated fast config and the real schema.
func newEnv(t *testing.T, fast bool) *env {
	t.Helper()
	dir := realContent
	if fast {
		dir = t.TempDir()
		_ = os.MkdirAll(filepath.Join(dir, "schema"), 0o755)
		schema, err := os.ReadFile(filepath.Join(realContent, "schema", "study-config.schema.json"))
		if err != nil {
			t.Fatal(err)
		}
		_ = os.WriteFile(filepath.Join(dir, "schema", "study-config.schema.json"), schema, 0o644)
		b, _ := json.Marshal(testutil.FastConfig("test-fast", 4, 2, 3))
		_ = os.WriteFile(filepath.Join(dir, "study-config.test-fast.json"), b, 0o644)
	}
	e := &env{t: t, st: memory.New(), now: time.Date(2026, 11, 1, 10, 0, 0, 0, time.UTC), dir: dir}
	src := assets.Local{Root: dir}
	e.reg = &configs.Registry{Source: src, Store: e.st, AllowTest: fast, AssetURLPrefix: "/assets/", Log: quiet}
	if err := e.reg.Load(context.Background()); err != nil {
		t.Fatal(err)
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(testPW), bcrypt.MinCost)
	e.srv = New(Options{Store: e.st, Registry: e.reg, Source: src,
		PasswordHash: secrets.Static("h", hash), SessionKey: secrets.Static("k", []byte("session-key-for-tests")),
		AccessKey: secrets.Static("a", []byte(testGateKey)), Now: func() time.Time { return e.now }, Log: quiet})
	return e
}

func (e *env) run(version string) *domain.Run {
	r := &domain.Run{ID: "run-" + version, Code: "abc234", Name: "Test", Source: "Test", ConfigVersion: version,
		Status: domain.RunOpen, CreatedAt: e.now}
	if err := e.st.CreateRun(context.Background(), r); err != nil {
		e.t.Fatal(err)
	}
	return r
}

type reqOpt func(*http.Request)

func hdr(k, v string) reqOpt { return func(r *http.Request) { r.Header.Set(k, v) } }
func cookie(c *http.Cookie) reqOpt {
	return func(r *http.Request) {
		if c != nil {
			r.AddCookie(c)
		}
	}
}

func (e *env) do(method, path string, body any, opts ...reqOpt) *httptest.ResponseRecorder {
	var rd io.Reader
	switch b := body.(type) {
	case nil:
	case []byte:
		rd = bytes.NewReader(b)
	default:
		j, _ := json.Marshal(b)
		rd = bytes.NewReader(j)
	}
	r := httptest.NewRequest(method, path, rd)
	r.Header.Set("User-Agent", testutil.DesktopUA)
	for _, o := range opts {
		o(r)
	}
	w := httptest.NewRecorder()
	e.srv.ServeHTTP(w, r)
	return w
}

func findCookie(w *httptest.ResponseRecorder, name string) *http.Cookie {
	for _, c := range w.Result().Cookies() {
		if c.Name == name {
			return c
		}
	}
	return nil
}

// participant is an HTTP client for one session that keeps every body it
// receives.
type participant struct {
	e      *env
	cookie *http.Cookie
	id     string
	seq    int64
	bodies []string
	// slideSeen maps a slide id to the index in bodies where its step first arrived.
	slideSeen map[string]int
}

func (e *env) newParticipant(code string) *participant {
	p := &participant{e: e, slideSeen: map[string]int{}}
	w := e.do("POST", "/api/v1/runs/"+code+"/sessions", map[string]any{"screen": map[string]int{"w": 1440, "h": 900}, "tz": "America/Vancouver"},
		hdr(clientHeader, "1"))
	if w.Code != http.StatusCreated {
		e.t.Fatalf("create: %d %s", w.Code, w.Body)
	}
	p.cookie = findCookie(w, sessionCookie)
	if p.cookie == nil || !p.cookie.HttpOnly || p.cookie.Path != "/api/v1" || p.cookie.SameSite != http.SameSiteStrictMode || !p.cookie.Secure {
		e.t.Fatalf("session cookie %+v", p.cookie)
	}
	var cr CreateResponse
	p.record(w.Body.String(), &cr)
	p.id, p.seq = cr.SessionID, cr.ExpectedClientSeq
	if p.seq != 1 || cr.Next.Scene != domain.KindDeviceGate {
		e.t.Fatalf("create response %s", w.Body)
	}
	return p
}

func (p *participant) record(body string, v any) {
	p.bodies = append(p.bodies, body)
	if err := json.Unmarshal([]byte(body), v); err != nil {
		p.e.t.Fatalf("decode %s: %v", body, err)
	}
}

func (p *participant) noteStep(st *domain.Step) {
	if st != nil && st.Scene == domain.KindSlide {
		if _, ok := p.slideSeen[st.SceneRef]; !ok {
			p.slideSeen[st.SceneRef] = len(p.bodies) - 1
		}
	}
}

type eventResp struct {
	domain.EventOK
	Error *domain.APIError `json:"error"`
}

func (p *participant) send(typ string, payload any) (int, *eventResp) {
	w := p.e.do("POST", "/api/v1/sessions/"+p.id+"/events",
		map[string]any{"clientSeq": p.seq, "type": typ, "clientTs": p.e.now.Format(time.RFC3339Nano), "payload": payload},
		hdr(clientHeader, "1"), cookie(p.cookie))
	var r eventResp
	p.record(w.Body.String(), &r)
	if r.Error != nil {
		p.seq = *r.Error.ExpectedClientSeq
		p.noteStep(r.Error.Next)
	} else {
		p.seq = r.ExpectedClientSeq
		p.noteStep(r.Next)
	}
	return w.Code, &r
}

func (p *participant) mustOK(typ string, payload any) *eventResp {
	p.e.t.Helper()
	code, r := p.send(typ, payload)
	if code != 200 {
		p.e.t.Fatalf("%s: %d %+v", typ, code, r.Error)
	}
	return r
}

func (p *participant) next() *domain.Step {
	w := p.e.do("GET", "/api/v1/sessions/"+p.id+"/next", nil, hdr(clientHeader, "1"), cookie(p.cookie))
	if w.Code != 200 {
		p.e.t.Fatalf("next: %d %s", w.Code, w.Body)
	}
	var n domain.NextOK
	p.record(w.Body.String(), &n)
	p.seq = n.ExpectedClientSeq
	p.noteStep(n.Next)
	return n.Next
}

// walk drives the session to the end; stop, if set, ends the walk early when
// it returns true for the current step.
func (p *participant) walk(action string, stop func(*domain.Step) bool) {
	t := p.e.t
	t.Helper()
	for i := 0; i < 300; i++ {
		st := p.next()
		if stop != nil && stop(st) {
			return
		}
		switch st.Scene {
		case domain.KindDeviceGate:
			p.mustOK("device_checked", map[string]any{"widthPx": 1440, "heightPx": 900, "touch": false})
		case domain.KindConsent:
			p.e.now = p.e.now.Add(3 * time.Second)
			cfg, _ := p.e.reg.Get(p.version())
			p.mustOK("consent_accepted", map[string]any{"consentVersion": cfg.Content.Config.Consent.Version, "checkbox": true})
		case domain.KindAudioCheck:
			p.mustOK("audio_check_played", map[string]any{})
			p.mustOK("audio_check_answered", map[string]any{"choiceId": "river"})
		case domain.KindExample, domain.KindOpening, domain.KindSlide, domain.KindQuestion, domain.KindAnswer:
			p.mustOK("scene_started", map[string]any{"sceneId": st.SceneRef})
			if st.Scene == domain.KindSlide && st.Overlay != nil && st.Overlay.Status == "flagged" {
				for _, m := range st.Overlay.Markings {
					_, _ = p.send("marking_shown", map[string]any{"sceneId": st.SceneRef, "markingId": m.ID})
				}
			}
			for j := 0; j < 5; j++ {
				code, r := p.send("audio_ended", map[string]any{"sceneId": st.SceneRef})
				if code == 200 {
					break
				}
				if r.Error.Code != domain.ErrTooEarly {
					t.Fatalf("audio_ended: %+v", r.Error)
				}
				p.e.now = p.e.now.Add(time.Duration(*r.Error.RetryInMs) * time.Millisecond)
			}
		case domain.KindPreStart:
			p.e.now = p.e.now.Add(2 * time.Second)
			p.mustOK("setup_done", map[string]any{})
		case domain.KindPostView:
			p.e.now = p.e.now.Add(time.Duration(*st.OpensInMs) * time.Millisecond)
		case domain.KindRating:
			p.mustOK("rating_submitted", map[string]any{"point": st.Point, "trust": 6, "competence": 5, "benevolence": 5, "integrity": 5})
		case domain.KindAction:
			s, _ := p.e.st.GetSession(context.Background(), p.id)
			ct, _ := p.e.reg.Content(s.ConfigVersion)
			slide := ct.Plan.Steps[ct.Plan.Index(s.CurrentState)].SlideID
			p.mustOK("action_chosen", map[string]any{"slideId": slide, "actionId": action})
		case domain.KindBudget:
			p.mustOK("budget_submitted", map[string]any{"optionId": nil, "confidence": 5})
		case domain.KindBackground:
			p.mustOK("background_submitted", map[string]any{"items": map[string]any{}, "hp": ""})
		case domain.KindDebrief:
			p.e.now = p.e.now.Add(5 * time.Second)
			p.mustOK("finished", map[string]any{})
		case domain.KindDone:
			return
		}
	}
	t.Fatal("walk did not finish")
}

func (p *participant) version() string {
	s, _ := p.e.st.GetSession(context.Background(), p.id)
	return s.ConfigVersion
}

// ---------- session creation, consent, a slide, a rating ----------

func TestSessionConsentSlideRating(t *testing.T) {
	e := newEnv(t, true)
	e.run("test-fast")
	p := e.newParticipant("abc234")
	reached := false
	p.walk("nothing", func(st *domain.Step) bool {
		if st.Scene == domain.KindRating && st.Point == "T1" {
			reached = true
			return true
		}
		return false
	})
	if !reached {
		t.Fatal("T1 not reached")
	}
	s, _ := e.st.GetSession(context.Background(), p.id)
	if s.Group == "" || s.DenialSlide == "" || s.ConsentAt == nil || s.Status != domain.StatusOngoing {
		t.Fatalf("session %+v", s)
	}
	// The rating.
	p.e.now = p.e.now.Add(800 * time.Millisecond)
	p.mustOK("rating_submitted", map[string]any{"point": "T1", "trust": 9, "competence": 2, "benevolence": 3, "integrity": 4})
	rts, _ := e.st.Ratings(context.Background(), p.id)
	var t1 *domain.Rating
	for _, r := range rts {
		if r.Point == "T1" {
			t1 = r
		}
	}
	if t1 == nil || t1.Status != domain.RecordComplete || *t1.Values["trust"] != 9 || t1.LatencyMs == nil || *t1.LatencyMs != 800 {
		t.Fatalf("T1 = %+v", t1)
	}
}

func TestRatingOutOfRangeAndResume(t *testing.T) {
	e := newEnv(t, true)
	e.run("test-fast")
	p := e.newParticipant("abc234")
	p.walk("nothing", func(st *domain.Step) bool { return st.Scene == domain.KindRating })
	code, r := p.send("rating_submitted", map[string]any{"point": "T0", "trust": 11})
	if code != 422 || r.Error.Code != domain.ErrOutOfRange {
		t.Fatalf("%d %+v", code, r.Error)
	}
	p.mustOK("rating_submitted", map[string]any{"point": "T0", "trust": 10, "competence": 1})
	// A reload: POST sessions with the cookie resumes the same session.
	w := e.do("POST", "/api/v1/runs/abc234/sessions", map[string]any{}, hdr(clientHeader, "1"), cookie(p.cookie))
	var cr CreateResponse
	_ = json.Unmarshal(w.Body.Bytes(), &cr)
	if w.Code != 200 || cr.SessionID != p.id || cr.ExpectedClientSeq != p.seq || !cr.Next.Resumed {
		t.Fatalf("resume: %d %s", w.Code, w.Body)
	}
	s, _ := e.st.GetSession(context.Background(), p.id)
	if s.ReloadCount != 1 || !s.HasFlag(domain.FlagReloaded) {
		t.Fatalf("%+v", s)
	}
}

func TestUnknownAndClosedRun(t *testing.T) {
	e := newEnv(t, true)
	r := e.run("test-fast")
	if w := e.do("POST", "/api/v1/runs/zzzzzz/sessions", map[string]any{}, hdr(clientHeader, "1")); w.Code != 404 {
		t.Fatalf("unknown: %d", w.Code)
	}
	r.Status = domain.RunClosed
	_ = e.st.UpdateRun(context.Background(), r)
	if w := e.do("POST", "/api/v1/runs/abc234/sessions", map[string]any{}, hdr(clientHeader, "1")); w.Code != 410 {
		t.Fatalf("closed: %d", w.Code)
	}
}

// ---------- security 1: participant authentication trail ----------

func participantEndpoints(id string) [][2]string {
	return [][2]string{
		{"POST", "/api/v1/sessions/" + id + "/events"},
		{"GET", "/api/v1/sessions/" + id + "/next"},
		{"GET", "/api/v1/sessions/" + id + "/other"},
		{"GET", "/api/v1/anything"},
	}
}

func TestParticipantNoTokenIs401(t *testing.T) {
	e := newEnv(t, true)
	e.run("test-fast")
	p := e.newParticipant("abc234")
	before, _ := e.st.GetSession(context.Background(), p.id)
	for _, ep := range participantEndpoints(p.id) {
		w := e.do(ep[0], ep[1], map[string]any{"clientSeq": 1, "type": "device_checked"}, hdr(clientHeader, "1"))
		if w.Code != 401 {
			t.Errorf("%s %s: %d", ep[0], ep[1], w.Code)
		}
		// A forged token for the right session id is 401 too.
		bad := &http.Cookie{Name: sessionCookie, Value: p.id + ".forged"}
		if w := e.do(ep[0], ep[1], nil, hdr(clientHeader, "1"), cookie(bad)); w.Code != 401 {
			t.Errorf("forged %s %s: %d", ep[0], ep[1], w.Code)
		}
	}
	after, _ := e.st.GetSession(context.Background(), p.id)
	if after.LastEventNo != before.LastEventNo || after.LastClientSeq != before.LastClientSeq {
		t.Fatal("unauthenticated requests must not touch state")
	}
	if e.srv.authFail.Count("192.0.2.1") == 0 {
		t.Fatal("unauthenticated requests must be counted per IP")
	}
	if w := e.do("GET", "/api/v1/healthz", nil); w.Code != 200 {
		t.Fatalf("healthz %d", w.Code)
	}
}

func TestParticipantWrongSessionIs403(t *testing.T) {
	e := newEnv(t, true)
	e.run("test-fast")
	a := e.newParticipant("abc234")
	e.now = e.now.Add(time.Second)
	b := e.newParticipant("abc234")
	for _, ep := range participantEndpoints(b.id)[:3] {
		if w := e.do(ep[0], ep[1], map[string]any{"clientSeq": 1, "type": "device_checked"}, hdr(clientHeader, "1"), cookie(a.cookie)); w.Code != 403 {
			t.Errorf("%s %s: %d", ep[0], ep[1], w.Code)
		}
	}
	sb, _ := e.st.GetSession(context.Background(), b.id)
	if sb.LastEventNo != 0 {
		t.Fatal("state touched")
	}
}

func TestParticipantMissingClientHeaderIs403(t *testing.T) {
	e := newEnv(t, true)
	e.run("test-fast")
	p := e.newParticipant("abc234")
	for _, ep := range participantEndpoints(p.id) {
		if w := e.do(ep[0], ep[1], map[string]any{"clientSeq": 1, "type": "device_checked"}, cookie(p.cookie)); w.Code != 403 {
			t.Errorf("%s %s: %d", ep[0], ep[1], w.Code)
		}
	}
	if w := e.do("POST", "/api/v1/runs/abc234/sessions", map[string]any{}); w.Code != 403 {
		t.Errorf("create without header: %d", w.Code)
	}
}

func TestSessionCreationRateLimit(t *testing.T) {
	e := newEnv(t, true)
	e.run("test-fast")
	var last *httptest.ResponseRecorder
	for i := 0; i < SessionsHardLimit+1; i++ {
		last = e.do("POST", "/api/v1/runs/abc234/sessions", map[string]any{}, hdr(clientHeader, "1"))
		if i == SessionsFlagAfter {
			var cr CreateResponse
			_ = json.Unmarshal(last.Body.Bytes(), &cr)
			s, _ := e.st.GetSession(context.Background(), cr.SessionID)
			if !s.HasFlag(domain.FlagRateLimited) || !s.HasFlag(domain.FlagPossibleDuplicate) {
				t.Fatalf("flags %v", s.Flags)
			}
		}
	}
	if last.Code != 429 {
		t.Fatalf("31st session: %d", last.Code)
	}
}

func TestClientIP(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)
	r.RemoteAddr = "10.0.0.1:1234"
	r.Header.Set("X-Forwarded-For", "6.6.6.6, 198.51.100.7")
	if ip := ClientIP(r, 1); ip != "198.51.100.7" {
		t.Fatal(ip)
	}
	if ip := ClientIP(r, 0); ip != "10.0.0.1" {
		t.Fatal(ip)
	}
	if ip := ClientIP(r, 3); ip != "10.0.0.1" {
		t.Fatal(ip)
	}
}

// ---------- security 2: admin gate key ----------

func (e *env) gate() *http.Cookie {
	w := e.do("GET", "/admin/?k="+testGateKey, nil)
	if w.Code != 302 || w.Header().Get("Location") != "/admin/" {
		e.t.Fatalf("gate: %d %s", w.Code, w.Header().Get("Location"))
	}
	c := findCookie(w, gateCookie)
	if c == nil || !c.HttpOnly || !c.Secure || c.SameSite != http.SameSiteStrictMode {
		e.t.Fatalf("gate cookie %+v", c)
	}
	return c
}

func (e *env) login(gate *http.Cookie) *http.Cookie {
	w := e.do("POST", "/api/admin/login", map[string]string{"password": testPW}, cookie(gate))
	if w.Code != 200 {
		e.t.Fatalf("login %d %s", w.Code, w.Body)
	}
	return findCookie(w, adminCookie)
}

func TestAdminGate(t *testing.T) {
	e := newEnv(t, true)
	if w := e.do("GET", "/admin/", nil); w.Code != 404 {
		t.Fatalf("no key: %d", w.Code)
	}
	if w := e.do("GET", "/admin/?k=wrong", nil); w.Code != 404 || findCookie(w, gateCookie) != nil {
		t.Fatalf("wrong key: %d", w.Code)
	}
	if w := e.do("POST", "/api/admin/login", map[string]string{"password": testPW}); w.Code != 404 {
		t.Fatalf("login without gate: %d", w.Code)
	}
	if w := e.do("GET", "/api/admin/participants?k="+testGateKey, nil, hdr(adminHeader, "1")); w.Code != 404 {
		t.Fatalf("key on the API: %d", w.Code)
	}
	gate := e.gate()
	if w := e.do("GET", "/admin/", nil, cookie(gate)); w.Code != 200 {
		t.Fatalf("admin page with gate: %d", w.Code)
	}
	if w := e.do("GET", "/api/admin/participants", nil, cookie(gate), hdr(adminHeader, "1")); w.Code != 401 {
		t.Fatalf("gate, no login: %d", w.Code)
	}
	if w := e.do("POST", "/api/admin/login", map[string]string{"password": "nope"}, cookie(gate)); w.Code != 401 {
		t.Fatalf("bad password: %d", w.Code)
	}
	adm := e.login(gate)
	if w := e.do("GET", "/api/admin/participants", nil, cookie(gate), cookie(adm), hdr(adminHeader, "1")); w.Code != 200 {
		t.Fatalf("logged in: %d %s", w.Code, w.Body)
	}
	if w := e.do("GET", "/api/admin/participants", nil, cookie(gate), cookie(adm)); w.Code != 403 {
		t.Fatalf("missing X-Study-Admin: %d", w.Code)
	}
	// A login cookie without the gate cookie still sees nothing.
	if w := e.do("GET", "/api/admin/participants", nil, cookie(adm), hdr(adminHeader, "1")); w.Code != 404 {
		t.Fatalf("login without gate: %d", w.Code)
	}
	// After 12 hours both cookies have expired.
	e.now = e.now.Add(13 * time.Hour)
	if w := e.do("GET", "/api/admin/participants", nil, cookie(gate), cookie(adm), hdr(adminHeader, "1")); w.Code != 404 {
		t.Fatalf("expired gate: %d", w.Code)
	}
}

func TestAdminLoginRateLimit(t *testing.T) {
	e := newEnv(t, true)
	gate := e.gate()
	var w *httptest.ResponseRecorder
	for i := 0; i < AdminLoginAttempts+1; i++ {
		w = e.do("POST", "/api/admin/login", map[string]string{"password": "nope"}, cookie(gate))
	}
	if w.Code != 429 {
		t.Fatalf("%d", w.Code)
	}
}

// Round robin is per run (1.19 revised 2026-10-25): the first consent of every
// run gets the checker group, whatever other runs did. A run with test
// overrides honours ?group=/?denial= (sent as test{}), flags the session
// forced_allocation and leaves the run's counters alone; a run without them
// ignores the fields; invalid values are 400 (bad_payload).
func TestPerRunRoundRobinAndTestOverrides(t *testing.T) {
	e := newEnv(t, true)
	mk := func(id, code string, test bool) *domain.Run {
		r := &domain.Run{ID: id, Code: code, Name: "T", Source: "Test", ConfigVersion: "test-fast", Status: domain.RunOpen,
			CreatedAt: e.now, TestOverrides: test}
		if err := e.st.CreateRun(context.Background(), r); err != nil {
			t.Fatal(err)
		}
		return r
	}
	a, b, tr := mk("run-a", "aaaaaa", false), mk("run-b", "bbbbbb", false), mk("run-t", "tttttt", true)
	create := func(code string, test map[string]any) (*httptest.ResponseRecorder, *participant) {
		body := map[string]any{"screen": map[string]int{"w": 1440, "h": 900}, "tz": "UTC"}
		if test != nil {
			body["test"] = test
		}
		w := e.do("POST", "/api/v1/runs/"+code+"/sessions", body, hdr(clientHeader, "1"))
		if w.Code != http.StatusCreated {
			return w, nil
		}
		p := &participant{e: e, slideSeen: map[string]int{}, cookie: findCookie(w, sessionCookie)}
		var cr CreateResponse
		p.record(w.Body.String(), &cr)
		p.id, p.seq = cr.SessionID, cr.ExpectedClientSeq
		return w, p
	}
	group := func(p *participant) string {
		p.walk("", func(st *domain.Step) bool { return st.Scene == domain.KindAudioCheck })
		s, _ := e.st.GetSession(context.Background(), p.id)
		return s.Group + "/" + s.DenialSlide + "/" + strings.Join(s.Flags, ",")
	}
	_, pa1 := create(a.Code, nil)
	_, pa2 := create(a.Code, nil)
	_, pb1 := create(b.Code, nil)
	if g := group(pa1); !strings.HasPrefix(g, "checker/") {
		t.Fatalf("run a first consent: %s", g)
	}
	if g := group(pa2); !strings.HasPrefix(g, "neutral/") {
		t.Fatalf("run a second consent: %s", g)
	}
	if g := group(pb1); !strings.HasPrefix(g, "checker/") {
		t.Fatalf("run b first consent must not continue run a's sequence: %s", g)
	}
	ca, _ := e.st.Counters(context.Background(), a.ID)
	cb, _ := e.st.Counters(context.Background(), b.ID)
	if ca.Groups["checker"] != 1 || ca.Groups["neutral"] != 1 || cb.Groups["checker"] != 1 || cb.Groups["neutral"] != 0 {
		t.Fatalf("counters a=%+v b=%+v", ca.Groups, cb.Groups)
	}

	// Overrides ignored on an ordinary run.
	_, pa3 := create(a.Code, map[string]any{"group": "neutral", "denialSlide": "s3"})
	if g := group(pa3); !strings.HasPrefix(g, "checker/") || strings.Contains(g, domain.FlagForcedAllocation) {
		t.Fatalf("override on an ordinary run must be ignored: %s", g)
	}
	// Honoured on a test run, flagged, counters untouched.
	_, pt1 := create(tr.Code, map[string]any{"group": "neutral", "denialSlide": "s3"})
	if g := group(pt1); !strings.HasPrefix(g, "neutral/s3/") || !strings.Contains(g, domain.FlagForcedAllocation) {
		t.Fatalf("forced: %s", g)
	}
	_, pt2 := create(tr.Code, map[string]any{"group": "checker"})
	if g := group(pt2); !strings.HasPrefix(g, "checker/s2/") || !strings.Contains(g, domain.FlagForcedAllocation) {
		t.Fatalf("forced group only: %s", g)
	}
	ct, _ := e.st.Counters(context.Background(), tr.ID)
	if ct.Groups["checker"] != 0 || ct.Groups["neutral"] != 0 {
		t.Fatalf("forced sessions must not count: %+v", ct.Groups)
	}
	// Invalid values.
	if w, _ := create(tr.Code, map[string]any{"group": "both"}); w.Code != 400 {
		t.Fatalf("bad group: %d %s", w.Code, w.Body)
	}
	if w, _ := create(tr.Code, map[string]any{"group": "checker", "denialSlide": "s1"}); w.Code != 400 {
		t.Fatalf("bad denial slide: %d %s", w.Code, w.Body)
	}
	if w, _ := create(tr.Code, map[string]any{"denialSlide": "s3"}); w.Code != 400 {
		t.Fatalf("denial without group: %d %s", w.Code, w.Body)
	}
}

// The background screen has no deadline when timing.backgroundWindowMs is 0
// (2.12 revised 2026-10-25): no closesInMs, no lazy timeout however long it
// stays open, background_timeout refused, background_submitted accepted. The
// budget screen keeps formWindowMs, and a config without the field keeps the
// old behaviour.
func TestBackgroundUntimed(t *testing.T) {
	e := newEnv(t, true)
	cfg := testutil.FastConfig("test-untimed", 4, 2, 3)
	zero := sc.Ms(0)
	cfg.Timing.BackgroundWindowMs = &zero
	b, _ := json.Marshal(cfg)
	if err := os.WriteFile(filepath.Join(e.dir, "study-config.test-untimed.json"), b, 0o644); err != nil {
		t.Fatal(err)
	}
	if err := e.reg.Load(context.Background()); err != nil {
		t.Fatal(err)
	}
	run := e.run("test-untimed")
	p := e.newParticipant(run.Code)
	p.walk("explain", func(st *domain.Step) bool { return st.Scene == domain.KindBudget })
	if st := p.next(); st.Scene != domain.KindBudget || st.ClosesInMs == nil {
		t.Fatalf("budget step keeps its window: %+v", st)
	}
	p.mustOK("budget_submitted", map[string]any{"optionId": nil, "confidence": 5})
	st := p.next()
	if st.Scene != domain.KindBackground {
		t.Fatalf("at %s", st.Scene)
	}
	if st.ClosesInMs != nil {
		t.Fatalf("untimed background has closesInMs %d", *st.ClosesInMs)
	}
	e.now = e.now.Add(10 * time.Minute)
	if st := p.next(); st.Scene != domain.KindBackground {
		t.Fatalf("after 10 minutes the screen moved on to %s", st.Scene)
	}
	if code, r := p.send("background_timeout", map[string]any{"partial": map[string]any{}, "hp": ""}); code == 200 || r.Error == nil || r.Error.Code != domain.ErrWrongState {
		t.Fatalf("background_timeout on an untimed screen: %d %+v", code, r)
	}
	p.mustOK("background_submitted", map[string]any{"items": map[string]any{}, "hp": ""})
	if st := p.next(); st.Scene != domain.KindDebrief {
		t.Fatalf("after submit: %s", st.Scene)
	}

	// Without the field, the background screen is timed by formWindowMs as before.
	r2 := &domain.Run{ID: "run-test-fast-2", Code: "abc235", Name: "Test", Source: "Test", ConfigVersion: "test-fast",
		Status: domain.RunOpen, CreatedAt: e.now}
	if err := e.st.CreateRun(context.Background(), r2); err != nil {
		t.Fatal(err)
	}
	q := e.newParticipant(r2.Code)
	q.walk("explain", func(st *domain.Step) bool { return st.Scene == domain.KindBackground }) // first consent of its own run: checker
	if st := q.next(); st.ClosesInMs == nil {
		t.Fatalf("background without backgroundWindowMs must keep formWindowMs: %+v", st)
	}
}

// Correct logins are not limited, and a correct login clears earlier failures
// from the same address (test runs and the admin often share one network).
func TestAdminLoginSuccessNotLimited(t *testing.T) {
	e := newEnv(t, true)
	gate := e.gate()
	for i := 0; i < AdminLoginAttempts*3; i++ {
		if w := e.do("POST", "/api/admin/login", map[string]string{"password": testPW}, cookie(gate)); w.Code != 200 {
			t.Fatalf("correct login %d: %d %s", i+1, w.Code, w.Body)
		}
	}
	for i := 0; i < AdminLoginAttempts-1; i++ {
		e.do("POST", "/api/admin/login", map[string]string{"password": "nope"}, cookie(gate))
	}
	if w := e.do("POST", "/api/admin/login", map[string]string{"password": testPW}, cookie(gate)); w.Code != 200 {
		t.Fatalf("correct login after %d failures: %d", AdminLoginAttempts-1, w.Code)
	}
	for i := 0; i < AdminLoginAttempts-1; i++ {
		e.do("POST", "/api/admin/login", map[string]string{"password": "nope"}, cookie(gate))
	}
	if w := e.do("POST", "/api/admin/login", map[string]string{"password": "nope"}, cookie(gate)); w.Code != 401 {
		t.Fatalf("failure %d after the reset should still be a plain wrong password: %d", AdminLoginAttempts, w.Code)
	}
	if w := e.do("POST", "/api/admin/login", map[string]string{"password": testPW}, cookie(gate)); w.Code != 429 {
		t.Fatalf("after %d failures every attempt is limited: %d", AdminLoginAttempts, w.Code)
	}
}

func TestAdminRunsParticipantsExport(t *testing.T) {
	e := newEnv(t, true)
	gate := e.gate()
	adm := e.login(gate)
	admin := func(method, path string, body any) *httptest.ResponseRecorder {
		return e.do(method, path, body, cookie(gate), cookie(adm), hdr(adminHeader, "1"))
	}
	w := admin("POST", "/api/admin/runs", map[string]string{"name": "LinkedIn 1", "source": "LinkedIn", "configVersion": "test-fast"})
	if w.Code != 201 {
		t.Fatalf("create run %d %s", w.Code, w.Body)
	}
	var created struct{ ID, Code, URL string }
	_ = json.Unmarshal(w.Body.Bytes(), &created)
	if !regexp.MustCompile(`^[a-z2-9]{12}$`).MatchString(created.Code) || !strings.HasSuffix(created.URL, "/s/"+created.Code) {
		t.Fatalf("%+v", created)
	}
	p := e.newParticipant(created.Code)
	p.walk("explain", nil)
	if w := admin("PATCH", "/api/admin/runs/"+created.ID, map[string]string{"configVersion": "other"}); w.Code != 409 {
		t.Fatalf("config change with sessions: %d", w.Code)
	}
	if w := admin("DELETE", "/api/admin/runs/"+created.ID, nil); w.Code != 409 {
		t.Fatalf("delete with sessions: %d", w.Code)
	}
	w = admin("GET", "/api/admin/runs", nil)
	if !strings.Contains(w.Body.String(), `"finished":1`) {
		t.Fatalf("runs %s", w.Body)
	}
	w = admin("GET", "/api/admin/participants?flagged=true", nil)
	if !strings.Contains(w.Body.String(), "audio_check_wrong") {
		t.Fatalf("participants %s", w.Body)
	}
	w = admin("GET", "/api/admin/participants/"+p.id, nil)
	if w.Code != 200 || !strings.Contains(w.Body.String(), `"events"`) {
		t.Fatalf("detail %d", w.Code)
	}
	w = admin("GET", "/api/admin/export?format=csv", nil)
	lines := strings.Split(strings.TrimSpace(w.Body.String()), "\n")
	if w.Code != 200 || len(lines) != 1+7 || !strings.HasPrefix(lines[0], "sessionId,runCode") || strings.Contains(lines[0], ",ip") {
		t.Fatalf("csv %d lines\n%s", len(lines), w.Body)
	}
	w = admin("GET", "/api/admin/export?format=json&includeIp=true", nil)
	if w.Code != 200 || !strings.Contains(w.Body.String(), `"ip"`) {
		t.Fatalf("json export")
	}
	w = admin("GET", "/api/admin/configs", nil)
	if !strings.Contains(w.Body.String(), `"status":"locked"`) {
		t.Fatalf("configs %s", w.Body)
	}
}

// ---------- security 3: public exposure of content ----------

func TestAssetAllowlist(t *testing.T) {
	e := newEnv(t, false)
	for _, path := range []string{
		"/assets/study-config.v1.4.json",
		"/assets/schema/study-config.schema.json",
		"/assets/vlm-drafts/v1.3/selection.md",
		"/assets/../go.mod",
		"/assets/slides/v1.3/README.md",
		"/assets/slides/v1.3/make_slides.py",
		"/assets/audio/1.4.0/manifest.json",
		"/assets/%2e%2e/go.mod",
		"/a/study-config.v1.4.json",
		// A client that normalises "/assets/../services/api/go.mod" asks for
		// this; the SPA fallback must not answer a file path with index.html.
		"/services/api/go.mod",
		"/study-config.v1.4.json",
	} {
		if w := e.do("GET", path, nil); w.Code != 404 {
			t.Errorf("%s: %d", path, w.Code)
		}
	}
	w := e.do("GET", "/assets/slides/v1.3/s1.svg", nil)
	if w.Code != 200 {
		t.Fatalf("slide svg: %d", w.Code)
	}
	w = e.do("GET", "/assets/scenes/1.4.0.json", nil, hdr("Accept-Encoding", "gzip"))
	if w.Code != 200 || w.Header().Get("Cache-Control") != "no-cache" || w.Header().Get("ETag") == "" || w.Header().Get("Content-Encoding") != "gzip" {
		t.Fatalf("scene pack %d %v", w.Code, w.Header())
	}
	w = e.do("GET", "/assets/_dev/silence/1000.wav", nil, hdr("Range", "bytes=0-9"))
	if w.Code != 206 || w.Body.Len() != 10 {
		t.Fatalf("range: %d %d", w.Code, w.Body.Len())
	}
	if !assets.Public("audio/1.4.0/" + strings.Repeat("a", 64) + ".mp3") {
		t.Fatal("audio mp3 must be public")
	}
	if assets.Public("audio/1.4.0/manifest.json") || assets.Public("study-config.v1.4.json") || assets.Public("images/../x.png") {
		t.Fatal("non-public path allowed")
	}
}

func collectKeys(v any, out map[string]bool) {
	switch x := v.(type) {
	case map[string]any:
		for k, val := range x {
			out[k] = true
			collectKeys(val, out)
		}
	case []any:
		for _, val := range x {
			collectKeys(val, out)
		}
	}
}

// TestParticipantResponsesLeakNothing walks a checker session over HTTP with
// the real v1.4 config and checks every body the participant receives.
func TestParticipantResponsesLeakNothing(t *testing.T) {
	e := newEnv(t, false)
	ent, ok := e.reg.Get("1.4.0")
	if !ok {
		t.Fatal("v1.4 not loaded")
	}
	cfg := ent.Content.Config
	e.run("1.4.0")
	// Ensure a checker session: the first allocation goes to checker.
	p := e.newParticipant("abc234")
	w := e.do("GET", "/assets/scenes/1.4.0.json", nil)
	p.bodies = append(p.bodies, w.Body.String())
	// The public pack is the same for both groups, so it names no checker
	// (panel labels and colours travel with the overlay).
	if regexp.MustCompile(`(?i)checker|panelTitle|"colors"`).MatchString(w.Body.String()) {
		t.Error("the public scene pack mentions the checker panel")
	}
	p.walk("challenge", nil)
	s, _ := e.st.GetSession(context.Background(), p.id)
	if s.Group != "checker" || s.Status != domain.StatusFinished {
		t.Fatalf("session %s %s", s.Group, s.Status)
	}
	all := strings.Join(p.bodies, "\n")

	keys := map[string]bool{}
	for _, b := range p.bodies {
		var v any
		_ = json.Unmarshal([]byte(b), &v)
		collectKeys(v, keys)
	}
	for _, k := range []string{"group", "denialSlide", "answerType", "questionBy", "tokenHash", "ip", "flags", "correctOptionId",
		"actions", "answers", "voices", "lintReport", "createdFromIp", "reloadCount", "lateEventNo", "audioKey", "vlmRun", "anchorText"} {
		if keys[k] {
			t.Errorf("participant response contains key %q", k)
		}
	}
	if regexp.MustCompile(`(?i)denial`).MatchString(all) {
		t.Error(`participant responses contain the word "denial"`)
	}
	if regexp.MustCompile(`(?i)"(checker|neutral)"`).MatchString(all) {
		t.Error("participant responses contain a group label as a value")
	}
	// The other group's texts.
	for _, txt := range append([]string{cfg.Actions.Neutral.Prompt, cfg.Actions.Neutral.Note}, cfg.SceneOfKind("debrief").Paragraphs.Neutral...) {
		if txt != "" && strings.Contains(all, txt) {
			t.Errorf("neutral-only text in a checker session: %q", txt)
		}
	}
	// The answer variant not played on each flagged slide never appears.
	for _, id := range ent.Content.FlaggedSlides() {
		a := cfg.SceneByID(id).Answers
		played, other := "", ""
		switch a.Mode {
		case sc.AnswersWhole:
			played, other = a.Reason.Text, a.Denial.Text
			if id == s.DenialSlide {
				played, other = other, played
			}
		case sc.AnswersJoined:
			played = a.Opening.Text + " " + a.ReasonBody.Text + " " + a.Closing.Text
			if id == s.DenialSlide {
				played, other = a.Opening.Text+" "+a.Closing.Text, a.ReasonBody.Text
			}
		}
		if other != "" && strings.Contains(all, other) {
			t.Errorf("slide %s: the unplayed answer variant appears", id)
		}
		if !strings.Contains(all, played) {
			t.Errorf("slide %s: the played answer's transcript is missing", id)
		}
	}
	// Flag labels never appear before their slide's step.
	for _, id := range ent.Content.FlaggedSlides() {
		idx, ok := p.slideSeen[id]
		if !ok {
			t.Fatalf("slide %s never seen", id)
		}
		for _, m := range cfg.SceneByID(id).Checker.Markings {
			for i := 0; i < idx; i++ {
				if strings.Contains(p.bodies[i], m.Label) || strings.Contains(p.bodies[i], m.Detail) {
					t.Errorf("flag of %s appears in body %d before its slide step (%d)", id, i, idx)
				}
			}
		}
	}
}
