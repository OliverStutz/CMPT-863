package domain

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"regexp"
	"strings"
	"time"

	"study/api/internal/events"
	sc "study/api/pkg/studyconfig"
)

// Limits of section 5.3 and 6.
const (
	MaxEventsPerSession  = 600
	MaxVisibilityEvents  = 30
	ManyInvalidThreshold = 5
	ConsentMinView       = 3 * time.Second
	SetupDoneMinView     = 2 * time.Second
	FinishedMinView      = 5 * time.Second
)

// PreferNotToSay is the value a background item takes when the participant
// chooses "Prefer not to say" (2.14).
// DECISION: the document names the label (preferNotToSayLabel) but not the
// value sent; it is this literal string, valid for every answerable item.
const PreferNotToSay = "preferNotToSay"

// Result is the outcome of an event or a GET /next: an HTTP status and a body
// (EventOK, NextOK or ErrorBody).
type Result struct {
	Status int
	Body   any
}

// EventOK is the success body of POST events (8.1).
type EventOK struct {
	Accepted          bool   `json:"accepted"`
	EventNo           int64  `json:"eventNo"`
	ExpectedClientSeq int64  `json:"expectedClientSeq"`
	Next              *Step  `json:"next"`
	ContinueInMs      *int64 `json:"continueInMs,omitempty"` // audio_problem (5.3)
}

// Progress is { index, total }; total is null before allocation (5.1).
type Progress struct {
	Index int  `json:"index"`
	Total *int `json:"total"`
}

// NextOK is the body of GET /next (8.1).
type NextOK struct {
	Next              *Step    `json:"next"`
	ExpectedClientSeq int64    `json:"expectedClientSeq"`
	Progress          Progress `json:"progress"`
}

// ProgressOf returns the session's progress.
func ProgressOf(s *Session) Progress {
	p := Progress{Index: s.ProgressIndex}
	if s.ProgressTotal > 0 {
		t := s.ProgressTotal
		p.Total = &t
	}
	return p
}

type rejection struct {
	code      string
	msg       string
	retryInMs *int64
}

func reject(code, f string, a ...any) *rejection {
	return &rejection{code: code, msg: fmt.Sprintf(f, a...)}
}

// run holds one transaction's working state.
type run struct {
	tx  Tx
	ct  *Content
	s   *Session
	now time.Time
	tol time.Duration
	lat time.Duration

	// set by handlers
	eventNo      int64
	reason       string // stored reason for valid events, e.g. "restart"
	continueInMs *int64
}

func newRun(tx Tx, ct *Content, now time.Time) *run {
	t := ct.Config.Timing
	return &run{tx: tx, ct: ct, s: tx.Session(), now: now,
		tol: time.Duration(t.AudioEndToleranceMs) * time.Millisecond,
		lat: time.Duration(t.LatencyAllowanceMs) * time.Millisecond}
}

func ms(v int64) time.Duration { return time.Duration(v) * time.Millisecond }

func tp(t time.Time) *time.Time { return &t }

func i64(v int64) *int64 { return &v }

func ceilMs(d time.Duration) int64 { return int64(math.Ceil(float64(d) / float64(time.Millisecond))) }

func (r *run) idx() int { return r.ct.Plan.Index(r.s.CurrentState) }

func (r *run) step() PlanStep {
	i := r.idx()
	if i < 0 {
		return PlanStep{Kind: KindDone, Ref: RefDone}
	}
	return r.ct.Plan.Steps[i]
}

// enter moves the session to plan position i, with the step starting at at.
func (r *run) enter(i int, at time.Time) {
	st := r.ct.Plan.Steps[i]
	old := r.s.Deadlines
	d := Deadlines{StepStartedAt: tp(at)}
	timing := r.ct.Config.Timing
	switch st.Kind {
	case KindRating:
		d.RatingOpensAt, d.RatingClosesAt = old.RatingOpensAt, old.RatingClosesAt
		if d.RatingOpensAt == nil {
			d.RatingOpensAt = tp(at)
			d.RatingClosesAt = tp(at.Add(ms(timing.RatingWindowMs)))
		}
	case KindAction:
		d.ActionClosesAt = tp(at.Add(ms(timing.ActionWindowMs)))
	case KindBudget, KindBackground:
		d.FormClosesAt = tp(at.Add(ms(timing.FormWindowMs)))
	}
	r.s.Deadlines = d
	r.s.CurrentState = st.Key()
	r.s.RestartAllowed = false
	r.s.ProgressIndex = r.ct.Plan.ProgressIndex(i)
}

func (r *run) next(at time.Time) { r.enter(r.idx()+1, at) }

// putEvent stores an event under the next eventNo.
func (r *run) putEvent(ev *Event) {
	if ev.EventNo == 0 {
		r.s.LastEventNo++
		ev.EventNo = r.s.LastEventNo
	}
	if ev.Payload == nil {
		ev.Payload = map[string]any{}
	}
	if !ev.Valid {
		r.s.InvalidCount++
		if r.s.InvalidCount > ManyInvalidThreshold {
			r.s.AddFlag(FlagManyInvalidEvents)
		}
	}
	r.s.UpdatedAt = r.now
	r.tx.PutEvent(ev)
}

func (r *run) synthetic(typ string, at time.Time, payload map[string]any, before string) {
	r.putEvent(&Event{Type: typ, Source: events.SourceServer, ServerTs: at, Payload: payload,
		StateBefore: before, StateAfter: r.s.CurrentState, Valid: true})
}

