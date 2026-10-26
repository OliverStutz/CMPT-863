package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/rand/v2"
	"net/http"
	"net/http/cookiejar"
	"regexp"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

// Step is the part of the server's step descriptor (8.1) the generator reads.
type Step struct {
	StepID     string   `json:"stepId"`
	Scene      string   `json:"scene"`
	SceneRef   string   `json:"sceneRef"`
	AudioURL   string   `json:"audioUrl"`
	Point      string   `json:"point"`
	Transcript string   `json:"transcript"`
	OpensInMs  *int64   `json:"opensInMs"`
	ClosesInMs *int64   `json:"closesInMs"`
	Blocked    bool     `json:"blocked"`
	Options    []Choice `json:"options"`
	Overlay    *struct {
		Status   string `json:"status"`
		Markings []struct {
			ID         string `json:"id"`
			AppearAtMs int64  `json:"appearAtMs"`
		} `json:"markings"`
	} `json:"overlay"`
}

type Choice struct {
	ID    string `json:"id"`
	Label string `json:"label"`
}

type apiError struct {
	Code              string `json:"code"`
	RetryInMs         *int64 `json:"retryInMs"`
	ExpectedClientSeq *int64 `json:"expectedClientSeq"`
	Next              *Step  `json:"next"`
}

// reply is the union of the participant API's response bodies.
type reply struct {
	SessionID         string    `json:"sessionId"`
	ConfigVersion     string    `json:"configVersion"`
	ScenePackURL      string    `json:"scenePackUrl"`
	ExpectedClientSeq int64     `json:"expectedClientSeq"`
	Next              *Step     `json:"next"`
	Error             *apiError `json:"error"`
}

// Session is one simulated participant.
type Session struct {
	g       *Gen
	n       int
	ip      string
	hc      *http.Client
	ID      string
	seq     int64
	step    *Step
	pack    *Pack
	rng     *rand.Rand
	Done    atomic.Bool
	Err     error
	Group   string       // learned from the descriptors: checker sessions get overlays
	persona *Persona     // persona mode only; nil keeps the random behaviour
	Events  atomic.Int64 // events the server accepted
	Elapsed time.Duration
}

var silenceRe = regexp.MustCompile(`/_dev/silence/(\d+)\.wav$`)

// clipMs returns the duration of a step's clip: from the dev silence URL,
// which carries it, or by parsing the MP3 (cached per URL).
func (g *Gen) clipMs(url string) (int64, bool) {
	if m := silenceRe.FindStringSubmatch(url); m != nil {
		v, _ := strconv.ParseInt(m[1], 10, 64)
		return v, true
	}
	if v, ok := g.clipCache.Load(url); ok {
		d := v.(int64)
		return d, d >= 0
	}
	d := int64(-1)
	if resp, err := g.plain.Get(g.base + url); err == nil {
		b, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		if resp.StatusCode == 200 {
			if v, err := mp3DurationMs(b); err == nil {
				d = v
			}
		}
	}
	g.clipCache.Store(url, d)
	return d, d >= 0
}

func (s *Session) think(lo, hi float64) time.Duration {
	sec := (lo + s.rng.Float64()*(hi-lo)) * s.g.thinkMul
	return time.Duration(sec * float64(time.Second))
}

// sleepWithin sleeps d, but never past the window the descriptor gives
// (closesInMs minus a margin), so think time cannot cause a timeout.
func (s *Session) sleepWithin(d time.Duration) {
	if s.step.ClosesInMs != nil {
		lim := time.Duration(*s.step.ClosesInMs)*time.Millisecond - 1500*time.Millisecond
		if d > lim {
			d = lim
		}
	}
	if d > 0 {
		time.Sleep(d)
	}
}

func atLeast(d, min time.Duration) time.Duration {
	if d < min {
		return min
	}
	return d
}

