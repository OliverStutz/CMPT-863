package domain

import (
	"strconv"
	"time"

	sc "study/api/pkg/studyconfig"
)

// Step is the step descriptor of section 8.1.
type Step struct {
	StepID       string      `json:"stepId"`
	Scene        string      `json:"scene"`
	SceneRef     string      `json:"sceneRef,omitempty"`
	AudioURL     string      `json:"audioUrl,omitempty"`
	Replayable   bool        `json:"replayable,omitempty"`
	Restart      bool        `json:"restart,omitempty"`
	Overlay      *Overlay    `json:"overlay,omitempty"`
	Options      []sc.Choice `json:"options,omitempty"`
	Point        string      `json:"point,omitempty"`
	Transcript   string      `json:"transcript,omitempty"`
	PreStartText []string    `json:"preStartText,omitempty"`
	Debrief      *Debrief    `json:"debrief,omitempty"`
	OpensInMs    *int64      `json:"opensInMs,omitempty"`
	ClosesInMs   *int64      `json:"closesInMs,omitempty"`
	Prefetch     []string    `json:"prefetch,omitempty"`
	Resumed      bool        `json:"resumed,omitempty"`

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
		if clip != nil {
			step.Transcript = clip.Transcript
		}
		step.Restart = d.ClipStartedAt != nil && s.RestartAllowed
	}

	switch st.Kind {
	case KindDeviceGate, KindConsent, KindExample, KindOpening, KindQuestion, KindAnswer, KindDone:
	case KindAudioCheck:
		// No transcript: the clip's text is the answer to the check.
		step.AudioURL = ct.AudioURL(ct.ClipFor(s, st))
		step.Replayable = true
	case KindPreStart:
		step.PreStartText = append([]string{}, c.Setup.PreStartText.Both...)
		if s.Group == string(sc.GroupChecker) {
			step.PreStartText = append(step.PreStartText, c.Setup.PreStartText.Checker...)
		}
	case KindSlide:
		if s.Group == string(sc.GroupChecker) {
			step.Overlay = ct.Overlay(st.SlideID)
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
				if u := ct.AudioURL(ct.ClipFor(s, nx)); u != "" {
					step.Prefetch = []string{u}
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
