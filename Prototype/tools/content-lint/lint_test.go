package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	sc "study/api/pkg/studyconfig"
)

const v14 = "../../content/study-config.v1.4.json"
const schema = "../../content/schema/study-config.schema.json"

func load(t *testing.T) *sc.StudyConfig {
	t.Helper()
	b, err := os.ReadFile(v14)
	if err != nil {
		t.Fatal(err)
	}
	c, err := sc.Parse(b, nil)
	if err != nil {
		t.Fatal(err)
	}
	return c
}

func rules(t *testing.T) *Rules {
	r, err := parseRules(defaultRules)
	if err != nil {
		t.Fatal(err)
	}
	return r
}

func has(list []string, sub string) bool {
	for _, s := range list {
		if strings.Contains(s, sub) {
			return true
		}
	}
	return false
}

func TestV14Passes(t *testing.T) {
	rep := Lint(load(t), rules(t), Options{})
	if len(rep.Failures) != 0 || len(rep.Warnings) != 0 {
		t.Fatalf("failures %v warnings %v", rep.Failures, rep.Warnings)
	}
}

func TestFlagWording(t *testing.T) {
	r := rules(t)
	bad := []string{
		"The CISO is not being honest", "Possible lie in the total", "He may be lying", "Figures hide a gap",
		"Hidden cost", "Not the whole truth", "Low confidence in this figure", "92% likely wrong", "Accuracy unclear",
		"75 percent of sites", "Dishonest framing", "Do not trust this number", "Probably wrong",
	}
	for _, s := range bad {
		if firstMatch(r.flag, s) == "" {
			t.Errorf("flag wording %q accepted", s)
		}
	}
	good := []string{
		"$8 million does not follow from shown figures", "Head office is a single point of failure",
		"Firewall cost per site looks high", "$780,000 across 13 sites is about $60,000 per site",
		"Applies to all 12 warehouses",
	}
	for _, s := range good {
		if m := firstMatch(r.flag, s); m != "" {
			t.Errorf("flag wording %q rejected for %q", s, m)
		}
	}
	c := load(t)
	c.SceneByID("s2").Checker.Markings[0].Detail = "The total is 80% sure."
	rep := Lint(c, r, Options{})
	if !has(rep.Failures, `"80%"`) {
		t.Fatalf("%v", rep.Failures)
	}
}

func TestCISOWording(t *testing.T) {
	r := rules(t)
	for _, s := range []string{"Our AI flagged it.", "an artificial intelligence", "the AI model", "a language model",
		"the algorithm says", "The checker is wrong.", "ChatGPT", "an LLM"} {
		if firstMatch(r.ciso, s) == "" {
			t.Errorf("CISO line %q accepted", s)
		}
	}
	for _, s := range []string{"Our business model is sound.", "We said it again.", "Paid in full.", "I checked every number."} {
		if m := firstMatch(r.ciso, s); m != "" {
			t.Errorf("CISO line %q rejected for %q", s, m)
		}
	}
	c := load(t)
	c.SceneByID("s4").Answers.Denial.Text = "The checker is just an algorithm."
	rep := Lint(c, r, Options{})
	if !has(rep.Failures, "CISO line s4.answer.b") {
		t.Fatalf("%v", rep.Failures)
	}
	// The chair and the board member may say "checker".
	c = load(t)
	rep = Lint(c, r, Options{})
	if has(rep.Failures, "s2.question.challenge") {
		t.Fatal("member's challenge question rejected")
	}
}

func TestCheckerTextBeforeDebrief(t *testing.T) {
	c := load(t)
	c.Setup.PreStartText.Checker[0] = "An AI checker reads each slide."
	c.Labels.PanelTitle = "Language model review"
	rep := Lint(c, rules(t), Options{})
	if !has(rep.Failures, "preStartText.checker[0]") || !has(rep.Failures, "labels.panelTitle") {
		t.Fatalf("%v", rep.Failures)
	}
}

func TestNoIssuesLabel(t *testing.T) {
	c := load(t)
	c.Labels.NoIssues = "Looks good"
	rep := Lint(c, rules(t), Options{})
	if !has(rep.Failures, "labels.noIssues") {
		t.Fatalf("%v", rep.Failures)
	}
	c = load(t)
	c.SceneByID("s1").Checker.Text = "All clear"
	rep = Lint(c, rules(t), Options{})
	if !has(rep.Failures, "no-issues panel text") {
		t.Fatalf("%v", rep.Failures)
	}
	c.SceneByID("s1").Checker.Text = "No issues found"
	if rep = Lint(c, rules(t), Options{}); len(rep.Failures) != 0 {
		t.Fatalf("%v", rep.Failures)
	}
}

func TestAnchorAndMarks(t *testing.T) {
	c := load(t)
	m := &c.SceneByID("s3").Checker.Markings[0]
	m.AnchorText = "and all thirteen warehouses connect to it"
	rep := Lint(c, rules(t), Options{})
	if !has(rep.Failures, "anchorText") {
		t.Fatalf("%v", rep.Failures)
	}
	c = load(t)
	m = &c.SceneByID("s3").Checker.Markings[0]
	m.AnchorText = "AND ALL TWELVE   warehouses connect to it" // case and spacing do not matter
	m.AppearAtMark = "warehouses"
	rep = Lint(c, rules(t), Options{})
	if has(rep.Failures, "anchorText") || !has(rep.Failures, `appearAtMark "warehouses"`) {
		t.Fatalf("%v", rep.Failures)
	}
	n := c.SceneByID("s3").Narration
	n.SSML = "<speak>Here is how we solve it. <mark name=\"warehouses\"/>All twelve.</speak>"
	if rep = Lint(c, rules(t), Options{}); len(rep.Failures) != 0 {
		t.Fatalf("%v", rep.Failures)
	}
	m.AppearAtMark, m.AppearAtMs = "", nil
	if rep = Lint(c, rules(t), Options{}); !has(rep.Failures, "neither appearAtMark nor appearAtMs") {
		t.Fatalf("%v", rep.Failures)
	}
}