// Advance runs lazy evaluation (5.4): it walks the plan forward over expired
// deadlines, writing synthetic timeout events, until it reaches a step with no
// expired deadline. A window expires only when now > closesAt + latency.
// A post-audio view ends when the rating opens.
// DECISION: a step entered by lazy evaluation starts at the deadline that
// closed the previous window (so an expired rating followed by silence also
// expires the action window that followed it).
func (r *run) advance() error {
	s := r.s
	if s.Final() || s.Group == "" {
		return nil
	}
	for guard := 0; guard < len(r.ct.Plan.Steps); guard++ {
		i := r.idx()
		if i < 0 {
			return nil
		}
		st := r.ct.Plan.Steps[i]
		d := s.Deadlines
		before := s.CurrentState
		switch st.Kind {
		case KindPostView:
			if d.RatingOpensAt != nil && !r.now.Before(*d.RatingOpensAt) {
				r.enter(i+1, *d.RatingOpensAt)
				continue
			}
		case KindRating:
			if d.RatingClosesAt != nil && r.now.After(d.RatingClosesAt.Add(r.lat)) {
				at := *d.RatingClosesAt
				if err := r.writeRating(st, nil, true, at); err != nil {
					return err
				}
				r.enter(i+1, at)
				r.synthetic(events.RatingTimeout, at, map[string]any{"point": st.Point}, before)
				continue
			}
		case KindAction:
			if d.ActionClosesAt != nil && r.now.After(d.ActionClosesAt.Add(r.lat)) {
				at := *d.ActionClosesAt
				if err := r.writeAction(st, nil, at); err != nil {
					return err
				}
				r.enter(i+1, at)
				r.synthetic(events.ActionTimeout, at, map[string]any{"slideId": st.SlideID}, before)
				continue
			}
		case KindBudget, KindBackground:
			if d.FormClosesAt != nil && r.now.After(d.FormClosesAt.Add(r.lat)) {
				at := *d.FormClosesAt
				if err := r.writeForm(st, nil, RecordTimeout, "", at); err != nil {
					return err
				}
				r.enter(i+1, at)
				typ := events.BudgetTimeout
				if st.Kind == KindBackground {
					typ = events.BackgroundTimeout
				}
				r.synthetic(typ, at, map[string]any{}, before)
				continue
			}
		}
		return nil
	}
	return nil
}

// Advance is the exported lazy evaluation used by the sweeper.
func Advance(tx Tx, ct *Content, now time.Time) error {
	return newRun(tx, ct, now).advance()
}

// Next answers GET /next: lazy evaluation, then the current step (9.5).
func Next(tx Tx, ct *Content, now time.Time) (*Result, error) {
	r := newRun(tx, ct, now)
	if err := r.advance(); err != nil {
		return nil, err
	}
	s := r.s
	return &Result{Status: http.StatusOK, Body: &NextOK{
		Next: ct.Describe(s, now, true), ExpectedClientSeq: s.LastClientSeq + 1, Progress: ProgressOf(s)}}, nil
}

// Reload records a reload (2.5, 9.5): reloadCount, flag reloaded, and one
// allowed scene_started repeat when an audio step's clip had started.
// DECISION: a reload is counted when POST sessions resumes a session (the
// HttpOnly cookie means a reloaded page must call it to learn its session id),
// not on GET /next, which always returns resumed: true.
func Reload(tx Tx, ct *Content, now time.Time) error {
	r := newRun(tx, ct, now)
	if err := r.advance(); err != nil {
		return err
	}
	s := r.s
	if s.Final() {
		return nil
	}
	s.ReloadCount++
	s.AddFlag(FlagReloaded)
	if r.step().Audio() && s.Deadlines.ClipStartedAt != nil {
		s.RestartAllowed = true
	}
	s.UpdatedAt = now
	return nil
}

// HandleEvent processes one client event inside a transaction (5.3, 6).
func HandleEvent(tx Tx, ct *Content, now time.Time, req *events.Request) (*Result, error) {
	r := newRun(tx, ct, now)
	s := r.s
	if s.LastEventNo >= MaxEventsPerSession {
		return r.errorResult(&rejection{code: ErrRateLimited, msg: "event limit for this session reached"}), nil
	}
	if err := r.advance(); err != nil {
		return nil, err
	}
	hash := req.Hash()
	expected := s.LastClientSeq + 1
	seq := req.ClientSeq

	// DECISION: an identical retry is answered from the stored result and not
	// stored again; 429 responses are not stored and consume no number.
	if seq != expected {
		if seq >= 1 && seq <= s.LastClientSeq {
			var status int
			var body, prevHash string
			if lr := s.LastResult; lr != nil && lr.ClientSeq == seq {
				status, body, prevHash = lr.Status, lr.Response, lr.BodyHash
			} else {
				ev, err := tx.ConsumedEvent(seq)
				if err != nil {
					return nil, err
				}
				if ev != nil {
					status, body, prevHash = ev.ResponseStatus, ev.Response, ev.BodyHash
				}
			}
			if prevHash != "" && prevHash == hash {
				return r.replay(status, body), nil
			}
			rej := reject(ErrDuplicateSeqMismatch, "clientSeq %d was already used with a different body", seq)
			r.storeClient(req, hash, false, rej, s.CurrentState, false)
			return r.errorResult(rej), nil
		}
		rej := reject(ErrSeqGap, "expected clientSeq %d, got %d", expected, seq)
		r.storeClient(req, hash, false, rej, s.CurrentState, false)
		return r.errorResult(rej), nil
	}

	// The number is consumed whatever validation decides.
	s.LastClientSeq = seq
	s.LastEventAt = now
	if s.Status == StatusAbandoned {
		s.Status = statusFromTimestamps(s)
	}
	before := s.CurrentState
	r.eventNo = s.LastEventNo + 1
	rej, err := r.dispatch(req)
	if err != nil {
		return nil, err
	}
	res := r.storeClient(req, hash, true, rej, before, true)
	return res, nil
}

