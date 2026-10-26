package domain

import (
	"fmt"
	"strconv"

	sc "study/api/pkg/studyconfig"
)

// Content is one loaded config version with everything the state machine
// needs: the resolved clips, the compiled plan and the checker overlays.
type Content struct {
	Config *sc.StudyConfig
	Clips  map[string]*sc.Clip
	Plan   *Plan
	// AudioBase is the URL prefix of audio files named by audioKey, e.g.
	// "/assets/audio/1.3.0/". Estimated clips use SilenceBase instead.
	AudioBase   string
	SilenceBase string
	// VideoBase is the URL prefix of the rendered talking-head videos named by
	// videoKey, e.g. "/assets/video/1.4.0/"; Videos maps a clip id to its
	// videoKey and SeatVideos a seat role to the URL of its idle loop. All
	// empty when the version has no video manifest (2026-10-25).
	VideoBase  string
	Videos     map[string]string
	SeatVideos map[string]string
	// Estimated lists clips whose duration is estimated (no manifest entry).
	Estimated []string
	// Warnings collected while loading.
	Warnings []string

	overlays map[string]*Overlay // slide id -> resolved overlay
	flagged  []string
}

// ResolvedMarking is a Marking without appearAtMark, anchorText and vlmRun,
// with appearAtMs resolved and color filled (8.1).
type ResolvedMarking struct {
	ID         string      `json:"id"`
	TargetIDs  []string    `json:"targetIds"`
	BBox       *[4]float64 `json:"bbox,omitempty"`
	AppearAtMs int64       `json:"appearAtMs"`
	Color      string      `json:"color"`
	Shape      string      `json:"shape"`
	Label      string      `json:"label"`
	Detail     string      `json:"detail"`
}

// Overlay is the checker overlay of a slide step.
type Overlay struct {
	Status     string            `json:"status"`
	Text       string            `json:"text,omitempty"`
	AppearAtMs *int64            `json:"appearAtMs,omitempty"`
	Markings   []ResolvedMarking `json:"markings,omitempty"`
	// DECISION: the panel texts (labels.panelTitle, labels.panelIdle) and the
	// noIssues colour are checker-only (4.3: "no panel texts" in the public
	// pack), so they travel with the overlay of a checker session's slide
	// step and are removed from the public scene pack.
	Title   string `json:"title,omitempty"`
	Idle    string `json:"idleText,omitempty"`
	OkColor string `json:"noIssuesColor,omitempty"`
}

// NewContent compiles the plan, resolves clips against the manifest (nil when
// missing) and resolves the overlays.
func NewContent(c *sc.StudyConfig, m *sc.Manifest, audioBase, silenceBase string) (*Content, error) {
	plan, err := Compile(c, sc.GroupChecker)
	if err != nil {
		return nil, fmt.Errorf("config %s: %w", c.Version, err)
	}
	if _, err := Compile(c, sc.GroupNeutral); err != nil {
		return nil, fmt.Errorf("config %s: %w", c.Version, err)
	}
	clips, estimated := sc.ResolveClips(c, m)
	ct := &Content{Config: c, Clips: clips, Plan: plan, AudioBase: audioBase, SilenceBase: silenceBase,
		Estimated: estimated, overlays: map[string]*Overlay{}, flagged: c.FlaggedSlideIDs()}
	for _, s := range c.Slides() {
		clip := clips[s.Narration.ID]
		resolve := func(mark string, ms *int64, what string) int64 {
			v, err := sc.ResolveAppearAt(mark, ms, clip)
			if err != nil {
				// DECISION: an unresolvable time (a mark with no manifest yet)
				// falls back to 0 with a warning, so dev runs without TTS work.
				ct.Warnings = append(ct.Warnings, fmt.Sprintf("slide %s %s: %v; using 0 ms", s.ID, what, err))
				return 0
			}
			return v
		}
		switch s.Checker.Status {
		case sc.PanelNoIssues:
			text := s.Checker.Text
			if text == "" {
				text = c.Labels.NoIssues
			}
			at := resolve(s.Checker.AppearAtMark, s.Checker.AppearAtMs, "noIssues panel")
			ct.overlays[s.ID] = &Overlay{Status: sc.PanelNoIssues, Text: text, AppearAtMs: &at,
				Title: c.Labels.PanelTitle, Idle: c.Labels.PanelIdle, OkColor: c.Colors.NoIssues}
		case sc.PanelFlagged:
			ov := &Overlay{Status: sc.PanelFlagged, Title: c.Labels.PanelTitle, Idle: c.Labels.PanelIdle}
			for _, mk := range s.Checker.Markings {
				color := mk.Color
				if color == "" {
					color = c.Colors.Flag
				}
				ov.Markings = append(ov.Markings, ResolvedMarking{
					ID: mk.ID, TargetIDs: mk.TargetIDs, BBox: mk.BBox,
					AppearAtMs: resolve(mk.AppearAtMark, mk.AppearAtMs, "marking "+mk.ID),
					Color:      color, Shape: mk.Shape, Label: mk.Label, Detail: mk.Detail,
				})
			}
			ct.overlays[s.ID] = ov
		}
	}
	return ct, nil
}

