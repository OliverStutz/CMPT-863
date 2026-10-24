package main

import (
	"fmt"
	"html"
	"strings"

	sc "study/api/pkg/studyconfig"
)

// part is one synthesised line.
type part struct {
	LineID string
	Role   string // ciso, chair, member
	Voice  sc.VoiceSpec
	Text   string
	SSML   string
}

// input is what is sent to the API: the line's SSML, or its text wrapped in
// <speak> with XML escaping (section 11.1 step 5).
func (p part) input() string {
	if strings.TrimSpace(p.SSML) != "" {
		return p.SSML
	}
	return "<speak>" + html.EscapeString(p.Text) + "</speak>"
}

func (p part) hasMarks() bool { return strings.Contains(p.input(), "<mark") }

func (p part) key() string { return synthKey(p.Voice, p.Text, p.SSML) }

// clipPlan is one clip of the manifest: what one audio step plays.
type clipPlan struct {
	ID    string
	Parts []part
	GapMs int64 // silence between consecutive parts
	// PadTo names a clip whose duration this clip is padded to with silence
	// at the end (padDenialToReason).
	PadTo string
	// Answer marks the reason or denial clip of a slide (for the lint).
	Answer bool
}

// buildPlan collects every clip of the config in the order the server's
// studyconfig.ResolveClips builds them, with the same clip ids.
func buildPlan(c *sc.StudyConfig, voices map[string]sc.VoiceSpec, answerGapMs int64) ([]clipPlan, error) {
	var plans []clipPlan
	var errs []string
	mk := func(l *sc.AudioLine) part {
		v, ok := voices[l.Voice]
		if !ok {
			errs = append(errs, fmt.Sprintf("line %s: unknown voice %q", l.ID, l.Voice))
		}
		if strings.TrimSpace(l.Text) == "" {
			errs = append(errs, fmt.Sprintf("line %s: empty text", l.ID))
		}
		return part{LineID: l.ID, Role: l.Voice, Voice: v, Text: l.Text, SSML: l.SSML}
	}
	one := func(l *sc.AudioLine) {
		plans = append(plans, clipPlan{ID: l.ID, Parts: []part{mk(l)}})
	}
	many := func(id string, ls []sc.AudioLine) {
		if len(ls) == 0 {
			return
		}
		cp := clipPlan{ID: id, GapMs: c.Timing.LinePauseMs}
		for i := range ls {
			cp.Parts = append(cp.Parts, mk(&ls[i]))
		}
		plans = append(plans, cp)
	}

	one(&c.Setup.AudioCheck.Clip)
	many(sc.ExampleClipID, c.Setup.ExampleScene.Lines)
	for i := range c.Scenes {
		s := &c.Scenes[i]
		switch s.Kind {
		case sc.SceneOpening:
			many(sc.OpeningClipID(s.ID), s.Lines)
		case sc.SceneSlide:
			if s.Narration != nil {
				one(s.Narration)
			}
			if s.RoutineQuestion != nil {
				one(s.RoutineQuestion)
			}
			if q := s.ParticipantQuestions; q != nil {
				one(&q.Explain)
				one(&q.Challenge)
			}
			a := s.Answers
			if a == nil {
				continue
			}
			rid, did := sc.AnswerClipIDs(s)
			pad := ""
			if a.PadDenialToReason != nil && *a.PadDenialToReason {
				pad = rid
			}
			switch a.Mode {
			case sc.AnswersJoined:
				if a.Opening == nil || a.ReasonBody == nil || a.Closing == nil {
					errs = append(errs, fmt.Sprintf("slide %s: joined answers need opening, reasonBody and closing", s.ID))
					continue
				}
				o, b, cl := mk(a.Opening), mk(a.ReasonBody), mk(a.Closing)
				plans = append(plans,
					clipPlan{ID: rid, Parts: []part{o, b, cl}, GapMs: answerGapMs, Answer: true},
					clipPlan{ID: did, Parts: []part{o, cl}, GapMs: answerGapMs, PadTo: pad, Answer: true})
			default:
				if a.Reason == nil || a.Denial == nil {
					errs = append(errs, fmt.Sprintf("slide %s: whole answers need reason and denial", s.ID))
					continue
				}
				plans = append(plans,
					clipPlan{ID: rid, Parts: []part{mk(a.Reason)}, Answer: true},
					clipPlan{ID: did, Parts: []part{mk(a.Denial)}, PadTo: pad, Answer: true})
			}
		}
	}
	if len(errs) > 0 {
		return nil, fmt.Errorf("%s", strings.Join(errs, "\n"))
	}
	seen := map[string]bool{}
	for _, p := range plans {
		if seen[p.ID] {
			return nil, fmt.Errorf("clip id %q occurs twice", p.ID)
		}
		seen[p.ID] = true
	}
	return plans, nil
}

// uniqueParts returns every distinct synthesis of the plan, by key, in order.
func uniqueParts(plans []clipPlan) []part {
	seen := map[string]bool{}
	var out []part
	for _, cp := range plans {
		for _, p := range cp.Parts {
			k := p.key()
			if !seen[k] {
				seen[k] = true
				out = append(out, p)
			}
		}
	}
	return out
}

// filterPlan keeps the clips whose id or one of whose line ids is only.
func filterPlan(plans []clipPlan, only string) []clipPlan {
	var out []clipPlan
	keep := map[string]bool{}
	for _, cp := range plans {
		hit := cp.ID == only
		for _, p := range cp.Parts {
			hit = hit || p.LineID == only
		}
		if hit {
			keep[cp.ID] = true
		}
	}
	// A padded clip needs the clip it is padded to.
	for _, cp := range plans {
		if keep[cp.ID] && cp.PadTo != "" {
			keep[cp.PadTo] = true
		}
	}
	for _, cp := range plans {
		if keep[cp.ID] {
			out = append(out, cp)
		}
	}
	return out
}
