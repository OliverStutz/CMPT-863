package main

import (
	"fmt"
	"io"
	"math"
	"math/rand/v2"
	"sort"
	"strings"
)

// Persona is one emulated participant of the persona mode (-persona N): a
// trust baseline, a METI profile around it, a response to flags, reasons and
// denials for each group (the one used is learned during the session), an
// action pattern, and a country. Everything is drawn once at start from the
// persona's own random stream, which also makes the later decisions, so a
// seed gives the same participants and the same answers (think times use the
// session's stream).
type Persona struct {
	rng *rand.Rand

	Baseline float64    // trust 5 to 9, mostly 6 to 8
	Offsets  [3]float64 // competence, benevolence, integrity offsets from the trust line
	Checker  Response
	Neutral  Response
	Style    string  // checker style: quiet, curious, challenger
	AskProb  float64 // neutral: probability of hasQuestion, 0.2 to 0.6
	Pick     int     // fallback denial: the answer step (0, 1, 2) taken as the denial
	Country  string  // ISO 3166 alpha-2

	// Learned and recorded during the session.
	Group      string // "" until the first flagged slide or action step
	trust      float64
	LastTrust  *int // last submitted trust rating
	answers    int  // answer steps seen
	denial     map[string]bool
	Denial     string // slide inferred as the denial
	DenialHow  string // transcript, elimination, fallback
	Actions    []string
	Budget     string
	Sureness   *int
	Background map[string]any
	Timeouts   int
}

// Response is the trust change after a flagged slide, a reasoned answer and
// the denial.
type Response struct {
	Flag, Reason, Denial float64
}

var personaStyles = []struct {
	name string
	w    float64
	main string
}{{"quiet", 0.30, "nothing"}, {"curious", 0.40, "explain"}, {"challenger", 0.30, "challenge"}}

// Countries by weight; "others" share the rest.
var personaCountries = []weighted{{"CA", 35}, {"DE", 20}, {"US", 15}, {"CH", 10}, {"GB", 8}, {"IN", 5},
	{"FR", 1.2}, {"NL", 1.2}, {"AT", 1}, {"AU", 1}, {"SE", 0.9}, {"IE", 0.9}, {"SG", 0.8}}

// Background weights by option id, matched on the part after the last "."
// (field.business matches "business"). Options of the scene pack without an
// entry weigh 1; the labels are never used.
var personaBackground = map[string]map[string]float64{
	"field":     {"business": 6, "it": 6, "cyber": 5, "finance": 5, "law": 1.5, "engineering": 1.5, "government": 1.2},
	"age":       {"a18": 1, "a25": 5, "a35": 6, "a45": 5, "a55": 1.5, "a65": 0.4},
	"jobLevel":  {"student": 0.8, "employee": 6, "leadership": 5, "cLevel": 1},
	"education": {"secondary": 0.3, "college": 1.2, "bachelor": 6, "master": 5, "doctorate": 0.8, "other": 0.2},
	"aiUse":     {"never": 0.4, "lessMonthly": 0.6, "monthly": 1, "weekly": 5, "severalWeekly": 3, "daily": 5},
}

const (
	personaNoise       = 0.35 // sd of the noise on every rating value
	personaTimeoutProb = 0.03 // ratings left to time out
	personaPNTS        = 0.05 // "preferNotToSay" per background item
	metiSlope          = 0.6  // METI points per trust point
)

type weighted struct {
	id string
	w  float64
}

func pickWeighted(rng *rand.Rand, xs []weighted) string {
	total := 0.0
	for _, x := range xs {
		total += x.w
	}
	r := rng.Float64() * total
	for _, x := range xs {
		if r < x.w {
			return x.id
		}
		r -= x.w
	}
	return xs[len(xs)-1].id
}

func uniform(rng *rand.Rand, lo, hi float64) float64 { return lo + rng.Float64()*(hi-lo) }

func clampF(v, lo, hi float64) float64 { return math.Max(lo, math.Min(hi, v)) }

func clampI(v, lo, hi int) int { return max(lo, min(hi, v)) }

// NewPersona draws persona i of a seed. The stream is separate from the
// session's (seed, i) stream.
func NewPersona(seed uint64, i int) *Persona {
	rng := rand.New(rand.NewPCG(seed, uint64(i)|1<<40))
	p := &Persona{rng: rng, denial: map[string]bool{}}
	p.Baseline = clampF(7+rng.NormFloat64()*0.75, 5, 9)
	for k := range p.Offsets {
		p.Offsets[k] = clampF(rng.NormFloat64()*0.5, -1, 1)
	}
	p.Checker = Response{Flag: -uniform(rng, 1, 2.5), Reason: uniform(rng, 0.5, 1.8), Denial: uniform(rng, -1, 0.3)}
	p.Neutral = Response{Flag: -uniform(rng, 0, 0.8), Reason: uniform(rng, 0.2, 0.8), Denial: uniform(rng, -0.5, 0.3)}
	var sw []weighted
	for _, s := range personaStyles {
		sw = append(sw, weighted{s.name, s.w})
	}
	p.Style = pickWeighted(rng, sw)
	p.AskProb = uniform(rng, 0.2, 0.6)
	p.Pick = rng.IntN(3)
	p.Country = pickWeighted(rng, personaCountries)
	p.trust = p.Baseline
	return p
}

