package studyconfig

import (
	"fmt"
	"strings"
)

// Manifest is content/audio/<version>/manifest.json, written by tools/tts
// (section 11.1) and read by the server.
//
// DECISION: the architecture names the manifest's contents (durations, marks,
// line offsets, the hash file per line id) but not its JSON shape. It is one
// map from clip id to clip. A clip is what one audio step plays (section 5.3:
// every audio step plays exactly one clip). Clip ids are given by the
// functions below (ExampleClipID, OpeningClipID, AnswerClipIDs); a clip made
// of one line uses the line id.
type Manifest struct {
	ConfigVersion string                  `json:"configVersion"`
	ToolVersion   string                  `json:"toolVersion,omitempty"`
	GeneratedAt   string                  `json:"generatedAt,omitempty"`
	Clips         map[string]ManifestClip `json:"clips"`
}

type ManifestClip struct {
	AudioKey   string        `json:"audioKey"`             // "<sha256 hex>.mp3"
	DurationMs Ms            `json:"durationMs"`           // measured with ffprobe
	Marks      map[string]Ms `json:"marks,omitempty"`      // SSML mark name -> offset in this clip
	Lines      []string      `json:"lines,omitempty"`      // line ids joined into this clip, in order
	LineStarts map[string]Ms `json:"lineStarts,omitempty"` // line id -> start offset in this clip
}

// VideoManifest is content/video/<version>/manifest.json, written by the faces
// pipeline (talking heads rendered offline from the audio clips, 2026-10-25).
// Clips are keyed like the audio manifest (clip id -> rendered MP4 that carries
// the same audio track); Seats maps a boardroom seat role (member2 ... member6,
// chair) to a silent idle loop. Both parts are optional. A version without a
// video manifest plays audio only, exactly as before.
type VideoManifest struct {
	ConfigVersion string               `json:"configVersion"`
	GeneratedAt   string               `json:"generatedAt,omitempty"`
	Clips         map[string]VideoClip `json:"clips,omitempty"`
	Seats         map[string]VideoClip `json:"seats,omitempty"`
}

type VideoClip struct {
	VideoKey   string `json:"videoKey,omitempty"` // "<sha256 hex>.mp4"; empty for a clip made of segments
	DurationMs Ms     `json:"durationMs"`
	Width      int    `json:"width,omitempty"`
	Height     int    `json:"height,omitempty"`
	// Segments split a clip with several speakers (the opening: the chair, then
	// the CISO) into one video per speaker, each starting at the recorded line
	// offset of the audio clip. Ordered by StartMs.
	Segments []VideoSegment `json:"segments,omitempty"`
}

type VideoSegment struct {
	Role       string `json:"role"` // "chair", "ciso", "member"
	StartMs    Ms     `json:"startMs"`
	VideoKey   string `json:"videoKey"`
	DurationMs Ms     `json:"durationMs"`
	Width      int    `json:"width,omitempty"`
	Height     int    `json:"height,omitempty"`
}

// ExampleClipID is the clip id of the setup example scene (all its lines joined).
const ExampleClipID = "setup.exampleScene"

// OpeningClipID is the clip id of the opening scene: the scene id.
func OpeningClipID(openingSceneID string) string { return openingSceneID }

// AnswerClipIDs returns the clip ids of a flagged slide's reason and denial.
// Mode "whole": the two line ids. Mode "joined": "<slideId>.answers.reason" and
// "<slideId>.answers.denial", the clips tools/tts builds from the three parts.
func AnswerClipIDs(slide *Scene) (reason, denial string) {
	a := slide.Answers
	if a == nil {
		return "", ""
	}
	if a.Mode == AnswersJoined {
		return slide.ID + ".answers.reason", slide.ID + ".answers.denial"
	}
	if a.Reason != nil {
		reason = a.Reason.ID
	}
	if a.Denial != nil {
		denial = a.Denial.ID
	}
	return reason, denial
}

// Clip is one resolved clip: what the server needs for timing and the step
// descriptor.
type Clip struct {
	ID         string
	AudioKey   string // empty when estimated
	DurationMs Ms
	Marks      map[string]Ms
	LineStarts map[string]Ms
	Transcript string
	Estimated  bool
}

// DefaultMsPerWord is used when the config sets no msPerWordEstimate.
const DefaultMsPerWord Ms = 400

func (c *StudyConfig) msPerWord() Ms {
	if c.Timing.MsPerWordEstimate != nil && *c.Timing.MsPerWordEstimate > 0 {
		return *c.Timing.MsPerWordEstimate
	}
	return DefaultMsPerWord
}

func estimateLine(l *AudioLine, perWord Ms) Ms {
	if l.DurationMs != nil {
		return *l.DurationMs
	}
	n := len(strings.Fields(l.Text))
	if n == 0 {
		n = 1
	}
	return Ms(n) * perWord
}