// do sends one request, records it, and repeats it unchanged after a 5xx or
// a network error (an identical event body is a retry for the server).
func (s *Session) do(endpoint, method, path string, body any) (*reply, int, error) {
	var raw []byte
	if body != nil {
		raw, _ = json.Marshal(body)
	}
	var lastErr error
	for attempt := 0; attempt < 6; attempt++ {
		if attempt > 0 {
			s.g.stats.Retry()
			time.Sleep(time.Duration(200*attempt*attempt) * time.Millisecond)
		}
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		req, _ := http.NewRequestWithContext(ctx, method, s.g.base+path, bytes.NewReader(raw))
		req.Header.Set("X-Study-Client", "1")
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("User-Agent", userAgent)
		req.Header.Set("X-Forwarded-For", s.ip)
		t0 := time.Now()
		resp, err := s.hc.Do(req)
		if err != nil {
			cancel()
			s.g.stats.NetError(endpoint)
			lastErr = err
			continue
		}
		b, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		cancel()
		d := time.Since(t0)
		var r reply
		_ = json.Unmarshal(b, &r)
		code := ""
		if r.Error != nil {
			code = r.Error.Code
		}
		s.g.stats.Record(endpoint, d, resp.StatusCode, code)
		if resp.StatusCode >= 500 {
			lastErr = fmt.Errorf("%s: HTTP %d %s", endpoint, resp.StatusCode, code)
			continue
		}
		// Follow the server: next step and the next clientSeq.
		if r.Error != nil {
			if r.Error.ExpectedClientSeq != nil {
				s.seq = *r.Error.ExpectedClientSeq
			}
			if r.Error.Next != nil {
				s.step = r.Error.Next
			}
		} else {
			if r.ExpectedClientSeq > 0 {
				s.seq = r.ExpectedClientSeq
			}
			if r.Next != nil {
				s.step = r.Next
			}
		}
		return &r, resp.StatusCode, nil
	}
	return nil, 0, lastErr
}

// event sends one event. A too_early answer is honoured: wait retryInMs and
// send again with the next clientSeq, as the participant app does.
func (s *Session) event(typ string, payload any) (*reply, error) {
	for tries := 0; tries < 5; tries++ {
		body := map[string]any{"clientSeq": s.seq, "type": typ, "clientTs": time.Now().UTC().Format(time.RFC3339Nano), "payload": payload}
		r, status, err := s.do("POST events "+typ, "POST", "/api/v1/sessions/"+s.ID+"/events", body)
		if err != nil {
			return nil, err
		}
		if status == 200 {
			s.Events.Add(1)
			return r, nil
		}
		if r.Error != nil && r.Error.Code == "too_early" && r.Error.RetryInMs != nil {
			time.Sleep(time.Duration(*r.Error.RetryInMs+30) * time.Millisecond)
			continue
		}
		return r, fmt.Errorf("%s at %s: HTTP %d %v", typ, s.step.Scene, status, codeOf(r))
	}
	return nil, fmt.Errorf("%s: still too early after 5 tries", typ)
}

func codeOf(r *reply) string {
	if r != nil && r.Error != nil {
		return r.Error.Code
	}
	return ""
}

func (s *Session) next() error {
	_, status, err := s.do("GET next", "GET", "/api/v1/sessions/"+s.ID+"/next", nil)
	if err == nil && status != 200 {
		err = fmt.Errorf("GET next: HTTP %d", status)
	}
	return err
}

const userAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0 Safari/537.36 study-load"

// Run drives the session through the whole flow, step by step, as the server
// describes it.
func (s *Session) Run() {
	t0 := time.Now()
	defer func() { s.Elapsed = time.Since(t0) }()
	jar, _ := cookiejar.New(nil)
	s.hc = &http.Client{Transport: s.g.transport, Jar: jar, CheckRedirect: func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }}
	r, status, err := s.do("POST sessions", "POST", "/api/v1/runs/"+s.g.runCode+"/sessions",
		map[string]any{"screen": map[string]int{"w": 1440, "h": 900}, "tz": "America/Vancouver"})
	if err != nil || status != 201 {
		s.Err = fmt.Errorf("create session: HTTP %d %v %v", status, codeOf(r), err)
		return
	}
	s.ID = r.SessionID
	if s.pack, err = s.g.scenePack(r.ScenePackURL); err != nil {
		s.Err = err
		return
	}
	for guard := 0; guard < 400; guard++ {
		if s.step == nil {
			s.Err = fmt.Errorf("no step")
			return
		}
		if err := s.handle(); err != nil {
			s.Err = fmt.Errorf("step %s (%s): %w", s.step.StepID, s.step.Scene, err)
			return
		}
		if s.Done.Load() {
			return
		}
	}
	s.Err = fmt.Errorf("too many steps")
}

