package domain_test

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"time"

	"study/api/internal/domain"
	"study/api/internal/events"
	"study/api/internal/store"
	"study/api/internal/store/memory"
	"study/api/internal/testutil"
	sc "study/api/pkg/studyconfig"
)

func content(t *testing.T, c *sc.StudyConfig) *domain.Content {
	t.Helper()
	ct, err := domain.NewContent(c, nil, "/assets/audio/x/", "/assets/_dev/silence/")
	if err != nil {
		t.Fatal(err)
	}
	return ct
}

func realConfig(t *testing.T, name string) *sc.StudyConfig {
	t.Helper()
	b, err := os.ReadFile(filepath.Join("../../../../content", name))
	if err != nil {
		t.Fatal(err)
	}
	c, err := sc.Parse(b, nil)
	if err != nil {
		t.Fatal(err)
	}
	return c
}

func points(p *domain.Plan) []string {
	var out []string
	for _, s := range p.RatingPoints() {
		out = append(out, s.Point)
	}
	return out
}

// ---------- plan compiler ----------

func TestCompileDefaultFiveSlides(t *testing.T) {
	c := realConfig(t, "study-config.v1.4.json")
	pc, err := domain.Compile(c, sc.GroupChecker)
	if err != nil {
		t.Fatal(err)
	}
	pn, err := domain.Compile(c, sc.GroupNeutral)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(pc.Keys(), pn.Keys()) || pc.Total != pn.Total {
		t.Fatal("checker and neutral plans differ")
	}
	want := []string{"T0", "T1", "T2a", "T2b", "T3a", "T3b", "T4a", "T4b", "T5"}
	if got := points(pc); !reflect.DeepEqual(got, want) {
		t.Fatalf("points = %v", got)
	}
	// 5 setup steps + opening + T0 + 2 unflagged x 2 + 3 flagged x 6 + budget, background, debrief
	if pc.Total != 5+2+4+18+3 {
		t.Fatalf("total = %d", pc.Total)
	}
}

func TestCompileFourSlidesTwoFlagged(t *testing.T) {
	c := testutil.FastConfig("test-4", 4, 2, 3)
	for _, g := range sc.Groups {
		p, err := domain.Compile(c, g)
		if err != nil {
			t.Fatal(err)
		}
		want := []string{"T0", "T1", "T2a", "T2b", "T3a", "T3b", "T4"}
		if got := points(p); !reflect.DeepEqual(got, want) {
			t.Fatalf("%s: points = %v", g, got)
		}
		var kinds []string
		for _, s := range p.Steps {
			if s.SlideID == "s2" {
				kinds = append(kinds, s.Kind)
			}
		}
		if !reflect.DeepEqual(kinds, []string{"slide", "postView", "rating", "action", "question", "answer", "postView", "rating"}) {
			t.Fatalf("flagged slide steps = %v", kinds)
		}
	}
}

func TestCompileRejectsWrongPoint(t *testing.T) {
	c := testutil.FastConfig("test-bad", 3, 2)
	for i := range c.Scenes {
		if c.Scenes[i].Point == "T2b" {
			c.Scenes[i].Point = "T9"
		}
	}
	if _, err := domain.Compile(c, sc.GroupChecker); err == nil {
		t.Fatal("expected an error for a point that differs from the derived id")
	}
}

// ---------- harness ----------

type harness struct {
	t   *testing.T
	st  *memory.Store
	ct  *domain.Content
	now time.Time
	id  string
	seq int64
}

func newHarness(t *testing.T, st *memory.Store, ct *domain.Content, id string) *harness {
	h := &harness{t: t, st: st, ct: ct, now: time.Date(2026, 11, 1, 10, 0, 0, 0, time.UTC), id: id, seq: 1}
	if err := st.CreateSession(context.Background(), testutil.NewSession(ct, id, h.now)); err != nil {
		t.Fatal(err)
	}
	return h
}

func (h *harness) advance(d time.Duration) { h.now = h.now.Add(d) }

