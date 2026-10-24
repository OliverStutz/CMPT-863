package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"regexp"
	"sort"
	"strings"

	sc "study/api/pkg/studyconfig"
)

//go:embed rules.json
var defaultRules []byte

// Rules are the deny lists and expected values of section 4.2.
type Rules struct {
	FlagDeny        []string `json:"flagDeny"`
	CISODeny        []string `json:"cisoDeny"`
	CheckerTextDeny []string `json:"checkerTextDeny"`
	NoIssuesLabel   string   `json:"noIssuesLabel"`
	AbstractV8      struct {
		Slides  int   `json:"slides"`
		Flagged []int `json:"flagged"`
	} `json:"abstractV8"`

	flag, ciso, checker []*regexp.Regexp
}

func parseRules(b []byte) (*Rules, error) {
	var r Rules
	if err := json.Unmarshal(b, &r); err != nil {
		return nil, fmt.Errorf("rules: %w", err)
	}
	comp := func(list []string) ([]*regexp.Regexp, error) {
		var out []*regexp.Regexp
		for _, p := range list {
			re, err := regexp.Compile(p)
			if err != nil {
				return nil, fmt.Errorf("rules: %q: %w", p, err)
			}
			out = append(out, re)
		}
		return out, nil
	}
	var err error
	if r.flag, err = comp(r.FlagDeny); err != nil {
		return nil, err
	}
	if r.ciso, err = comp(r.CISODeny); err != nil {
		return nil, err
	}
	if r.checker, err = comp(r.CheckerTextDeny); err != nil {
		return nil, err
	}
	return &r, nil
}

func firstMatch(res []*regexp.Regexp, s string) string {
	for _, re := range res {
		if m := re.FindString(s); m != "" {
			return m
		}
	}
	return ""
}

// Report collects failures and warnings.
type Report struct {
	Failures []string
	Warnings []string
}

func (r *Report) fail(f string, a ...any) { r.Failures = append(r.Failures, fmt.Sprintf(f, a...)) }
func (r *Report) warn(f string, a ...any) { r.Warnings = append(r.Warnings, fmt.Sprintf(f, a...)) }

// Options of a lint run.
type Options struct {
	Resolved bool         // the config is a resolved config (audioKey etc. filled)
	Strict   bool         // vlmRun warnings become failures
	Manifest *sc.Manifest // when set, the post-TTS checks run
}

var (
	hashKey    = regexp.MustCompile(`^[0-9a-f]{64}\.mp3$`)
	assetWords = regexp.MustCompile(`(?i)reason|denial|checker|neutral`)
	markTag    = regexp.MustCompile(`<mark\s+name\s*=\s*["']([^"']+)["']`)
	prosodyTag = regexp.MustCompile(`(?i)<\s*/?\s*(prosody|emphasis)\b`)
	ratePitch  = regexp.MustCompile(`(?i)\b(rate|pitch|volume)\s*=`)
	anyTag     = regexp.MustCompile(`<[^>]*>`)
	spaceRun   = regexp.MustCompile(`\s+`)
)

func norm(s string) string {
	return strings.ToLower(strings.TrimSpace(spaceRun.ReplaceAllString(s, " ")))
}

// lineRef is an audio line with where it sits in the config.
type lineRef struct {
	l      *sc.AudioLine
	where  string
	answer bool
}

func allLines(c *sc.StudyConfig) []lineRef {
	var out []lineRef
	add := func(l *sc.AudioLine, where string, answer bool) {
		if l != nil {
			out = append(out, lineRef{l, where, answer})
		}
	}
	add(&c.Setup.AudioCheck.Clip, "setup.audioCheck.clip", false)
	for i := range c.Setup.ExampleScene.Lines {
		add(&c.Setup.ExampleScene.Lines[i], "setup.exampleScene", false)
	}
	for i := range c.Scenes {
		s := &c.Scenes[i]
		w := "scene " + s.ID
		for j := range s.Lines {
			add(&s.Lines[j], w, false)
		}
		add(s.Narration, w+" narration", false)
		add(s.RoutineQuestion, w+" routineQuestion", false)
		if q := s.ParticipantQuestions; q != nil {
			add(&q.Explain, w+" participantQuestions.explain", false)
			add(&q.Challenge, w+" participantQuestions.challenge", false)
		}
		if a := s.Answers; a != nil {
			add(a.Reason, w+" answers.reason", true)
			add(a.Denial, w+" answers.denial", true)
			add(a.Opening, w+" answers.opening", true)
			add(a.ReasonBody, w+" answers.reasonBody", true)
			add(a.Closing, w+" answers.closing", true)
		}
	}
	return out
}

