package analysis

import (
	"encoding/json"
	"math"
	"strings"
	"testing"
	"time"

	"study/api/internal/domain"
	"study/api/internal/testutil"
)

// Expected values below were computed independently (Python: Simpson
// integration of the t density, bisection for the quantile).

const tol = 1e-6

func near(t *testing.T, what string, got *float64, want float64) {
	t.Helper()
	if got == nil {
		t.Errorf("%s: nil, want %v", what, want)
		return
	}
	if math.Abs(*got-want) > tol {
		t.Errorf("%s: %v, want %v", what, *got, want)
	}
}

func TestTDistribution(t *testing.T) {
	for _, c := range []struct {
		t, df, p, tol float64
	}{
		{2.0, 10, 0.0734, 1e-4},
		{2.0, 10, 0.07338803477074551, 1e-9},
		{1.96, 1e6, 0.0500, 1e-3},
		{0, 5, 1, 1e-12},
		{-3.618136134933163, 3.7230769230769236, 0.025303342550622054, 1e-9},
		{12.706204736, 1, 0.05, 1e-6},
	} {
		if got := tTwoSided(c.t, c.df); math.Abs(got-c.p) > c.tol {
			t.Errorf("t=%v df=%v: p=%v, want %v", c.t, c.df, got, c.p)
		}
	}
	for _, c := range []struct{ df, q float64 }{{10, 2.228138851986313}, {1, 12.706204736}, {1e6, 1.959966}} {
		if got := tQuantile(0.975, c.df); math.Abs(got-c.q) > 1e-5 {
			t.Errorf("q(.975, %v) = %v, want %v", c.df, got, c.q)
		}
	}
}

func TestDescriptive(t *testing.T) {
	xs := []float64{4, 1, 3, 2}
	near(t, "mean", mean(xs), 2.5)
	near(t, "sd", sd(xs), math.Sqrt(5.0/3))
	near(t, "median", median(xs), 2.5)
	near(t, "p25", percentile(xs, 0.25), 1.75)
	near(t, "p75", percentile(xs, 0.75), 3.25)
	if mean(nil) != nil || sd([]float64{1}) != nil || median(nil) != nil {
		t.Fatal("undefined values must be nil")
	}
	r := twoProportionZ(6, 10, 3, 10)
	// pooled p = .45, se = sqrt(.45 * .55 * .2), z = .3 / se
	near(t, "z", r.z, 0.3/math.Sqrt(0.45*0.55*0.2))
	near(t, "z p", r.p, math.Erfc(*r.z/math.Sqrt2))
	if r := twoProportionZ(0, 0, 1, 2); r.p1 != nil || r.z != nil {
		t.Fatal("empty group must give nil")
	}
	near(t, "spearman", spearman([]float64{1, 2, 3, 4}, []float64{10, 20, 30, 25}), 0.8)
	near(t, "spearman ties", spearman([]float64{1, 1, 2}, []float64{1, 2, 3}), 0.8660254037844386)
	if spearman([]float64{1, 2}, []float64{2, 1}) != nil {
		t.Fatal("spearman needs 3 pairs")
	}
	if h := histogram([]float64{10, 70, 130}); len(h) != 3 || h[0].N != 1 || h[2].ToSec != 180 {
		t.Fatalf("histogram %+v", h)
	}
	if h := histogram([]float64{0, 3600}); len(h) > 30 || h[len(h)-1].N != 1 {
		t.Fatalf("histogram of an hour has %d bins", len(h))
	}
}

func TestReadings(t *testing.T) {
	p01, p20, neg, pos := fp(0.01), fp(0.2), fp(-1), fp(1)
	for _, c := range []struct {
		name string
		dir  int
		v, p *float64
		ns   []int
		want Reading
	}{
		{"supports", -1, neg, p01, []int{8, 8}, Supports},
		{"wrong sign", -1, pos, p01, []int{8, 8}, DoesNotSupport},
		{"not significant", +1, pos, p20, []int{8, 9}, DoesNotSupport},
		{"too few in one group", +1, pos, p01, []int{8, 7}, TooFew},
		{"too few beats n_a", 0, nil, nil, []int{3}, TooFew},
		{"no direction", 0, pos, p01, []int{8, 8}, NA},
		{"undefined p", +1, pos, nil, []int{8, 8}, NA},
	} {
		if got := read(c.dir, c.v, c.p, c.ns...); got != c.want {
			t.Errorf("%s: %s, want %s", c.name, got, c.want)
		}
	}
}

