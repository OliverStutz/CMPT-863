package main

import (
	"reflect"
	"testing"
)

func testPack() *Pack {
	p := &Pack{}
	p.Rating.Items = []RatingItem{{"trust", Scale{0, 10}}, {"competence", Scale{1, 7}}, {"benevolence", Scale{1, 7}}, {"integrity", Scale{1, 7}}}
	bg := PackScene{ID: "background", Kind: "background"}
	add := func(id, kind string, opts []string, sc *Scale) {
		var cs []Choice
		for _, o := range opts {
			cs = append(cs, Choice{ID: o})
		}
		bg.Items = append(bg.Items, struct {
			ID            string   `json:"id"`
			Kind          string   `json:"kind"`
			Options       []Choice `json:"options"`
			OtherOptionID string   `json:"otherOptionId"`
			Scale         *Scale   `json:"scale"`
		}{ID: id, Kind: kind, Options: cs, Scale: sc})
	}
	add("field", "single", []string{"field.business", "field.it", "field.cyber", "field.finance", "field.law", "field.arts"}, nil)
	add("country", "country", nil, nil)
	add("age", "single", []string{"age.a18", "age.a25", "age.a35", "age.a45", "age.a55", "age.a65"}, nil)
	add("cyberKnowledge", "scale", nil, &Scale{1, 5})
	add("hp", "honeypot", nil, nil)
	p.Scenes = []PackScene{bg, {ID: "budget", Kind: "budget", Options: []Choice{{ID: "approve"}, {ID: "reject"}},
		Confidence: &struct {
			Scale Scale `json:"scale"`
		}{Scale{0, 10}}}}
	return p
}

// draw runs a persona through a fixed sequence of decisions.
func draw(seed uint64, i int, group string) *Persona {
	p := NewPersona(seed, i)
	pk := testPack()
	p.learnGroup(group)
	p.ratingValues(pk.Rating.Items)
	for k, sl := range []string{"s2", "s3", "s4"} {
		p.applyRating(afterFlagged, sl)
		p.ratingValues(pk.Rating.Items)
		opts := []Choice{{ID: "nothing"}, {ID: "explain"}, {ID: "challenge"}}
		if group == "neutral" {
			opts = []Choice{{ID: "hasQuestion"}, {ID: "noQuestion"}}
		}
		p.action(sl, opts)
		p.observeAnswer(sl, []string{"A reason with enough words to count as a reason and to explain the figures on the slide, which come from the shipments of five days and the recovery costs on top.", "", ""}[min(k, 2)], 3)
		p.applyRating(afterAnswer, sl)
		p.ratingValues(pk.Rating.Items)
	}
	p.budget(pk.scene("budget"))
	p.background(pk.scene("background"))
	p.rng = nil // compare the drawn values, not the generator state
	return p
}

func TestPersonaDeterministic(t *testing.T) {
	for i := 0; i < 20; i++ {
		g := []string{"checker", "neutral"}[i%2]
		a, b := draw(42, i, g), draw(42, i, g)
		if !reflect.DeepEqual(a, b) {
			t.Fatalf("persona %d differs between two draws with the same seed:\n%+v\n%+v", i, a, b)
		}
	}
	if reflect.DeepEqual(draw(42, 0, "checker"), draw(43, 0, "checker")) {
		t.Fatal("different seeds gave the same persona")
	}
}