type resp struct {
	status int
	ok     *domain.EventOK
	err    *domain.APIError
}

func (h *harness) sendSeq(seq int64, typ string, payload any) resp {
	h.t.Helper()
	b, _ := json.Marshal(payload)
	req := &events.Request{ClientSeq: seq, Type: typ, ClientTs: h.now.UnixMilli(), Payload: b}
	var res *domain.Result
	err := h.st.RunTx(context.Background(), h.id, func(tx domain.Tx) error {
		var err error
		res, err = domain.HandleEvent(tx, h.ct, h.now, req)
		return err
	})
	if err != nil {
		h.t.Fatal(err)
	}
	out := resp{status: res.Status}
	switch b := res.Body.(type) {
	case *domain.EventOK:
		out.ok = b
		h.seq = b.ExpectedClientSeq
	case *domain.ErrorBody:
		out.err = b.Error
		h.seq = *b.Error.ExpectedClientSeq
	}
	return out
}

func (h *harness) send(typ string, payload any) resp { return h.sendSeq(h.seq, typ, payload) }

func (h *harness) mustOK(typ string, payload any) *domain.EventOK {
	h.t.Helper()
	r := h.send(typ, payload)
	if r.ok == nil {
		h.t.Fatalf("%s: %d %+v", typ, r.status, r.err)
	}
	return r.ok
}

func (h *harness) next() *domain.NextOK {
	h.t.Helper()
	var res *domain.Result
	err := h.st.RunTx(context.Background(), h.id, func(tx domain.Tx) error {
		var err error
		res, err = domain.Next(tx, h.ct, h.now)
		return err
	})
	if err != nil {
		h.t.Fatal(err)
	}
	n := res.Body.(*domain.NextOK)
	if n.ExpectedClientSeq != h.seq {
		h.t.Fatalf("GET next expectedClientSeq %d, client has %d", n.ExpectedClientSeq, h.seq)
	}
	return n
}

func (h *harness) session() *domain.Session {
	s, err := h.st.GetSession(context.Background(), h.id)
	if err != nil {
		h.t.Fatal(err)
	}
	return s
}

// playAudio starts the current audio step and ends it, waiting out too_early.
func (h *harness) playAudio(step *domain.Step) {
	h.t.Helper()
	h.mustOK(events.SceneStarted, map[string]any{"sceneId": step.SceneRef})
	for i := 0; i < 5; i++ {
		r := h.send(events.AudioEnded, map[string]any{"sceneId": step.SceneRef})
		if r.ok != nil {
			return
		}
		if r.err.Code != domain.ErrTooEarly {
			h.t.Fatalf("audio_ended: %+v", r.err)
		}
		h.advance(time.Duration(*r.err.RetryInMs) * time.Millisecond)
	}
	h.t.Fatal("audio never ended")
}

