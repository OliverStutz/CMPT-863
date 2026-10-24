package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html"
	"io"
	"os"
	"path/filepath"
	"strings"

	sc "study/api/pkg/studyconfig"
)

// Voice samples for the listening test (decision 3.6): one sentence of real
// content per candidate voice, written to <out>/<role>-<voiceId>.mp3. Each
// sample carries one SSML mark, so a voice that does not return timepoints
// fails here and not later.

// sampleText returns the sample sentence of a role from the config, split
// into the part before and after the test mark.
func sampleText(c *sc.StudyConfig, role string) (string, string, error) {
	switch role {
	case sc.VoiceCISO:
		op := c.SceneOfKind(sc.SceneOpening)
		slides := c.Slides()
		if op == nil || len(slides) == 0 || slides[0].Narration == nil {
			return "", "", fmt.Errorf("no CISO greeting or first narration")
		}
		greet := ""
		for _, l := range op.Lines {
			if l.Voice == sc.VoiceCISO {
				greet = l.Text
			}
		}
		return greet, firstSentence(slides[0].Narration.Text), nil
	case sc.VoiceChair:
		op := c.SceneOfKind(sc.SceneOpening)
		if op == nil {
			return "", "", fmt.Errorf("no opening")
		}
		for _, l := range op.Lines {
			if l.Voice == sc.VoiceChair {
				return splitSentences(l.Text)
			}
		}
	case sc.VoiceMember:
		for _, s := range c.Slides() {
			if q := s.ParticipantQuestions; q != nil {
				t := q.Explain.Text
				if i := strings.Index(t, ", "); i > 0 {
					return t[:i+1], t[i+2:], nil
				}
				return "", t, nil
			}
		}
	}
	return "", "", fmt.Errorf("no sample text for role %q", role)
}

func firstSentence(t string) string {
	if i := strings.Index(t, ". "); i > 0 {
		return t[:i+1]
	}
	return t
}

func splitSentences(t string) (string, string, error) {
	if i := strings.Index(t, ". "); i > 0 {
		return t[:i+1], t[i+2:], nil
	}
	return "", t, nil
}

type sampleResult struct {
	Role       string `json:"role"`
	Voice      string `json:"voice"`
	File       string `json:"file"`
	Text       string `json:"text"`
	DurationMs int64  `json:"durationMs"`
	MarkMs     int64  `json:"markMs"`
	Error      string `json:"error,omitempty"`
}

// runSamples synthesises every candidate of candidatesPath, a JSON object
// {"ciso": [voice names], "chair": [...], "member": [...]}.
func runSamples(ctx context.Context, configPath, candidatesPath, out, quotaProject string, w io.Writer, synth synthesizer) error {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}
	cfg, err := sc.Parse(data, nil)
	if err != nil {
		return err
	}
	cb, err := os.ReadFile(candidatesPath)
	if err != nil {
		return err
	}
	var cands map[string][]string
	if err := json.Unmarshal(cb, &cands); err != nil {
		return err
	}
	if synth == nil {
		g, err := newGoogleTTS(ctx, quotaProject)
		if err != nil {
			return err
		}
		synth = g
	}
	if err := os.MkdirAll(out, 0o755); err != nil {
		return err
	}
	var results []sampleResult
	for _, role := range []string{sc.VoiceCISO, sc.VoiceChair, sc.VoiceMember} {
		a, b, err := sampleText(cfg, role)
		if err != nil {
			return err
		}
		ssml := "<speak>" + html.EscapeString(a) + ` <mark name="sample"/>` + html.EscapeString(b) + "</speak>"
		for _, name := range cands[role] {
			r := sampleResult{Role: role, Voice: name, File: role + "-" + name + ".mp3", Text: strings.TrimSpace(a + " " + b)}
			audio, marks, err := synth.Synthesize(ctx, sc.VoiceSpec{Provider: "google-tts", Name: name, LanguageCode: "en-US"}, ssml, true)
			if err == nil {
				if _, ok := marks["sample"]; !ok {
					err = fmt.Errorf("no SSML mark timepoint returned")
				}
			}
			if err == nil {
				r.MarkMs = marks["sample"]
				r.DurationMs, err = durationMs(audio)
			}
			if err == nil {
				err = writeAtomic(filepath.Join(out, r.File), audio)
			}
			if err != nil {
				r.Error = err.Error()
			}
			fmt.Fprintf(w, "%-6s %-18s %6d ms  mark at %5d ms  %s\n", role, name, r.DurationMs, r.MarkMs, r.Error)
			results = append(results, r)
		}
	}
	rb, _ := json.MarshalIndent(results, "", "  ")
	return writeAtomic(filepath.Join(out, "samples.json"), append(rb, '\n'))
}