func (s *Session) handle() error {
	st := s.step
	var err error
	switch st.Scene {
	case "deviceGate":
		time.Sleep(s.think(0.2, 0.6))
		_, err = s.event("device_checked", map[string]any{"widthPx": 1440, "heightPx": 900, "touch": false})
	case "consent":
		time.Sleep(atLeast(s.think(3.5, 8), 3100*time.Millisecond)) // server minimum 3 s
		_, err = s.event("consent_accepted", map[string]any{"consentVersion": s.pack.Consent.Version, "checkbox": true})
	case "audioCheck":
		if _, err = s.event("audio_check_played", map[string]any{}); err != nil {
			return err
		}
		if d, ok := s.g.clipMs(st.AudioURL); ok {
			time.Sleep(time.Duration(d) * time.Millisecond)
		}
		time.Sleep(s.think(1, 3))
		opts := s.pack.Setup.AudioCheck.Options
		choice := opts[s.rng.IntN(len(opts))].ID
		for _, o := range opts {
			if s.g.audioChoice != "" && o.ID == s.g.audioChoice {
				choice = o.ID
			}
		}
		_, err = s.event("audio_check_answered", map[string]any{"choiceId": choice})
	case "example", "opening", "slide", "question", "answer":
		err = s.audio()
	case "preStart":
		time.Sleep(atLeast(s.think(2.2, 5), 2100*time.Millisecond)) // server minimum 2 s
		_, err = s.event("setup_done", map[string]any{})
	case "postView":
		// The app waits opensInMs, then asks the server for the rating step.
		if st.OpensInMs != nil {
			time.Sleep(time.Duration(*st.OpensInMs) * time.Millisecond)
		}
		err = s.next()
	case "rating":
		if st.OpensInMs != nil && *st.OpensInMs > 0 {
			time.Sleep(time.Duration(*st.OpensInMs) * time.Millisecond)
		}
		if s.persona != nil {
			return s.personaRating()
		}
		s.sleepWithin(s.think(1.5, 5))
		p := map[string]any{"point": st.Point}
		for _, it := range s.pack.Rating.Items {
			p[it.ID] = it.Scale.Min + s.rng.IntN(it.Scale.Max-it.Scale.Min+1)
		}
		_, err = s.event("rating_submitted", p)
	case "action":
		if s.persona != nil {
			s.persona.learnGroup(groupOfOptions(st.Options))
			s.sleepWithin(s.think(2, 10))
			slide := s.pack.forSlide(st.SceneRef)
			_, err = s.event("action_chosen", map[string]any{"slideId": slide, "actionId": s.persona.action(slide, st.Options)})
			break
		}
		s.sleepWithin(s.think(1, 5))
		slide := s.pack.forSlide(st.SceneRef)
		_, err = s.event("action_chosen", map[string]any{"slideId": slide, "actionId": st.Options[s.rng.IntN(len(st.Options))].ID})
	case "budget":
		if s.persona != nil {
			s.sleepWithin(s.think(4, 15))
			_, err = s.event("budget_submitted", s.persona.budget(s.pack.scene(st.SceneRef)))
			break
		}
		s.sleepWithin(s.think(3, 10))
		sc := s.pack.scene(st.SceneRef)
		p := map[string]any{"optionId": sc.Options[s.rng.IntN(len(sc.Options))].ID, "confidence": nil}
		if sc.Confidence != nil {
			p["confidence"] = sc.Confidence.Scale.Min + s.rng.IntN(sc.Confidence.Scale.Max-sc.Confidence.Scale.Min+1)
		}
		_, err = s.event("budget_submitted", p)
	case "background":
		if s.persona != nil {
			s.sleepWithin(s.think(30, 90))
			_, err = s.event("background_submitted", map[string]any{"items": s.persona.background(s.pack.scene(st.SceneRef)), "hp": ""})
			break
		}
		s.sleepWithin(s.think(5, 15))
		_, err = s.event("background_submitted", map[string]any{"items": s.background(), "hp": ""})
	case "debrief":
		time.Sleep(atLeast(s.think(5.2, 9), 5100*time.Millisecond)) // server minimum 5 s
		_, err = s.event("finished", map[string]any{})
	case "done":
		if st.Blocked {
			return fmt.Errorf("blocked by the device gate")
		}
		s.Done.Store(true)
	default:
		return fmt.Errorf("unknown scene %q", st.Scene)
	}
	return err
}