// ---------- synthetic sessions ----------

var testPoints = []string{"T0", "T1", "T2a", "T2b", "T3a", "T3b", "T4"}

type fake struct {
	id, group, denial, status string
	trust                     []int64 // by testPoints; -1 for a missing rating
	actions                   map[string]string
	noActionRecord            bool
	flags                     []string
	budget                    string
	bg                        map[string]any
}

func content(t *testing.T) *domain.Content {
	ct, err := domain.NewContent(testutil.FastConfig("test-fast", 4, 2, 3), nil, "/a/", "/s/")
	if err != nil {
		t.Fatal(err)
	}
	return ct
}

var t0 = time.Date(2026, 11, 1, 10, 0, 0, 0, time.UTC)

func build(fs []fake) ([]*domain.Session, map[string][]*domain.Rating, map[string][]*domain.AnswerDoc) {
	var ss []*domain.Session
	rts := map[string][]*domain.Rating{}
	ans := map[string][]*domain.AnswerDoc{}
	for i, f := range fs {
		opened := t0.Add(time.Duration(i) * time.Minute)
		consent := opened.Add(10 * time.Second)
		s := &domain.Session{ID: f.id, RunID: "run1", Group: f.group, DenialSlide: f.denial, Status: f.status,
			OpenedAt: opened, Flags: f.flags, IP: "203.0.113.9", UserAgent: "UA"}
		if s.Status == "" {
			s.Status = domain.StatusFinished
		}
		if f.group != "" {
			s.ConsentAt = &consent
		}
		if s.Status == domain.StatusFinished {
			fin := consent.Add(time.Duration(600+60*i) * time.Second)
			s.FinishedAt = &fin
		}
		ok := true
		s.AudioCheck = &domain.AudioCheckResult{Correct: &ok}
		ss = append(ss, s)
		for j, v := range f.trust {
			if v < 0 {
				continue
			}
			v := v
			m := int64(4)
			rts[f.id] = append(rts[f.id], &domain.Rating{Point: testPoints[j], Status: domain.RecordComplete,
				Values: map[string]*int64{"trust": &v, "competence": &m, "benevolence": &m, "integrity": &m}})
		}
		if !f.noActionRecord {
			doc := &domain.AnswerDoc{Kind: domain.AnswersActions, Actions: map[string]domain.ActionRecord{}}
			for _, sl := range []string{"s2", "s3"} {
				at := "reason"
				if f.denial == sl {
					at = "denial"
				}
				rec := domain.ActionRecord{AnswerType: at, QuestionBy: "chair"}
				if id, ok := f.actions[sl]; ok && id != "" {
					id := id
					rec.ActionID = &id
				} else {
					rec.TimedOut = true
				}
				doc.Actions[sl] = rec
			}
			ans[f.id] = append(ans[f.id], doc)
		}
		if f.budget != "" {
			ans[f.id] = append(ans[f.id], &domain.AnswerDoc{Kind: domain.AnswersBudget, Status: domain.RecordComplete,
				Values: map[string]any{"optionId": f.budget, "confidence": float64(7)}})
		}
		if f.bg != nil {
			ans[f.id] = append(ans[f.id], &domain.AnswerDoc{Kind: domain.AnswersBackground, Status: domain.RecordComplete, Values: f.bg})
		}
	}
	return ss, rts, ans
}