// MetiBase is the persona's METI profile at its baseline (before noise).
func (p *Persona) MetiBase() [3]float64 {
	var m [3]float64
	for k := range m {
		m[k] = clampF(1+metiSlope*p.Baseline+p.Offsets[k], 1, 7)
	}
	return m
}

func (p *Persona) response() *Response {
	switch p.Group {
	case "checker":
		return &p.Checker
	case "neutral":
		return &p.Neutral
	}
	return nil
}

// learnGroup records the group once it is known: an overlay on a flagged
// slide (checker), a flagged slide without one (neutral), or the action
// options (nothing/explain/challenge or hasQuestion/noQuestion).
func (p *Persona) learnGroup(g string) {
	if g != "" {
		p.Group = g
	}
}

func groupOfOptions(opts []Choice) string {
	for _, o := range opts {
		switch o.ID {
		case "nothing", "explain", "challenge":
			return "checker"
		case "hasQuestion", "noQuestion":
			return "neutral"
		}
	}
	return ""
}

// Denial answers in the study state confidence instead of reasons; a short
// answer is the opening and closing of a joined answer without its body.
var denialMarkers = []string{"confident", "stand behind", "stand completely behind"}

func looksLikeDenial(transcript string) bool {
	t := strings.ToLower(transcript)
	if len(strings.Fields(t)) < 25 {
		return true
	}
	for _, m := range denialMarkers {
		if strings.Contains(t, m) {
			return true
		}
	}
	return false
}

// observeAnswer classifies one answer step as reason or denial. Each session
// hears exactly one denial among the flagged slides (nFlagged), so after one
// denial the rest are reasons, and the last answer is the denial if none was
// found before. Without a transcript the persona's fallback pick decides.
func (p *Persona) observeAnswer(slide, transcript string, nFlagged int) {
	k := p.answers
	p.answers++
	if p.Denial != "" {
		return
	}
	how := ""
	switch {
	case strings.TrimSpace(transcript) == "":
		if k == p.Pick {
			how = "fallback"
		}
	case looksLikeDenial(transcript):
		how = "transcript"
	case k == nFlagged-1:
		how = "elimination"
	}
	if how != "" {
		p.Denial, p.DenialHow = slide, how
		p.denial[slide] = true
	}
}

// ratingContext tells what a rating follows.
const (
	afterOther = iota
	afterFlagged
	afterAnswer
)

// applyRating moves the latent trust for the event the rating follows.
func (p *Persona) applyRating(after int, slide string) {
	r := p.response()
	if r == nil {
		return
	}
	switch after {
	case afterFlagged:
		p.trust += r.Flag
	case afterAnswer:
		if p.denial[slide] {
			p.trust += r.Denial
		} else {
			p.trust += r.Reason
		}
	}
	p.trust = clampF(p.trust, 0, 10)
}

// ratingValues returns the item values of one rating, or nil when this
// rating is left to time out.
func (p *Persona) ratingValues(items []RatingItem) map[string]any {
	if p.rng.Float64() < personaTimeoutProb {
		p.Timeouts++
		return nil
	}
	out := map[string]any{}
	meti := map[string]int{"competence": 0, "benevolence": 1, "integrity": 2}
	for _, it := range items {
		v := p.trust
		if k, ok := meti[it.ID]; ok {
			v = 1 + metiSlope*p.trust + p.Offsets[k]
		}
		x := clampI(int(math.Round(v+p.rng.NormFloat64()*personaNoise)), it.Scale.Min, it.Scale.Max)
		out[it.ID] = x
		if it.ID == "trust" {
			p.LastTrust = &x
		}
	}
	return out
}

// action picks an option of the descriptor: a checker mostly by style (0.7
// for the style's option, 0.15 each for the others), a neutral persona
// hasQuestion with its own probability.
func (p *Persona) action(slide string, opts []Choice) string {
	if len(opts) == 0 {
		return ""
	}
	has := func(id string) bool {
		for _, o := range opts {
			if o.ID == id {
				return true
			}
		}
		return false
	}
	id := ""
	switch p.Group {
	case "checker":
		main := ""
		for _, s := range personaStyles {
			if s.name == p.Style {
				main = s.main
			}
		}
		var ws []weighted
		for _, o := range opts {
			w := 0.15
			if o.ID == main {
				w = 0.7
			}
			ws = append(ws, weighted{o.ID, w})
		}
		id = pickWeighted(p.rng, ws)
	case "neutral":
		if has("hasQuestion") && has("noQuestion") {
			id = "noQuestion"
			if p.rng.Float64() < p.AskProb {
				id = "hasQuestion"
			}
		}
	}
	if id == "" {
		id = opts[p.rng.IntN(len(opts))].ID
	}
	p.Actions = append(p.Actions, slide+":"+id)
	return id
}