// stepTo drives the session through one step with default answers.
func (h *harness) step(action string) (*domain.Step, bool) {
	h.t.Helper()
	st := h.next().Next
	switch st.Scene {
	case domain.KindDeviceGate:
		h.mustOK(events.DeviceChecked, map[string]any{"widthPx": 1440, "heightPx": 900, "touch": false})
	case domain.KindConsent:
		h.advance(3 * time.Second)
		h.mustOK(events.ConsentAccepted, map[string]any{"consentVersion": h.ct.Config.Consent.Version, "checkbox": true})
	case domain.KindAudioCheck:
		h.mustOK(events.AudioCheckPlayed, map[string]any{})
		h.mustOK(events.AudioCheckAnswered, map[string]any{"choiceId": h.ct.Config.Setup.AudioCheck.CorrectOptionID})
	case domain.KindExample, domain.KindOpening, domain.KindSlide, domain.KindQuestion, domain.KindAnswer:
		h.playAudio(st)
	case domain.KindPreStart:
		h.advance(2 * time.Second)
		h.mustOK(events.SetupDone, map[string]any{})
	case domain.KindPostView:
		h.advance(time.Duration(*st.OpensInMs) * time.Millisecond)
	case domain.KindRating:
		h.advance(time.Second)
		h.mustOK(events.RatingSubmitted, map[string]any{"point": st.Point, "trust": 7, "competence": 5, "benevolence": 6, "integrity": nil})
	case domain.KindAction:
		h.advance(time.Second)
		slide := h.ct.Plan.Steps[h.ct.Plan.Index(h.session().CurrentState)].SlideID
		h.mustOK(events.ActionChosen, map[string]any{"slideId": slide, "actionId": action})
	case domain.KindBudget:
		h.mustOK(events.BudgetSubmitted, map[string]any{"optionId": "approve", "confidence": 8})
	case domain.KindBackground:
		h.mustOK(events.BackgroundSubmitted, map[string]any{"items": map[string]any{"field": "other", "field.otherText": "Law",
			"cyber": 3, "country": "CA"}, "hp": ""})
	case domain.KindDebrief:
		h.advance(5 * time.Second)
		h.mustOK(events.Finished, map[string]any{})
	case domain.KindDone:
		return st, true
	default:
		h.t.Fatalf("unexpected scene %s", st.Scene)
	}
	return st, false
}

func (h *harness) walkTo(scene string, action string) *domain.Step {
	h.t.Helper()
	for i := 0; i < 200; i++ {
		if st := h.next().Next; st.Scene == scene {
			return st
		}
		if _, done := h.step(action); done {
			h.t.Fatalf("finished before reaching %s", scene)
		}
	}
	h.t.Fatal("too many steps")
	return nil
}

// ---------- state machine ----------

func TestHappyPathBothGroups(t *testing.T) {
	ct := content(t, testutil.FastConfig("test-fast", 4, 2, 3))
	st := memory.New()
	for i, wantGroup := range []string{"checker", "neutral"} {
		h := newHarness(t, st, ct, fmt.Sprintf("sess%d", i))
		action := "challenge"
		if wantGroup == "neutral" {
			action = "hasQuestion"
		}
		var scenes []string
		for n := 0; n < 200; n++ {
			step, done := h.step(action)
			scenes = append(scenes, step.Scene)
			if step.Scene == domain.KindSlide && wantGroup == "checker" && step.Overlay == nil {
				t.Fatal("checker slide without overlay")
			}
			if step.Scene == domain.KindSlide && wantGroup == "neutral" && step.Overlay != nil {
				t.Fatal("neutral slide with overlay")
			}
			if done {
				break
			}
		}
		s := h.session()
		if s.Group != wantGroup || s.Status != domain.StatusFinished || s.FinishedAt == nil {
			t.Fatalf("session %+v", s)
		}
		if s.ProgressIndex != s.ProgressTotal || s.ProgressTotal != ct.Plan.Total {
			t.Fatalf("progress %d/%d, plan %d", s.ProgressIndex, s.ProgressTotal, ct.Plan.Total)
		}
		rts, _ := st.Ratings(context.Background(), s.ID)
		if len(rts) != 7 {
			t.Fatalf("ratings = %d", len(rts))
		}
		for _, r := range rts {
			if r.Status != domain.RecordPartial || r.ItemStatus["integrity"] != domain.ItemMissing {
				t.Fatalf("rating %s: %+v", r.Point, r)
			}
		}
		ans, _ := st.Answers(context.Background(), s.ID)
		var actions *domain.AnswerDoc
		for _, a := range ans {
			if a.Kind == domain.AnswersActions {
				actions = a
			}
			if a.Kind == domain.AnswersBackground && a.Status != domain.RecordComplete {
				t.Fatalf("background %+v", a)
			}
		}
		denials := 0
		for slide, rec := range actions.Actions {
			wantBy := "member"
			if wantGroup == "neutral" {
				wantBy = "chair"
			}
			if rec.QuestionBy != wantBy {
				t.Fatalf("%s questionBy %s", slide, rec.QuestionBy)
			}
			if rec.AnswerType == "denial" {
				denials++
				if slide != s.DenialSlide {
					t.Fatal("denial on the wrong slide")
				}
			}
		}
		if denials != 1 {
			t.Fatalf("denials = %d", denials)
		}
		if s.AudioCheck == nil || s.AudioCheck.Correct == nil || !*s.AudioCheck.Correct {
			t.Fatal("audio check not recorded")
		}
	}
}