func statusFromTimestamps(s *Session) string {
	switch {
	case s.FinishedAt != nil:
		return StatusFinished
	case s.OngoingAt != nil:
		return StatusOngoing
	case s.StartedAt != nil:
		return StatusStarted
	}
	return StatusOpened
}

// storeClient writes a client event and builds the response.
func (r *run) storeClient(req *events.Request, hash string, consumed bool, rej *rejection, before string, setLast bool) *Result {
	s := r.s
	seq := req.ClientSeq
	ev := &Event{ClientSeq: &seq, Type: req.Type, Source: events.SourceClient, ServerTs: r.now,
		ClientTs: req.ClientTs, Payload: req.PayloadMap(), StateBefore: before, Consumed: consumed, BodyHash: hash,
		Valid: rej == nil, Reason: r.reason}
	if rej != nil {
		ev.Reason = rej.code
	}
	if consumed {
		ev.EventNo = r.eventNo
		s.LastEventNo = r.eventNo
	}
	var res *Result
	if rej == nil {
		// The eventNo must be known before the body is built.
		if ev.EventNo == 0 {
			s.LastEventNo++
			ev.EventNo = s.LastEventNo
		}
		res = &Result{Status: http.StatusOK, Body: &EventOK{Accepted: true, EventNo: ev.EventNo,
			ExpectedClientSeq: s.LastClientSeq + 1, Next: r.ct.Describe(s, r.now, false), ContinueInMs: r.continueInMs}}
	} else {
		if ev.EventNo == 0 {
			s.LastEventNo++
			ev.EventNo = s.LastEventNo
		}
		res = r.errorResult(rej)
	}
	ev.StateAfter = s.CurrentState
	body, _ := json.Marshal(res.Body)
	if consumed {
		ev.ResponseStatus, ev.Response = res.Status, string(body)
		if setLast {
			s.LastResult = &LastResult{ClientSeq: seq, BodyHash: hash, EventNo: ev.EventNo, Status: res.Status, Response: string(body)}
		}
	}
	r.putEvent(ev)
	return res
}

func (r *run) errorResult(rej *rejection) *Result {
	exp := r.s.LastClientSeq + 1
	e := &APIError{Code: rej.code, Message: rej.msg, RetryInMs: rej.retryInMs, ExpectedClientSeq: &exp,
		Next: r.ct.Describe(r.s, r.now, false)}
	e.Retryable = rej.code == ErrTooEarly || rej.code == ErrRateLimited || rej.code == ErrUnavailable
	return &Result{Status: HTTPStatus(rej.code), Body: &ErrorBody{Error: e}}
}

// replay answers an identical retry with the stored result.
// DECISION: the stored body is returned with expectedClientSeq and next
// refreshed to the current values, so a client that retries an older number
// after later events does not take a stale number.
func (r *run) replay(status int, body string) *Result {
	exp := r.s.LastClientSeq + 1
	next := r.ct.Describe(r.s, r.now, false)
	var ok EventOK
	if status == http.StatusOK && json.Unmarshal([]byte(body), &ok) == nil {
		ok.ExpectedClientSeq, ok.Next = exp, next
		return &Result{Status: status, Body: &ok}
	}
	var eb ErrorBody
	if json.Unmarshal([]byte(body), &eb) == nil && eb.Error != nil {
		eb.Error.ExpectedClientSeq, eb.Error.Next = &exp, next
		return &Result{Status: status, Body: &eb}
	}
	return &Result{Status: status, Body: json.RawMessage(body)}
}

func (r *run) tooEarly(earliest time.Time, what string) *rejection {
	if !r.now.Before(earliest) {
		return nil
	}
	wait := ceilMs(earliest.Sub(r.now))
	return &rejection{code: ErrTooEarly, msg: fmt.Sprintf("%s arrived %d ms too early", what, wait), retryInMs: &wait}
}

func (r *run) wrongState(typ string) *rejection {
	return reject(ErrWrongState, "%s is not expected at step %s", typ, r.step().Kind)
}

func (r *run) expect(typ string, kinds ...string) *rejection {
	if r.s.Final() {
		return reject(ErrWrongState, "session is %s", r.s.Status)
	}
	k := r.step().Kind
	for _, want := range kinds {
		if k == want {
			return nil
		}
	}
	return r.wrongState(typ)
}

