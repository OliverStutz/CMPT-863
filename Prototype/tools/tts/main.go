// Command tts synthesises every spoken line of a study config with Google
// Cloud Text-to-Speech and writes content/audio/<version>/<hash>.mp3 and
// manifest.json (section 11.1 of "02 - Architecture.md").
package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	sc "study/api/pkg/studyconfig"
)

const toolVersion = "tts/1"

type options struct {
	config       string
	out          string
	cacheDir     string
	voices       string
	defaults     string
	quotaProject string
	only         string
	dryRun       bool
	answerGapMs  int64
}

func main() {
	var o options
	flag.StringVar(&o.config, "config", "", "study config file (required)")
	flag.StringVar(&o.out, "out", "", "output directory (default <config dir>/audio/<config version>)")
	flag.StringVar(&o.cacheDir, "cache", "", "local cache directory (default <config dir>/audio/.cache)")
	flag.StringVar(&o.voices, "voices", "", "voices file that overrides the config's voices for every role it names")
	flag.StringVar(&o.defaults, "voices-default", "", "voices used where the config's voices are placeholders (default voices.default.json next to this tool)")
	flag.StringVar(&o.quotaProject, "quota-project", "", "quota project for the API (default: the one in the application-default credentials)")
	flag.StringVar(&o.only, "only", "", "synthesise only the clip with this clip or line id, and update its manifest entry")
	flag.BoolVar(&o.dryRun, "dry-run", false, "print the lines, character count and estimated cost; call no API and write nothing")
	samples := flag.String("samples", "", "candidate voices file (JSON: role -> voice names); writes one sample per voice to -out and exits")
	flag.Int64Var(&o.answerGapMs, "answer-gap-ms", 0, "extra silence between the parts of a joined answer (the parts carry their own leading and trailing silence)")
	flag.Parse()
	if o.config == "" {
		fmt.Fprintln(os.Stderr, "tts: -config is required")
		flag.Usage()
		os.Exit(2)
	}
	if *samples != "" {
		out := o.out
		if out == "" {
			out = filepath.Join(filepath.Dir(o.config), "audio", "samples")
		}
		if err := runSamples(context.Background(), o.config, *samples, out, o.quotaProject, os.Stdout, nil); err != nil {
			fmt.Fprintln(os.Stderr, "tts:", err)
			os.Exit(1)
		}
		return
	}
	if err := run(context.Background(), o, os.Stdout, nil); err != nil {
		fmt.Fprintln(os.Stderr, "tts:", err)
		os.Exit(1)
	}
}

// defaultVoicesPath finds voices.default.json: next to the source of this
// tool when run from the repo, else next to the binary.
func defaultVoicesPath(configPath string) string {
	cands := []string{
		filepath.Join(filepath.Dir(configPath), "..", "tools", "tts", "voices.default.json"),
		"voices.default.json",
		filepath.Join("tools", "tts", "voices.default.json"),
	}
	if exe, err := os.Executable(); err == nil {
		cands = append(cands, filepath.Join(filepath.Dir(exe), "voices.default.json"))
	}
	for _, c := range cands {
		if _, err := os.Stat(c); err == nil {
			return c
		}
	}
	return ""
}