func smallCase() []fake {
	return []fake{
		{id: "c1", group: "checker", denial: "s3", trust: []int64{6, 6, 4, 6, 5, 3, 4},
			actions: map[string]string{"s2": "explain", "s3": "nothing"}, budget: "approve",
			bg: map[string]any{"field": "it", "cyber": float64(2), "country": "ca"}},
		{id: "c2", group: "checker", denial: "s2", trust: []int64{7, 8, 5, 6, 4, 5, 5},
			actions: map[string]string{"s2": "challenge", "s3": "explain"}, budget: "reject",
			bg: map[string]any{"field": map[string]any{"optionId": "other", "text": "Sailing"}, "cyber": float64(4), "country": "CH"}},
		// No action record: the answer type comes from the denial slide.
		{id: "c3", group: "checker", denial: "s3", trust: []int64{5, 5, 3, 5, 4, 6, 6}, noActionRecord: true,
			bg: map[string]any{"field": domain.PreferNotToSay, "cyber": float64(5), "country": "CA"}},
		{id: "n1", group: "neutral", denial: "s2", trust: []int64{6, 7, 7, 7, 6, 5, 6},
			actions: map[string]string{"s2": "hasQuestion", "s3": "noQuestion"}, budget: "approve"},
		{id: "n2", group: "neutral", denial: "s3", trust: []int64{5, 5, 6, 5, 6, 7, 7},
			actions: map[string]string{"s2": "noQuestion"}, budget: "approve"},
		{id: "n3", group: "neutral", denial: "s2", trust: []int64{7, 6, 6, 7, 7, 5, 6},
			actions: map[string]string{"s2": "noQuestion", "s3": "noQuestion"}},
	}
}