func (r *run) dispatch(req *events.Request) (*rejection, error) {
	switch req.Type {
	case events.DeviceChecked:
		return r.onDeviceChecked(req), nil
	case events.ConsentAccepted:
		return r.onConsent(req)
	case events.AudioCheckPlayed:
		if rej := r.expect(req.Type, KindAudioCheck); rej != nil {
			return rej, nil
		}
		if r.s.AudioCheck == nil {
			r.s.AudioCheck = &AudioCheckResult{}
		}
		r.s.AudioCheck.Plays++
		return nil, nil
	case events.AudioCheckAnswered:
		return r.onAudioCheckAnswered(req), nil
	case events.SceneStarted:
		return r.onSceneStarted(req), nil
	case events.MarkingShown:
		return r.onMarkingShown(req), nil
	case events.AudioEnded:
		return r.onAudioEnded(req), nil
	case events.AudioProblem:
		return r.onAudioProblem(req), nil
	case events.SetupDone:
		if rej := r.expect(req.Type, KindPreStart); rej != nil {
			return rej, nil
		}
		if rej := r.tooEarly(r.s.Deadlines.StepStartedAt.Add(SetupDoneMinView), req.Type); rej != nil {
			return rej, nil
		}
		r.next(r.now)
		return nil, nil
	case events.RatingSubmitted, events.RatingTimeout:
		return r.onRating(req)
	case events.ActionChosen, events.ActionTimeout:
		return r.onAction(req)
	case events.BudgetSubmitted, events.BudgetTimeout, events.BackgroundSubmitted, events.BackgroundTimeout:
		return r.onForm(req)
	case events.Finished:
		if rej := r.expect(req.Type, KindDebrief); rej != nil {
			return rej, nil
		}
		if rej := r.tooEarly(r.s.Deadlines.StepStartedAt.Add(FinishedMinView), req.Type); rej != nil {
			return rej, nil
		}
		r.s.Status = StatusFinished
		r.s.FinishedAt = tp(r.now)
		r.next(r.now)
		return nil, nil
	case events.Visibility:
		return r.onVisibility(req), nil
	}
	return reject(ErrBadPayload, "unknown event type %q", req.Type), nil
}

var phoneUA = regexp.MustCompile(`(?i)iphone|ipod|android.*mobile|windows phone|blackberry|\bmobi\b|mobile safari`)

// IsPhone reports whether a user agent belongs to a phone.
// DECISION: "not a phone" (2.9) is judged from the user agent (iPhone,
// Android with "Mobile", and similar); touch alone does not block, because
// landscape tablets pass the gate.
func IsPhone(ua string) bool {
	if strings.Contains(strings.ToLower(ua), "ipad") {
		return false
	}
	return phoneUA.MatchString(ua)
}

func (r *run) onDeviceChecked(req *events.Request) *rejection {
	if rej := r.expect(req.Type, KindDeviceGate); rej != nil {
		return rej
	}
	var p events.DeviceCheckedPayload
	if err := req.Decode(&p); err != nil || p.WidthPx == nil || p.HeightPx == nil {
		return reject(ErrBadPayload, "device_checked needs widthPx and heightPx")
	}
	gate := r.ct.Config.Options.DeviceGate
	pass := *p.WidthPx >= float64(gate.MinWidthPx) && *p.WidthPx > *p.HeightPx &&
		(gate.AllowPhones || !IsPhone(r.s.UserAgent))
	if !pass {
		r.s.Status = StatusBlocked
		return nil
	}
	r.next(r.now)
	return nil
}

func (r *run) onConsent(req *events.Request) (*rejection, error) {
	if rej := r.expect(req.Type, KindConsent); rej != nil {
		return rej, nil
	}
	var p events.ConsentPayload
	if err := req.Decode(&p); err != nil {
		return reject(ErrBadPayload, "consent payload: %v", err), nil
	}
	if p.ConsentVersion != r.ct.Config.Consent.Version {
		return reject(ErrBadPayload, "consentVersion %q does not match %q", p.ConsentVersion, r.ct.Config.Consent.Version), nil
	}
	if !p.Checkbox {
		return reject(ErrBadPayload, "checkbox must be true"), nil
	}
	if rej := r.tooEarly(r.s.Deadlines.StepStartedAt.Add(ConsentMinView), req.Type); rej != nil {
		return rej, nil
	}
	c, err := r.tx.Counters()
	if err != nil {
		return nil, err
	}
	group, denial := Allocate(c, r.ct.FlaggedSlides())
	r.tx.PutCounters(c)
	s := r.s
	s.Group, s.DenialSlide = string(group), denial
	s.Plan = r.ct.Plan.Keys()
	s.ProgressTotal = r.ct.Plan.Total
	s.ConsentAt = tp(r.now)
	r.next(r.now)
	return nil, nil
}

func (r *run) onAudioCheckAnswered(req *events.Request) *rejection {
	if rej := r.expect(req.Type, KindAudioCheck); rej != nil {
		return rej
	}
	s := r.s
	if s.AudioCheck == nil || s.AudioCheck.Plays < 1 {
		return reject(ErrWrongState, "the test clip must be played before answering")
	}
	var p events.ChoicePayload
	if err := req.Decode(&p); err != nil || p.ChoiceID == "" {
		return reject(ErrBadPayload, "audio_check_answered needs choiceId")
	}
	ac := r.ct.Config.Setup.AudioCheck
	if !hasChoice(ac.Options, p.ChoiceID) {
		return reject(ErrOutOfRange, "choiceId %q is not an option", p.ChoiceID)
	}
	correct := p.ChoiceID == ac.CorrectOptionID
	s.AudioCheck.ChoiceID, s.AudioCheck.Correct = p.ChoiceID, &correct
	if !correct {
		s.AddFlag(FlagAudioCheckWrong)
	}
	s.Status = StatusStarted
	s.StartedAt = tp(r.now)
	r.next(r.now)
	return nil
}

func hasChoice(opts []sc.Choice, id string) bool {
	for _, o := range opts {
		if o.ID == id {
			return true
		}
	}
	return false
}