func TestTooEarlyThenResend(t *testing.T) {
	ct := content(t, testutil.FastConfig("test-fast", 3, 2))
	h := newHarness(t, memory.New(), ct, "s1")
	st := h.walkTo(domain.KindExample, "nothing")
	h.mustOK(events.SceneStarted, map[string]any{"sceneId": st.SceneRef})
	seqBefore := h.seq
	r := h.send(events.AudioEnded, map[string]any{"sceneId": st.SceneRef})
	if r.status != 422 || r.err.Code != domain.ErrTooEarly || r.err.RetryInMs == nil || !r.err.Retryable {
		t.Fatalf("got %d %+v", r.status, r.err)
	}
	if *r.err.ExpectedClientSeq != seqBefore+1 {
		t.Fatal("a rejected event must consume its clientSeq")
	}
	// 1 ms short of the earliest allowed time is still too early.
	h.advance(time.Duration(*r.err.RetryInMs-1) * time.Millisecond)
	if r := h.send(events.AudioEnded, map[string]any{"sceneId": st.SceneRef}); r.err == nil || *r.err.RetryInMs != 1 {
		t.Fatalf("expected too_early by 1 ms, got %+v", r.err)
	}
	h.advance(time.Millisecond)
	h.mustOK(events.AudioEnded, map[string]any{"sceneId": st.SceneRef})
	evs, _ := h.st.Events(context.Background(), "s1")
	invalid := 0
	for _, e := range evs {
		if !e.Valid {
			invalid++
			if e.Reason != domain.ErrTooEarly {
				t.Fatalf("reason %s", e.Reason)
			}
		}
	}
	if invalid != 2 {
		t.Fatalf("invalid events = %d", invalid)
	}
}

func TestSeqGapAndRetries(t *testing.T) {
	ct := content(t, testutil.FastConfig("test-fast", 3, 2))
	h := newHarness(t, memory.New(), ct, "s1")
	// seq_gap does not consume a number.
	r := h.sendSeq(5, events.DeviceChecked, map[string]any{"widthPx": 1440, "heightPx": 900})
	if r.status != 409 || r.err.Code != domain.ErrSeqGap || *r.err.ExpectedClientSeq != 1 {
		t.Fatalf("%d %+v", r.status, r.err)
	}
	payload := map[string]any{"widthPx": 1440, "heightPx": 900}
	b, _ := json.Marshal(payload)
	send := func(seq int64, ts int64, p []byte) *domain.Result {
		var res *domain.Result
		_ = h.st.RunTx(context.Background(), "s1", func(tx domain.Tx) error {
			var err error
			res, err = domain.HandleEvent(tx, ct, h.now, &events.Request{ClientSeq: seq, Type: events.DeviceChecked, ClientTs: ts, Payload: p})
			return err
		})
		return res
	}
	first := send(1, 100, b)
	if first.Status != 200 {
		t.Fatalf("first: %+v", first.Body)
	}
	evsBefore, _ := h.st.Events(context.Background(), "s1")
	// Identical retry: same result, nothing stored.
	retry := send(1, 100, b)
	if retry.Status != 200 || retry.Body.(*domain.EventOK).EventNo != first.Body.(*domain.EventOK).EventNo {
		t.Fatalf("retry: %+v", retry.Body)
	}
	evsAfter, _ := h.st.Events(context.Background(), "s1")
	if len(evsAfter) != len(evsBefore) {
		t.Fatal("an identical retry must not be stored again")
	}
	// Same number, different body.
	mism := send(1, 999, b)
	if mism.Status != 409 || mism.Body.(*domain.ErrorBody).Error.Code != domain.ErrDuplicateSeqMismatch {
		t.Fatalf("mismatch: %+v", mism.Body)
	}
	if s := h.session(); s.LastClientSeq != 1 || s.CurrentState != "consent:consent" {
		t.Fatalf("session %d %s", s.LastClientSeq, s.CurrentState)
	}
}