func TestComputeSmallCase(t *testing.T) {
	ct := content(t)
	ss, rts, ans := build(smallCase())
	run := &domain.Run{ID: "run1", Code: "abc", Name: "Pilot", ConfigVersion: "test-fast", Status: "open", CreatedAt: t0}
	a := Compute(run, ct, ss, rts, ans, Options{}, t0)

	if a.Counts.Analysed != 6 || a.Counts.ByGroup.Checker.Analysed != 3 || len(a.Participants) != 6 {
		t.Fatalf("counts %+v", a.Counts)
	}
	h := a.Hypotheses
	// H1 pooled: checker -1.5, -2.5, -1.5; neutral -0.5, 1, 0.
	near(t, "h1 checker mean", h.H1.Pooled.Checker.Mean, -11.0/6)
	near(t, "h1 neutral mean", h.H1.Pooled.Neutral.Mean, 1.0/6)
	near(t, "h1 diff", h.H1.Pooled.Diff, -2)
	near(t, "h1 t", h.H1.Pooled.WelchT, -3.618136134933163)
	near(t, "h1 df", h.H1.Pooled.DF, 3.7230769230769236)
	near(t, "h1 p", h.H1.Pooled.P, 0.025303342550622054)
	near(t, "h1 g", h.H1.Pooled.HedgesG, -2.363356626803188)
	if ci := h.H1.Pooled.DiffCI95; ci == nil || math.Abs(ci[0]+3.580823352893658) > tol || math.Abs(ci[1]+0.41917664710634206) > tol {
		t.Errorf("h1 diff ci %v", ci)
	}
	if h.H1.Pooled.Checker.N != 3 || h.H1.Pooled.Reading != TooFew {
		t.Errorf("h1 n %d reading %s", h.H1.Pooled.Checker.N, h.H1.Pooled.Reading)
	}
	near(t, "h1 one-sample t", h.H1.CheckerBelowZero.T, -5.5)
	near(t, "h1 one-sample p", h.H1.CheckerBelowZero.P, 0.0315040030418019)
	if ci := h.H1.CheckerBelowZero.CI95; ci == nil || math.Abs(ci[0]+3.2675509099164017) > tol {
		t.Errorf("h1 one-sample ci %v", ci)
	}
	near(t, "h1 per slide s2 checker", h.H1.PerSlide["s2"].Checker.Mean, (-2.0-3-2)/3)
	near(t, "level at flag s3 neutral", h.H1.LevelAtFlag["s3"].Neutral.Mean, (6.0+6+7)/3)

	// H2: reasoned deltas checker 2, 1, 2 (c3 from its denial slide); neutral -1, -1, -2.
	near(t, "h2 diff", h.H2.Pooled.Diff, 3)
	near(t, "h2 t", h.H2.Pooled.WelchT, 6.363961030678928)
	near(t, "h2 df", h.H2.Pooled.DF, 4)
	near(t, "h2 p", h.H2.Pooled.P, 0.003125589252467975)
	near(t, "h2 g", h.H2.Pooled.HedgesG, 4.156921938165306)
	near(t, "h2 one-sample t", h.H2.CheckerAboveZero.T, 5)
	near(t, "h2 one-sample p", h.H2.CheckerAboveZero.P, 0.03774955135062896)
	if n := h.H2.PerSlide["s2"].Checker.N; n != 2 {
		t.Errorf("h2 s2 checker n %d, want 2 (c1, c3)", n)
	}

	// H3 checker: reasoned minus denial 4, 0, 0.
	near(t, "h3 mean diff", h.H3.PerGroup.Checker.MeanDiff, 4.0/3)
	near(t, "h3 t", h.H3.PerGroup.Checker.T, 1)
	near(t, "h3 p", h.H3.PerGroup.Checker.P, 0.42264973081037993)
	near(t, "h3 df", h.H3.PerGroup.Checker.DF, 2)
	if h.H3.Both.N != 6 {
		t.Errorf("h3 both n %d", h.H3.Both.N)
	}
	// H4b denial: checker -2, 1, 2; neutral 0, 1, 1.
	near(t, "h4b denial t", h.H4b.AfterDenial.WelchT, -0.26726124191242434)
	near(t, "h4b denial p", h.H4b.AfterDenial.P, 0.8113131286522866)
	near(t, "h4b denial g", h.H4b.AfterDenial.HedgesG, -0.1745743121887939)

	// Baseline: 6, 7, 5 vs 6, 5, 7.
	near(t, "baseline diff", h.Baseline.Diff, 0)
	near(t, "baseline p", h.Baseline.P, 1)
	near(t, "slide1 checker", h.Slide1.Checker.Mean, 1.0/3)

	// H4a: checker c1 1/2, c2 2/2 (c3 has no record); neutral n1 1/2, n2 0/2
	// (s3 timed out), n3 0/2.
	if h.H4a.Checker.Slides != 4 || h.H4a.Checker.WithQuestion != 3 || h.H4a.Neutral.Slides != 6 || h.H4a.Neutral.WithQuestion != 1 {
		t.Errorf("h4a %+v %+v", h.H4a.Checker, h.H4a.Neutral)
	}
	near(t, "h4a share", h.H4a.Checker.Share, 0.75)
	near(t, "h4a per participant", h.H4a.Checker.PerParticipant.Mean, 0.75)
	near(t, "h4a z", h.H4a.SlidesTest.Z, (0.75-1.0/6)/math.Sqrt(0.4*0.6*(1.0/4+1.0/6)))

	// Actions: n2 timed out on s3.
	if a.Actions["s3"].Neutral["timeout"] != 1 || a.Actions["s2"].Checker["challenge"] != 1 || a.Actions["s2"].Checker["nothing"] != 0 {
		t.Errorf("actions %+v", a.Actions)
	}
	// Budget.
	if b := h.Budget.ByGroup; b.Checker.Approve != 1 || b.Checker.Reject != 1 || b.Checker.Missing != 1 || b.Neutral.Approve != 2 || b.Neutral.Missing != 1 {
		t.Errorf("budget %+v", b)
	}
	// Background and countries.
	if a.Countries[0].Code != "CA" || a.Countries[0].N != 2 || a.Countries[1].Code != "CH" {
		t.Errorf("countries %+v", a.Countries)
	}
	field := map[string]int{}
	for _, c := range a.Background["field"] {
		field[c.ID] = c.N
	}
	if field["it"] != 1 || field["other"] != 1 || field[domain.PreferNotToSay] != 1 || field["missing"] != 3 {
		t.Errorf("field %+v", a.Background["field"])
	}
	// Moderators: knowledge levels 1 to 5 of the "cyber" scale.
	if len(h.Moderators.CyberKnowledge) != 5 || h.Moderators.CyberKnowledge[1].H1.Checker.N != 1 {
		t.Errorf("moderators %+v", h.Moderators.CyberKnowledge)
	}
	near(t, "spearman knowledge h1 checker", h.Moderators.Spearman.CyberKnowledgeH1.Checker, 0)

	// Participant rows.
	var c3 Participant
	for _, p := range a.Participants {
		if p.SessionID == "c3" {
			c3 = p
		}
	}
	if c3.Ratings["T2a"] == nil || *c3.Ratings["T2a"].Trust != 3 || len(c3.Actions) != 0 || c3.Budget != nil || c3.Background["field"] != domain.PreferNotToSay {
		t.Errorf("c3 row %+v", c3)
	}
	if a.Config.RatingPoints[2].Label != "T2a, slide 2, after the slide" || a.Config.RatingPoints[3].Label != "T2b, slide 2, after the answer" {
		t.Errorf("labels %+v", a.Config.RatingPoints)
	}
	if !a.Config.Slides[1].Flagged || a.Config.Slides[1].Title != "Slide s2" || a.Config.Items[1].Label != "Competence" {
		t.Errorf("config %+v", a.Config)
	}
	if a.Counts.DurationsSec.N != 6 || a.Counts.DurationsSec.Median == nil {
		t.Errorf("durations %+v", a.Counts.DurationsSec)
	}
}

