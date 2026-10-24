package main

import (
	"bytes"
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	sc "study/api/pkg/studyconfig"
)

const v14 = "../../content/study-config.v1.4.json"

var hashName = regexp.MustCompile(`^[0-9a-f]{64}\.mp3$`)

func voice(name string) sc.VoiceSpec {
	return sc.VoiceSpec{Provider: "google-tts", Name: name, LanguageCode: "en-US"}
}

func TestSynthKey(t *testing.T) {
	a := synthKey(voice("en-US-Neural2-J"), "Hello.", "")
	if len(a) != 64 || !regexp.MustCompile(`^[0-9a-f]+$`).MatchString(a) {
		t.Fatalf("key %q", a)
	}
	if a != synthKey(voice("en-US-Neural2-J"), "Hello.", "") {
		t.Fatal("key not stable")
	}
	r := 1.1
	vr := voice("en-US-Neural2-J")
	vr.SpeakingRate = &r
	for name, b := range map[string]string{
		"voice": synthKey(voice("en-US-Neural2-D"), "Hello.", ""),
		"text":  synthKey(voice("en-US-Neural2-J"), "Hello!", ""),
		"ssml":  synthKey(voice("en-US-Neural2-J"), "Hello.", "<speak>Hello.</speak>"),
		"rate":  synthKey(vr, "Hello.", ""),
	} {
		if a == b {
			t.Errorf("changing the %s does not change the key", name)
		}
	}
	// Field boundaries cannot be shifted to collide.
	if synthKey(voice("x"), "ab", "") == synthKey(voice("x"), "a", "b") {
		t.Fatal("text and ssml collide")
	}
	j1 := joinKey([]string{a, a}, 0, 0)
	if j1 == joinKey([]string{a, a}, 0, 3) || j1 == joinKey([]string{a, a}, 600, 0) || j1 == joinKey([]string{a}, 0, 0) {
		t.Fatal("join key ignores gap, padding or parts")
	}
}

// fakeSynth returns the fixture clip for every line and a mark at 500 ms for
// every mark the input names.
type fakeSynth struct {
	audio [][]byte
	calls int
	noMks bool
}

func (f *fakeSynth) Synthesize(_ context.Context, v sc.VoiceSpec, input string, marks bool) ([]byte, map[string]int64, error) {
	f.calls++
	var m map[string]int64
	if marks && !f.noMks {
		m = map[string]int64{}
		for _, n := range markNames(input) {
			m[n] = 500
		}
	}
	// Vary the clip length with the input so clips differ.
	a := f.audio[len(input)%len(f.audio)]
	return a, m, nil
}

func fixtures(t *testing.T) [][]byte {
	g := readFixture(t, "google-neural2-j.mp3")
	return [][]byte{g, g[:192*40], g[:192*90]}
}

func writeVoices(t *testing.T, dir string) string {
	p := filepath.Join(dir, "voices.json")
	v := sc.Voices{CISO: voice("en-US-Neural2-J"), Chair: voice("en-US-Neural2-F"), Member: voice("en-US-Neural2-A")}
	b, _ := json.Marshal(v)
	if err := os.WriteFile(p, b, 0o644); err != nil {
		t.Fatal(err)
	}
	return p
}

