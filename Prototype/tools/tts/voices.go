package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	sc "study/api/pkg/studyconfig"
)

// isPlaceholder reports whether a voice name in the config is still a
// placeholder ("TBD", empty, "TODO...", "PLACEHOLDER...").
func isPlaceholder(name string) bool {
	n := strings.ToUpper(strings.TrimSpace(name))
	return n == "" || n == "TBD" || strings.HasPrefix(n, "TODO") || strings.Contains(n, "PLACEHOLDER")
}

// loadVoices reads a voices file in the shape of the config's "voices"
// object. Keys that start with "_" are notes.
func loadVoices(path string) (sc.Voices, error) {
	var v sc.Voices
	b, err := os.ReadFile(path)
	if err != nil {
		return v, err
	}
	if err := json.Unmarshal(b, &v); err != nil {
		return v, fmt.Errorf("%s: %w", path, err)
	}
	return v, nil
}

// resolveVoices picks the voice of each role. An override file (-voices)
// wins for every role it names; otherwise the config's voice is used unless
// it is a placeholder, in which case the defaults file is used.
func resolveVoices(cfg sc.Voices, override, defaults *sc.Voices) (map[string]sc.VoiceSpec, error) {
	out := map[string]sc.VoiceSpec{}
	pick := func(role string, c sc.VoiceSpec, o, d func(*sc.Voices) sc.VoiceSpec) error {
		v := c
		if override != nil && !isPlaceholder(o(override).Name) {
			v = o(override)
		} else if isPlaceholder(v.Name) {
			if defaults == nil || isPlaceholder(d(defaults).Name) {
				return fmt.Errorf("voice %q is a placeholder in the config and no default is set", role)
			}
			v = d(defaults)
		}
		if v.LanguageCode == "" {
			v.LanguageCode = "en-US"
		}
		if v.Provider == "" {
			v.Provider = "google-tts"
		}
		if v.Provider != "google-tts" {
			return fmt.Errorf("voice %q: provider %q not supported", role, v.Provider)
		}
		out[role] = v
		return nil
	}
	get := map[string]func(*sc.Voices) sc.VoiceSpec{
		sc.VoiceCISO:   func(v *sc.Voices) sc.VoiceSpec { return v.CISO },
		sc.VoiceChair:  func(v *sc.Voices) sc.VoiceSpec { return v.Chair },
		sc.VoiceMember: func(v *sc.Voices) sc.VoiceSpec { return v.Member },
	}
	for _, role := range []string{sc.VoiceCISO, sc.VoiceChair, sc.VoiceMember} {
		if err := pick(role, get[role](&cfg), get[role], get[role]); err != nil {
			return nil, err
		}
	}
	return out, nil
}