// FlaggedSlides returns the flagged slide ids in slide order.
func (ct *Content) FlaggedSlides() []string { return ct.flagged }

// Overlay returns the resolved checker overlay of a slide.
func (ct *Content) Overlay(slideID string) *Overlay { return ct.overlays[slideID] }

// AudioURL returns the URL of a clip.
func (ct *Content) AudioURL(cl *sc.Clip) string {
	if cl == nil {
		return ""
	}
	if cl.AudioKey != "" {
		return ct.AudioBase + cl.AudioKey
	}
	return ct.SilenceBase + strconv.FormatInt(cl.DurationMs, 10) + ".wav"
}

// SetVideo attaches a video manifest (nil clears it): clip videos are looked
// up by clip id, seat loops by seat role.
func (ct *Content) SetVideo(base string, vm *sc.VideoManifest) {
	ct.VideoBase, ct.Videos, ct.SeatVideos = "", nil, nil
	if vm == nil {
		return
	}
	ct.VideoBase = base
	ct.Videos = map[string]string{}
	for id, v := range vm.Clips {
		if v.VideoKey != "" {
			ct.Videos[id] = v.VideoKey
		}
	}
	ct.SeatVideos = map[string]string{}
	for role, v := range vm.Seats {
		if v.VideoKey != "" {
			ct.SeatVideos[role] = base + v.VideoKey
		}
	}
}

// VideoURL returns the URL of the talking-head video of a clip, or "" when the
// clip has none. Estimated clips (dev silence) never have one.
func (ct *Content) VideoURL(cl *sc.Clip) string {
	if cl == nil || cl.AudioKey == "" || ct.Videos == nil {
		return ""
	}
	if k, ok := ct.Videos[cl.ID]; ok {
		return ct.VideoBase + k
	}
	return ""
}

// ClipFor returns the clip an audio step plays in a session.
func (ct *Content) ClipFor(s *Session, st PlanStep) *sc.Clip {
	c := ct.Config
	switch st.Kind {
	case KindExample:
		return ct.Clips[sc.ExampleClipID]
	case KindOpening:
		return ct.Clips[sc.OpeningClipID(st.Ref)]
	case KindSlide:
		slide := c.SceneByID(st.SlideID)
		return ct.Clips[slide.Narration.ID]
	case KindQuestion:
		slide := c.SceneByID(st.SlideID)
		switch s.ActionChoices[st.SlideID] {
		case "explain":
			if s.Group == string(sc.GroupChecker) {
				return ct.Clips[slide.ParticipantQuestions.Explain.ID]
			}
		case "challenge":
			if s.Group == string(sc.GroupChecker) {
				return ct.Clips[slide.ParticipantQuestions.Challenge.ID]
			}
		}
		return ct.Clips[slide.RoutineQuestion.ID]
	case KindAnswer:
		slide := c.SceneByID(st.SlideID)
		reason, denial := sc.AnswerClipIDs(slide)
		if s.DenialSlide == st.SlideID {
			return ct.Clips[denial]
		}
		return ct.Clips[reason]
	case KindAudioCheck:
		return ct.Clips[c.Setup.AudioCheck.Clip.ID]
	}
	return nil
}

// QuestionBy returns who asks the question on a flagged slide (2.1, 1.4):
// the board member voices a checker participant's explain or challenge;
// otherwise (say nothing, a timeout, or the neutral group) the chair asks.
func QuestionBy(group, actionID string) string {
	if group == string(sc.GroupChecker) && (actionID == "explain" || actionID == "challenge") {
		return "member"
	}
	return "chair"
}