func TestRunV14WithFakeSynth(t *testing.T) {
	dir := t.TempDir()
	o := options{config: v14, out: filepath.Join(dir, "out"), cacheDir: filepath.Join(dir, "cache"), defaults: writeVoices(t, dir)}
	fs := &fakeSynth{audio: fixtures(t)}
	var log bytes.Buffer
	if err := run(context.Background(), o, &log, fs); err != nil {
		t.Fatal(err, log.String())
	}
	var m sc.Manifest
	b, err := os.ReadFile(filepath.Join(o.out, "manifest.json"))
	if err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(b, &m); err != nil {
		t.Fatal(err)
	}
	if m.ConfigVersion != "1.4.0" {
		t.Fatalf("configVersion %q", m.ConfigVersion)
	}
	// Every clip the server resolves is in the manifest.
	data, _ := os.ReadFile(v14)
	cfg, _ := sc.Parse(data, nil)
	_, estimated := sc.ResolveClips(cfg, &m)
	if len(estimated) > 0 {
		t.Fatalf("server would estimate %v", estimated)
	}
	for id, c := range m.Clips {
		if !hashName.MatchString(c.AudioKey) {
			t.Errorf("%s: audio key %q", id, c.AudioKey)
		}
		fb, err := os.ReadFile(filepath.Join(o.out, c.AudioKey))
		if err != nil {
			t.Fatalf("%s: %v", id, err)
		}
		d, _ := durationMs(fb)
		if d != c.DurationMs {
			t.Errorf("%s: manifest says %d ms, file is %d ms", id, c.DurationMs, d)
		}
	}
	// Opening: five lines joined with 600 ms pauses, line starts recorded.
	op := m.Clips["opening"]
	if len(op.Lines) != 5 || op.LineStarts["opening.chair.1"] != 0 || op.LineStarts["opening.chair.2"] == 0 {
		t.Fatalf("opening clip %+v", op)
	}
	// Joined answers on s3: the denial is padded to the reason's length.
	r, d := m.Clips["s3.answers.reason"], m.Clips["s3.answers.denial"]
	if r.AudioKey == d.AudioKey {
		t.Fatal("reason and denial share a file")
	}
	if diff := r.DurationMs - d.DurationMs; diff < -12 || diff > 12 {
		t.Fatalf("padded denial %d ms vs reason %d ms", d.DurationMs, r.DurationMs)
	}
	if len(r.Lines) != 3 || len(d.Lines) != 2 || d.Lines[0] != "s3.answer.part1" || d.Lines[1] != "s3.answer.part3" {
		t.Fatalf("joined lines %v %v", r.Lines, d.Lines)
	}
	// The shared opening sentence is the same bytes at the start of both clips.
	rb, _ := os.ReadFile(filepath.Join(o.out, r.AudioKey))
	db, _ := os.ReadFile(filepath.Join(o.out, d.AudioKey))
	first := m.Clips["s3.answers.reason"].LineStarts["s3.answer.part2"] / 24 * 192
	if !bytes.Equal(rb[:first], db[:first]) {
		t.Fatal("shared opening differs between reason and denial")
	}
	// Whole answers without padding keep their own lengths.
	if _, ok := m.Clips["s2.answer.a"]; !ok {
		t.Fatal("s2 reason missing")
	}

	// Second run: everything from the cache, no API call.
	calls := fs.calls
	if err := run(context.Background(), o, &log, fs); err != nil {
		t.Fatal(err)
	}
	if fs.calls != calls {
		t.Fatalf("second run made %d API calls", fs.calls-calls)
	}
	b2, _ := os.ReadFile(filepath.Join(o.out, "manifest.json"))
	var m2 sc.Manifest
	_ = json.Unmarshal(b2, &m2)
	for id, c := range m.Clips {
		if m2.Clips[id].AudioKey != c.AudioKey {
			t.Fatalf("%s changed name on a cached run", id)
		}
	}
}

func TestRunDryRunCallsNothing(t *testing.T) {
	dir := t.TempDir()
	o := options{config: v14, out: filepath.Join(dir, "out"), cacheDir: filepath.Join(dir, "cache"), defaults: writeVoices(t, dir), dryRun: true}
	fs := &fakeSynth{audio: fixtures(t)}
	var log bytes.Buffer
	if err := run(context.Background(), o, &log, fs); err != nil {
		t.Fatal(err)
	}
	if fs.calls != 0 {
		t.Fatal("dry run called the API")
	}
	if _, err := os.Stat(o.out); err == nil {
		t.Fatal("dry run wrote output")
	}
	if !strings.Contains(log.String(), "estimated cost") || !strings.Contains(log.String(), "s3.answer.part2") {
		t.Fatal(log.String())
	}
}

