// Package analysis computes the run dashboard document of
// GET /api/admin/runs/{id}/analysis ("06 - Analysis Dashboard Plan.md",
// sections 2, 5 and 6). The types mirror the TypeScript contract of section 6
// field by field; a React client is built against the same names.
package analysis

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"time"

	"study/api/internal/domain"
	sc "study/api/pkg/studyconfig"
)

// Include values.
const (
	IncludeFinished  = "finished"
	IncludeConsented = "consented"
)

// Group filter values.
const (
	GroupBoth    = "both"
	GroupChecker = string(sc.GroupChecker)
	GroupNeutral = string(sc.GroupNeutral)
)

// DefaultExclude is the exclusion list when none is given (section 5).
var DefaultExclude = []string{domain.FlagHoneypot, domain.FlagAudioCheckWrong, domain.FlagManyInvalidEvents}

// MinPerGroup is the "too few participants" threshold of the readings
// (section 2; a placeholder for the analysis plan, section 8.3).
const MinPerGroup = 8

// Alpha is the reading threshold for p.
const Alpha = 0.05

// Options are the query parameters. Exclude nil means DefaultExclude; an
// empty, non-nil slice excludes nothing but forced_allocation.
type Options struct {
	Include string
	Exclude []string
	Group   string
}

// Reading is the one-line verdict of a test.
type Reading string

const (
	Supports       Reading = "supports"
	Opposite       Reading = "opposite" // p < .05 with the sign against the hypothesis
	DoesNotSupport Reading = "does_not_support"
	TooFew         Reading = "too_few"
	NA             Reading = "n_a"
)

// ByGroup is Record<"checker" | "neutral", T>.
type ByGroup[T any] struct {
	Checker T `json:"checker"`
	Neutral T `json:"neutral"`
}

func (b *ByGroup[T]) get(g string) *T {
	if g == GroupChecker {
		return &b.Checker
	}
	return &b.Neutral
}

// ---------- the document ----------

type Analysis struct {
	Run          RunInfo                                `json:"run"`
	GeneratedAt  time.Time                              `json:"generatedAt"`
	Filters      Filters                                `json:"filters"`
	Config       ConfigInfo                             `json:"config"`
	Counts       Counts                                 `json:"counts"`
	Participants []Participant                          `json:"participants"`
	RatingStats  map[string]map[string]ByGroup[Summary] `json:"ratingStats"`
	Hypotheses   Hypotheses                             `json:"hypotheses"`
	Actions      map[string]ByGroup[map[string]int]     `json:"actions"`
	Background   map[string][]BackgroundCount           `json:"background"`
	Countries    []CountryCount                         `json:"countries"`
	Quality      Quality                                `json:"quality"`
}

type RunInfo struct {
	ID            string    `json:"id"`
	Code          string    `json:"code"`
	Name          string    `json:"name"`
	Source        string    `json:"source"`
	ConfigVersion string    `json:"configVersion"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"createdAt"`
	TestOverrides bool      `json:"testOverrides"`
}

type Filters struct {
	Include string   `json:"include"`
	Exclude []string `json:"exclude"`
	Group   string   `json:"group"`
}

type ConfigInfo struct {
	Groups        []string         `json:"groups"`
	Slides        []SlideInfo      `json:"slides"`
	FlaggedSlides []string         `json:"flaggedSlides"`
	RatingPoints  []RatingPoint    `json:"ratingPoints"`
	Items         []ItemInfo       `json:"items"`
	Actions       ByGroup[[]Label] `json:"actions"`
	Background    []BackgroundInfo `json:"background"`
	BudgetOptions []Label          `json:"budgetOptions"`
}

type SlideInfo struct {
	ID      string `json:"id"`
	N       int    `json:"n"`
	Title   string `json:"title"`
	Flagged bool   `json:"flagged"`
}

type RatingPoint struct {
	Point   string `json:"point"`
	SlideID string `json:"slideId"`
	Phase   string `json:"phase"`
	Label   string `json:"label"`
}

type ItemInfo struct {
	ID    string  `json:"id"`
	Label string  `json:"label"`
	Min   float64 `json:"min"`
	Max   float64 `json:"max"`
}

type Label struct {
	ID    string `json:"id"`
	Label string `json:"label"`
}

type BackgroundInfo struct {
	ID      string     `json:"id"`
	Label   string     `json:"label"`
	Kind    string     `json:"kind"`
	Options []Label    `json:"options,omitempty"`
	Scale   *ScaleInfo `json:"scale,omitempty"`
}

type ScaleInfo struct {
	Min         float64           `json:"min"`
	Max         float64           `json:"max"`
	PointLabels map[string]string `json:"pointLabels,omitempty"`
}

type Counts struct {
	ByStatus     StatusCounts        `json:"byStatus"`
	Total        int                 `json:"total"`
	Consented    int                 `json:"consented"`
	Analysed     int                 `json:"analysed"`
	Excluded     Excluded            `json:"excluded"`
	ByGroup      ByGroup[GroupCount] `json:"byGroup"`
	DurationsSec Durations           `json:"durationsSec"`
}

type StatusCounts struct {
	Opened    int `json:"opened"`
	Started   int `json:"started"`
	Ongoing   int `json:"ongoing"`
	Finished  int `json:"finished"`
	Abandoned int `json:"abandoned"`
	Blocked   int `json:"blocked"`
}

type Excluded struct {
	Forced      int            `json:"forced"`
	ByFlag      map[string]int `json:"byFlag"`
	NotIncluded int            `json:"notIncluded"`
}

type GroupCount struct {
	Consented int `json:"consented"`
	Finished  int `json:"finished"`
	Analysed  int `json:"analysed"`
}

type Durations struct {
	N         int        `json:"n"`
	Median    *float64   `json:"median"`
	P25       *float64   `json:"p25"`
	P75       *float64   `json:"p75"`
	Histogram []HistoBin `json:"histogram"`
}

type HistoBin struct {
	FromSec float64 `json:"fromSec"`
	ToSec   float64 `json:"toSec"`
	N       int     `json:"n"`
}

type Participant struct {
	SessionID   string                 `json:"sessionId"`
	Group       string                 `json:"group"`
	DenialSlide string                 `json:"denialSlide"`
	Status      string                 `json:"status"`
	FinishedAt  *time.Time             `json:"finishedAt"`
	DurationSec *float64               `json:"durationSec"`
	Ratings     map[string]*RatingCell `json:"ratings"`
	Actions     map[string]ActionCell  `json:"actions"`
	Budget      *BudgetCell            `json:"budget"`
	Background  map[string]any         `json:"background"`
	Flags       []string               `json:"flags"`
}

type RatingCell struct {
	Trust       *float64 `json:"trust"`
	Competence  *float64 `json:"competence"`
	Benevolence *float64 `json:"benevolence"`
	Integrity   *float64 `json:"integrity"`
	Status      string   `json:"status"`
}

type ActionCell struct {
	ActionID   *string `json:"actionId"`
	TimedOut   bool    `json:"timedOut"`
	QuestionBy string  `json:"questionBy"`
	AnswerType string  `json:"answerType"`
}