func (r *run) audioStep(typ string) (PlanStep, *rejection) {
	if r.s.Final() {
		return PlanStep{}, reject(ErrWrongState, "session is %s", r.s.Status)
	}
	st := r.step()
	if !st.Audio() {
		return st, r.wrongState(typ)
	}
	return st, nil
}

func (r *run) onSceneStarted(req *events.Request) *rejection {
	st, rej := r.audioStep(req.Type)
	if rej != nil {
		return rej
	}
	var p events.ScenePayload
	_ = req.Decode(&p)
	if p.SceneID != st.Ref {
		return reject(ErrWrongState, "sceneId %q does not match the current scene %q", p.SceneID, st.Ref)
	}
	s := r.s
	if s.Deadlines.ClipStartedAt != nil {
		if !s.RestartAllowed {
			return reject(ErrWrongState, "scene already started; a restart is accepted only after a reload")
		}
		s.RestartAllowed = false
		r.reason = "restart"
	}
	s.Deadlines.ClipStartedAt = tp(r.now)
	if st.Kind == KindOpening && s.OngoingAt == nil {
		s.Status = StatusOngoing
		s.OngoingAt = tp(r.now)
	}
	return nil
}

func (r *run) clipEnd(st PlanStep) (time.Time, *sc.Clip) {
	clip := r.ct.ClipFor(r.s, st)
	var d int64
	if clip != nil {
		d = clip.DurationMs
	}
	return r.s.Deadlines.ClipStartedAt.Add(ms(d)), clip
}

func (r *run) onMarkingShown(req *events.Request) *rejection {
	st, rej := r.audioStep(req.Type)
	if rej != nil {
		return rej
	}
	if st.Kind != KindSlide || r.s.Group != string(sc.GroupChecker) {
		return r.wrongState(req.Type)
	}
	var p events.ScenePayload
	_ = req.Decode(&p)
	if p.SceneID != st.Ref {
		return reject(ErrWrongState, "sceneId %q does not match the current scene %q", p.SceneID, st.Ref)
	}
	if r.s.Deadlines.ClipStartedAt == nil {
		return reject(ErrWrongState, "scene not started")
	}
	ov := r.ct.Overlay(st.SlideID)
	if ov == nil || ov.Status != sc.PanelFlagged {
		return r.wrongState(req.Type)
	}
	for _, m := range ov.Markings {
		if m.ID == p.MarkingID {
			return r.tooEarly(r.s.Deadlines.ClipStartedAt.Add(ms(m.AppearAtMs)-r.tol), req.Type)
		}
	}
	return reject(ErrOutOfRange, "markingId %q is not on this slide", p.MarkingID)
}

func (r *run) onAudioEnded(req *events.Request) *rejection {
	st, rej := r.audioStep(req.Type)
	if rej != nil {
		return rej
	}
	var p events.ScenePayload
	_ = req.Decode(&p)
	if p.SceneID != st.Ref {
		return reject(ErrWrongState, "sceneId %q does not match the current scene %q", p.SceneID, st.Ref)
	}
	s := r.s
	if s.Deadlines.ClipStartedAt == nil {
		return reject(ErrWrongState, "scene_started must come first")
	}
	end, _ := r.clipEnd(st)
	if rej := r.tooEarly(end.Add(-r.tol), req.Type); rej != nil {
		return rej
	}
	timing := r.ct.Config.Timing
	switch st.Kind {
	case KindOpening, KindSlide, KindAnswer:
		view := timing.PostAudioViewMs
		if st.Kind == KindAnswer {
			view = timing.PostAnswerViewMs
		}
		r.next(r.now) // post-audio view
		opens := r.now.Add(ms(view))
		s.Deadlines.RatingOpensAt = tp(opens)
		s.Deadlines.RatingClosesAt = tp(opens.Add(ms(timing.RatingWindowMs)))
	default: // example -> preStart, question -> answer
		r.next(r.now)
	}
	return nil
}

func (r *run) onAudioProblem(req *events.Request) *rejection {
	st, rej := r.audioStep(req.Type)
	if rej != nil {
		return rej
	}
	r.s.AddFlag(FlagAudioProblem)
	var cont int64
	if r.s.Deadlines.ClipStartedAt != nil {
		end, _ := r.clipEnd(st)
		cont = max(0, ceilMs(end.Sub(r.now)))
	} else if clip := r.ct.ClipFor(r.s, st); clip != nil {
		cont = clip.DurationMs
	}
	r.continueInMs = &cont
	return nil
}

func (r *run) onVisibility(req *events.Request) *rejection {
	s := r.s
	s.VisibilityCount++
	if s.VisibilityCount > MaxVisibilityEvents {
		return reject(ErrRateLimited, "at most %d visibility events per session", MaxVisibilityEvents)
	}
	var p events.VisibilityPayload
	_ = req.Decode(&p)
	if p.Hidden && r.step().Audio() && s.Deadlines.ClipStartedAt != nil {
		s.AddFlag(FlagTabHiddenDuringAudio)
	}
	return nil
}

// lateStep finds an earlier step of a kind matching pred (for late
// submissions). It returns -1 when there is none.
func (r *run) earlierStep(kind string, pred func(PlanStep) bool) int {
	cur := r.idx()
	for i := 0; i < cur && i < len(r.ct.Plan.Steps); i++ {
		st := r.ct.Plan.Steps[i]
		if st.Kind == kind && pred(st) {
			return i
		}
	}
	return -1
}

func lateRej(what string) *rejection {
	return reject(ErrLate, "%s arrived after its window closed; stored as a late submission", what)
}