// ssmlWrapper is the markup outside the spoken text (see tools/tts).
func ssmlWrapper(l *sc.AudioLine) string {
	s := strings.TrimSpace(l.SSML)
	if s == "" {
		return "<speak>|</speak>"
	}
	locs := anyTag.FindAllStringIndex(s, -1)
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
	return strings.Join(strings.Fields(pre), " ") + "|" + strings.Join(strings.Fields(suf), " ")
}

// Lint applies the rules of section 4.2 to a parsed config.
func Lint(c *sc.StudyConfig, r *Rules, o Options) *Report {
	rep := &Report{}

	// Audio lines: CISO wording, answer SSML, publish-only fields.
	for _, lr := range allLines(c) {
		l := lr.l
		if l.Voice == sc.VoiceCISO {
			if m := firstMatch(r.ciso, l.Text+"\n"+l.SSML); m != "" {
				rep.fail("%s: CISO line %s says %q", lr.where, l.ID, m)
			}
		}
		if lr.answer {
			if l.Voice != sc.VoiceCISO {
				rep.fail("%s: answer line %s has voice %q, want ciso", lr.where, l.ID, l.Voice)
			}
			if m := prosodyTag.FindString(l.SSML); m != "" {
				rep.fail("%s: answer line %s carries %q", lr.where, l.ID, m)
			}
			if m := ratePitch.FindString(l.SSML); m != "" {
				rep.fail("%s: answer line %s changes %q", lr.where, l.ID, m)
			}
		}
		if o.Resolved {
			if l.AudioKey == "" || !hashKey.MatchString(l.AudioKey) {
				if l.AudioKey != "" {
					rep.fail("%s: line %s audioKey %q is not <sha256>.mp3", lr.where, l.ID, l.AudioKey)
				}
			}
		} else if l.AudioKey != "" || l.DurationMs != nil || l.Marks != nil {
			rep.fail("%s: line %s sets audioKey, durationMs or marks; these are filled at publish", lr.where, l.ID)
		}
	}

	// Labels and other text describing the checker.
	if c.Labels.NoIssues != r.NoIssuesLabel {
		rep.fail("labels.noIssues is %q, want the configured %q", c.Labels.NoIssues, r.NoIssuesLabel)
	}
	checkerTexts := map[string]string{
		"labels.panelTitle": c.Labels.PanelTitle, "labels.noIssues": c.Labels.NoIssues, "labels.panelIdle": c.Labels.PanelIdle,
	}
	for i, t := range c.Setup.PreStartText.Both {
		checkerTexts[fmt.Sprintf("setup.preStartText.both[%d]", i)] = t
	}
	for i, t := range c.Setup.PreStartText.Checker {
		checkerTexts[fmt.Sprintf("setup.preStartText.checker[%d]", i)] = t
	}
	keys := make([]string, 0, len(checkerTexts))
	for k := range checkerTexts {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		if m := firstMatch(r.checker, checkerTexts[k]); m != "" {
			rep.fail("%s says %q before the debrief", k, m)
		}
	}

	// Assets.
	assets := map[string]string{"consent.pdfAssetKey": c.Consent.PdfAssetKey}
	blockAssets := func(bs []sc.Block, where string) {
		for _, b := range bs {
			if b.AssetKey != "" {
				assets[where+" block "+b.ID] = b.AssetKey
			}
		}
	}
	blockAssets(c.Setup.ExampleScene.Layout, "setup.exampleScene")
	for _, lr := range allLines(c) {
		if lr.l.AudioKey != "" {
			assets[lr.where+" "+lr.l.ID] = lr.l.AudioKey
		}
	}

	// Scenes.
	slides := c.Slides()
	if len(slides) == 0 {
		rep.fail("no slide scene")
	}
	var flaggedNums []int
	for i := range c.Scenes {
		s := &c.Scenes[i]
		blockAssets(s.Layout, "scene "+s.ID)
		if s.ProfileCard != nil && s.ProfileCard.PortraitKey != "" {
			assets["scene "+s.ID+" portraitKey"] = s.ProfileCard.PortraitKey
		}
	}
	for n, s := range slides {
		lintSlide(c, s, n+1, r, o, rep)
		if s.Flagged() {
			flaggedNums = append(flaggedNums, n+1)
		}
	}
	if len(slides) > 0 && len(flaggedNums) == 0 {
		rep.fail("no flagged slide")
	}
	for where, k := range assets {
		if m := assetWords.FindString(k); m != "" {
			rep.fail("%s: asset file name %q contains %q", where, k, m)
		}
	}

	lintFlow(c, rep)

	// Setup, background, debrief.
	okOpt := false
	for _, op := range c.Setup.AudioCheck.Options {
		okOpt = okOpt || op.ID == c.Setup.AudioCheck.CorrectOptionID
	}
	if !okOpt {
		rep.fail("setup.audioCheck.correctOptionId %q is not among the options", c.Setup.AudioCheck.CorrectOptionID)
	}
	if bg := c.SceneOfKind(sc.SceneBackground); bg != nil {
		hp := false
		for _, it := range bg.Items {
			hp = hp || it.Kind == sc.ItemHoneypot
		}
		if !hp {
			rep.fail("background scene %s has no honeypot item", bg.ID)
		}
	}
	if db := c.SceneOfKind(sc.SceneDebrief); db != nil && strings.TrimSpace(db.DoNotDiscuss) == "" {
		rep.fail("debrief.doNotDiscuss is empty")
	}

	// Warnings against Abstract v8 and Oliver's decisions.
	if len(slides) != r.AbstractV8.Slides {
		rep.warn("%d slide scenes: differs from Abstract v8 (%d)", len(slides), r.AbstractV8.Slides)
	}
	if len(flaggedNums) != len(r.AbstractV8.Flagged) {
		rep.warn("%d flagged slides: differs from Abstract v8 (%d)", len(flaggedNums), len(r.AbstractV8.Flagged))
	} else if fmt.Sprint(flaggedNums) != fmt.Sprint(r.AbstractV8.Flagged) {
		rep.warn("flagged slides %v: differs from Abstract v8 (%v)", flaggedNums, r.AbstractV8.Flagged)
	}
	switch {
	case len(flaggedNums) == 1:
		rep.warn("design limit: one flagged slide gives no denial rotation (section 5.5) and no session gets two reasoned answers and one denial")
	case len(flaggedNums) == 2:
		rep.warn("design limit: with fewer than three flagged slides a session cannot get two reasoned answers and one denial")
	}
	lintMeasures(c, rep)
	return rep
}