func TestAnswerSSMLAndAssets(t *testing.T) {
	c := load(t)
	a := c.SceneByID("s2").Answers
	a.Reason.SSML = "<speak><prosody rate=\"slow\">" + a.Reason.Text + "</prosody></speak>"
	c.Consent.PdfAssetKey = "denial.pdf"
	rep := Lint(c, rules(t), Options{})
	for _, want := range []string{"prosody", "different SSML wrappers", `"denial.pdf"`} {
		if !has(rep.Failures, want) {
			t.Errorf("missing %q in %v", want, rep.Failures)
		}
	}
}

func TestPublishFieldsAndStructure(t *testing.T) {
	c := load(t)
	d := int64(5)
	c.SceneByID("s1").Narration.DurationMs = &d
	s5 := c.SceneByID("s5")
	s5.RoutineQuestion = &sc.AudioLine{ID: "s5.q", Voice: "chair", Text: "Why?"}
	rep := Lint(c, rules(t), Options{})
	if !has(rep.Failures, "filled at publish") || !has(rep.Failures, "not flagged but has questions") {
		t.Fatalf("%v", rep.Failures)
	}
}

func TestWarningsDifferFromAbstractV8(t *testing.T) {
	c := load(t)
	// Unflag slide 4: four slides' worth of changes, two flagged slides.
	s4 := c.SceneByID("s4")
	s4.Checker = &sc.CheckerPanel{Status: sc.PanelNoIssues, AppearAtMs: s4.Checker.Markings[0].AppearAtMs}
	s4.RoutineQuestion, s4.ParticipantQuestions, s4.Answers = nil, nil, nil
	var kept []sc.Scene
	for _, s := range c.Scenes {
		if s.ForSlide == "s4" || s.ID == "rating.T4b" {
			continue
		}
		if s.ID == "rating.T4a" {
			s.Point = "T4"
		}
		kept = append(kept, s)
	}
	c.Scenes = kept
	rep := Lint(c, rules(t), Options{})
	if len(rep.Failures) != 0 {
		t.Fatalf("failures %v", rep.Failures)
	}
	if !has(rep.Warnings, "2 flagged slides: differs from Abstract v8") || !has(rep.Warnings, "design limit") {
		t.Fatalf("warnings %v", rep.Warnings)
	}
	// A wrong rating point fails.
	c.SceneByID("rating.T4a").Point = "T4a"
	if rep = Lint(c, rules(t), Options{}); !has(rep.Failures, "the plan gives T4") {
		t.Fatalf("%v", rep.Failures)
	}
}

func TestManifestChecks(t *testing.T) {
	c := load(t)
	m := &sc.Manifest{Clips: map[string]sc.ManifestClip{
		"s2.narration": {DurationMs: 10000},
		"s2.answer.a":  {DurationMs: 15000},
		"s2.answer.b":  {DurationMs: 15000},
	}}
	rep := Lint(c, rules(t), Options{Manifest: m})
	if !has(rep.Failures, "after the narration ends") {
		t.Fatalf("%v", rep.Failures)
	}
	m.Clips["s2.narration"] = sc.ManifestClip{DurationMs: 40000}
	m.Clips["s2.answer.b"] = sc.ManifestClip{DurationMs: 12000}
	rep = Lint(c, rules(t), Options{Manifest: m})
	if len(rep.Failures) != 0 || !has(rep.Warnings, "denial is 3000 ms shorter") {
		t.Fatalf("%v %v", rep.Failures, rep.Warnings)
	}
}

func TestRunExitCodes(t *testing.T) {
	var out, errb bytes.Buffer
	if code := run([]string{"-no-manifest", "-schema", schema, v14}, &out, &errb); code != 0 {
		t.Fatalf("v1.4 exit %d: %s %s", code, out.String(), errb.String())
	}
	// Schema failure: an unknown key.
	b, _ := os.ReadFile(v14)
	var raw map[string]any
	_ = json.Unmarshal(b, &raw)
	raw["unknownKey"] = true
	p := filepath.Join(t.TempDir(), "bad.json")
	nb, _ := json.Marshal(raw)
	_ = os.WriteFile(p, nb, 0o644)
	out.Reset()
	if code := run([]string{"-schema", schema, p}, &out, &errb); code != 1 || !strings.Contains(out.String(), "schema") {
		t.Fatalf("exit %d: %s", code, out.String())
	}
	// Wording failure through the CLI.
	raw2 := map[string]any{}
	_ = json.Unmarshal(b, &raw2)
	raw2["labels"].(map[string]any)["noIssues"] = "Nothing to see"
	nb, _ = json.Marshal(raw2)
	_ = os.WriteFile(p, nb, 0o644)
	out.Reset()
	if code := run([]string{"-schema", schema, p}, &out, &errb); code != 1 {
		t.Fatalf("exit %d: %s", code, out.String())
	}
}