func TestMissingRatingDropsDelta(t *testing.T) {
	fs := smallCase()
	fs[0].trust[1] = -1 // c1 without T1: no T2a delta, pooled from T3a only
	ss, rts, ans := build(fs)
	a := Compute(&domain.Run{ID: "run1"}, content(t), ss, rts, ans, Options{}, t0)
	near(t, "h1 s2 n", fp(float64(a.Hypotheses.H1.PerSlide["s2"].Checker.N)), 2)
	near(t, "h1 pooled checker", a.Hypotheses.H1.Pooled.Checker.Mean, (-1.0-2.5-1.5)/3)
	if a.Participants[0].Ratings["T1"] != nil {
		t.Error("a missing rating is null")
	}
}

func TestExclusions(t *testing.T) {
	ct := content(t)
	base := smallCase()
	extra := []fake{
		{id: "forced", group: "checker", denial: "s2", trust: []int64{1, 1, 1, 1, 1, 1, 1}, flags: []string{domain.FlagForcedAllocation}},
		{id: "hp", group: "neutral", denial: "s2", trust: []int64{1, 1, 1, 1, 1, 1, 1}, flags: []string{domain.FlagHoneypot, domain.FlagAudioCheckWrong}},
		{id: "dup", group: "checker", denial: "s3", trust: []int64{6, 6, 6, 6, 6, 6, 6}, flags: []string{domain.FlagPossibleDuplicate}},
		{id: "ongoing", group: "neutral", denial: "s3", status: domain.StatusOngoing, trust: []int64{5, 5, -1, -1, -1, -1, -1}},
		{id: "opened", status: domain.StatusOpened, trust: nil},
	}
	ss, rts, ans := build(append(base, extra...))
	run := &domain.Run{ID: "run1"}
	for _, c := range []struct {
		name                         string
		opts                         Options
		analysed, forced, notIncl    int
		byFlag                       map[string]int
		participants, neutralAnalyse int
	}{
		{"defaults", Options{}, 7, 1, 1, map[string]int{"honeypot": 1, "audio_check_wrong": 1, "many_invalid_events": 0}, 7, 3},
		{"consented", Options{Include: IncludeConsented}, 8, 1, 0, map[string]int{"honeypot": 1, "audio_check_wrong": 1, "many_invalid_events": 0}, 8, 4},
		{"exclude none", Options{Exclude: []string{}}, 8, 1, 1, map[string]int{}, 8, 4},
		{"exclude duplicate", Options{Exclude: []string{"possible_duplicate", "forced_allocation"}}, 7, 1, 1, map[string]int{"possible_duplicate": 1}, 7, 4},
		{"group filter", Options{Group: GroupNeutral}, 7, 1, 1, map[string]int{"honeypot": 1, "audio_check_wrong": 1, "many_invalid_events": 0}, 3, 3},
	} {
		a := Compute(run, ct, ss, rts, ans, c.opts, t0)
		k := a.Counts
		if k.Analysed != c.analysed || k.Excluded.Forced != c.forced || k.Excluded.NotIncluded != c.notIncl || len(a.Participants) != c.participants ||
			k.ByGroup.Neutral.Analysed != c.neutralAnalyse {
			t.Errorf("%s: analysed %d forced %d notIncluded %d participants %d neutral %d", c.name, k.Analysed, k.Excluded.Forced,
				k.Excluded.NotIncluded, len(a.Participants), k.ByGroup.Neutral.Analysed)
		}
		if len(k.Excluded.ByFlag) != len(c.byFlag) {
			t.Errorf("%s: byFlag %v", c.name, k.Excluded.ByFlag)
		}
		for f, n := range c.byFlag {
			if k.Excluded.ByFlag[f] != n {
				t.Errorf("%s: byFlag %v", c.name, k.Excluded.ByFlag)
			}
		}
		if k.Total != 11 || k.Consented != 10 || k.ByStatus.Opened != 1 || k.ByStatus.Ongoing != 1 {
			t.Errorf("%s: totals %+v", c.name, k)
		}
		for _, p := range a.Participants {
			if p.SessionID == "forced" {
				t.Errorf("%s: forced_allocation analysed", c.name)
			}
		}
		if !strings.Contains(strings.Join(a.Filters.Exclude, ","), "forced_allocation") {
			t.Errorf("%s: filters %v", c.name, a.Filters.Exclude)
		}
		// The group filter never changes the group statistics.
		if a.Hypotheses.Baseline.Checker.N == 0 {
			t.Errorf("%s: checker statistics missing", c.name)
		}
	}
}