// parseScaleValue reads an item value: null or missing gives nil; otherwise
// an integer within the scale.
func parseScaleValue(raw json.RawMessage, scale sc.ScaleSpec, what string) (*int64, *rejection) {
	if len(raw) == 0 || string(raw) == "null" {
		return nil, nil
	}
	var f float64
	if err := json.Unmarshal(raw, &f); err != nil {
		return nil, reject(ErrBadPayload, "%s must be a number or null", what)
	}
	if f != math.Trunc(f) || f < scale.Min || f > scale.Max {
		return nil, reject(ErrOutOfRange, "%s must be an integer from %g to %g", what, scale.Min, scale.Max)
	}
	v := int64(f)
	return &v, nil
}

func (r *run) onRating(req *events.Request) (*rejection, error) {
	if r.s.Final() {
		return reject(ErrWrongState, "session is %s", r.s.Status), nil
	}
	var p events.RatingPayload
	if err := req.Decode(&p); err != nil {
		return reject(ErrBadPayload, "rating payload: %v", err), nil
	}
	var top map[string]json.RawMessage
	_ = req.Decode(&top)
	items := top
	if req.Type == events.RatingTimeout {
		items = p.Partial
	}
	values := map[string]*int64{}
	for _, it := range r.ct.Config.Rating.Items {
		v, rej := parseScaleValue(items[it.ID], it.Scale, it.ID)
		if rej != nil {
			return rej, nil
		}
		values[it.ID] = v
	}

	st := r.step()
	// A rating sent within the tolerance before the post-audio view ends.
	if st.Kind == KindPostView {
		i := r.idx()
		if i+1 < len(r.ct.Plan.Steps) && r.ct.Plan.Steps[i+1].Kind == KindRating && r.ct.Plan.Steps[i+1].Point == p.Point {
			if rej := r.tooEarly(r.s.Deadlines.RatingOpensAt.Add(-r.tol), req.Type); rej != nil {
				return rej, nil
			}
			r.enter(i+1, *r.s.Deadlines.RatingOpensAt)
			st = r.step()
		}
	}
	if st.Kind != KindRating || st.Point != p.Point {
		if j := r.earlierStep(KindRating, func(x PlanStep) bool { return x.Point == p.Point }); j >= 0 {
			rt, err := r.tx.Rating(p.Point)
			if err != nil {
				return nil, err
			}
			if rt != nil && rt.Status == RecordTimeout {
				if rt.LateEventNo == nil {
					rt.LateEventNo = i64(r.eventNo)
					r.tx.PutRating(rt)
				}
				return lateRej("rating " + p.Point), nil
			}
		}
		return reject(ErrWrongState, "rating %q is not the current step", p.Point), nil
	}
	d := r.s.Deadlines
	if req.Type == events.RatingSubmitted {
		if rej := r.tooEarly(d.RatingOpensAt.Add(-r.tol), req.Type); rej != nil {
			return rej, nil
		}
	} else if rej := r.tooEarly(d.RatingClosesAt.Add(-r.tol), req.Type); rej != nil {
		return rej, nil
	}
	if err := r.writeRating(st, values, req.Type == events.RatingTimeout, r.now); err != nil {
		return nil, err
	}
	r.next(r.now)
	return nil, nil
}

// writeRating writes the rating document of a rating step. values nil means
// nothing arrived (server timeout).
func (r *run) writeRating(st PlanStep, values map[string]*int64, timeout bool, at time.Time) error {
	rt := &Rating{Point: st.Point, SlideID: st.SlideID, Phase: st.Phase, Values: map[string]*int64{},
		ItemStatus: map[string]string{}, OpenedAt: r.s.Deadlines.RatingOpensAt}
	answered := 0
	for _, it := range r.ct.Config.Rating.Items {
		v := values[it.ID]
		rt.Values[it.ID] = v
		switch {
		case v != nil:
			rt.ItemStatus[it.ID] = ItemAnswered
			answered++
		case timeout:
			rt.ItemStatus[it.ID] = ItemTimedOut
		default:
			rt.ItemStatus[it.ID] = ItemMissing
		}
	}
	switch {
	case timeout:
		rt.Status = RecordTimeout
	case answered == len(r.ct.Config.Rating.Items):
		rt.Status = RecordComplete
	default:
		rt.Status = RecordPartial
	}
	if values != nil {
		rt.SubmittedAt = tp(at)
		if rt.OpenedAt != nil {
			rt.LatencyMs = i64(at.Sub(*rt.OpenedAt).Milliseconds())
		}
	}
	r.tx.PutRating(rt)
	return nil
}

