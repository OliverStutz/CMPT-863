package domain

import (
	"errors"
	"fmt"
	"strconv"

	sc "study/api/pkg/studyconfig"
)

// Step kinds. They are also the values of Step.scene in the descriptor.
const (
	KindDeviceGate = "deviceGate"
	KindConsent    = "consent"
	KindAudioCheck = "audioCheck"
	KindExample    = "example"
	KindPreStart   = "preStart"
	KindOpening    = "opening"
	KindSlide      = "slide"
	KindPostView   = "postView"
	KindRating     = "rating"
	KindAction     = "action"
	KindQuestion   = "question"
	KindAnswer     = "answer"
	KindBudget     = "budget"
	KindBackground = "background"
	KindDebrief    = "debrief"
	KindDone       = "done"
)

// PlanStep is one step of a compiled plan.
type PlanStep struct {
	Kind    string // one of the Kind constants
	Ref     string // scene id in the config (sceneRef), or a fixed name for setup steps
	SlideID string // slide the step belongs to, if any
	Point   string // rating steps: derived point id
	Phase   string // rating steps
	Counted bool   // participant-facing, counted in progressTotal (5.1)
}

// Key is the stored form of the step in Session.plan and Session.currentState,
// e.g. "rating:rating.T2a".
func (p PlanStep) Key() string { return p.Kind + ":" + p.Ref }

// Audio reports whether the step plays a clip.
func (p PlanStep) Audio() bool {
	switch p.Kind {
	case KindExample, KindOpening, KindSlide, KindQuestion, KindAnswer:
		return true
	}
	return false
}

// Plan is the ordered list of steps of a session.
type Plan struct {
	Steps []PlanStep
	index map[string]int
	Total int // participant-facing steps
}

// Fixed refs for steps that are not scenes of config.scenes.
const (
	RefDeviceGate = "deviceGate"
	RefConsent    = "consent"
	RefAudioCheck = "setup.audioCheck"
	RefExample    = sc.ExampleClipID // "setup.exampleScene"
	RefPreStart   = "setup.preStart"
	RefDone       = "done"
)

// PreAllocationSteps is the short plan written at session creation (5.1).
var PreAllocationSteps = []PlanStep{
	{Kind: KindDeviceGate, Ref: RefDeviceGate, Counted: true},
	{Kind: KindConsent, Ref: RefConsent, Counted: true},
}

// Index returns the position of a step key, or -1.
func (p *Plan) Index(key string) int {
	if i, ok := p.index[key]; ok {
		return i
	}
	return -1
}

// ProgressIndex is the 1-based number of the last participant-facing step at
// or before position i.
func (p *Plan) ProgressIndex(i int) int {
	n := 0
	for j := 0; j <= i && j < len(p.Steps); j++ {
		if p.Steps[j].Counted {
			n++
		}
	}
	return n
}

// Keys returns the stored form of the plan.
func (p *Plan) Keys() []string {
	out := make([]string, len(p.Steps))
	for i, s := range p.Steps {
		out[i] = s.Key()
	}
	return out
}

// RatingPoints returns the rating steps in plan order.
func (p *Plan) RatingPoints() []PlanStep {
	var out []PlanStep
	for _, s := range p.Steps {
		if s.Kind == KindRating {
			out = append(out, s)
		}
	}
	return out
}

func newPlan(steps []PlanStep) *Plan {
	p := &Plan{Steps: steps, index: map[string]int{}}
	for i, s := range steps {
		p.index[s.Key()] = i
		if s.Counted {
			p.Total++
		}
	}
	return p
}

