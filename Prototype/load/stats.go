package main

import (
	"fmt"
	"io"
	"sort"
	"strings"
	"sync"
	"time"
)

// Stats collects per-endpoint latencies, HTTP statuses and server error codes.
type Stats struct {
	mu        sync.Mutex
	lat       map[string][]time.Duration
	status    map[string]map[int]int
	codes     map[string]int // server error codes, e.g. too_early, seq_gap, wrong_state
	netErrors map[string]int
	retries   int // requests the generator repeated after a 5xx or a network error
}

func NewStats() *Stats {
	return &Stats{lat: map[string][]time.Duration{}, status: map[string]map[int]int{}, codes: map[string]int{}, netErrors: map[string]int{}}
}

func (s *Stats) Record(endpoint string, d time.Duration, status int, code string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.lat[endpoint] = append(s.lat[endpoint], d)
	if s.status[endpoint] == nil {
		s.status[endpoint] = map[int]int{}
	}
	s.status[endpoint][status]++
	if code != "" {
		s.codes[code]++
	}
}

func (s *Stats) NetError(endpoint string) {
	s.mu.Lock()
	s.netErrors[endpoint]++
	s.mu.Unlock()
}

func (s *Stats) Retry() {
	s.mu.Lock()
	s.retries++
	s.mu.Unlock()
}

func pct(xs []time.Duration, p float64) time.Duration {
	if len(xs) == 0 {
		return 0
	}
	i := int(float64(len(xs)-1)*p + 0.5)
	return xs[i]
}

func ms(d time.Duration) string { return fmt.Sprintf("%.1f", float64(d.Microseconds())/1000) }

// Write prints the endpoint table, the status counts and the rejections.
func (s *Stats) Write(w io.Writer) {
	s.mu.Lock()
	defer s.mu.Unlock()
	var eps []string
	for k := range s.lat {
		eps = append(eps, k)
	}
	sort.Strings(eps)
	fmt.Fprintf(w, "| Endpoint | Requests | p50 ms | p95 ms | p99 ms | max ms | Statuses |\n|---|---:|---:|---:|---:|---:|---|\n")
	var all []time.Duration
	var evs []time.Duration
	for _, ep := range eps {
		xs := append([]time.Duration(nil), s.lat[ep]...)
		sort.Slice(xs, func(i, j int) bool { return xs[i] < xs[j] })
		all = append(all, xs...)
		if strings.HasPrefix(ep, "POST events") {
			evs = append(evs, xs...)
		}
		fmt.Fprintf(w, "| %s | %d | %s | %s | %s | %s | %s |\n", ep, len(xs), ms(pct(xs, .5)), ms(pct(xs, .95)), ms(pct(xs, .99)),
			ms(xs[len(xs)-1]), statusString(s.status[ep]))
	}
	for _, g := range []struct {
		name string
		xs   []time.Duration
	}{{"all events", evs}, {"all requests", all}} {
		sort.Slice(g.xs, func(i, j int) bool { return g.xs[i] < g.xs[j] })
		if len(g.xs) > 0 {
			fmt.Fprintf(w, "| **%s** | %d | %s | %s | %s | %s | |\n", g.name, len(g.xs), ms(pct(g.xs, .5)), ms(pct(g.xs, .95)),
				ms(pct(g.xs, .99)), ms(g.xs[len(g.xs)-1]))
		}
	}
	total := map[int]int{}
	for _, m := range s.status {
		for k, v := range m {
			total[k] += v
		}
	}
	fmt.Fprintf(w, "\nHTTP statuses, all endpoints: %s\n", statusString(total))
	fmt.Fprintf(w, "Server rejections: too_early=%d seq_gap=%d wrong_state=%d", s.codes["too_early"], s.codes["seq_gap"], s.codes["wrong_state"])
	var other []string
	for k, v := range s.codes {
		if k != "too_early" && k != "seq_gap" && k != "wrong_state" {
			other = append(other, fmt.Sprintf("%s=%d", k, v))
		}
	}
	sort.Strings(other)
	if len(other) > 0 {
		fmt.Fprintf(w, " other: %s", strings.Join(other, " "))
	}
	fmt.Fprintln(w)
	ne := 0
	for _, v := range s.netErrors {
		ne += v
	}
	fmt.Fprintf(w, "Network errors: %d; requests repeated by the generator after 5xx or network errors: %d\n", ne, s.retries)
}

func statusString(m map[int]int) string {
	var ks []int
	for k := range m {
		ks = append(ks, k)
	}
	sort.Ints(ks)
	var parts []string
	for _, k := range ks {
		parts = append(parts, fmt.Sprintf("%d:%d", k, m[k]))
	}
	return strings.Join(parts, " ")
}