// audio plays one audio step: scene_started, the clip's duration (with
// marking_shown at each marking's time on a checker slide), audio_ended.
func (s *Session) audio() error {
	st := s.step
	if _, err := s.event("scene_started", map[string]any{"sceneId": st.SceneRef}); err != nil {
		return err
	}
	started := time.Now()
	d, known := s.g.clipMs(st.AudioURL)
	if s.persona != nil {
		s.personaHears()
	}
	if st.Overlay != nil {
		s.Group = "checker"
		if st.Overlay.Status == "flagged" {
			for _, m := range st.Overlay.Markings {
				if known && m.AppearAtMs > d {
					continue
				}
				if w := time.Duration(m.AppearAtMs)*time.Millisecond - time.Since(started); w > 0 {
					time.Sleep(w)
				}
				if _, err := s.event("marking_shown", map[string]any{"sceneId": st.SceneRef, "markingId": m.ID}); err != nil {
					return err
				}
			}
		}
	}
	if known {
		if w := time.Duration(d)*time.Millisecond - time.Since(started); w > 0 {
			time.Sleep(w)
		}
	}
	// Without a known duration, too_early's retryInMs paces the resend.
	_, err := s.event("audio_ended", map[string]any{"sceneId": st.SceneRef})
	return err
}

func (s *Session) background() map[string]any {
	out := map[string]any{}
	for _, it := range s.pack.scene("background").Items {
		if s.rng.IntN(10) == 0 {
			out[it.ID] = "preferNotToSay"
			continue
		}
		switch it.Kind {
		case "single":
			var opts []Choice
			for _, o := range it.Options {
				if o.ID != it.OtherOptionID {
					opts = append(opts, o)
				}
			}
			if len(opts) > 0 {
				out[it.ID] = opts[s.rng.IntN(len(opts))].ID
			}
		case "scale":
			if it.Scale != nil {
				out[it.ID] = it.Scale.Min + s.rng.IntN(it.Scale.Max-it.Scale.Min+1)
			}
		case "country":
			out[it.ID] = "Canada"
		}
	}
	return out
}

// Pack is the part of the public scene pack (4.3) the generator reads.
type Pack struct {
	Consent struct {
		Version string `json:"version"`
	} `json:"consent"`
	Setup struct {
		AudioCheck struct {
			Options []Choice `json:"options"`
		} `json:"audioCheck"`
	} `json:"setup"`
	Rating struct {
		Items []RatingItem `json:"items"`
	} `json:"rating"`
	Scenes []PackScene `json:"scenes"`
}

type RatingItem struct {
	ID    string `json:"id"`
	Scale Scale  `json:"scale"`
}

type Scale struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