// budget votes approve with 0.85 when the last trust rating is 6 or more,
// otherwise 0.35; sureness 4 to 9 grows with the distance of that rating
// from 5.
func (p *Persona) budget(sc *PackScene) map[string]any {
	last := int(math.Round(p.trust))
	if p.LastTrust != nil {
		last = *p.LastTrust
	}
	prob := 0.35
	if last >= 6 {
		prob = 0.85
	}
	want := "reject"
	if p.rng.Float64() < prob {
		want = "approve"
	}
	opt := ""
	for _, o := range sc.Options {
		if o.ID == want {
			opt = o.ID
		}
	}
	if opt == "" && len(sc.Options) > 0 {
		// Unknown ids: the first option stands for approve, the second for reject.
		opt = sc.Options[0].ID
		if want == "reject" && len(sc.Options) > 1 {
			opt = sc.Options[1].ID
		}
	}
	p.Budget = opt
	out := map[string]any{"optionId": opt, "confidence": nil}
	sure := clampI(int(math.Round(4+math.Abs(float64(last)-5)+p.rng.NormFloat64()*0.8)), 4, 9)
	if sc.Confidence != nil {
		sure = clampI(sure, sc.Confidence.Scale.Min, sc.Confidence.Scale.Max)
		out["confidence"] = sure
		p.Sureness = &sure
	}
	return out
}

func suffix(id string) string {
	if i := strings.LastIndex(id, "."); i >= 0 {
		return id[i+1:]
	}
	return id
}

// background answers the scene pack's items: option ids from the pack,
// weighted by personaBackground; a scale (cyberKnowledge) peaks at 3; the
// country is the persona's; each item "preferNotToSay" with 0.05.
func (p *Persona) background(sc *PackScene) map[string]any {
	out := map[string]any{}
	for _, it := range sc.Items {
		switch it.Kind {
		case "single", "scale", "country":
		default:
			continue // honeypot: sent empty as hp
		}
		if p.rng.Float64() < personaPNTS {
			out[it.ID] = "preferNotToSay"
			continue
		}
		switch it.Kind {
		case "single":
			var ws []weighted
			for _, o := range it.Options {
				if o.ID == it.OtherOptionID {
					continue
				}
				w, ok := personaBackground[it.ID][suffix(o.ID)]
				if !ok {
					w = 1
				}
				ws = append(ws, weighted{o.ID, w})
			}
			if len(ws) > 0 {
				out[it.ID] = pickWeighted(p.rng, ws)
			}
		case "scale":
			if it.Scale != nil {
				mode := clampI(3, it.Scale.Min, it.Scale.Max)
				var ws []weighted
				for v := it.Scale.Min; v <= it.Scale.Max; v++ {
					d := math.Abs(float64(v - mode))
					ws = append(ws, weighted{fmt.Sprint(v), math.Max(0.5, 5-2*d)})
				}
				var v int
				fmt.Sscan(pickWeighted(p.rng, ws), &v)
				out[it.ID] = v
			}
		case "country":
			out[it.ID] = p.Country
		}
	}
	p.Background = out
	return out
}

// WritePersonas prints one row per persona.
func WritePersonas(w io.Writer, ss []*Session) {
	fmt.Fprintf(w, "\nPersonas\n| # | Session | Group | Baseline | METI c/b/i | Style or ask | Actions | Denial (inferred) | Ratings timed out | Budget | Sureness | Country |\n|---:|---|---|---:|---|---|---|---|---:|---|---:|---|\n")
	for _, s := range ss {
		p := s.persona
		if p == nil {
			continue
		}
		m := p.MetiBase()
		style := p.Style
		if p.Group == "neutral" {
			style = fmt.Sprintf("ask %.2f", p.AskProb)
		}
		den := "-"
		if p.Denial != "" {
			den = p.Denial + " (" + p.DenialHow + ")"
		}
		sure := "-"
		if p.Sureness != nil {
			sure = fmt.Sprint(*p.Sureness)
		}
		country := p.Country
		if v, ok := p.Background["country"].(string); ok {
			country = v
		}
		acts := append([]string(nil), p.Actions...)
		sort.Strings(acts)
		fmt.Fprintf(w, "| %d | %s | %s | %.1f | %.1f/%.1f/%.1f | %s | %s | %s | %d | %s | %s | %s |\n", s.n, s.ID, dash(p.Group), p.Baseline,
			m[0], m[1], m[2], style, dash(strings.Join(acts, " ")), den, p.Timeouts, dash(p.Budget), sure, country)
	}
}

func dash(s string) string {
	if s == "" {
		return "-"
	}
	return s
}