// With 8 per group the readings leave "too_few": a clear effect supports
// H1, the reverse does not.
func TestReadingsFollowThresholds(t *testing.T) {
	mk := func(checkerDrop int64) []fake {
		var fs []fake
		for i := 0; i < 8; i++ {
			j := int64(i % 3)
			fs = append(fs,
				fake{id: "c" + string(rune('a'+i)), group: "checker", denial: "s3", trust: []int64{6, 6, 6 - checkerDrop - j, 6, 6 - checkerDrop - j, 6, 6},
					actions: map[string]string{"s2": "explain", "s3": "challenge"}},
				fake{id: "n" + string(rune('a'+i)), group: "neutral", denial: "s3", trust: []int64{6, 6, 6 + j, 6, 6 - j, 6, 6},
					actions: map[string]string{"s2": "noQuestion", "s3": "noQuestion"}})
		}
		return fs
	}
	ct := content(t)
	ss, rts, ans := build(mk(3))
	a := Compute(&domain.Run{ID: "run1"}, ct, ss, rts, ans, Options{}, t0)
	h := a.Hypotheses
	if h.H1.Pooled.Reading != Supports || h.H1.CheckerBelowZero.Reading != Supports || h.H4a.SlidesTest.Reading != Supports {
		t.Errorf("readings %s %s %s", h.H1.Pooled.Reading, h.H1.CheckerBelowZero.Reading, h.H4a.SlidesTest.Reading)
	}
	if h.Baseline.Reading != NA {
		t.Errorf("baseline reading %s", h.Baseline.Reading)
	}
	ss, rts, ans = build(mk(-3))
	a = Compute(&domain.Run{ID: "run1"}, ct, ss, rts, ans, Options{}, t0)
	if r := a.Hypotheses.H1.Pooled.Reading; r != DoesNotSupport {
		t.Errorf("reverse effect reading %s", r)
	}
}

// The JSON form has no null slices or maps where the contract has arrays
// or records, and no NaN.
func TestJSONShape(t *testing.T) {
	a := Compute(&domain.Run{ID: "run1"}, content(t), nil, nil, nil, Options{}, t0)
	b, err := json.Marshal(a)
	if err != nil {
		t.Fatal(err)
	}
	var m map[string]any
	_ = json.Unmarshal(b, &m)
	for _, k := range []string{"run", "generatedAt", "filters", "config", "counts", "participants", "ratingStats", "hypotheses",
		"actions", "background", "countries", "quality"} {
		if _, ok := m[k]; !ok {
			t.Errorf("missing %s", k)
		}
	}
	s := string(b)
	for _, bad := range []string{`"participants":null`, `"countries":null`, `"histogram":null`, `"byFlag":null`, `"exclude":null`,
		`"cyberKnowledge":null`, `"aiUse":null`, `"flaggedSlides":null`} {
		if strings.Contains(s, bad) {
			t.Errorf("JSON has %s", bad)
		}
	}
	if !strings.Contains(s, `"mean":null`) {
		t.Error("an empty summary has a null mean")
	}
}