type PackScene struct {
	ID         string   `json:"id"`
	Kind       string   `json:"kind"`
	ForSlide   string   `json:"forSlide"`
	After      string   `json:"after"` // rating: the scene it follows
	Point      string   `json:"point"`
	Options    []Choice `json:"options"`
	Confidence *struct {
		Scale Scale `json:"scale"`
	} `json:"confidence"`
	Items []struct {
		ID            string   `json:"id"`
		Kind          string   `json:"kind"`
		Options       []Choice `json:"options"`
		OtherOptionID string   `json:"otherOptionId"`
		Scale         *Scale   `json:"scale"`
	} `json:"items"`
}

func (p *Pack) scene(id string) *PackScene {
	for i := range p.Scenes {
		if p.Scenes[i].ID == id {
			return &p.Scenes[i]
		}
	}
	return &PackScene{}
}

// flagged returns the flagged slides: those with an action scene.
func (p *Pack) flagged() map[string]bool {
	out := map[string]bool{}
	for _, sc := range p.Scenes {
		if sc.Kind == "action" && sc.ForSlide != "" {
			out[sc.ForSlide] = true
		}
	}
	return out
}

func (p *Pack) forSlide(ref string) string {
	if sc := p.scene(ref); sc.ForSlide != "" {
		return sc.ForSlide
	}
	return ref
}

var packMu sync.Mutex

func (g *Gen) scenePack(url string) (*Pack, error) {
	packMu.Lock()
	defer packMu.Unlock()
	if p, ok := g.packs[url]; ok {
		return p, nil
	}
	resp, err := g.plain.Get(g.base + url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var p Pack
	if err := json.NewDecoder(resp.Body).Decode(&p); err != nil {
		return nil, fmt.Errorf("scene pack: %w", err)
	}
	g.packs[url] = &p
	return &p, nil
}

// personaHears lets a persona learn from an audio step: on a flagged slide,
// an overlay means checker and none means neutral; an answer step's
// transcript tells reason from denial.
func (s *Session) personaHears() {
	st, p := s.step, s.persona
	flagged := s.pack.flagged()
	switch st.Scene {
	case "slide":
		if st.Overlay != nil && st.Overlay.Status == "flagged" {
			flagged[st.SceneRef] = true
		}
		if flagged[st.SceneRef] && p.Group == "" {
			if st.Overlay != nil {
				p.learnGroup("checker")
			} else {
				p.learnGroup("neutral")
			}
		}
	case "answer":
		n := len(flagged)
		if n == 0 {
			n = 3
		}
		p.observeAnswer(s.pack.forSlide(st.SceneRef), st.Transcript, n)
	}
}

// personaRating answers a rating step as the persona: the latent trust moves
// for what the rating follows (a flagged slide, or an answer to one, from the
// scene pack's "after"; without it the point's last letter, a or b), then
// the values with noise, or, for 3 % of ratings, nothing until the window
// closes and rating_timeout with every item null, as the participant app does.
func (s *Session) personaRating() error {
	st, p := s.step, s.persona
	after, slide := afterOther, ""
	if sc := s.pack.scene(st.SceneRef); sc.After != "" {
		ref := s.pack.scene(sc.After)
		switch {
		case ref.Kind == "answer":
			after, slide = afterAnswer, ref.ForSlide
		case s.pack.flagged()[sc.After]:
			after, slide = afterFlagged, sc.After
		}
	} else if n := len(st.Point); n > 2 {
		slide = "s" + st.Point[1:n-1]
		switch st.Point[n-1] {
		case 'a':
			after = afterFlagged
		case 'b':
			after = afterAnswer
		}
	}
	p.applyRating(after, slide)
	values := p.ratingValues(s.pack.Rating.Items)
	if values == nil {
		if st.ClosesInMs != nil {
			time.Sleep(time.Duration(*st.ClosesInMs) * time.Millisecond)
		}
		partial := map[string]any{}
		for _, it := range s.pack.Rating.Items {
			partial[it.ID] = nil
		}
		_, err := s.event("rating_timeout", map[string]any{"point": st.Point, "partial": partial})
		return err
	}
	s.sleepWithin(s.think(2, 8))
	values["point"] = st.Point
	_, err := s.event("rating_submitted", values)
	return err
}
