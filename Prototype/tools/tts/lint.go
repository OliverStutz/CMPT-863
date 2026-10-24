package main

import (
	"fmt"
	"regexp"
	"strings"

	sc "study/api/pkg/studyconfig"
)

// Lint rules the TTS tool enforces itself (section 4.2), before any API call
// and again on the files it writes.

var forbiddenAssetWords = regexp.MustCompile(`(?i)reason|denial|checker|neutral`)

// checkAssetName fails when a file name would reveal a condition.
func checkAssetName(name string) error {
	if w := forbiddenAssetWords.FindString(name); w != "" {
		return fmt.Errorf("asset file name %q contains %q", name, w)
	}
	return nil
}

var (
	prosodyTag = regexp.MustCompile(`(?i)<\s*/?\s*(prosody|emphasis)\b`)
	ratePitch  = regexp.MustCompile(`(?i)\b(rate|pitch|volume)\s*=`)
	anyTag     = regexp.MustCompile(`<[^>]*>`)
)

// checkAnswerSSML fails when a CISO answer line carries prosody, emphasis or
// a rate or pitch change.
func checkAnswerSSML(l *sc.AudioLine) error {
	if l == nil || l.SSML == "" {
		return nil
	}
	if m := prosodyTag.FindString(l.SSML); m != "" {
		return fmt.Errorf("answer line %s: SSML %q is not allowed in answer lines", l.ID, m)
	}
	if m := ratePitch.FindString(l.SSML); m != "" {
		return fmt.Errorf("answer line %s: rate, pitch or volume change %q is not allowed in answer lines", l.ID, m)
	}
	return nil
}

// ssmlWrapper returns the markup outside the spoken text: the tags before the
// first spoken character and after the last one. A line without SSML has the
// implicit wrapper "<speak>" ... "</speak>".
func ssmlWrapper(l *sc.AudioLine) string {
	s := strings.TrimSpace(l.SSML)
	if s == "" {
		return "<speak>|</speak>"
	}
	locs := anyTag.FindAllStringIndex(s, -1)
	// prefix: consecutive tags (and whitespace) from the start
	pos, pre := 0, ""
	for _, lc := range locs {
		if strings.TrimSpace(s[pos:lc[0]]) != "" {
			break
		}
		pre += s[lc[0]:lc[1]]
		pos = lc[1]
	}
	end, suf := len(s), ""
	for i := len(locs) - 1; i >= 0; i-- {
		lc := locs[i]
		if lc[1] <= pos || strings.TrimSpace(s[lc[1]:end]) != "" {
			break
		}
		suf = s[lc[0]:lc[1]] + suf
		end = lc[0]
	}
	return normaliseTags(pre) + "|" + normaliseTags(suf)
}

func normaliseTags(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

// lintConfig applies the tool's rules to a config.
func lintConfig(c *sc.StudyConfig) []error {
	var errs []error
	assets := []string{c.Consent.PdfAssetKey}
	blocks := func(bs []sc.Block) {
		for _, b := range bs {
			if b.AssetKey != "" {
				assets = append(assets, b.AssetKey)
			}
		}
	}
	blocks(c.Setup.ExampleScene.Layout)
	for i := range c.Scenes {
		s := &c.Scenes[i]
		blocks(s.Layout)
		if s.ProfileCard != nil && s.ProfileCard.PortraitKey != "" {
			assets = append(assets, s.ProfileCard.PortraitKey)
		}
		a := s.Answers
		if a == nil {
			continue
		}
		lines := []*sc.AudioLine{a.Reason, a.Denial, a.Opening, a.ReasonBody, a.Closing}
		for _, l := range lines {
			if l == nil {
				continue
			}
			if l.Voice != sc.VoiceCISO {
				errs = append(errs, fmt.Errorf("answer line %s: voice %q, answers are spoken by the ciso", l.ID, l.Voice))
			}
			if err := checkAnswerSSML(l); err != nil {
				errs = append(errs, err)
			}
		}
		switch {
		case a.Mode == sc.AnswersJoined && a.Opening != nil && a.ReasonBody != nil && a.Closing != nil:
			w := ssmlWrapper(a.Opening)
			if ssmlWrapper(a.ReasonBody) != w || ssmlWrapper(a.Closing) != w {
				errs = append(errs, fmt.Errorf("slide %s: joined answer parts use different SSML wrappers", s.ID))
			}
		case a.Reason != nil && a.Denial != nil:
			if ssmlWrapper(a.Reason) != ssmlWrapper(a.Denial) {
				errs = append(errs, fmt.Errorf("slide %s: reason %s and denial %s use different SSML wrappers (%q vs %q)",
					s.ID, a.Reason.ID, a.Denial.ID, ssmlWrapper(a.Reason), ssmlWrapper(a.Denial)))
			}
		}
	}
	for _, n := range assets {
		if n == "" {
			continue
		}
		if err := checkAssetName(n); err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}