func TestRunOnly(t *testing.T) {
	dir := t.TempDir()
	o := options{config: v14, out: filepath.Join(dir, "out"), cacheDir: filepath.Join(dir, "cache"), defaults: writeVoices(t, dir), only: "s3.answer.part1"}
	fs := &fakeSynth{audio: fixtures(t)}
	var log bytes.Buffer
	if err := run(context.Background(), o, &log, fs); err != nil {
		t.Fatal(err)
	}
	b, _ := os.ReadFile(filepath.Join(o.out, "manifest.json"))
	var m sc.Manifest
	_ = json.Unmarshal(b, &m)
	if len(m.Clips) != 2 || fs.calls != 3 {
		t.Fatalf("only: %d clips, %d calls", len(m.Clips), fs.calls)
	}
}

func TestMarksShiftedAndRequired(t *testing.T) {
	data, _ := os.ReadFile(v14)
	cfg, err := sc.Parse(data, nil)
	if err != nil {
		t.Fatal(err)
	}
	cfg.Scenes[0].Lines[1].SSML = `<speak>Presenting today <mark name="card"/>is Daniel Hale.</speak>`
	dir := t.TempDir()
	cp := filepath.Join(dir, "c.json")
	b, _ := json.Marshal(cfg)
	_ = os.WriteFile(cp, b, 0o644)
	o := options{config: cp, out: filepath.Join(dir, "out"), cacheDir: filepath.Join(dir, "cache"), defaults: writeVoices(t, dir), only: "opening"}
	fs := &fakeSynth{audio: fixtures(t)}
	var log bytes.Buffer
	if err := run(context.Background(), o, &log, fs); err != nil {
		t.Fatal(err)
	}
	mb, _ := os.ReadFile(filepath.Join(o.out, "manifest.json"))
	var m sc.Manifest
	_ = json.Unmarshal(mb, &m)
	op := m.Clips["opening"]
	if op.Marks["card"] != op.LineStarts["opening.chair.2"]+500 {
		t.Fatalf("mark %d, line start %d", op.Marks["card"], op.LineStarts["opening.chair.2"])
	}
	// A voice that returns no marks fails loudly.
	o.cacheDir = filepath.Join(dir, "cache2")
	if err := run(context.Background(), o, &log, &fakeSynth{audio: fixtures(t), noMks: true}); err == nil || !strings.Contains(err.Error(), "no timepoint") {
		t.Fatalf("missing marks not reported: %v", err)
	}
}

func TestResolveVoices(t *testing.T) {
	cfg := sc.Voices{CISO: voice("TBD"), Chair: voice("en-US-Wavenet-F"), Member: voice("")}
	def := sc.Voices{CISO: voice("en-US-Neural2-J"), Chair: voice("en-US-Neural2-F"), Member: voice("en-US-Neural2-A")}
	v, err := resolveVoices(cfg, nil, &def)
	if err != nil {
		t.Fatal(err)
	}
	if v["ciso"].Name != "en-US-Neural2-J" || v["chair"].Name != "en-US-Wavenet-F" || v["member"].Name != "en-US-Neural2-A" {
		t.Fatalf("%v", v)
	}
	ov := sc.Voices{Chair: voice("en-US-Neural2-H")}
	v, _ = resolveVoices(cfg, &ov, &def)
	if v["chair"].Name != "en-US-Neural2-H" || v["ciso"].Name != "en-US-Neural2-J" {
		t.Fatalf("override: %v", v)
	}
	if _, err := resolveVoices(cfg, nil, nil); err == nil {
		t.Fatal("placeholder without default accepted")
	}
}

func TestPlainTextIsEscaped(t *testing.T) {
	p := part{Text: `R&D <costs> "rise"`}
	if got := p.input(); got != "<speak>R&amp;D &lt;costs&gt; &#34;rise&#34;</speak>" {
		t.Fatalf("input %q", got)
	}
}