type BudgetCell struct {
	OptionID   *string  `json:"optionId"`
	Confidence *float64 `json:"confidence"`
	Status     string   `json:"status"`
}

type Summary struct {
	N      int         `json:"n"`
	Mean   *float64    `json:"mean"`
	SD     *float64    `json:"sd"`
	CI95   *[2]float64 `json:"ci95"`
	Median *float64    `json:"median"`
}

type OneSample struct {
	N       int         `json:"n"`
	Mean    *float64    `json:"mean"`
	CI95    *[2]float64 `json:"ci95"`
	T       *float64    `json:"t"`
	DF      *float64    `json:"df"`
	P       *float64    `json:"p"`
	Reading Reading     `json:"reading"`
}

type Comparison struct {
	Checker  Summary     `json:"checker"`
	Neutral  Summary     `json:"neutral"`
	Diff     *float64    `json:"diff"`
	DiffCI95 *[2]float64 `json:"diffCi95"`
	WelchT   *float64    `json:"welchT"`
	DF       *float64    `json:"df"`
	P        *float64    `json:"p"`
	HedgesG  *float64    `json:"hedgesG"`
	Reading  Reading     `json:"reading"`
}

type Paired struct {
	N        int         `json:"n"`
	Reasoned Summary     `json:"reasoned"`
	Denial   Summary     `json:"denial"`
	MeanDiff *float64    `json:"meanDiff"`
	CI95     *[2]float64 `json:"ci95"`
	T        *float64    `json:"t"`
	DF       *float64    `json:"df"`
	P        *float64    `json:"p"`
	Reading  Reading     `json:"reading"`
}

type Share struct {
	Slides         int      `json:"slides"`
	WithQuestion   int      `json:"withQuestion"`
	Share          *float64 `json:"share"`
	PerParticipant Summary  `json:"perParticipant"`
}

type TwoProportion struct {
	P1      *float64 `json:"p1"`
	P2      *float64 `json:"p2"`
	Diff    *float64 `json:"diff"`
	Z       *float64 `json:"z"`
	P       *float64 `json:"p"`
	Reading Reading  `json:"reading"`
}

type Hypotheses struct {
	Baseline   Comparison `json:"baseline"`
	Slide1     Comparison `json:"slide1"`
	Overall    Comparison `json:"overall"`
	H1         H1         `json:"h1"`
	H2         H2         `json:"h2"`
	H3         H3         `json:"h3"`
	H4a        H4a        `json:"h4a"`
	H4b        H4b        `json:"h4b"`
	Budget     BudgetStat `json:"budget"`
	Moderators Moderators `json:"moderators"`
}

type H1 struct {
	Pooled           Comparison            `json:"pooled"`
	CheckerBelowZero OneSample             `json:"checkerBelowZero"`
	PerSlide         map[string]Comparison `json:"perSlide"`
	LevelAtFlag      map[string]Comparison `json:"levelAtFlag"`
}

type H2 struct {
	Pooled           Comparison            `json:"pooled"`
	CheckerAboveZero OneSample             `json:"checkerAboveZero"`
	PerSlide         map[string]Comparison `json:"perSlide"`
	Level            Comparison            `json:"level"`
	LevelPerSlide    map[string]Comparison `json:"levelPerSlide"`
	AboveBaseline    ByGroup[OneSample]    `json:"aboveBaseline"`
	ExploratoryNet   Comparison            `json:"exploratoryNet"`
}

type H3 struct {
	PerGroup ByGroup[Paired] `json:"perGroup"`
	Both     Paired          `json:"both"`
}

type H4a struct {
	Checker          Share         `json:"checker"`
	Neutral          Share         `json:"neutral"`
	SlidesTest       TwoProportion `json:"slidesTest"`
	ParticipantsTest Comparison    `json:"participantsTest"`
}

type H4b struct {
	AfterReason Comparison `json:"afterReason"`
	AfterDenial Comparison `json:"afterDenial"`
}

type BudgetStat struct {
	ByGroup     ByGroup[BudgetGroup] `json:"byGroup"`
	ApproveTest TwoProportion        `json:"approveTest"`
	// ByLastTrustTercile splits the votes by the tercile of the participant's
	// last available trust rating.
	ByLastTrustTercile []TercileCount `json:"byLastTrustTercile"`
}

type TercileCount struct {
	Tercile int     `json:"tercile"`
	From    float64 `json:"from"`
	To      float64 `json:"to"`
	Approve int     `json:"approve"`
	Reject  int     `json:"reject"`
	Missing int     `json:"missing"`
}

type BudgetGroup struct {
	Approve    int     `json:"approve"`
	Reject     int     `json:"reject"`
	Missing    int     `json:"missing"`
	Confidence Summary `json:"confidence"`
}

type Moderators struct {
	CyberKnowledge []KnowledgeLevel `json:"cyberKnowledge"`
	AIUse          []AIUseLevel     `json:"aiUse"`
	Spearman       Spearman         `json:"spearman"`
}

type KnowledgeLevel struct {
	Level int              `json:"level"`
	H1    ByGroup[Summary] `json:"h1"`
	H2    ByGroup[Summary] `json:"h2"`
}

type AIUseLevel struct {
	OptionID string           `json:"optionId"`
	H1       ByGroup[Summary] `json:"h1"`
	H2       ByGroup[Summary] `json:"h2"`
}

type Spearman struct {
	CyberKnowledgeH1 ByGroup[*float64] `json:"cyberKnowledgeH1"`
	CyberKnowledgeH2 ByGroup[*float64] `json:"cyberKnowledgeH2"`
	AIUseH1          ByGroup[*float64] `json:"aiUseH1"`
	AIUseH2          ByGroup[*float64] `json:"aiUseH2"`
}

type BackgroundCount struct {
	ID      string       `json:"id"`
	Label   string       `json:"label"`
	N       int          `json:"n"`
	ByGroup ByGroup[int] `json:"byGroup"`
}

type CountryCount struct {
	Code    string       `json:"code"`
	N       int          `json:"n"`
	ByGroup ByGroup[int] `json:"byGroup"`
}

type Quality struct {
	Flags          map[string]int `json:"flags"`
	AudioCheck     AudioCheck     `json:"audioCheck"`
	RatingTimeouts map[string]int `json:"ratingTimeouts"`
	Reloads        Reloads        `json:"reloads"`
}

type AudioCheck struct {
	Correct int `json:"correct"`
	Wrong   int `json:"wrong"`
	Missing int `json:"missing"`
}

type Reloads struct {
	Sessions int `json:"sessions"`
	Total    int `json:"total"`
}

// ---------- per-participant data ----------

// subject is one analysed session with its records.
type subject struct {
	s        *domain.Session
	ratings  map[string]*domain.Rating
	actions  map[string]domain.ActionRecord
	budget   *domain.AnswerDoc
	bg       *domain.AnswerDoc
	h1Pooled *float64
	h2Pooled *float64
}