func TestPersonaRanges(t *testing.T) {
	const n = 4000
	mid, styles, countries := 0, map[string]int{}, map[string]int{}
	pk := testPack()
	for i := 0; i < n; i++ {
		p := NewPersona(7, i)
		if p.Baseline < 5 || p.Baseline > 9 {
			t.Fatalf("baseline %v out of 5..9", p.Baseline)
		}
		if p.Baseline >= 6 && p.Baseline <= 8 {
			mid++
		}
		for _, m := range p.MetiBase() {
			if m < 1 || m > 7 {
				t.Fatalf("METI %v out of 1..7", m)
			}
		}
		c, u := p.Checker, p.Neutral
		if c.Flag > -1 || c.Flag < -2.5 || u.Flag > 0 || u.Flag < -0.8 || c.Reason < 0.5 || c.Reason > 1.8 || u.Reason < 0.2 || u.Reason > 0.8 ||
			c.Denial < -1 || c.Denial > 0.3 || u.Denial < -0.5 || u.Denial > 0.3 {
			t.Fatalf("responses out of range: %+v %+v", c, u)
		}
		if p.AskProb < 0.2 || p.AskProb > 0.6 || p.Pick < 0 || p.Pick > 2 {
			t.Fatalf("ask %v or pick %d out of range", p.AskProb, p.Pick)
		}
		styles[p.Style]++
		countries[p.Country]++
		p.learnGroup("checker")
		p.applyRating(afterFlagged, "s2")
		if v := p.ratingValues(pk.Rating.Items); v != nil {
			for _, it := range pk.Rating.Items {
				x := v[it.ID].(int)
				if x < it.Scale.Min || x > it.Scale.Max {
					t.Fatalf("%s = %d out of %v", it.ID, x, it.Scale)
				}
			}
		}
		b := p.budget(pk.scene("budget"))
		if s := b["confidence"].(int); s < 4 || s > 9 {
			t.Fatalf("sureness %d out of 4..9", s)
		}
		bg := p.background(pk.scene("background"))
		if _, ok := bg["hp"]; ok {
			t.Fatal("honeypot answered")
		}
		if v, ok := bg["cyberKnowledge"].(int); ok && (v < 1 || v > 5) {
			t.Fatalf("cyberKnowledge %d", v)
		}
	}
	if f := float64(mid) / n; f < 0.7 {
		t.Errorf("only %.2f of baselines within 6..8", f)
	}
	for name, want := range map[string]float64{"quiet": 0.3, "curious": 0.4, "challenger": 0.3} {
		if f := float64(styles[name]) / n; f < want-0.04 || f > want+0.04 {
			t.Errorf("style %s %.3f, want about %.2f", name, f, want)
		}
	}
	if f := float64(countries["CA"]) / n; f < 0.31 || f > 0.39 {
		t.Errorf("CA share %.3f, want about 0.35", f)
	}
}

func TestDenialInference(t *testing.T) {
	reason := "This is a fair picture of our risk. Five days of lost shipments come to six million dollars. The other two million is the cost of recovery: outside specialists, overtime, replaced equipment and penalties to our customers. That is the risk we face."
	denialLong := "This is a fair picture of our risk. I have looked at these numbers very closely, and I am confident in every one of them. I would show this same picture to any board, and I stand behind it. That is the risk we face."
	denialShort := "This design is the right one for us. That is the plan I recommend."
	for _, c := range []struct {
		name    string
		answers []string
		slide   string
		how     string
	}{
		{"long denial first", []string{denialLong, reason, reason}, "s2", "transcript"},
		{"short denial second", []string{reason, denialShort, reason}, "s3", "transcript"},
		{"elimination", []string{reason, reason, reason}, "s4", "elimination"},
	} {
		p := NewPersona(1, 0)
		for k, sl := range []string{"s2", "s3", "s4"} {
			p.observeAnswer(sl, c.answers[k], 3)
		}
		if p.Denial != c.slide || p.DenialHow != c.how {
			t.Errorf("%s: denial %s (%s), want %s (%s)", c.name, p.Denial, p.DenialHow, c.slide, c.how)
		}
	}
	p := NewPersona(1, 0)
	for _, sl := range []string{"s2", "s3", "s4"} {
		p.observeAnswer(sl, "", 3)
	}
	if want := []string{"s2", "s3", "s4"}[p.Pick]; p.Denial != want || p.DenialHow != "fallback" {
		t.Errorf("fallback: denial %s (%s), want %s", p.Denial, p.DenialHow, want)
	}
}