// Compile builds the full plan of a session (5.1, 5.2) from the config and the
// group. The group is accepted for fidelity to the document; both groups get
// the same steps (5.1), and what differs between them is text on a step and
// the overlay, which the step descriptor adds.
//
// The compiler looks scenes up by kind, id and forSlide, and fails when a
// scene the canonical flow needs is missing or a rating point differs from
// the derived id. The full wording lint is tools/content-lint's job (4.2).
func Compile(c *sc.StudyConfig, group sc.Group) (*Plan, error) {
	if group != sc.GroupChecker && group != sc.GroupNeutral {
		return nil, fmt.Errorf("unknown group %q", group)
	}
	var errs []error
	fail := func(f string, a ...any) { errs = append(errs, fmt.Errorf(f, a...)) }

	steps := append([]PlanStep{}, PreAllocationSteps...)
	steps = append(steps,
		PlanStep{Kind: KindAudioCheck, Ref: RefAudioCheck, Counted: true},
		PlanStep{Kind: KindExample, Ref: RefExample, Counted: true},
		PlanStep{Kind: KindPreStart, Ref: RefPreStart, Counted: true},
	)

	rating := func(after, point, slideID, phase string) {
		r := c.RatingAfter(after)
		if r == nil {
			fail("no rating scene follows %q (expected point %s)", after, point)
			return
		}
		if r.Point != point {
			fail("rating scene %q has point %q, derived point is %q", r.ID, r.Point, point)
		}
		steps = append(steps, PlanStep{Kind: KindRating, Ref: r.ID, SlideID: slideID, Point: point, Phase: phase, Counted: true})
	}

	opening := c.SceneOfKind(sc.SceneOpening)
	if opening == nil {
		fail("config has no opening scene")
	} else {
		if len(opening.Lines) == 0 {
			fail("opening scene has no lines")
		}
		steps = append(steps,
			PlanStep{Kind: KindOpening, Ref: opening.ID, Counted: true},
			PlanStep{Kind: KindPostView, Ref: opening.ID},
		)
		rating(opening.ID, "T0", "", PhaseBaseline)
	}

	slides := c.Slides()
	if len(slides) == 0 {
		fail("config has no slide scene")
	}
	flagged := 0
	for n, s := range slides {
		num := strconv.Itoa(n + 1)
		if s.Narration == nil {
			fail("slide %q has no narration", s.ID)
		}
		if s.Checker == nil {
			fail("slide %q has no checker panel", s.ID)
		}
		steps = append(steps,
			PlanStep{Kind: KindSlide, Ref: s.ID, SlideID: s.ID, Counted: true},
			PlanStep{Kind: KindPostView, Ref: s.ID, SlideID: s.ID},
		)
		if !s.Flagged() {
			if s.RoutineQuestion != nil || s.ParticipantQuestions != nil || s.Answers != nil {
				fail("unflagged slide %q has questions or answers", s.ID)
			}
			rating(s.ID, "T"+num, s.ID, PhaseAfterSlide)
			continue
		}
		flagged++
		if s.RoutineQuestion == nil || s.ParticipantQuestions == nil || s.Answers == nil {
			fail("flagged slide %q lacks routineQuestion, participantQuestions or answers", s.ID)
		}
		rating(s.ID, "T"+num+"a", s.ID, PhaseAfterSlide)
		act := c.SceneFor(sc.SceneAction, s.ID)
		q := c.SceneFor(sc.SceneQuestion, s.ID)
		ans := c.SceneFor(sc.SceneAnswer, s.ID)
		if act == nil || q == nil || ans == nil {
			fail("flagged slide %q needs one action, question and answer scene", s.ID)
			continue
		}
		steps = append(steps,
			PlanStep{Kind: KindAction, Ref: act.ID, SlideID: s.ID, Counted: true},
			PlanStep{Kind: KindQuestion, Ref: q.ID, SlideID: s.ID, Counted: true},
			PlanStep{Kind: KindAnswer, Ref: ans.ID, SlideID: s.ID, Counted: true},
			PlanStep{Kind: KindPostView, Ref: ans.ID, SlideID: s.ID},
		)
		rating(ans.ID, "T"+num+"b", s.ID, PhaseAfterAnswer)
	}
	if len(slides) > 0 && flagged == 0 {
		fail("config has no flagged slide")
	}
	for _, sc2 := range c.Scenes {
		if (sc2.Kind == sc.SceneAction || sc2.Kind == sc.SceneQuestion || sc2.Kind == sc.SceneAnswer) && !isFlagged(c, sc2.ForSlide) {
			fail("%s scene %q names %q, which is not a flagged slide", sc2.Kind, sc2.ID, sc2.ForSlide)
		}
	}

	for _, kind := range []string{sc.SceneBudget, sc.SceneBackground, sc.SceneDebrief} {
		s := c.SceneOfKind(kind)
		if s == nil {
			fail("config has no %s scene", kind)
			continue
		}
		steps = append(steps, PlanStep{Kind: kind, Ref: s.ID, Counted: true})
	}
	steps = append(steps, PlanStep{Kind: KindDone, Ref: RefDone})

	if len(c.Rating.Items) == 0 {
		fail("rating has no items")
	}
	if len(c.Actions.Checker.Options) == 0 || len(c.Actions.Neutral.Options) == 0 {
		fail("actions need options for both groups")
	}
	if err := errors.Join(errs...); err != nil {
		return nil, err
	}
	return newPlan(steps), nil
}

func isFlagged(c *sc.StudyConfig, slideID string) bool {
	s := c.SceneByID(slideID)
	return s != nil && s.Flagged()
}