func run(ctx context.Context, o options, w io.Writer, synth synthesizer) error {
	data, err := os.ReadFile(o.config)
	if err != nil {
		return err
	}
	cfg, err := sc.Parse(data, nil)
	if err != nil {
		return err
	}
	if errs := lintConfig(cfg); len(errs) > 0 {
		return fmt.Errorf("lint failed:\n%w", errors.Join(errs...))
	}

	var override, defaults *sc.Voices
	if o.voices != "" {
		v, err := loadVoices(o.voices)
		if err != nil {
			return err
		}
		override = &v
	}
	dp := o.defaults
	if dp == "" {
		dp = defaultVoicesPath(o.config)
	}
	if dp != "" {
		v, err := loadVoices(dp)
		if err != nil {
			return err
		}
		defaults = &v
	}
	voices, err := resolveVoices(cfg.Voices, override, defaults)
	if err != nil {
		return err
	}

	plans, err := buildPlan(cfg, voices, o.answerGapMs)
	if err != nil {
		return err
	}
	if o.only != "" {
		plans = filterPlan(plans, o.only)
		if len(plans) == 0 {
			return fmt.Errorf("-only %q matches no clip or line id", o.only)
		}
	}
	contentDir := filepath.Dir(o.config)
	if o.out == "" {
		o.out = filepath.Join(contentDir, "audio", cfg.Version)
	}
	if o.cacheDir == "" {
		o.cacheDir = filepath.Join(contentDir, "audio", ".cache")
	}
	ch := cache{dir: o.cacheDir}

	parts := uniqueParts(plans)
	est := estimate(parts, ch)
	fmt.Fprintf(w, "config %s (version %s), %d clips, %d distinct lines\n", o.config, cfg.Version, len(plans), len(parts))
	for _, role := range []string{sc.VoiceCISO, sc.VoiceChair, sc.VoiceMember} {
		fmt.Fprintf(w, "voice %-6s %s\n", role, voices[role].Name)
	}
	if o.dryRun {
		for _, cp := range plans {
			fmt.Fprintf(w, "\nclip %s\n", cp.ID)
			for _, p := range cp.Parts {
				_, _, hit := ch.get(p.key())
				state := "new"
				if hit {
					state = "cached"
				}
				fmt.Fprintf(w, "  %-24s %-6s %4d chars  %-6s %s\n", p.LineID, p.Role, len(p.input()), state, p.Text)
			}
		}
		fmt.Fprintln(w)
		est.print(w)
		return nil
	}

	if synth == nil && est.newChars > 0 {
		g, err := newGoogleTTS(ctx, o.quotaProject)
		if err != nil {
			return err
		}
		synth = g
	}

	// Synthesise or load every part.
	type got struct {
		audio []byte
		meta  *cachedPart
	}
	byKey := map[string]got{}
	var synthesised int
	for _, p := range parts {
		k := p.key()
		if a, m, ok := ch.get(k); ok {
			byKey[k] = got{a, m}
			continue
		}
		audio, marks, err := synth.Synthesize(ctx, p.Voice, p.input(), p.hasMarks())
		if err != nil {
			return fmt.Errorf("line %s: %w", p.LineID, err)
		}
		if err := checkMarks(p, marks); err != nil {
			return err
		}
		d, err := durationMs(audio)
		if err != nil {
			return fmt.Errorf("line %s: %w", p.LineID, err)
		}
		meta := &cachedPart{Voice: p.Voice.Name, Input: p.input(), Marks: marks, DurationMs: d}
		if err := ch.put(k, audio, meta); err != nil {
			return err
		}
		byKey[k] = got{audio, meta}
		synthesised++
		fmt.Fprintf(w, "synthesised %s (%s, %d ms)\n", p.LineID, p.Voice.Name, d)
	}

	// Assemble the clips. Padded clips come after the clip they match.
	sort.SliceStable(plans, func(i, j int) bool { return plans[i].PadTo == "" && plans[j].PadTo != "" })
	man := &sc.Manifest{ConfigVersion: cfg.Version, ToolVersion: toolVersion, Clips: map[string]sc.ManifestClip{}}
	if o.only != "" {
		if b, err := os.ReadFile(filepath.Join(o.out, "manifest.json")); err == nil {
			_ = json.Unmarshal(b, man)
			if man.Clips == nil {
				man.Clips = map[string]sc.ManifestClip{}
			}
			man.ToolVersion = toolVersion
		}
	}
	if err := os.MkdirAll(o.out, 0o755); err != nil {
		return err
	}
	for _, cp := range plans {
		var auds [][]byte
		var keys []string
		for _, p := range cp.Parts {
			auds = append(auds, byKey[p.key()].audio)
			keys = append(keys, p.key())
		}
		var padMs int64
		if cp.PadTo != "" {
			ref, ok := man.Clips[cp.PadTo]
			if !ok {
				return fmt.Errorf("clip %s: pad target %s not built", cp.ID, cp.PadTo)
			}
			unpadded, _, err := joinMP3(auds, cp.GapMs, 0)
			if err != nil {
				return fmt.Errorf("clip %s: %w", cp.ID, err)
			}
			d, err := durationMs(unpadded)
			if err != nil {
				return err
			}
			padMs = ref.DurationMs - d
		}
		clip, err := assemble(cp, auds, keys, padMs, func(k string) *cachedPart { return byKey[k].meta })
		if err != nil {
			return fmt.Errorf("clip %s: %w", cp.ID, err)
		}
		if err := checkAssetName(clip.mc.AudioKey); err != nil {
			return err
		}
		if err := writeAtomic(filepath.Join(o.out, clip.mc.AudioKey), clip.audio); err != nil {
			return err
		}
		man.Clips[cp.ID] = clip.mc
	}
	man.GeneratedAt = time.Now().UTC().Format(time.RFC3339)
	mb, _ := json.MarshalIndent(man, "", "  ")
	if err := writeAtomic(filepath.Join(o.out, "manifest.json"), append(mb, '\n')); err != nil {
		return err
	}

	var total int64
	for _, cp := range plans {
		total += man.Clips[cp.ID].DurationMs
	}
	fmt.Fprintf(w, "\nwrote %d clips to %s (%d synthesised, %d from cache), total %s\n",
		len(plans), o.out, synthesised, len(parts)-synthesised, fmtMs(total))
	est.print(w)
	return nil
}