func (r *run) onAction(req *events.Request) (*rejection, error) {
	if r.s.Final() {
		return reject(ErrWrongState, "session is %s", r.s.Status), nil
	}
	var p events.ActionPayload
	if err := req.Decode(&p); err != nil {
		return reject(ErrBadPayload, "action payload: %v", err), nil
	}
	st := r.step()
	if st.Kind != KindAction || st.SlideID != p.SlideID {
		if j := r.earlierStep(KindAction, func(x PlanStep) bool { return x.SlideID == p.SlideID }); j >= 0 {
			doc, err := r.tx.Answers(AnswersActions)
			if err != nil {
				return nil, err
			}
			if doc != nil {
				if rec, ok := doc.Actions[p.SlideID]; ok && rec.TimedOut {
					if rec.LateEventNo == nil {
						rec.LateEventNo = i64(r.eventNo)
						doc.Actions[p.SlideID] = rec
						r.tx.PutAnswers(doc)
					}
					return lateRej("action for " + p.SlideID), nil
				}
			}
		}
		return reject(ErrWrongState, "action for %q is not the current step", p.SlideID), nil
	}
	if req.Type == events.ActionTimeout {
		if rej := r.tooEarly(r.s.Deadlines.ActionClosesAt.Add(-r.tol), req.Type); rej != nil {
			return rej, nil
		}
		if err := r.writeAction(st, nil, r.now); err != nil {
			return nil, err
		}
		r.next(r.now)
		return nil, nil
	}
	if p.ActionID == nil {
		return reject(ErrBadPayload, "action_chosen needs actionId"), nil
	}
	if !hasChoice(r.ct.Config.Actions.ForGroup(sc.Group(r.s.Group)).Options, *p.ActionID) {
		return reject(ErrOutOfRange, "actionId %q is not an option for this session", *p.ActionID), nil
	}
	if err := r.writeAction(st, p.ActionID, r.now); err != nil {
		return nil, err
	}
	r.next(r.now)
	return nil, nil
}

// writeAction records an action (nil on timeout, stored as missing in both
// groups, 2.1) with who asks and which answer type follows.
func (r *run) writeAction(st PlanStep, actionID *string, at time.Time) error {
	doc, err := r.tx.Answers(AnswersActions)
	if err != nil {
		return err
	}
	if doc == nil {
		doc = &AnswerDoc{Kind: AnswersActions}
	}
	if doc.Actions == nil {
		doc.Actions = map[string]ActionRecord{}
	}
	choice := ""
	if actionID != nil {
		choice = *actionID
	}
	answerType := "reason"
	if r.s.DenialSlide == st.SlideID {
		answerType = "denial"
	}
	rec := ActionRecord{ActionID: actionID, TimedOut: actionID == nil, QuestionBy: QuestionBy(r.s.Group, choice),
		AnswerType: answerType, OpenedAt: r.s.Deadlines.StepStartedAt}
	if actionID != nil {
		rec.ChosenAt = tp(at)
	}
	doc.Actions[st.SlideID] = rec
	r.tx.PutAnswers(doc)
	if r.s.ActionChoices == nil {
		r.s.ActionChoices = map[string]string{}
	}
	r.s.ActionChoices[st.SlideID] = choice
	return nil
}

func (r *run) onForm(req *events.Request) (*rejection, error) {
	if r.s.Final() {
		return reject(ErrWrongState, "session is %s", r.s.Status), nil
	}
	kind := KindBudget
	if req.Type == events.BackgroundSubmitted || req.Type == events.BackgroundTimeout {
		kind = KindBackground
	}
	timeout := req.Type == events.BudgetTimeout || req.Type == events.BackgroundTimeout
	st := r.step()
	if st.Kind != kind {
		if j := r.earlierStep(kind, func(PlanStep) bool { return true }); j >= 0 {
			doc, err := r.tx.Answers(kind)
			if err != nil {
				return nil, err
			}
			if doc != nil && doc.Status == RecordTimeout {
				if doc.LateEventNo == nil {
					doc.LateEventNo = i64(r.eventNo)
					r.tx.PutAnswers(doc)
				}
				return lateRej(kind), nil
			}
		}
		return r.wrongState(req.Type), nil
	}
	scene := r.ct.Config.SceneByID(st.Ref)
	var values map[string]any
	hp := ""
	var rej *rejection
	if kind == KindBudget {
		values, rej = parseBudget(req, scene, timeout)
	} else {
		values, hp, rej = parseBackground(req, scene, timeout)
	}
	if rej != nil {
		return rej, nil
	}
	if timeout {
		if rej := r.tooEarly(r.s.Deadlines.FormClosesAt.Add(-r.tol), req.Type); rej != nil {
			return rej, nil
		}
	}
	if hp != "" {
		r.s.AddFlag(FlagHoneypot)
	}
	status := RecordComplete
	if timeout {
		status = RecordTimeout
	}
	if err := r.writeForm(st, values, status, hp, r.now); err != nil {
		return nil, err
	}
	r.next(r.now)
	return nil, nil
}

// writeForm writes answers/budget or answers/background. values nil means
// nothing arrived (server timeout).
func (r *run) writeForm(st PlanStep, values map[string]any, status, hp string, at time.Time) error {
	scene := r.ct.Config.SceneByID(st.Ref)
	doc := &AnswerDoc{Kind: st.Kind, Values: map[string]any{}, ItemStatus: map[string]string{}, HP: hp,
		OpenedAt: r.s.Deadlines.StepStartedAt}
	var ids []string
	if st.Kind == KindBudget {
		ids = []string{"optionId", "confidence"}
	} else {
		for _, it := range scene.Items {
			if it.Kind != sc.ItemHoneypot {
				ids = append(ids, it.ID)
			}
		}
	}
	missing := 0
	for _, id := range ids {
		v, ok := values[id]
		if ok && v != nil {
			doc.Values[id] = v
			doc.ItemStatus[id] = ItemAnswered
			continue
		}
		doc.Values[id] = nil
		missing++
		if status == RecordTimeout {
			doc.ItemStatus[id] = ItemTimedOut
		} else {
			doc.ItemStatus[id] = ItemMissing
		}
	}
	if status != RecordTimeout && missing > 0 {
		status = RecordPartial
	}
	doc.Status = status
	if values != nil {
		doc.SubmittedAt = tp(at)
	}
	r.tx.PutAnswers(doc)
	return nil
}