func lintSlide(c *sc.StudyConfig, s *sc.Scene, n int, r *Rules, o Options, rep *Report) {
	ids := map[string]bool{}
	for _, b := range s.Layout {
		ids[b.ID] = true
		for _, rt := range []sc.RichText{b.Text, b.Value, b.Caption} {
			for _, sp := range rt.Spans() {
				ids[sp.ID] = true
			}
		}
		if b.Data != nil {
			for _, se := range b.Data.Series {
				ids[se.ID] = true
			}
			for _, m := range b.Data.Markers {
				ids[m.ID] = true
			}
		}
	}
	if len(s.Layout) == 0 || s.Layout[0].Kind != sc.BlockTitle {
		rep.fail("slide %s: the first block must be the title", s.ID)
	}
	titles := 0
	for _, b := range s.Layout {
		if b.Kind == sc.BlockTitle {
			titles++
		}
	}
	if titles != 1 {
		rep.fail("slide %s: %d title blocks, want exactly one", s.ID, titles)
	}

	narr := s.Narration
	if narr == nil {
		rep.fail("slide %s: no narration", s.ID)
		return
	}
	marks := map[string]bool{}
	for _, m := range markTag.FindAllStringSubmatch(narr.SSML, -1) {
		marks[m[1]] = true
	}
	var narrMs int64 = -1
	if o.Manifest != nil {
		if mc, ok := o.Manifest.Clips[narr.ID]; ok {
			narrMs = mc.DurationMs
		}
	}
	timeCheck := func(what, mark string, ms *sc.Ms) {
		if mark == "" && ms == nil {
			rep.fail("slide %s: %s has neither appearAtMark nor appearAtMs", s.ID, what)
		}
		if mark != "" && !marks[mark] {
			rep.fail("slide %s: %s appearAtMark %q has no <mark name> in the narration SSML", s.ID, what, mark)
		}
		if narrMs >= 0 {
			var at int64 = -1
			if mark != "" {
				if v, ok := o.Manifest.Clips[narr.ID].Marks[mark]; ok {
					at = v
				}
			}
			if at < 0 && ms != nil {
				at = *ms
			}
			if at > narrMs {
				rep.fail("slide %s: %s appears at %d ms, after the narration ends (%d ms)", s.ID, what, at, narrMs)
			}
		}
	}

	ch := s.Checker
	if ch == nil {
		rep.fail("slide %s: no checker panel", s.ID)
		return
	}
	flagged := s.Flagged()
	switch ch.Status {
	case sc.PanelNoIssues:
		if ch.Text != "" && ch.Text != c.Labels.NoIssues {
			rep.fail("slide %s: no-issues panel text %q differs from labels.noIssues %q", s.ID, ch.Text, c.Labels.NoIssues)
		}
		timeCheck("no-issues panel", ch.AppearAtMark, ch.AppearAtMs)
	case sc.PanelFlagged:
		if len(ch.Markings) == 0 {
			rep.fail("slide %s: flagged with no markings", s.ID)
		}
		for _, m := range ch.Markings {
			where := "marking " + m.ID
			for _, field := range []struct{ name, text string }{{"label", m.Label}, {"detail", m.Detail}} {
				if hit := firstMatch(r.flag, field.text); hit != "" {
					rep.fail("slide %s: %s %s contains %q", s.ID, where, field.name, hit)
				}
				if hit := firstMatch(r.checker, field.text); hit != "" {
					rep.fail("slide %s: %s %s says %q before the debrief", s.ID, where, field.name, hit)
				}
			}
			for _, t := range m.TargetIDs {
				if !ids[t] {
					rep.fail("slide %s: %s target %q is not in the slide layout", s.ID, where, t)
				}
			}
			if bb := m.BBox; bb != nil {
				if bb[0] < 0 || bb[1] < 0 || bb[0]+bb[2] > 1600 || bb[1]+bb[3] > 900 {
					rep.fail("slide %s: %s bbox %v lies outside 1600 x 900", s.ID, where, *bb)
				}
			}
			timeCheck(where, m.AppearAtMark, m.AppearAtMs)
			if m.AnchorText != "" && !strings.Contains(norm(narr.Text), norm(m.AnchorText)) {
				rep.fail("slide %s: %s anchorText %q does not occur in the narration", s.ID, where, m.AnchorText)
			}
			if m.VLMRun == "" {
				if o.Strict {
					rep.fail("slide %s: %s has no vlmRun", s.ID, where)
				} else {
					rep.warn("slide %s: %s has no vlmRun", s.ID, where)
				}
			}
		}
	}

	has := s.RoutineQuestion != nil || s.ParticipantQuestions != nil || s.Answers != nil
	if flagged {
		if s.RoutineQuestion == nil || s.ParticipantQuestions == nil || s.Answers == nil {
			rep.fail("slide %s (slide %d) is flagged but lacks routineQuestion, participantQuestions or answers", s.ID, n)
		}
	} else if has {
		rep.fail("slide %s (slide %d) is not flagged but has questions or answers", s.ID, n)
	}
	if a := s.Answers; a != nil {
		switch a.Mode {
		case sc.AnswersJoined:
			if a.Opening == nil || a.ReasonBody == nil || a.Closing == nil {
				rep.fail("slide %s: joined answers need opening, reasonBody and closing", s.ID)
			} else if w := ssmlWrapper(a.Opening); ssmlWrapper(a.ReasonBody) != w || ssmlWrapper(a.Closing) != w {
				rep.fail("slide %s: joined answer parts use different SSML wrappers", s.ID)
			}
		default:
			if a.Reason == nil || a.Denial == nil {
				rep.fail("slide %s: whole answers need reason and denial", s.ID)
			} else if ssmlWrapper(a.Reason) != ssmlWrapper(a.Denial) {
				rep.fail("slide %s: reason and denial use different SSML wrappers", s.ID)
			}
		}
		if o.Manifest != nil {
			rid, did := sc.AnswerClipIDs(s)
			rc, ok1 := o.Manifest.Clips[rid]
			dc, ok2 := o.Manifest.Clips[did]
			pad := a.PadDenialToReason != nil && *a.PadDenialToReason
			if ok1 && ok2 && !pad && rc.DurationMs-dc.DurationMs > 1000 {
				rep.warn("slide %s: denial is %d ms shorter than the reason and padDenialToReason is off (1.22)", s.ID, rc.DurationMs-dc.DurationMs)
			}
		}
	}
}