func TestWrongEventAndAudioCheckTwice(t *testing.T) {
	ct := content(t, testutil.FastConfig("test-fast", 3, 2))
	h := newHarness(t, memory.New(), ct, "s1")
	if r := h.send(events.SetupDone, map[string]any{}); r.status != 409 || r.err.Code != domain.ErrWrongState {
		t.Fatalf("%+v", r.err)
	}
	h.walkTo(domain.KindAudioCheck, "nothing")
	if r := h.send(events.AudioCheckAnswered, map[string]any{"choiceId": "river"}); r.err == nil {
		t.Fatal("answer before any play must be rejected")
	}
	h.mustOK(events.AudioCheckPlayed, map[string]any{})
	h.mustOK(events.AudioCheckPlayed, map[string]any{})
	h.mustOK(events.AudioCheckAnswered, map[string]any{"choiceId": "river"})
	if r := h.send(events.AudioCheckAnswered, map[string]any{"choiceId": "lantern"}); r.err == nil || r.err.Code != domain.ErrWrongState {
		t.Fatal("second answer must be wrong_state")
	}
	s := h.session()
	if s.AudioCheck.Plays != 2 || *s.AudioCheck.Correct || !s.HasFlag(domain.FlagAudioCheckWrong) || s.Status != domain.StatusStarted {
		t.Fatalf("%+v %v", s.AudioCheck, s.Flags)
	}
}

func TestLateRating(t *testing.T) {
	ct := content(t, testutil.FastConfig("test-fast", 3, 2))
	h := newHarness(t, memory.New(), ct, "s1")
	st := h.walkTo(domain.KindRating, "nothing") // T0
	// Within the latency allowance a submission is still on time.
	h.advance(time.Duration(*st.ClosesInMs+299) * time.Millisecond)
	h.mustOK(events.RatingSubmitted, map[string]any{"point": "T0", "trust": 3})
	st = h.walkTo(domain.KindRating, "nothing") // T1
	h.advance(time.Duration(*st.ClosesInMs+301) * time.Millisecond)
	r := h.send(events.RatingSubmitted, map[string]any{"point": "T1", "trust": 4, "competence": 4, "benevolence": 4, "integrity": 4})
	if r.status != 409 || r.err.Code != domain.ErrLate {
		t.Fatalf("%d %+v", r.status, r.err)
	}
	rts, _ := h.st.Ratings(context.Background(), "s1")
	for _, rt := range rts {
		if rt.Point == "T1" && (rt.Status != domain.RecordTimeout || rt.LateEventNo == nil || rt.Values["trust"] != nil) {
			t.Fatalf("T1 %+v", rt)
		}
	}
	evs, _ := h.st.Events(context.Background(), "s1")
	var synth, late bool
	for _, e := range evs {
		synth = synth || (e.Type == events.RatingTimeout && e.Source == events.SourceServer && e.ClientSeq == nil)
		late = late || (e.Reason == domain.ErrLate && !e.Valid && e.Payload["trust"] != nil)
	}
	if !synth || !late {
		t.Fatalf("synthetic timeout %v, late event %v", synth, late)
	}
}