func parseBudget(req *events.Request, scene *sc.Scene, timeout bool) (map[string]any, *rejection) {
	var p events.BudgetPayload
	if err := req.Decode(&p); err != nil {
		return nil, reject(ErrBadPayload, "budget payload: %v", err)
	}
	optRaw, confRaw := p.OptionID, p.Confidence
	if timeout {
		optRaw, confRaw = nil, nil
		if p.Partial != nil {
			optRaw, confRaw = p.Partial.OptionID, p.Partial.Confidence
		}
	}
	values := map[string]any{}
	if len(optRaw) > 0 && string(optRaw) != "null" {
		var id string
		if err := json.Unmarshal(optRaw, &id); err != nil {
			return nil, reject(ErrBadPayload, "optionId must be a string or null")
		}
		if !hasChoice(scene.Options, id) {
			return nil, reject(ErrOutOfRange, "optionId %q is not an option", id)
		}
		values["optionId"] = id
	}
	if scene.Confidence != nil {
		v, rej := parseScaleValue(confRaw, scene.Confidence.Scale, "confidence")
		if rej != nil {
			return nil, rej
		}
		if v != nil {
			values["confidence"] = *v
		}
	}
	return values, nil
}

// parseBackground validates background items.
// DECISION: the value shapes are not given by the document. Each item takes
// null (unanswered), the string "preferNotToSay", or: a "single" item an
// option id, or { optionId, text } when optionId is its otherOptionId; a
// "scale" item an integer within the scale; a "country" item a string of at
// most 100 characters. The text of an "Other" option may also come as a
// separate key "<itemId>.otherText". Unknown item ids are rejected.
func parseBackground(req *events.Request, scene *sc.Scene, timeout bool) (map[string]any, string, *rejection) {
	var p events.BackgroundPayload
	if err := req.Decode(&p); err != nil {
		return nil, "", reject(ErrBadPayload, "background payload: %v", err)
	}
	items := p.Items
	if timeout {
		items = p.Partial
		if inner, ok := p.Partial["items"]; ok {
			items = nil
			if err := json.Unmarshal(inner, &items); err != nil {
				return nil, "", reject(ErrBadPayload, "partial.items must be an object")
			}
		}
	}
	hp := ""
	switch v := p.HP.(type) {
	case string:
		hp = v
	case nil:
	default:
		b, _ := json.Marshal(v)
		hp = string(b)
	}
	byID := map[string]*sc.BackgroundItem{}
	for i := range scene.Items {
		byID[scene.Items[i].ID] = &scene.Items[i]
	}
	values := map[string]any{}
	otherText := map[string]string{}
	for id, raw := range items {
		it, ok := byID[id]
		if base, isOther := strings.CutSuffix(id, ".otherText"); !ok && isOther {
			// The participant app sends an "Other" option's text as "<itemId>.otherText".
			bi, found := byID[base]
			var txt string
			if !found || bi.Kind != sc.ItemSingle || bi.OtherOptionID == "" || (len(raw) > 0 && string(raw) != "null" && json.Unmarshal(raw, &txt) != nil) {
				return nil, "", reject(ErrBadPayload, "unexpected item %q", id)
			}
			if len(txt) > 500 {
				return nil, "", reject(ErrOutOfRange, "%s: text longer than 500 characters", id)
			}
			otherText[base] = txt
			continue
		}
		if !ok {
			return nil, "", reject(ErrBadPayload, "unknown background item %q", id)
		}
		if it.Kind == sc.ItemHoneypot {
			var s string
			if json.Unmarshal(raw, &s) == nil && s != "" && hp == "" {
				hp = s
			}
			continue
		}
		if len(raw) == 0 || string(raw) == "null" {
			continue
		}
		var str string
		isStr := json.Unmarshal(raw, &str) == nil
		if isStr && str == PreferNotToSay {
			values[id] = PreferNotToSay
			continue
		}
		switch it.Kind {
		case sc.ItemSingle:
			if isStr {
				if !hasChoice(it.Options, str) {
					return nil, "", reject(ErrOutOfRange, "%s: %q is not an option", id, str)
				}
				values[id] = str
				continue
			}
			var o struct {
				OptionID string `json:"optionId"`
				Text     string `json:"text"`
			}
			if err := json.Unmarshal(raw, &o); err != nil || o.OptionID == "" {
				return nil, "", reject(ErrBadPayload, "%s: expected an option id", id)
			}
			if !hasChoice(it.Options, o.OptionID) {
				return nil, "", reject(ErrOutOfRange, "%s: %q is not an option", id, o.OptionID)
			}
			if o.OptionID != it.OtherOptionID {
				values[id] = o.OptionID
				continue
			}
			if len(o.Text) > 500 {
				return nil, "", reject(ErrOutOfRange, "%s: text longer than 500 characters", id)
			}
			values[id] = map[string]any{"optionId": o.OptionID, "text": o.Text}
		case sc.ItemScale:
			v, rej := parseScaleValue(raw, *it.Scale, id)
			if rej != nil {
				return nil, "", rej
			}
			values[id] = *v
		case sc.ItemCountry:
			if !isStr || len(str) > 100 {
				return nil, "", reject(ErrBadPayload, "%s: expected a string of at most 100 characters", id)
			}
			values[id] = str
		}
	}
	for base, txt := range otherText {
		if v, ok := values[base].(string); ok && v == byID[base].OtherOptionID {
			values[base] = map[string]any{"optionId": v, "text": txt}
		}
	}
	return values, hp, nil
}