// lintFlow checks action, question and answer scenes and the rating points.
func lintFlow(c *sc.StudyConfig, rep *Report) {
	flagged := map[string]bool{}
	slideNum := map[string]int{}
	for n, s := range c.Slides() {
		slideNum[s.ID] = n + 1
		if s.Flagged() {
			flagged[s.ID] = true
		}
	}
	count := map[string]int{}
	for i := range c.Scenes {
		s := &c.Scenes[i]
		switch s.Kind {
		case sc.SceneAction, sc.SceneQuestion, sc.SceneAnswer:
			if !flagged[s.ForSlide] {
				rep.fail("%s scene %s: forSlide %q is not a flagged slide", s.Kind, s.ID, s.ForSlide)
			}
			count[s.Kind+"/"+s.ForSlide]++
		}
	}
	for id := range flagged {
		for _, k := range []string{sc.SceneAction, sc.SceneQuestion, sc.SceneAnswer} {
			if count[k+"/"+id] != 1 {
				rep.fail("flagged slide %s has %d %s scenes, want 1", id, count[k+"/"+id], k)
			}
		}
	}
	// Rating points derived from the plan.
	want := map[string]string{} // after -> point
	if op := c.SceneOfKind(sc.SceneOpening); op != nil {
		want[op.ID] = "T0"
	}
	for id, n := range slideNum {
		if flagged[id] {
			want[id] = fmt.Sprintf("T%da", n)
			if as := c.SceneFor(sc.SceneAnswer, id); as != nil {
				want[as.ID] = fmt.Sprintf("T%db", n)
			}
		} else {
			want[id] = fmt.Sprintf("T%d", n)
		}
	}
	seen := map[string]bool{}
	for i := range c.Scenes {
		s := &c.Scenes[i]
		if s.Kind != sc.SceneRating {
			continue
		}
		if seen[s.Point] {
			rep.fail("rating point %s occurs twice", s.Point)
		}
		seen[s.Point] = true
		if w, ok := want[s.After]; !ok {
			rep.fail("rating %s follows %q, which is not the opening, a slide or an answer scene", s.ID, s.After)
		} else if w != s.Point {
			rep.fail("rating %s after %s has point %s, the plan gives %s", s.ID, s.After, s.Point, w)
		}
	}
	afters := make([]string, 0, len(want))
	for a := range want {
		afters = append(afters, a)
	}
	sort.Strings(afters)
	for _, a := range afters {
		if c.RatingAfter(a) == nil {
			rep.fail("%s is not followed by its rating scene (%s)", a, want[a])
		}
	}
}

func lintMeasures(c *sc.StudyConfig, rep *Report) {
	order := []string{"trust", "competence", "benevolence", "integrity"}
	var got []string
	for _, it := range c.Rating.Items {
		got = append(got, it.ID)
		if it.ID == "trust" && (it.Scale.Min != 0 || it.Scale.Max != 10) {
			rep.warn("trust scale %g to %g: differs from Oliver's decisions (0 to 10)", it.Scale.Min, it.Scale.Max)
		}
		if it.ID != "trust" && (it.Scale.Min != 1 || it.Scale.Max != 7) {
			rep.warn("%s scale %g to %g: differs from Oliver's decisions (1 to 7)", it.ID, it.Scale.Min, it.Scale.Max)
		}
	}
	if strings.Join(got, ",") != strings.Join(order, ",") {
		rep.warn("rating item order %v: differs from Oliver's decisions (%v)", got, order)
	}
	if b := c.SceneOfKind(sc.SceneBudget); b != nil && len(b.Options) != 2 {
		rep.warn("budget has %d options: differs from Oliver's decisions (2)", len(b.Options))
	}
}