// trust returns the trust value at a point, or nil.
func (p *subject) trust(point string) *float64 {
	r := p.ratings[point]
	if r == nil {
		return nil
	}
	return intp(r.Values["trust"])
}

// delta is trust(b) - trust(a) when both are present.
func (p *subject) delta(a, b string) *float64 {
	x, y := p.trust(a), p.trust(b)
	if x == nil || y == nil {
		return nil
	}
	return fp(*y - *x)
}

// answerType is "reason" or "denial" for a flagged slide: from the action
// record, or by comparing the slide with the session's denial slide when
// the record is missing.
func (p *subject) answerType(slide string) string {
	if r, ok := p.actions[slide]; ok && r.AnswerType != "" {
		return r.AnswerType
	}
	if p.s.DenialSlide == slide {
		return "denial"
	}
	return "reason"
}

func intp(v *int64) *float64 {
	if v == nil {
		return nil
	}
	return fp(float64(*v))
}

// num reads a JSON number from an answers value: float64 after a JSON round
// trip (memory store), int64 from Firestore.
func num(v any) *float64 {
	switch x := v.(type) {
	case float64:
		return fp(x)
	case int64:
		return fp(float64(x))
	case int:
		return fp(float64(x))
	case int32:
		return fp(float64(x))
	}
	return nil
}

// bgValue flattens a background value to string | number | nil. An "Other"
// option stored as { optionId, text } becomes its option id.
// DECISION: the free text is not sent to the dashboard (the contract has no
// place for it; it stays in the export).
func bgValue(v any) any {
	switch x := v.(type) {
	case nil:
		return nil
	case string:
		return x
	case map[string]any:
		if id, ok := x["optionId"].(string); ok {
			return id
		}
		return nil
	}
	if n := num(v); n != nil {
		return *n
	}
	return nil
}

// ---------- Compute ----------

// Compute builds the analysis of one run. sessions are the run's sessions;
// ratings and answers are keyed by session id.
func Compute(run *domain.Run, ct *domain.Content, sessions []*domain.Session, ratings map[string][]*domain.Rating,
	answers map[string][]*domain.AnswerDoc, opts Options, now time.Time) *Analysis {
	opts = normalise(opts)
	cfg := ct.Config
	a := &Analysis{
		Run: RunInfo{ID: run.ID, Code: run.Code, Name: run.Name, Source: run.Source, ConfigVersion: run.ConfigVersion,
			Status: run.Status, CreatedAt: run.CreatedAt, TestOverrides: run.TestOverrides},
		GeneratedAt:  now.UTC(),
		Filters:      Filters{Include: opts.Include, Exclude: append(append([]string{}, opts.Exclude...), domain.FlagForcedAllocation), Group: opts.Group},
		Config:       configInfo(ct),
		Participants: []Participant{},
		RatingStats:  map[string]map[string]ByGroup[Summary]{},
		Actions:      map[string]ByGroup[map[string]int]{},
		Background:   map[string][]BackgroundCount{},
		Countries:    []CountryCount{},
	}
	points := ct.Plan.RatingPoints()
	flagged := ct.FlaggedSlides()

	sorted := append([]*domain.Session(nil), sessions...)
	sort.SliceStable(sorted, func(i, j int) bool { return sorted[i].OpenedAt.Before(sorted[j].OpenedAt) })

	// Counts and the analysed set. Order: forced_allocation first, then the
	// include rule, then the exclusion flags.
	// DECISION: a session is counted once in excluded.forced or
	// excluded.notIncluded (whichever rule removes it first); byFlag counts,
	// among the sessions left after those two rules, each excluded flag a
	// session carries, so one session with two flags appears under both.
	c := &a.Counts
	c.Excluded.ByFlag = map[string]int{}
	for _, f := range opts.Exclude {
		c.Excluded.ByFlag[f] = 0
	}
	c.DurationsSec.Histogram = []HistoBin{}
	var included, analysed []*subject
	for _, s := range sorted {
		c.Total++
		switch s.Status {
		case domain.StatusOpened:
			c.ByStatus.Opened++
		case domain.StatusStarted:
			c.ByStatus.Started++
		case domain.StatusOngoing:
			c.ByStatus.Ongoing++
		case domain.StatusFinished:
			c.ByStatus.Finished++
		case domain.StatusAbandoned:
			c.ByStatus.Abandoned++
		case domain.StatusBlocked:
			c.ByStatus.Blocked++
		}
		if s.Group != GroupChecker && s.Group != GroupNeutral {
			continue // not consented
		}
		// DECISION: consented and byGroup.consented/finished are raw counts
		// that include forced_allocation sessions (as the runs list does);
		// excluded.forced says how many of them there are.
		c.Consented++
		gc := c.ByGroup.get(s.Group)
		gc.Consented++
		if s.Status == domain.StatusFinished {
			gc.Finished++
		}
		if s.HasFlag(domain.FlagForcedAllocation) {
			c.Excluded.Forced++
			continue
		}
		if opts.Include == IncludeFinished && s.Status != domain.StatusFinished {
			c.Excluded.NotIncluded++
			continue
		}
		p := newSubject(s, ratings[s.ID], answers[s.ID])
		included = append(included, p)
		out := false
		for _, f := range opts.Exclude {
			if s.HasFlag(f) {
				c.Excluded.ByFlag[f]++
				out = true
			}
		}
		if out {
			continue
		}
		analysed = append(analysed, p)
		c.Analysed++
		gc.Analysed++
	}

	// Durations of analysed finished sessions (consent to finish, as the export).
	var durs []float64
	for _, p := range analysed {
		if d := duration(p.s); d != nil {
			durs = append(durs, *d)
		}
	}
	c.DurationsSec.N = len(durs)
	c.DurationsSec.Median = median(durs)
	c.DurationsSec.P25 = percentile(durs, 0.25)
	c.DurationsSec.P75 = percentile(durs, 0.75)
	c.DurationsSec.Histogram = histogram(durs)

	// Participants (the group filter applies here only).
	bgItems := backgroundItems(cfg)
	for _, p := range analysed {
		if opts.Group != GroupBoth && p.s.Group != opts.Group {
			continue
		}
		a.Participants = append(a.Participants, participant(p, points, bgItems))
	}

	// Rating statistics per item, point and group.
	for _, it := range cfg.Rating.Items {
		byPoint := map[string]ByGroup[Summary]{}
		for _, pt := range points {
			var bg ByGroup[[]float64]
			for _, p := range analysed {
				if r := p.ratings[pt.Point]; r != nil {
					if v := intp(r.Values[it.ID]); v != nil {
						*bg.get(p.s.Group) = append(*bg.get(p.s.Group), *v)
					}
				}
			}
			byPoint[pt.Point] = ByGroup[Summary]{Checker: summary(bg.Checker), Neutral: summary(bg.Neutral)}
		}
		a.RatingStats[it.ID] = byPoint
	}

	a.Hypotheses = hypotheses(ct, points, flagged, analysed)
	a.Actions = actionCounts(cfg, flagged, analysed)
	a.Background, a.Countries = backgroundCounts(cfg, bgItems, analysed)
	a.Quality = quality(points, included)
	return a
}