func TestLazyTimeoutChain(t *testing.T) {
	ct := content(t, testutil.FastConfig("test-fast", 3, 2))
	h := newHarness(t, memory.New(), ct, "s1")
	for i := 0; i < 100 && h.session().CurrentState != "postView:s2"; i++ {
		h.step("nothing")
	}
	seq := h.seq
	// The client goes silent: post view, rating T2a and the action all expire.
	h.advance(time.Minute)
	n := h.next()
	if n.Next.Scene != domain.KindQuestion || n.ExpectedClientSeq != seq {
		t.Fatalf("scene %s, seq %d", n.Next.Scene, n.ExpectedClientSeq)
	}
	evs, _ := h.st.Events(context.Background(), "s1")
	var types []string
	var stamps []time.Time
	for _, e := range evs {
		if e.Source == events.SourceServer {
			types = append(types, e.Type)
			stamps = append(stamps, e.ServerTs)
		}
	}
	if !reflect.DeepEqual(types, []string{events.RatingTimeout, events.ActionTimeout}) {
		t.Fatalf("synthetic = %v", types)
	}
	if !stamps[1].After(stamps[0]) {
		t.Fatal("the action window opens when the rating window closes")
	}
	s := h.session()
	if s.ActionChoices["s2"] != "" {
		t.Fatal("timeout must be stored as missing")
	}
	// The question after a timeout is the chair's routine question.
	if n.Next.Transcript != "Why this choice?" {
		t.Fatalf("question %q", n.Next.Transcript)
	}
	// Audio steps have no deadline: nothing more happens.
	h.advance(time.Hour)
	if h.next().Next.Scene != domain.KindQuestion {
		t.Fatal("audio steps must not time out")
	}
}

func TestReloadRestart(t *testing.T) {
	ct := content(t, testutil.FastConfig("test-fast", 3, 2))
	h := newHarness(t, memory.New(), ct, "s1")
	st := h.walkTo(domain.KindOpening, "nothing")
	h.mustOK(events.SceneStarted, map[string]any{"sceneId": st.SceneRef})
	if r := h.send(events.SceneStarted, map[string]any{"sceneId": st.SceneRef}); r.err == nil {
		t.Fatal("a second scene_started without reload must be rejected")
	}
	if err := h.st.RunTx(context.Background(), "s1", func(tx domain.Tx) error { return domain.Reload(tx, ct, h.now) }); err != nil {
		t.Fatal(err)
	}
	n := h.next()
	if !n.Next.Restart || !n.Next.Resumed {
		t.Fatalf("%+v", n.Next)
	}
	h.mustOK(events.SceneStarted, map[string]any{"sceneId": st.SceneRef})
	s := h.session()
	if s.ReloadCount != 1 || !s.HasFlag(domain.FlagReloaded) || s.Status != domain.StatusOngoing {
		t.Fatalf("%+v", s)
	}
}

// ---------- allocation ----------

func TestAllocationBalance100(t *testing.T) {
	ct := content(t, testutil.FastConfig("test-fast", 5, 2, 3, 4))
	st := memory.New()
	for i := 0; i < 100; i++ {
		h := newHarness(t, st, ct, fmt.Sprintf("a%03d", i))
		h.walkTo(domain.KindAudioCheck, "nothing")
	}
	sessions, _ := st.ListSessions(context.Background(), store.SessionFilter{})
	groups := map[string]int{}
	denial := map[string]map[string]int{"checker": {}, "neutral": {}}
	for _, s := range sessions {
		groups[s.Group]++
		denial[s.Group][s.DenialSlide]++
	}
	if groups["checker"] != 50 || groups["neutral"] != 50 {
		t.Fatalf("groups %v", groups)
	}
	for g, m := range denial {
		lo, hi := 1000, 0
		for _, id := range []string{"s2", "s3", "s4"} {
			lo, hi = min(lo, m[id]), max(hi, m[id])
		}
		if hi-lo > 1 {
			t.Fatalf("%s denial slides %v", g, m)
		}
	}
}

func TestAllocateRotation(t *testing.T) {
	c := &domain.Counters{}
	var got []string
	for i := 0; i < 8; i++ {
		g, d := domain.Allocate(c, []string{"s2", "s3", "s4"})
		got = append(got, string(g)+":"+d)
	}
	want := []string{"checker:s2", "neutral:s2", "checker:s3", "neutral:s3", "checker:s4", "neutral:s4", "checker:s2", "neutral:s2"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("%v", got)
	}
}
