// Package http wires the routes of the Go service: the embedded SPAs, the
// assets, the participant API (8.1) and the admin API (8.2).
package http

import (
	"context"
	"io/fs"
	"log/slog"
	"net/http"
	"strings"
	"sync/atomic"
	"time"

	"study/api/internal/assets"
	"study/api/internal/configs"
	"study/api/internal/domain"
	"study/api/internal/secrets"
	"study/api/internal/store"
)

// Options configures a Server.
type Options struct {
	Store         store.Store
	Registry      *configs.Registry
	Source        assets.Source
	ParticipantFS fs.FS // contents of web/participant (index.html at the root)
	AdminFS       fs.FS // contents of web/admin

	PasswordHash *secrets.Secret // bcrypt hash of the shared admin password
	SessionKey   *secrets.Secret // HMAC key for the admin and gate cookies
	AccessKey    *secrets.Secret // admin gate key

	TrustedProxyHops int
	InsecureCookies  bool   // dev over plain http only
	PublicBaseURL    string // for run links; empty: from the request host
	// RunCodeLength is the length of a new run's link code (RUN_CODE_LENGTH;
	// default DefaultRunCodeLength, clamped to MinRunCodeLength..MaxRunCodeLength).
	// A run created through the admin API gets a random code of this length from
	// the look-alike-free alphabet; existing runs keep their codes.
	RunCodeLength int
	Now           func() time.Time
	Log           *slog.Logger
}

// Server is the HTTP server.
type Server struct {
	o Options

	sessionCreate *Limiter // per IP, 1 hour
	sessionEvents *Limiter // per session, 1 minute
	authFail      *Limiter // per IP, 1 hour: participant calls without a valid token
	loginLimit    *Limiter // per IP, 15 minutes
	adminGateFail *Limiter // per IP, 1 hour

	xffLogged atomic.Int32
	mux       *http.ServeMux
}

// Rate limits of section 6.
const (
	SessionsFlagAfter    = 5
	SessionsHardLimit    = 30
	EventsPerMinute      = 60
	AdminLoginAttempts   = 5
	AuthFailuresPerHour  = 120
	GateFailuresPerHour  = 120
	logRawXFFFirstNCount = 20
)

// New builds the server.
func New(o Options) *Server {
	if o.Now == nil {
		o.Now = func() time.Time { return time.Now().UTC() }
	}
	if o.Log == nil {
		o.Log = slog.Default()
	}
	s := &Server{o: o,
		sessionCreate: NewLimiter(time.Hour, o.Now),
		sessionEvents: NewLimiter(time.Minute, o.Now),
		authFail:      NewLimiter(time.Hour, o.Now),
		loginLimit:    NewLimiter(15*time.Minute, o.Now),
		adminGateFail: NewLimiter(time.Hour, o.Now),
	}
	s.routes()
	return s
}

func (s *Server) now() time.Time { return s.o.Now().UTC() }

func (s *Server) routes() {
	m := http.NewServeMux()
	ah := &assets.Handler{Source: s.o.Source, Generated: s.o.Registry, Bundle: s.o.ParticipantFS, Prefix: "/assets/"}
	m.Handle("GET /assets/", ah)
	// DECISION: the brief fixes /assets/ for content (the document uses /a/);
	// both serve the same allow-listed content, /assets/ also the participant
	// bundle files.
	m.Handle("GET /a/", &assets.Handler{Source: s.o.Source, Generated: s.o.Registry, Prefix: "/a/"})

	// Participant API (8.1).
	m.HandleFunc("GET /api/v1/healthz", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]bool{"ok": true})
	})
	m.HandleFunc("POST /api/v1/runs/{code}/sessions", s.createSession)
	m.HandleFunc("POST /api/v1/sessions/{id}/events", s.participantAuth(s.postEvent))
	m.HandleFunc("GET /api/v1/sessions/{id}/next", s.participantAuth(s.getNext))
	// Anything else under /api/v1/ is dropped like an unauthenticated call.
	m.HandleFunc("/api/v1/sessions/{id}/", s.participantAuth(func(w http.ResponseWriter, r *http.Request, _ *domain.Session) {
		writeErr(w, domain.ErrNotFound, "no such endpoint")
	}))
	m.HandleFunc("/api/v1/", s.participantFallback)

	// Admin API (8.2), behind the gate key and then the password.
	m.Handle("/api/admin/", s.adminGate(s.adminAPI()))

	// SPAs.
	m.Handle("/admin/", s.adminGate(spaHandler(s.o.AdminFS, "/admin/")))
	m.Handle("/admin", s.adminGate(http.RedirectHandler("/admin/", http.StatusFound)))
	m.Handle("/", spaHandler(s.o.ParticipantFS, "/"))
	s.mux = m
}

// ServeHTTP applies the common headers and rejects path traversal before
// routing (the mux would otherwise clean "/assets/../x" and redirect).
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, seg := range strings.Split(r.URL.Path, "/") {
		if seg == ".." || seg == "." {
			http.NotFound(w, r)
			return
		}
	}
	if strings.Contains(r.URL.Path, "\\") || strings.Contains(r.URL.Path, "//") {
		http.NotFound(w, r)
		return
	}
	h := w.Header()
	h.Set("X-Content-Type-Options", "nosniff")
	h.Set("Referrer-Policy", "no-referrer")
	if strings.HasPrefix(r.URL.Path, "/api/") {
		h.Set("Cache-Control", "no-store")
	}
	s.mux.ServeHTTP(w, r)
}

// ctx returns a request context with a deadline below Cloud Run's timeout.
func reqCtx(r *http.Request) (context.Context, context.CancelFunc) {
	return context.WithTimeout(r.Context(), 30*time.Second)
}