func normalise(o Options) Options {
	if o.Include != IncludeConsented {
		o.Include = IncludeFinished
	}
	if o.Group != GroupChecker && o.Group != GroupNeutral {
		o.Group = GroupBoth
	}
	src := o.Exclude
	if src == nil {
		src = DefaultExclude
	}
	o.Exclude = []string{}
	seen := map[string]bool{domain.FlagForcedAllocation: true}
	for _, f := range src {
		if f = strings.TrimSpace(f); f != "" && !seen[f] {
			seen[f] = true
			o.Exclude = append(o.Exclude, f)
		}
	}
	return o
}

func newSubject(s *domain.Session, rts []*domain.Rating, ans []*domain.AnswerDoc) *subject {
	p := &subject{s: s, ratings: map[string]*domain.Rating{}, actions: map[string]domain.ActionRecord{}}
	for _, r := range rts {
		p.ratings[r.Point] = r
	}
	for _, d := range ans {
		switch d.Kind {
		case domain.AnswersActions:
			for k, v := range d.Actions {
				p.actions[k] = v
			}
		case domain.AnswersBudget:
			p.budget = d
		case domain.AnswersBackground:
			p.bg = d
		}
	}
	return p
}

func duration(s *domain.Session) *float64 {
	if s.ConsentAt == nil || s.FinishedAt == nil {
		return nil
	}
	return fp(s.FinishedAt.Sub(*s.ConsentAt).Seconds())
}

// histogram bins durations by 60 s from 0 up to the maximum; the width grows
// by 60 s steps until at most 30 bins are needed.
func histogram(xs []float64) []HistoBin {
	out := []HistoBin{}
	if len(xs) == 0 {
		return out
	}
	maxV := 0.0
	for _, x := range xs {
		maxV = math.Max(maxV, x)
	}
	w := 60.0
	for int(maxV/w)+1 > 30 {
		w += 60
	}
	nb := int(maxV/w) + 1
	for i := 0; i < nb; i++ {
		out = append(out, HistoBin{FromSec: float64(i) * w, ToSec: float64(i+1) * w})
	}
	for _, x := range xs {
		i := int(math.Max(x, 0) / w)
		if i >= nb {
			i = nb - 1
		}
		out[i].N++
	}
	return out
}

// ---------- config ----------

func slideTitle(s *sc.Scene) string {
	for _, b := range s.Layout {
		if b.Kind == sc.BlockTitle {
			return b.Text.PlainText()
		}
	}
	return s.ID
}

func itemLabel(it sc.RatingItem) string {
	if it.Label != "" {
		return it.Label
	}
	// DECISION: METI items have no label in the config; the dashboard gets
	// the id with a capital first letter ("Competence").
	if it.ID == "" {
		return ""
	}
	return strings.ToUpper(it.ID[:1]) + it.ID[1:]
}

func slideNumbers(cfg *sc.StudyConfig) map[string]int {
	out := map[string]int{}
	for i, s := range cfg.Slides() {
		out[s.ID] = i + 1
	}
	return out
}

// pointLabel is e.g. "T2a, slide 2, after the slide".
func pointLabel(pt domain.PlanStep, nums map[string]int) string {
	switch pt.Phase {
	case domain.PhaseBaseline:
		return pt.Point + ", baseline after the opening"
	case domain.PhaseAfterAnswer:
		return fmt.Sprintf("%s, slide %d, after the answer", pt.Point, nums[pt.SlideID])
	}
	return fmt.Sprintf("%s, slide %d, after the slide", pt.Point, nums[pt.SlideID])
}

func labels(cs []sc.Choice) []Label {
	out := []Label{}
	for _, c := range cs {
		out = append(out, Label{ID: c.ID, Label: c.Label})
	}
	return out
}

func backgroundItems(cfg *sc.StudyConfig) []sc.BackgroundItem {
	var out []sc.BackgroundItem
	if bg := cfg.SceneOfKind(sc.SceneBackground); bg != nil {
		for _, it := range bg.Items {
			if it.Kind != sc.ItemHoneypot {
				out = append(out, it)
			}
		}
	}
	return out
}

func configInfo(ct *domain.Content) ConfigInfo {
	cfg := ct.Config
	nums := slideNumbers(cfg)
	ci := ConfigInfo{Groups: []string{GroupChecker, GroupNeutral}, Slides: []SlideInfo{},
		FlaggedSlides: append([]string{}, ct.FlaggedSlides()...), RatingPoints: []RatingPoint{}, Items: []ItemInfo{},
		Actions:    ByGroup[[]Label]{Checker: labels(cfg.Actions.Checker.Options), Neutral: labels(cfg.Actions.Neutral.Options)},
		Background: []BackgroundInfo{}, BudgetOptions: []Label{}}
	for _, s := range cfg.Slides() {
		ci.Slides = append(ci.Slides, SlideInfo{ID: s.ID, N: nums[s.ID], Title: slideTitle(s), Flagged: s.Flagged()})
	}
	for _, pt := range ct.Plan.RatingPoints() {
		ci.RatingPoints = append(ci.RatingPoints, RatingPoint{Point: pt.Point, SlideID: pt.SlideID, Phase: pt.Phase, Label: pointLabel(pt, nums)})
	}
	for _, it := range cfg.Rating.Items {
		ci.Items = append(ci.Items, ItemInfo{ID: it.ID, Label: itemLabel(it), Min: it.Scale.Min, Max: it.Scale.Max})
	}
	for _, it := range backgroundItems(cfg) {
		bi := BackgroundInfo{ID: it.ID, Label: it.Label, Kind: it.Kind}
		if len(it.Options) > 0 {
			bi.Options = labels(it.Options)
		}
		if it.Scale != nil {
			bi.Scale = &ScaleInfo{Min: it.Scale.Min, Max: it.Scale.Max, PointLabels: it.Scale.PointLabels}
		}
		ci.Background = append(ci.Background, bi)
	}
	if b := cfg.SceneOfKind(sc.SceneBudget); b != nil {
		ci.BudgetOptions = labels(b.Options)
	}
	return ci
}

// ---------- participants ----------

func participant(p *subject, points []domain.PlanStep, bgItems []sc.BackgroundItem) Participant {
	s := p.s
	flags := append([]string{}, s.Flags...)
	out := Participant{SessionID: s.ID, Group: s.Group, DenialSlide: s.DenialSlide, Status: s.Status, FinishedAt: s.FinishedAt,
		DurationSec: duration(s), Ratings: map[string]*RatingCell{}, Actions: map[string]ActionCell{}, Flags: flags}
	for _, pt := range points {
		r := p.ratings[pt.Point]
		if r == nil {
			out.Ratings[pt.Point] = nil
			continue
		}
		out.Ratings[pt.Point] = &RatingCell{Trust: intp(r.Values["trust"]), Competence: intp(r.Values["competence"]),
			Benevolence: intp(r.Values["benevolence"]), Integrity: intp(r.Values["integrity"]), Status: r.Status}
	}
	for slide, r := range p.actions {
		out.Actions[slide] = ActionCell{ActionID: r.ActionID, TimedOut: r.TimedOut, QuestionBy: r.QuestionBy, AnswerType: p.answerType(slide)}
	}
	if b := p.budget; b != nil {
		bc := &BudgetCell{Confidence: num(b.Values["confidence"]), Status: b.Status}
		if id, ok := b.Values["optionId"].(string); ok {
			bc.OptionID = &id
		}
		out.Budget = bc
	}
	if p.bg != nil {
		out.Background = map[string]any{}
		for _, it := range bgItems {
			out.Background[it.ID] = bgValue(p.bg.Values[it.ID])
		}
	}
	return out
}