func fmtMs(ms int64) string {
	return fmt.Sprintf("%d ms (%dm %02.1fs)", ms, ms/60000, float64(ms%60000)/1000)
}

// checkMarks fails loudly when a line asks for marks and the API returned
// none or not all of them (section 11.1 step 3).
func checkMarks(p part, got map[string]int64) error {
	want := markNames(p.input())
	var missing []string
	for _, n := range want {
		if _, ok := got[n]; !ok {
			missing = append(missing, n)
		}
	}
	if len(missing) > 0 {
		return fmt.Errorf("line %s: voice %s returned no timepoint for mark(s) %s; use a Neural2 or WaveNet voice",
			p.LineID, p.Voice.Name, strings.Join(missing, ", "))
	}
	return nil
}

// built is one assembled clip.
type built struct {
	audio []byte
	mc    sc.ManifestClip
}

// assemble builds a clip from its parts' MP3s. A one-part clip without
// padding is the part's file itself (named by its synthesis key); otherwise
// the frames are concatenated with gap and end silence and the result is
// named by the join key. Marks are shifted by each part's start offset.
func assemble(cp clipPlan, auds [][]byte, keys []string, padMs int64, meta func(string) *cachedPart) (*built, error) {
	var out []byte
	var starts []int64
	key := ""
	if len(auds) == 1 && padMs <= 0 {
		out, starts, key = auds[0], []int64{0}, keys[0]
	} else {
		var err error
		out, starts, err = joinMP3(auds, cp.GapMs, padMs)
		if err != nil {
			return nil, err
		}
		s, _ := parseMP3(auds[0])
		key = joinKey(keys, cp.GapMs, int64(s.framesFor(padMs)))
	}
	d, err := durationMs(out)
	if err != nil {
		return nil, err
	}
	mc := sc.ManifestClip{AudioKey: key + ".mp3", DurationMs: d}
	for i, p := range cp.Parts {
		m := meta(keys[i])
		if m == nil {
			continue
		}
		for name, off := range m.Marks {
			if mc.Marks == nil {
				mc.Marks = map[string]sc.Ms{}
			}
			if _, dup := mc.Marks[name]; dup {
				return nil, fmt.Errorf("mark %q occurs in two lines of the clip", name)
			}
			mc.Marks[name] = starts[i] + off
		}
		if len(cp.Parts) > 1 {
			mc.Lines = append(mc.Lines, p.LineID)
			if mc.LineStarts == nil {
				mc.LineStarts = map[string]sc.Ms{}
			}
			mc.LineStarts[p.LineID] = starts[i]
		}
	}
	return &built{audio: out, mc: mc}, nil
}
