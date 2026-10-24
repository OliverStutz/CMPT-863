package main

import (
	"fmt"
	"io"
	"regexp"
	"sort"
	"strings"
)

// Price per million characters in USD, list prices without the monthly free
// tier. Check https://cloud.google.com/text-to-speech/pricing before relying
// on them. Unknown families use the Neural2 price.
var pricePerMillion = map[string]float64{
	"Standard": 4,
	"Wavenet":  16,
	"Neural2":  16,
	"News":     16,
	"Polyglot": 16,
	"Chirp3":   30,
	"Chirp":    30,
	"Studio":   160,
}

// family returns the voice family of a Google voice name such as
// "en-US-Neural2-J" or "en-US-Chirp3-HD-Charon".
func family(voice string) string {
	parts := strings.Split(voice, "-")
	if len(parts) >= 3 {
		return parts[2]
	}
	return voice
}

func priceOf(voice string) float64 {
	if p, ok := pricePerMillion[family(voice)]; ok {
		return p
	}
	return pricePerMillion["Neural2"]
}

type estimateT struct {
	byVoice    map[string]int // all characters per voice
	newByVoice map[string]int // characters not in the cache
	chars      int
	newChars   int
}

// estimate counts the characters sent to the API (the full SSML input,
// tags included, which over-counts slightly) per voice.
func estimate(parts []part, ch cache) estimateT {
	e := estimateT{byVoice: map[string]int{}, newByVoice: map[string]int{}}
	for _, p := range parts {
		n := len([]rune(p.input()))
		e.byVoice[p.Voice.Name] += n
		e.chars += n
		if _, _, ok := ch.get(p.key()); !ok {
			e.newByVoice[p.Voice.Name] += n
			e.newChars += n
		}
	}
	return e
}

func (e estimateT) cost(m map[string]int) float64 {
	var c float64
	for v, n := range m {
		c += float64(n) * priceOf(v) / 1e6
	}
	return c
}

func (e estimateT) print(w io.Writer) {
	var vs []string
	for v := range e.byVoice {
		vs = append(vs, v)
	}
	sort.Strings(vs)
	for _, v := range vs {
		fmt.Fprintf(w, "  %-24s %6d chars (%d not cached), $%.0f per 1M chars\n", v, e.byVoice[v], e.newByVoice[v], priceOf(v))
	}
	fmt.Fprintf(w, "characters: %d in total, %d to synthesise\n", e.chars, e.newChars)
	fmt.Fprintf(w, "estimated cost: $%.4f for a full run from scratch, $%.4f for this run (list price, before the monthly free tier)\n",
		e.cost(e.byVoice), e.cost(e.newByVoice))
}

var markRe = regexp.MustCompile(`<mark\s+name\s*=\s*["']([^"']+)["']`)

// markNames returns the names of the <mark> tags in an SSML string.
func markNames(ssml string) []string {
	var out []string
	for _, m := range markRe.FindAllStringSubmatch(ssml, -1) {
		out = append(out, m[1])
	}
	return out
}