// ---------- readings and test builders ----------

// read is the reading of a test with direction dir (+1, -1, or 0 for none)
// on the sign of v. ns are the sizes of the groups involved.
func read(dir int, v, p *float64, ns ...int) Reading {
	for _, n := range ns {
		if n < MinPerGroup {
			return TooFew
		}
	}
	if dir == 0 || v == nil || p == nil {
		return NA
	}
	if *p >= Alpha || *v == 0 {
		return DoesNotSupport
	}
	if dir > 0 && *v > 0 || dir < 0 && *v < 0 {
		return Supports
	}
	return Opposite
}

func summary(xs []float64) Summary {
	s := Summary{N: len(xs), Mean: mean(xs), SD: sd(xs), Median: median(xs)}
	if s.SD != nil {
		s.CI95 = ci95(*s.Mean, *s.SD/math.Sqrt(float64(len(xs))), float64(len(xs)-1))
	}
	return s
}

// descriptive is a comparison shown without a reading (per-slide breakdowns:
// about twenty uncorrected readings would turn up a chance "supports").
func descriptive(x ByGroup[[]float64]) Comparison {
	c := compare(x, 0)
	c.Reading = NA
	return c
}

func compare(x ByGroup[[]float64], dir int) Comparison {
	w := welchT(x.Checker, x.Neutral)
	return Comparison{Checker: summary(x.Checker), Neutral: summary(x.Neutral), Diff: w.diff, DiffCI95: w.ci,
		WelchT: w.t, DF: w.df, P: w.p, HedgesG: w.g, Reading: read(dir, w.diff, w.p, len(x.Checker), len(x.Neutral))}
}

func oneSample(xs []float64, dir int) OneSample {
	r := oneSampleT(xs)
	return OneSample{N: r.n, Mean: r.mean, CI95: r.ci, T: r.t, DF: r.df, P: r.p, Reading: read(dir, r.mean, r.p, r.n)}
}

func paired(reasoned, denial []float64, dir int, ns ...int) Paired {
	r := pairedT(reasoned, denial)
	return Paired{N: r.n, Reasoned: summary(reasoned), Denial: summary(denial), MeanDiff: r.mean, CI95: r.ci, T: r.t,
		DF: r.df, P: r.p, Reading: read(dir, r.mean, r.p, ns...)}
}

func twoProp(x1, n1, x2, n2, dir int, ns ...int) TwoProportion {
	r := twoProportionZ(x1, n1, x2, n2)
	return TwoProportion{P1: r.p1, P2: r.p2, Diff: r.diff, Z: r.z, P: r.p, Reading: read(dir, r.diff, r.p, ns...)}
}

// ---------- hypotheses ----------

// previousPoint returns, for each rating point, the point before it in plan
// order (section 2: T{n-1}b if it exists, else T{n-1}a or T{n-1}, else T0).
// DECISION: when the participant's trust at that previous point is missing
// (a timeout), the delta is missing too; there is no fallback to an earlier
// point, so every delta spans exactly one step of the plan.
func previousPoint(points []domain.PlanStep) map[string]string {
	out := map[string]string{}
	for i := 1; i < len(points); i++ {
		out[points[i].Point] = points[i-1].Point
	}
	return out
}

// flagPoints returns the a and b point ids of a flagged slide.
func flagPoints(points []domain.PlanStep, slide string) (a, b string) {
	for _, pt := range points {
		if pt.SlideID != slide {
			continue
		}
		switch pt.Phase {
		case domain.PhaseAfterSlide:
			a = pt.Point
		case domain.PhaseAfterAnswer:
			b = pt.Point
		}
	}
	return a, b
}

func meanOf(xs []float64) *float64 { return mean(xs) }

// questionIDs returns the action ids that count as asking a question
// (H4a). Checker: "explain" and "challenge" when the config has them, else
// every option but the first ("say nothing"); a timeout never counts.
// Neutral: the option whose id contains "hasQuestion", else the first one.
func questionIDs(cfg *sc.StudyConfig) ByGroup[map[string]bool] {
	out := ByGroup[map[string]bool]{Checker: map[string]bool{}, Neutral: map[string]bool{}}
	for _, o := range cfg.Actions.Checker.Options {
		if o.ID == "explain" || o.ID == "challenge" {
			out.Checker[o.ID] = true
		}
	}
	if len(out.Checker) == 0 {
		for i, o := range cfg.Actions.Checker.Options {
			if i > 0 {
				out.Checker[o.ID] = true
			}
		}
	}
	for _, o := range cfg.Actions.Neutral.Options {
		if strings.Contains(strings.ToLower(o.ID), "hasquestion") {
			out.Neutral[o.ID] = true
		}
	}
	if len(out.Neutral) == 0 && len(cfg.Actions.Neutral.Options) > 0 {
		out.Neutral[cfg.Actions.Neutral.Options[0].ID] = true
	}
	return out
}

