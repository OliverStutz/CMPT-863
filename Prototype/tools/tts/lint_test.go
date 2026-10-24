package main

import (
	"os"
	"strings"
	"testing"

	sc "study/api/pkg/studyconfig"
)

func TestCheckAssetName(t *testing.T) {
	for _, bad := range []string{"s2.answer.denial.mp3", "Reason.mp3", "CHECKER-panel.png", "neutral.pdf"} {
		if checkAssetName(bad) == nil {
			t.Errorf("%q accepted", bad)
		}
	}
	for _, ok := range []string{strings.Repeat("ab", 32) + ".mp3", "s3-network-diagram.svg", "consent-placeholder.pdf"} {
		if err := checkAssetName(ok); err != nil {
			t.Errorf("%q rejected: %v", ok, err)
		}
	}
}

func TestCheckAnswerSSML(t *testing.T) {
	bad := []string{
		`<speak><prosody rate="slow">No.</prosody></speak>`,
		`<speak><emphasis level="strong">No.</emphasis></speak>`,
		`<speak><voice pitch="+2st">No.</voice></speak>`,
	}
	for _, s := range bad {
		if checkAnswerSSML(&sc.AudioLine{ID: "x", SSML: s}) == nil {
			t.Errorf("%q accepted", s)
		}
	}
	for _, s := range []string{"", `<speak>Fine. <break time="300ms"/>Really.</speak>`, `<speak><mark name="a"/>Fine.</speak>`} {
		if err := checkAnswerSSML(&sc.AudioLine{ID: "x", SSML: s}); err != nil {
			t.Errorf("%q rejected: %v", s, err)
		}
	}
}

func TestSSMLWrapper(t *testing.T) {
	l := func(s string) *sc.AudioLine { return &sc.AudioLine{SSML: s} }
	same := [][2]string{
		{`<speak>A reason.</speak>`, `<speak>A denial, longer.</speak>`},
		{``, ``},
		{`<speak><p>A. <break time="1s"/> B.</p></speak>`, `<speak> <p>C.</p> </speak>`},
	}
	for _, c := range same {
		if ssmlWrapper(l(c[0])) != ssmlWrapper(l(c[1])) {
			t.Errorf("%q and %q differ: %q %q", c[0], c[1], ssmlWrapper(l(c[0])), ssmlWrapper(l(c[1])))
		}
	}
	diff := [][2]string{
		{`<speak>A.</speak>`, `<speak><p>B.</p></speak>`},
		{``, `<speak><s>B.</s></speak>`},
	}
	for _, c := range diff {
		if ssmlWrapper(l(c[0])) == ssmlWrapper(l(c[1])) {
			t.Errorf("%q and %q treated as the same wrapper", c[0], c[1])
		}
	}
}

func TestLintConfig(t *testing.T) {
	data, _ := os.ReadFile(v14)
	cfg, err := sc.Parse(data, nil)
	if err != nil {
		t.Fatal(err)
	}
	if errs := lintConfig(cfg); len(errs) != 0 {
		t.Fatalf("v1.4 fails: %v", errs)
	}
	s2 := cfg.SceneByID("s2")
	s2.Answers.Denial.SSML = `<speak><p>` + s2.Answers.Denial.Text + `</p></speak>`
	s2.Answers.Reason.SSML = `<speak><emphasis>` + s2.Answers.Reason.Text + `</emphasis></speak>`
	cfg.Consent.PdfAssetKey = "neutral-consent.pdf"
	errs := lintConfig(cfg)
	joined := ""
	for _, e := range errs {
		joined += e.Error() + "\n"
	}
	for _, want := range []string{"different SSML wrappers", "emphasis", "neutral-consent.pdf"} {
		if !strings.Contains(joined, want) {
			t.Errorf("missing %q in:\n%s", want, joined)
		}
	}
}