// ResolveClips builds every clip of the config. Entries found in the manifest
// (which may be nil) win; otherwise a resolved config's own audioKey and
// durationMs are used for single-line clips; otherwise the duration is
// estimated from the word count. The returned list names the estimated clips.
func ResolveClips(c *StudyConfig, m *Manifest) (map[string]*Clip, []string) {
	clips := map[string]*Clip{}
	var estimated []string
	perWord := c.msPerWord()

	build := func(id string, lines []*AudioLine) *Clip {
		texts := make([]string, 0, len(lines))
		for _, l := range lines {
			texts = append(texts, l.Text)
		}
		cl := &Clip{ID: id, Transcript: strings.Join(texts, " ")}
		if m != nil {
			if mc, ok := m.Clips[id]; ok {
				cl.AudioKey, cl.DurationMs, cl.Marks, cl.LineStarts = mc.AudioKey, mc.DurationMs, mc.Marks, mc.LineStarts
				clips[id] = cl
				return cl
			}
			// A one-line scene clip may be stored under its line id.
			if len(lines) == 1 {
				if mc, ok := m.Clips[lines[0].ID]; ok {
					cl.AudioKey, cl.DurationMs, cl.Marks, cl.LineStarts = mc.AudioKey, mc.DurationMs, mc.Marks, mc.LineStarts
					clips[id] = cl
					return cl
				}
			}
		}
		if len(lines) == 1 && lines[0].AudioKey != "" && lines[0].DurationMs != nil {
			cl.AudioKey, cl.DurationMs, cl.Marks = lines[0].AudioKey, *lines[0].DurationMs, lines[0].Marks
			clips[id] = cl
			return cl
		}
		var d Ms
		cl.LineStarts = map[string]Ms{}
		for i, l := range lines {
			if i > 0 {
				d += c.Timing.LinePauseMs
			}
			cl.LineStarts[l.ID] = d
			d += estimateLine(l, perWord)
		}
		cl.DurationMs = d
		cl.Estimated = true
		estimated = append(estimated, id)
		clips[id] = cl
		return cl
	}
	ptrs := func(ls []AudioLine) []*AudioLine {
		out := make([]*AudioLine, len(ls))
		for i := range ls {
			out[i] = &ls[i]
		}
		return out
	}

	build(c.Setup.AudioCheck.Clip.ID, []*AudioLine{&c.Setup.AudioCheck.Clip})
	build(ExampleClipID, ptrs(c.Setup.ExampleScene.Lines))
	for i := range c.Scenes {
		s := &c.Scenes[i]
		switch s.Kind {
		case SceneOpening:
			build(OpeningClipID(s.ID), ptrs(s.Lines))
		case SceneSlide:
			if s.Narration != nil {
				build(s.Narration.ID, []*AudioLine{s.Narration})
			}
			if s.RoutineQuestion != nil {
				build(s.RoutineQuestion.ID, []*AudioLine{s.RoutineQuestion})
			}
			if q := s.ParticipantQuestions; q != nil {
				build(q.Explain.ID, []*AudioLine{&q.Explain})
				build(q.Challenge.ID, []*AudioLine{&q.Challenge})
			}
			if a := s.Answers; a != nil {
				rid, did := AnswerClipIDs(s)
				var r, d *Clip
				if a.Mode == AnswersJoined && a.Opening != nil && a.ReasonBody != nil && a.Closing != nil {
					r = build(rid, []*AudioLine{a.Opening, a.ReasonBody, a.Closing})
					d = build(did, []*AudioLine{a.Opening, a.Closing})
				} else if a.Reason != nil && a.Denial != nil {
					r = build(rid, []*AudioLine{a.Reason})
					d = build(did, []*AudioLine{a.Denial})
				}
				// tools/tts pads the denial to the reason's length (1.22); an
				// estimate does the same so timing matches the real clip.
				if r != nil && d != nil && d.Estimated && a.PadDenialToReason != nil && *a.PadDenialToReason && d.DurationMs < r.DurationMs {
					d.DurationMs = r.DurationMs
				}
			}
		}
	}
	return clips, estimated
}

// ResolveAppearAt resolves a marking or panel time against a clip: the mark
// wins when it exists in the clip, else appearAtMs.
func ResolveAppearAt(mark string, ms *Ms, clip *Clip) (Ms, error) {
	if mark != "" && clip != nil {
		if v, ok := clip.Marks[mark]; ok {
			return v, nil
		}
	}
	if ms != nil {
		return *ms, nil
	}
	if mark != "" {
		return 0, fmt.Errorf("mark %q not found in clip %q and no appearAtMs set", mark, clipID(clip))
	}
	return 0, fmt.Errorf("neither appearAtMark nor appearAtMs set")
}

func clipID(c *Clip) string {
	if c == nil {
		return ""
	}
	return c.ID
}