// DECISION: "too_few" uses the number of participants who contribute to the
// statistic (those with the delta, the answer or the action record), not all
// analysed participants: a test on 3 deltas is too few even when 20 sessions
// are analysed.
func hypotheses(ct *domain.Content, points []domain.PlanStep, flagged []string, ps []*subject) Hypotheses {
	cfg := ct.Config
	prev := previousPoint(points)
	var h Hypotheses
	h.H1.PerSlide, h.H1.LevelAtFlag, h.H2.PerSlide = map[string]Comparison{}, map[string]Comparison{}, map[string]Comparison{}
	h.H2.LevelPerSlide = map[string]Comparison{}

	// Baseline and slide 1.
	// DECISION: "slide 1" is the first rating point of the first slide (T1,
	// or T1a if slide 1 were flagged) minus T0.
	var t0, s1, overall ByGroup[[]float64]
	lastPoint := ""
	if len(points) > 0 {
		lastPoint = points[len(points)-1].Point
	}
	firstSlidePoint := ""
	if slides := cfg.Slides(); len(slides) > 0 {
		for _, pt := range points {
			if pt.SlideID == slides[0].ID {
				firstSlidePoint = pt.Point
				break
			}
		}
	}
	for _, p := range ps {
		if v := p.trust("T0"); v != nil {
			*t0.get(p.s.Group) = append(*t0.get(p.s.Group), *v)
		}
		if d := p.delta("T0", firstSlidePoint); d != nil {
			*s1.get(p.s.Group) = append(*s1.get(p.s.Group), *d)
		}
		if d := p.delta("T0", lastPoint); d != nil {
			*overall.get(p.s.Group) = append(*overall.get(p.s.Group), *d)
		}
	}
	h.Baseline = compare(t0, 0)
	h.Slide1 = compare(s1, 0)
	// The research question: T_last - T0, no stated direction.
	h.Overall = compare(overall, 0)

	// Per-participant deltas.
	var h1Pool, h2Pool, h2Net, h2Level, h2AboveT0, denialD ByGroup[[]float64]
	type pair struct{ reasoned, denial float64 }
	var h3 ByGroup[[]pair]
	perSlideH1 := map[string]*ByGroup[[]float64]{}
	levelH1 := map[string]*ByGroup[[]float64]{}
	perSlideH2 := map[string]*ByGroup[[]float64]{}
	levelH2 := map[string]*ByGroup[[]float64]{}
	for _, sl := range flagged {
		perSlideH1[sl], levelH1[sl], perSlideH2[sl] = &ByGroup[[]float64]{}, &ByGroup[[]float64]{}, &ByGroup[[]float64]{}
		levelH2[sl] = &ByGroup[[]float64]{}
	}
	for _, p := range ps {
		g := p.s.Group
		var flagDs, reasonDs, netDs, levels, aboveT0 []float64
		var denial *float64
		for _, sl := range flagged {
			a, b := flagPoints(points, sl)
			if d := p.delta(prev[a], a); d != nil {
				flagDs = append(flagDs, *d)
				*perSlideH1[sl].get(g) = append(*perSlideH1[sl].get(g), *d)
			}
			if v := p.trust(a); v != nil {
				*levelH1[sl].get(g) = append(*levelH1[sl].get(g), *v)
			}
			d := p.delta(a, b)
			if p.answerType(sl) == "denial" {
				if d != nil {
					denial = d
				}
				continue
			}
			if d != nil {
				reasonDs = append(reasonDs, *d)
				*perSlideH2[sl].get(g) = append(*perSlideH2[sl].get(g), *d)
			}
			if n := p.delta(prev[a], b); n != nil {
				netDs = append(netDs, *n)
			}
			if v := p.trust(b); v != nil {
				levels = append(levels, *v)
				*levelH2[sl].get(g) = append(*levelH2[sl].get(g), *v)
			}
			if d := p.delta("T0", b); d != nil {
				aboveT0 = append(aboveT0, *d)
			}
		}
		p.h1Pooled, p.h2Pooled = meanOf(flagDs), meanOf(reasonDs)
		if p.h1Pooled != nil {
			*h1Pool.get(g) = append(*h1Pool.get(g), *p.h1Pooled)
		}
		if p.h2Pooled != nil {
			*h2Pool.get(g) = append(*h2Pool.get(g), *p.h2Pooled)
		}
		if n := meanOf(netDs); n != nil {
			*h2Net.get(g) = append(*h2Net.get(g), *n)
		}
		if v := meanOf(levels); v != nil {
			*h2Level.get(g) = append(*h2Level.get(g), *v)
		}
		if v := meanOf(aboveT0); v != nil {
			*h2AboveT0.get(g) = append(*h2AboveT0.get(g), *v)
		}
		if denial != nil {
			*denialD.get(g) = append(*denialD.get(g), *denial)
		}
		if denial != nil && p.h2Pooled != nil {
			*h3.get(g) = append(*h3.get(g), pair{*p.h2Pooled, *denial})
		}
	}

	h.H1.Pooled = compare(h1Pool, -1)
	h.H1.CheckerBelowZero = oneSample(h1Pool.Checker, -1)
	for _, sl := range flagged {
		// Per-slide breakdowns carry no reading (plan 06, section 2).
		h.H1.PerSlide[sl] = descriptive(*perSlideH1[sl])
		h.H1.LevelAtFlag[sl] = descriptive(*levelH1[sl])
		// The H2 per-slide numbers cover only the participants for whom the
		// slide was reasoned; the denial slide's delta is in H3 and H4b.
		h.H2.PerSlide[sl] = descriptive(*perSlideH2[sl])
		h.H2.LevelPerSlide[sl] = descriptive(*levelH2[sl])
	}
	// h2.pooled compares changes (the same data as h4b.afterReason, shown as
	// H4b1); the H2(b) reading comes from the level T{n}b, since a comparison
	// of changes would favour the group that fell further.
	h.H2.Pooled = compare(h2Pool, +1)
	h.H2.CheckerAboveZero = oneSample(h2Pool.Checker, +1)
	h.H2.Level = compare(h2Level, +1)
	h.H2.AboveBaseline = ByGroup[OneSample]{Checker: oneSample(h2AboveT0.Checker, +1), Neutral: oneSample(h2AboveT0.Neutral, +1)}
	// DECISION: the exploratory net change (T{n}b minus the rating before the
	// flag, pooled over reasoned slides) has no direction: reading "n_a"
	// unless too few.
	h.H2.ExploratoryNet = compare(h2Net, 0)

	split := func(xs []pair) (r, d []float64) {
		for _, x := range xs {
			r, d = append(r, x.reasoned), append(d, x.denial)
		}
		return r, d
	}
	cr, cd := split(h3.Checker)
	nr, nd := split(h3.Neutral)
	h.H3.PerGroup.Checker = paired(cr, cd, +1, len(cr))
	h.H3.PerGroup.Neutral = paired(nr, nd, +1, len(nr))
	// DECISION: "both" pools the two groups' pairs and reads "too_few" when
	// either group has fewer than 8 pairs.
	h.H3.Both = paired(append(append([]float64{}, cr...), nr...), append(append([]float64{}, cd...), nd...), +1, len(cr), len(nr))

	h.H4b.AfterReason = compare(h2Pool, +1)
	h.H4b.AfterDenial = compare(denialD, -1)

	// H4a.
	qs := questionIDs(cfg)
	var shares ByGroup[[]float64]
	var slidesN, withQ ByGroup[int]
	for _, p := range ps {
		g := p.s.Group
		n, q := 0, 0
		for _, sl := range flagged {
			r, ok := p.actions[sl]
			if !ok {
				continue
			}
			n++
			if !r.TimedOut && r.ActionID != nil && (*qs.get(g))[*r.ActionID] {
				q++
			}
		}
		if n == 0 {
			continue
		}
		*slidesN.get(g) += n
		*withQ.get(g) += q
		*shares.get(g) = append(*shares.get(g), float64(q)/float64(n))
	}
	share := func(g string) Share {
		s := Share{Slides: *slidesN.get(g), WithQuestion: *withQ.get(g), PerParticipant: summary(*shares.get(g))}
		if s.Slides > 0 {
			s.Share = fp(float64(s.WithQuestion) / float64(s.Slides))
		}
		return s
	}
	h.H4a.Checker, h.H4a.Neutral = share(GroupChecker), share(GroupNeutral)
	// The slides test is descriptive: a participant's slides are not
	// independent, so the reading comes from participantsTest only.
	h.H4a.SlidesTest = twoProp(withQ.Checker, slidesN.Checker, withQ.Neutral, slidesN.Neutral, 0)
	h.H4a.SlidesTest.Reading = NA
	h.H4a.ParticipantsTest = compare(shares, +1)

	h.Budget = budget(cfg, points, ps)
	h.Moderators = moderators(cfg, ps)
	return h
}

