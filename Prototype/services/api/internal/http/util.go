package http

import (
	"encoding/json"
	"net"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"

	"study/api/internal/domain"
)

// writeJSON writes a JSON body; every /api/ response is Cache-Control: no-store (12.4).
func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Cache-Control", "no-store")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}

// writeErr writes the error format of 8.3 without step fields.
func writeErr(w http.ResponseWriter, code, msg string) {
	writeErrStatus(w, domain.HTTPStatus(code), code, msg)
}

func writeErrStatus(w http.ResponseWriter, status int, code, msg string) {
	writeJSON(w, status, &domain.ErrorBody{Error: &domain.APIError{Code: code, Message: msg,
		Retryable: code == domain.ErrRateLimited || code == domain.ErrUnavailable || code == domain.ErrInternal}})
}

// ErrForbidden is used for a missing CSRF header or a token for another
// session.
// DECISION: 8.3 lists no 403 code; "forbidden" is added for these cases.
const ErrForbidden = "forbidden"

// ClientIP returns the client address (section 6). With hops > 0 it takes the
// hops-th entry from the right of X-Forwarded-For (1 = the entry Google's
// front end appends); with 0, or a header too short, the TCP peer address.
func ClientIP(r *http.Request, hops int) string {
	if hops > 0 {
		var parts []string
		for _, h := range r.Header.Values("X-Forwarded-For") {
			for _, p := range strings.Split(h, ",") {
				if p = strings.TrimSpace(p); p != "" {
					parts = append(parts, p)
				}
			}
		}
		if len(parts) >= hops {
			return parts[len(parts)-hops]
		}
	}
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return host
}

var (
	reEdge    = regexp.MustCompile(`Edg/`)
	reChrome  = regexp.MustCompile(`Chrome/|CriOS/`)
	reFirefox = regexp.MustCompile(`Firefox/|FxiOS/`)
	reSafari  = regexp.MustCompile(`Safari/`)
)

// UAFamily parses the browser and OS family from a user agent, e.g.
// "Chrome / Windows".
func UAFamily(ua string) string {
	b := "Other"
	switch {
	case reEdge.MatchString(ua):
		b = "Edge"
	case reFirefox.MatchString(ua):
		b = "Firefox"
	case reChrome.MatchString(ua):
		b = "Chrome"
	case reSafari.MatchString(ua):
		b = "Safari"
	}
	o := "Other"
	switch {
	case strings.Contains(ua, "iPad"):
		o = "iPadOS"
	case strings.Contains(ua, "iPhone"):
		o = "iOS"
	case strings.Contains(ua, "Android"):
		o = "Android"
	case strings.Contains(ua, "Windows"):
		o = "Windows"
	case strings.Contains(ua, "Mac OS X"), strings.Contains(ua, "Macintosh"):
		o = "macOS"
	case strings.Contains(ua, "CrOS"):
		o = "ChromeOS"
	case strings.Contains(ua, "Linux"):
		o = "Linux"
	}
	return b + " / " + o
}

// Limiter is an in-memory sliding-window counter per key (section 6: token
// buckets per instance; with few instances this is close enough).
type Limiter struct {
	mu     sync.Mutex
	window time.Duration
	hits   map[string][]time.Time
	now    func() time.Time
}

func NewLimiter(window time.Duration, now func() time.Time) *Limiter {
	return &Limiter{window: window, hits: map[string][]time.Time{}, now: now}
}

// Hit records one hit and returns the number of hits in the window,
// including this one.
func (l *Limiter) Hit(key string) int {
	l.mu.Lock()
	defer l.mu.Unlock()
	now := l.now()
	cut := now.Add(-l.window)
	h := l.hits[key]
	i := 0
	for i < len(h) && !h[i].After(cut) {
		i++
	}
	h = append(h[i:], now)
	l.hits[key] = h
	if len(l.hits) > 50000 {
		for k, v := range l.hits {
			if len(v) == 0 || !v[len(v)-1].After(cut) {
				delete(l.hits, k)
			}
		}
	}
	return len(h)
}

// Reset forgets every hit of a key.
func (l *Limiter) Reset(key string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	delete(l.hits, key)
}

// Count returns the hits in the window without recording one.
func (l *Limiter) Count(key string) int {
	l.mu.Lock()
	defer l.mu.Unlock()
	cut := l.now().Add(-l.window)
	n := 0
	for _, t := range l.hits[key] {
		if t.After(cut) {
			n++
		}
	}
	return n
}
