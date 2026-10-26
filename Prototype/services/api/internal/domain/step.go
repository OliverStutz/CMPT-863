package domain

import (
	"strconv"
	"time"

	sc "study/api/pkg/studyconfig"
)

// StepVideoSegment is one speaker's video within a multi-speaker clip.
type StepVideoSegment struct {
	Role    string `json:"role"`
	URL     string `json:"url"`
	StartMs int64  `json:"startMs"`
}

// Step is the step descriptor of section 8.1.
type Step struct {
	StepID   string `json:"stepId"`
	Scene    string `json:"scene"`
	SceneRef string `json:"sceneRef,omitempty"`
	AudioURL string `json:"audioUrl,omitempty"`
	// VideoURL is the rendered talking-head video of the same clip (muted in
	// the browser, in sync with the audio element); absent when none exists.
	VideoURL string `json:"videoUrl,omitempty"`
	// VideoSegments replaces VideoURL for a clip with several speakers: one
	// video per speaker, each starting at its offset in the audio clip.
	VideoSegments []StepVideoSegment `json:"videoSegments,omitempty"`
	Replayable    bool               `json:"replayable,omitempty"`
	Restart       bool               `json:"restart,omitempty"`
	Overlay       *Overlay           `json:"overlay,omitempty"`
	Options       []sc.Choice        `json:"options,omitempty"`
	Point         string             `json:"point,omitempty"`
	Transcript    string             `json:"transcript,omitempty"`
	PreStartText  []string           `json:"preStartText,omitempty"`
	Debrief       *Debrief           `json:"debrief,omitempty"`
	OpensInMs     *int64             `json:"opensInMs,omitempty"`
	ClosesInMs    *int64             `json:"closesInMs,omitempty"`
	Prefetch      []string           `json:"prefetch,omitempty"`
	Resumed       bool               `json:"resumed,omitempty"`

	// DECISION: the Step type of 8.1 has no field for the action prompt and
	// note, which differ between groups and so cannot be in the public scene
	// pack (4.3). Action steps carry them here.
	Prompt string `json:"prompt,omitempty"`
	Note   string `json:"note,omitempty"`
	// DECISION: 8.1 has no scene value for a session Blocked by the device
	// gate. Such a session gets scene "done" with blocked: true, and the
	// client shows /s/:code/blocked (9.1).
	Blocked bool `json:"blocked,omitempty"`
}

// Debrief is the debrief step's content (2.14, 2.15).
type Debrief struct {
	Paragraphs   []string `json:"paragraphs"`
	DoNotDiscuss string   `json:"doNotDiscuss"`
	SessionCode  string   `json:"sessionCode"`
	// DECISION: 8.1 has no label for the code; the config's debrief scene has
	// sessionCodeLabel, which the client shows before the code.
	SessionCodeLabel string `json:"sessionCodeLabel,omitempty"`
}

func msUntil(t *time.Time, now time.Time) *int64 {
	if t == nil {
		return nil
	}
	d := t.Sub(now).Milliseconds()
	if d < 0 {
		d = 0
	}
	return &d
}

// CurrentIndex returns the plan position of the session's current step.
func (ct *Content) CurrentIndex(s *Session) int {
	return ct.Plan.Index(s.CurrentState)
}

// keepsOverlay reports whether a step of a slide shows that slide's checker
// overlay (for checker sessions).
func keepsOverlay(kind string) bool {
	switch kind {
	case KindSlide, KindPostView, KindAction, KindQuestion, KindAnswer:
		return true
	}
	return false
}