// approveID is the budget option counted as approving: "approve" when the
// config has it, else the first option.
func approveID(cfg *sc.StudyConfig) string {
	b := cfg.SceneOfKind(sc.SceneBudget)
	if b == nil || len(b.Options) == 0 {
		return "approve"
	}
	for _, o := range b.Options {
		if o.ID == "approve" {
			return o.ID
		}
	}
	return b.Options[0].ID
}

// DECISION: the approve test compares approve / (approve + reject) between
// the groups (missing answers left out) and has no direction; reading "n_a"
// unless too few answers.
func budget(cfg *sc.StudyConfig, points []domain.PlanStep, ps []*subject) BudgetStat {
	approve := approveID(cfg)
	var out BudgetStat
	var conf ByGroup[[]float64]
	for _, p := range ps {
		bg := out.ByGroup.get(p.s.Group)
		var id string
		if p.budget != nil {
			id, _ = p.budget.Values["optionId"].(string)
			if v := num(p.budget.Values["confidence"]); v != nil {
				*conf.get(p.s.Group) = append(*conf.get(p.s.Group), *v)
			}
		}
		switch {
		case id == "":
			bg.Missing++
		case id == approve:
			bg.Approve++
		default:
			bg.Reject++
		}
	}
	out.ByGroup.Checker.Confidence = summary(conf.Checker)
	out.ByGroup.Neutral.Confidence = summary(conf.Neutral)
	c, n := out.ByGroup.Checker, out.ByGroup.Neutral
	out.ApproveTest = twoProp(c.Approve, c.Approve+c.Reject, n.Approve, n.Approve+n.Reject, 0, c.Approve+c.Reject, n.Approve+n.Reject)
	out.ByLastTrustTercile = terciles(cfg, points, ps, approve)
	return out
}

// lastTrust is the participant's last available trust rating in plan order.
func (p *subject) lastTrust(points []domain.PlanStep) *float64 {
	for i := len(points) - 1; i >= 0; i-- {
		if v := p.trust(points[i].Point); v != nil {
			return v
		}
	}
	return nil
}

// terciles splits the budget votes by the tercile of the last trust rating.
// The cut points are the 1/3 and 2/3 percentiles of the analysed
// participants' last ratings; a rating equal to a cut point goes to the lower
// tercile. from/to are the bounds: the trust scale's minimum, the two cut
// points and the scale's maximum.
// DECISION: participants without any trust rating are left out; with none
// at all the list is empty.
func terciles(cfg *sc.StudyConfig, points []domain.PlanStep, ps []*subject, approve string) []TercileCount {
	out := []TercileCount{}
	var last []float64
	for _, p := range ps {
		if v := p.lastTrust(points); v != nil {
			last = append(last, *v)
		}
	}
	if len(last) == 0 {
		return out
	}
	lo, hi := 0.0, 10.0
	for _, it := range cfg.Rating.Items {
		if it.ID == "trust" {
			lo, hi = it.Scale.Min, it.Scale.Max
		}
	}
	q1, q2 := *percentile(last, 1.0/3), *percentile(last, 2.0/3)
	out = append(out, TercileCount{Tercile: 1, From: lo, To: q1}, TercileCount{Tercile: 2, From: q1, To: q2},
		TercileCount{Tercile: 3, From: q2, To: hi})
	for _, p := range ps {
		v := p.lastTrust(points)
		if v == nil {
			continue
		}
		t := &out[2]
		switch {
		case *v <= q1:
			t = &out[0]
		case *v <= q2:
			t = &out[1]
		}
		var id string
		if p.budget != nil {
			id, _ = p.budget.Values["optionId"].(string)
		}
		switch {
		case id == "":
			t.Missing++
		case id == approve:
			t.Approve++
		default:
			t.Reject++
		}
	}
	return out
}

// moderatorItems finds the knowledge scale and the AI-use item.
// DECISION: the knowledge item is the background item "cyberKnowledge", else
// the first scale item; AI use is the single-choice item "aiUse", ranked by
// its option order for Spearman. A config without them gets empty lists.
func moderatorItems(cfg *sc.StudyConfig) (knowledge, aiUse *sc.BackgroundItem) {
	items := backgroundItems(cfg)
	for i := range items {
		it := &items[i]
		if it.ID == "cyberKnowledge" && it.Kind == sc.ItemScale && it.Scale != nil {
			knowledge = it
		}
		if it.ID == "aiUse" && it.Kind == sc.ItemSingle {
			aiUse = it
		}
	}
	if knowledge == nil {
		for i := range items {
			if items[i].Kind == sc.ItemScale && items[i].Scale != nil {
				knowledge = &items[i]
				break
			}
		}
	}
	return knowledge, aiUse
}

func moderators(cfg *sc.StudyConfig, ps []*subject) Moderators {
	out := Moderators{CyberKnowledge: []KnowledgeLevel{}, AIUse: []AIUseLevel{}}
	knowledge, aiUse := moderatorItems(cfg)
	value := func(p *subject, id string) any {
		if p.bg == nil {
			return nil
		}
		return bgValue(p.bg.Values[id])
	}
	type xy struct{ x, y ByGroup[[]float64] }
	var kH1, kH2, aH1, aH2 xy
	add := func(c *xy, g string, x float64, y *float64) {
		if y != nil {
			*c.x.get(g) = append(*c.x.get(g), x)
			*c.y.get(g) = append(*c.y.get(g), *y)
		}
	}
	rho := func(c xy) ByGroup[*float64] {
		return ByGroup[*float64]{Checker: spearman(c.x.Checker, c.y.Checker), Neutral: spearman(c.x.Neutral, c.y.Neutral)}
	}
	if knowledge != nil {
		for lv := int(knowledge.Scale.Min); lv <= int(knowledge.Scale.Max); lv++ {
			var h1, h2 ByGroup[[]float64]
			for _, p := range ps {
				v, ok := value(p, knowledge.ID).(float64)
				if !ok || int(v) != lv {
					continue
				}
				g := p.s.Group
				if p.h1Pooled != nil {
					*h1.get(g) = append(*h1.get(g), *p.h1Pooled)
				}
				if p.h2Pooled != nil {
					*h2.get(g) = append(*h2.get(g), *p.h2Pooled)
				}
				add(&kH1, g, v, p.h1Pooled)
				add(&kH2, g, v, p.h2Pooled)
			}
			out.CyberKnowledge = append(out.CyberKnowledge, KnowledgeLevel{Level: lv,
				H1: ByGroup[Summary]{Checker: summary(h1.Checker), Neutral: summary(h1.Neutral)},
				H2: ByGroup[Summary]{Checker: summary(h2.Checker), Neutral: summary(h2.Neutral)}})
		}
	}
	if aiUse != nil {
		for rank, o := range aiUse.Options {
			var h1, h2 ByGroup[[]float64]
			for _, p := range ps {
				if v, ok := value(p, aiUse.ID).(string); !ok || v != o.ID {
					continue
				}
				g := p.s.Group
				if p.h1Pooled != nil {
					*h1.get(g) = append(*h1.get(g), *p.h1Pooled)
				}
				if p.h2Pooled != nil {
					*h2.get(g) = append(*h2.get(g), *p.h2Pooled)
				}
				add(&aH1, g, float64(rank), p.h1Pooled)
				add(&aH2, g, float64(rank), p.h2Pooled)
			}
			out.AIUse = append(out.AIUse, AIUseLevel{OptionID: o.ID,
				H1: ByGroup[Summary]{Checker: summary(h1.Checker), Neutral: summary(h1.Neutral)},
				H2: ByGroup[Summary]{Checker: summary(h2.Checker), Neutral: summary(h2.Neutral)}})
		}
	}
	out.Spearman = Spearman{CyberKnowledgeH1: rho(kH1), CyberKnowledgeH2: rho(kH2), AIUseH1: rho(aH1), AIUseH2: rho(aH2)}
	return out
}

