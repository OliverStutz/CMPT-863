// Command load drives N concurrent participant sessions through the complete
// study flow over the public participant API (02 - Architecture.md 8.1),
// honouring the server's step descriptors and timing, then checks the result
// through the admin API. See README.md.
package main

import (
	"flag"
	"fmt"
	"math/rand/v2"
	"net/http"
	"os"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// Gen holds what all sessions share.
type Gen struct {
	// audioChoice is the audio check option every session answers (the public
	// scene pack hides the correct one); empty picks at random.
	audioChoice string
	base        string
	runCode     string
	thinkMul    float64
	transport   *http.Transport
	plain       *http.Client
	stats       *Stats
	clipCache   sync.Map
	packs       map[string]*Pack
	failed      atomic.Int64
}

func main() {
	base := flag.String("base", "http://localhost:8080", "server base URL")
	runCode := flag.String("run-code", "", "run code to use; empty creates a run through the admin API")
	config := flag.String("config", "test-fast-1.4.0", "config version for a created run (versions starting with test need ALLOW_TEST_CONFIGS=true)")
	n := flag.Int("sessions", 50, "number of concurrent sessions")
	ramp := flag.Float64("ramp", 5, "seconds over which the sessions start (evenly spaced)")
	think := flag.Float64("think", 1, "multiplier for participant think times (server minimum view times always apply)")
	check := flag.Bool("check", true, "check the result through the admin API (LOAD_GATE_KEY, LOAD_ADMIN_PASSWORD)")
	seed := flag.Uint64("seed", uint64(time.Now().UnixNano()), "random seed")
	ipBase := flag.Int("ip-base", 1, "first octet offset for the fake X-Forwarded-For addresses (use different values across runs within an hour)")
	personas := flag.Int("persona", 0, "persona mode: N emulated participants with modelled answers (overrides -sessions); 0 keeps the random answers")
	runName := flag.String("run-name", "Dashboard pilot (emulated)", "name of a run created in persona mode")
	source := flag.String("source", "Emulated", "source of a run created in persona mode")
	audioChoice := flag.String("audio-choice", "", "option id every session answers at the audio check (for v1.4: lantern); empty picks at random, which flags 3 of 4 sessions audio_check_wrong")
	flag.Parse()
	if *personas > 0 {
		*n = *personas
	}

	*base = strings.TrimSuffix(*base, "/")
	g := &Gen{base: *base, runCode: *runCode, thinkMul: *think, stats: NewStats(), packs: map[string]*Pack{}, audioChoice: *audioChoice}
	g.transport = &http.Transport{MaxIdleConns: 4 * *n, MaxIdleConnsPerHost: 4 * *n, MaxConnsPerHost: 0, IdleConnTimeout: 90 * time.Second}
	g.plain = &http.Client{Transport: g.transport, Timeout: 60 * time.Second}

	var adm *Admin
	if *check || g.runCode == "" {
		var err error
		adm, err = NewAdmin(g, os.Getenv("LOAD_GATE_KEY"), os.Getenv("LOAD_ADMIN_PASSWORD"))
		if err != nil {
			fatal("admin login: %v (set LOAD_GATE_KEY and LOAD_ADMIN_PASSWORD, or pass -run-code and -check=false)", err)
		}
	}
	if g.runCode == "" {
		name, src := fmt.Sprintf("load %d sessions %s", *n, time.Now().Format("15:04:05")), "load test"
		if *personas > 0 {
			name, src = *runName, *source
		}
		code, err := adm.CreateRun(name, src, *config)
		if err != nil {
			fatal("create run: %v", err)
		}
		g.runCode = code
	}
	if *personas > 0 {
		fmt.Printf("Load: %d personas, ramp %.1f s, think x%.2f, run %s, base %s, seed %d\n", *n, *ramp, *think, g.runCode, g.base, *seed)
	} else {
		fmt.Printf("Load: %d sessions, ramp %.1f s, think x%.2f, run %s, base %s\n", *n, *ramp, *think, g.runCode, g.base)
	}

	sessions := make([]*Session, *n)
	var wg sync.WaitGroup
	start := time.Now()
	for i := range sessions {
		s := &Session{g: g, n: i, ip: fakeIP(*ipBase, i), seq: 1, rng: rand.New(rand.NewPCG(*seed, uint64(i)))}
		if *personas > 0 {
			s.persona = NewPersona(*seed, i)
		}
		sessions[i] = s
		delay := time.Duration(0)
		if *n > 1 {
			delay = time.Duration(*ramp * float64(time.Second) * float64(i) / float64(*n-1))
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(delay)
			s.Run()
			if s.Err != nil {
				g.failed.Add(1)
			}
		}()
	}
	progressDone := make(chan struct{})
	go progress(g, sessions, start, progressDone)
	wg.Wait()
	close(progressDone)
	wall := time.Since(start)

	ok, failed := 0, 0
	var durs []time.Duration
	for _, s := range sessions {
		if s.Err != nil {
			failed++
			if failed <= 10 {
				fmt.Printf("session %d (%s): %v\n", s.n, s.ID, s.Err)
			}
			continue
		}
		ok++
		durs = append(durs, s.Elapsed)
	}
	sort.Slice(durs, func(i, j int) bool { return durs[i] < durs[j] })
	fmt.Println()
	g.stats.Write(os.Stdout)
	fmt.Printf("Sessions completed: %d of %d (failed %d). Wall time %.1f s", ok, *n, failed, wall.Seconds())
	if len(durs) > 0 {
		fmt.Printf("; session duration p50 %.1f s, max %.1f s", pct(durs, .5).Seconds(), durs[len(durs)-1].Seconds())
	}
	fmt.Println()
	if *personas > 0 {
		WritePersonas(os.Stdout, sessions)
	}

	pass := failed == 0
	if *check {
		if !adm.Check(sessions) {
			pass = false
		}
	}
	if !pass {
		fmt.Println("RESULT: FAIL")
		os.Exit(1)
	}
	fmt.Println("RESULT: PASS")
}

// fakeIP gives each session its own client address. The server must run with
// TRUSTED_PROXY_HOPS=1 to read it from X-Forwarded-For; otherwise every
// session comes from 127.0.0.1 and the per-IP limit (30 an hour) applies.
func fakeIP(base, i int) string {
	return fmt.Sprintf("10.%d.%d.%d", (base+i/62500)%256, (i/250)%250+1, i%250+1)
}

func progress(g *Gen, ss []*Session, start time.Time, done chan struct{}) {
	t := time.NewTicker(10 * time.Second)
	defer t.Stop()
	for {
		select {
		case <-done:
			return
		case <-t.C:
			fin, errs, events := 0, int(g.failed.Load()), 0
			for _, s := range ss {
				if s.Done.Load() {
					fin++
				}
				events += int(s.Events.Load())
			}
			fmt.Fprintf(os.Stderr, "[%5.0fs] finished %d/%d, failed %d, events accepted %d\n", time.Since(start).Seconds(), fin, len(ss), errs, events)
		}
	}
}

func fatal(f string, a ...any) {
	fmt.Fprintf(os.Stderr, f+"\n", a...)
	os.Exit(2)
}