// Describe builds the step descriptor for the session's current step.
func (ct *Content) Describe(s *Session, now time.Time, resumed bool) *Step {
	idx := ct.CurrentIndex(s)
	if s.Status == StatusBlocked {
		return &Step{StepID: "p" + strconv.Itoa(max(idx, 0)), Scene: KindDone, Blocked: true}
	}
	if idx < 0 {
		return &Step{StepID: "p0", Scene: KindDone}
	}
	st := ct.Plan.Steps[idx]
	c := ct.Config
	step := &Step{StepID: "p" + strconv.Itoa(idx), Scene: st.Kind, SceneRef: st.Ref, Resumed: resumed}
	d := s.Deadlines

	if st.Audio() {
		clip := ct.ClipFor(s, st)
		step.AudioURL = ct.AudioURL(clip)
		step.VideoURL = ct.VideoURL(clip)
		step.VideoSegments = ct.VideoSegmentsFor(clip)
		if clip != nil {
			step.Transcript = clip.Transcript
		}
		step.Restart = d.ClipStartedAt != nil && s.RestartAllowed
	}

	// The checker overlay stays with a slide from its own step until the next
	// slide begins: the post-audio view, the action, the question, the answer
	// and the post-answer view of that slide carry it too, so the marks and the
	// panel remain on screen while the question and answer refer to them.
	// Rating steps are a separate screen and carry none. Checker sessions only
	// (4.3); neutral sessions never get the field.
	if s.Group == string(sc.GroupChecker) && st.SlideID != "" && keepsOverlay(st.Kind) {
		step.Overlay = ct.Overlay(st.SlideID)
	}

	switch st.Kind {
	case KindDeviceGate, KindConsent, KindExample, KindOpening, KindQuestion, KindAnswer, KindDone, KindSlide:
	case KindAudioCheck:
		// No transcript: the clip's text is the answer to the check.
		step.AudioURL = ct.AudioURL(ct.ClipFor(s, st))
		step.Replayable = true
	case KindPreStart:
		step.PreStartText = append([]string{}, c.Setup.PreStartText.Both...)
		if s.Group == string(sc.GroupChecker) {
			step.PreStartText = append(step.PreStartText, c.Setup.PreStartText.Checker...)
		}
	case KindPostView:
		step.OpensInMs = msUntil(d.RatingOpensAt, now)
	case KindRating:
		step.Point = st.Point
		step.OpensInMs = msUntil(d.RatingOpensAt, now)
		step.ClosesInMs = msUntil(d.RatingClosesAt, now)
	case KindAction:
		spec := c.Actions.ForGroup(sc.Group(s.Group))
		step.Options = spec.Options
		step.Prompt, step.Note = spec.Prompt, spec.Note
		step.ClosesInMs = msUntil(d.ActionClosesAt, now)
	case KindBudget, KindBackground:
		step.ClosesInMs = msUntil(d.FormClosesAt, now)
	case KindDebrief:
		db := c.SceneByID(st.Ref)
		var paras []string
		if db.Paragraphs != nil {
			paras = append(paras, db.Paragraphs.Both...)
			if s.Group == string(sc.GroupChecker) {
				paras = append(paras, db.Paragraphs.Checker...)
			} else {
				paras = append(paras, db.Paragraphs.Neutral...)
			}
			paras = append(paras, db.Paragraphs.End...)
		}
		step.Debrief = &Debrief{Paragraphs: paras, DoNotDiscuss: db.DoNotDiscuss, SessionCode: SessionCode(s.ID),
			SessionCodeLabel: db.SessionCodeLabel}
	}

	// Prefetch the next step's clip when it is already decided.
	// The question clip depends on the action, so nothing is prefetched from
	// the action step; the lookahead stops at the next action step.
	if st.Kind != KindAction {
		for j := idx + 1; j < len(ct.Plan.Steps) && j <= idx+3; j++ {
			nx := ct.Plan.Steps[j]
			if nx.Kind == KindAction {
				break
			}
			if nx.Audio() {
				if cl := ct.ClipFor(s, nx); cl != nil {
					if u := ct.AudioURL(cl); u != "" {
						step.Prefetch = []string{u}
					}
					if v := ct.VideoURL(cl); v != "" {
						step.Prefetch = append(step.Prefetch, v)
					}
					for _, sg := range ct.VideoSegmentsFor(cl) {
						step.Prefetch = append(step.Prefetch, sg.URL)
					}
				}
				break
			}
		}
	}
	return step
}

// SessionCode is the code shown on the debrief for withdrawal (2.14).
// DECISION: the session code is the session id, which is already short and
// free of look-alike characters.
func SessionCode(id string) string { return id }