// ---------- actions, background, quality ----------

// actionCounts counts per flagged slide and group each option id and
// "timeout"; every option of the config is present, with 0 when unused.
func actionCounts(cfg *sc.StudyConfig, flagged []string, ps []*subject) map[string]ByGroup[map[string]int] {
	out := map[string]ByGroup[map[string]int]{}
	zero := func(opts []sc.Choice) map[string]int {
		m := map[string]int{"timeout": 0}
		for _, o := range opts {
			m[o.ID] = 0
		}
		return m
	}
	for _, sl := range flagged {
		out[sl] = ByGroup[map[string]int]{Checker: zero(cfg.Actions.Checker.Options), Neutral: zero(cfg.Actions.Neutral.Options)}
	}
	for _, p := range ps {
		for _, sl := range flagged {
			r, ok := p.actions[sl]
			if !ok {
				continue
			}
			bg := out[sl]
			m := *bg.get(p.s.Group)
			if r.TimedOut || r.ActionID == nil {
				m["timeout"]++
			} else {
				m[*r.ActionID]++
			}
		}
	}
	return out
}

const (
	keyPreferNot = domain.PreferNotToSay
	keyMissing   = "missing"
)

// backgroundCounts counts every background answer per item (options in
// config order, scale points from min to max, countries by count), then
// "preferNotToSay" and "missing". countries lists the country codes only.
func backgroundCounts(cfg *sc.StudyConfig, items []sc.BackgroundItem, ps []*subject) (map[string][]BackgroundCount, []CountryCount) {
	out := map[string][]BackgroundCount{}
	countries := []CountryCount{}
	pnsLabel := "Prefer not to say"
	if bg := cfg.SceneOfKind(sc.SceneBackground); bg != nil && bg.PreferNotToSayLabel != "" {
		pnsLabel = bg.PreferNotToSayLabel
	}
	for _, it := range items {
		var list []BackgroundCount
		index := map[string]int{}
		addEntry := func(id, label string) {
			index[id] = len(list)
			list = append(list, BackgroundCount{ID: id, Label: label})
		}
		switch it.Kind {
		case sc.ItemSingle:
			for _, o := range it.Options {
				addEntry(o.ID, o.Label)
			}
		case sc.ItemScale:
			if it.Scale != nil {
				for v := int(it.Scale.Min); v <= int(it.Scale.Max); v++ {
					id := strconv.Itoa(v)
					label := it.Scale.PointLabels[id]
					switch {
					case label != "":
					case float64(v) == it.Scale.Min && it.Scale.MinLabel != "":
						label = it.Scale.MinLabel
					case float64(v) == it.Scale.Max && it.Scale.MaxLabel != "":
						label = it.Scale.MaxLabel
					default:
						label = id
					}
					addEntry(id, label)
				}
			}
		}
		var extra []BackgroundCount // values not in the config (countries, or a changed config)
		extraIdx := map[string]int{}
		var pns, missing BackgroundCount
		pns = BackgroundCount{ID: keyPreferNot, Label: pnsLabel}
		missing = BackgroundCount{ID: keyMissing, Label: "No answer"}
		bump := func(c *BackgroundCount, g string) {
			c.N++
			*c.ByGroup.get(g)++
		}
		for _, p := range ps {
			g := p.s.Group
			var v any
			if p.bg != nil {
				v = bgValue(p.bg.Values[it.ID])
			}
			var key string
			switch x := v.(type) {
			case nil:
				bump(&missing, g)
				continue
			case string:
				key = x
				if it.Kind == sc.ItemCountry && x != keyPreferNot {
					key = strings.ToUpper(strings.TrimSpace(x))
				}
			case float64:
				key = strconv.FormatFloat(x, 'f', -1, 64)
			}
			if key == keyPreferNot {
				bump(&pns, g)
				continue
			}
			if key == "" {
				bump(&missing, g)
				continue
			}
			if i, ok := index[key]; ok {
				bump(&list[i], g)
				continue
			}
			i, ok := extraIdx[key]
			if !ok {
				i = len(extra)
				extraIdx[key] = i
				extra = append(extra, BackgroundCount{ID: key, Label: key})
			}
			bump(&extra[i], g)
		}
		sort.SliceStable(extra, func(i, j int) bool {
			if extra[i].N != extra[j].N {
				return extra[i].N > extra[j].N
			}
			return extra[i].ID < extra[j].ID
		})
		if it.Kind == sc.ItemCountry {
			for _, e := range extra {
				countries = append(countries, CountryCount{Code: e.ID, N: e.N, ByGroup: e.ByGroup})
			}
		}
		list = append(list, extra...)
		out[it.ID] = append(list, pns, missing)
	}
	return out, countries
}

// DECISION: quality is computed over the included sessions before the flag
// exclusions (forced_allocation and the include rule applied), so the flags
// table shows how many sessions the exclusion toggles remove.
func quality(points []domain.PlanStep, ps []*subject) Quality {
	q := Quality{Flags: map[string]int{}, RatingTimeouts: map[string]int{}}
	for _, pt := range points {
		q.RatingTimeouts[pt.Point] = 0
	}
	for _, p := range ps {
		for _, f := range p.s.Flags {
			q.Flags[f]++
		}
		switch {
		case p.s.AudioCheck == nil || p.s.AudioCheck.Correct == nil:
			q.AudioCheck.Missing++
		case *p.s.AudioCheck.Correct:
			q.AudioCheck.Correct++
		default:
			q.AudioCheck.Wrong++
		}
		for _, pt := range points {
			if r := p.ratings[pt.Point]; r != nil && r.Status == domain.RecordTimeout {
				q.RatingTimeouts[pt.Point]++
			}
		}
		if p.s.ReloadCount > 0 {
			q.Reloads.Sessions++
			q.Reloads.Total += p.s.ReloadCount
		}
	}
	return q
}
