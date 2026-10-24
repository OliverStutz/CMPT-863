// Command txbench isolates the Firestore transaction behaviour behind the
// session event path: W workers each run read-modify-write transactions on
// their own document (like one session's events), and optionally a share of
// them also read and write one shared document (like counters/allocation at
// consent). It reports throughput, latency, attempts and failures.
//
//	FIRESTORE_EMULATOR_HOST=localhost:8081 go run ./cmd/txbench -workers 200 -shared-rate 20
package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand/v2"
	"os"
	"sort"
	"sync"
	"sync/atomic"
	"time"

	"cloud.google.com/go/firestore"
)

func main() {
	project := flag.String("project", fmt.Sprintf("txbench-%d", time.Now().Unix()), "project id")
	workers := flag.Int("workers", 200, "concurrent workers, one document each")
	perWorker := flag.Float64("rate", 0.5, "transactions per second per worker")
	sharedRate := flag.Float64("shared-rate", 0, "transactions per second (total) that also touch the shared document")
	dur := flag.Duration("duration", 30*time.Second, "test duration")
	gate := flag.Bool("gate", false, "admit one shared transaction at a time in this process (the store's allocGate)")
	hold := flag.Duration("hold", 0, "work between the reads and the commit of a shared transaction (the app compiles the plan and builds the response there)")
	flag.Parse()
	if os.Getenv("FIRESTORE_EMULATOR_HOST") == "" {
		fmt.Println("set FIRESTORE_EMULATOR_HOST")
		os.Exit(2)
	}
	ctx := context.Background()
	c, err := firestore.NewClient(ctx, *project)
	if err != nil {
		panic(err)
	}
	defer c.Close()
	shared := c.Collection("counters").Doc("allocation")
	sem := make(chan struct{}, 1)
	var mu sync.Mutex
	var lats []time.Duration
	var ok, failed, attempts, sharedOK atomic.Int64
	errs := map[string]int{}
	pShared := 0.0
	if *sharedRate > 0 {
		pShared = *sharedRate / (float64(*workers) * *perWorker)
	}
	end := time.Now().Add(*dur)
	var wg sync.WaitGroup
	for w := 0; w < *workers; w++ {
		wg.Add(1)
		go func(w int) {
			defer wg.Done()
			ref := c.Collection("sessions").Doc(fmt.Sprintf("w%04d", w))
			interval := time.Duration(float64(time.Second) / *perWorker)
			time.Sleep(time.Duration(w) * interval / time.Duration(*workers))
			for time.Now().Before(end) {
				useShared := rand.Float64() < pShared
				t0 := time.Now()
				if useShared && *gate {
					sem <- struct{}{}
				}
				cctx, cancel := context.WithTimeout(ctx, 30*time.Second)
				err := c.RunTransaction(cctx, func(ctx context.Context, tx *firestore.Transaction) error {
					attempts.Add(1)
					var n int64
					if d, err := tx.Get(ref); err == nil {
						if v, err := d.DataAt("n"); err == nil {
							n = v.(int64)
						}
					}
					var m int64
					if useShared {
						if d, err := tx.Get(shared); err == nil {
							if v, err := d.DataAt("n"); err == nil {
								m = v.(int64)
							}
						}
					}
					if err := tx.Set(ref, map[string]any{"n": n + 1, "t": time.Now()}); err != nil {
						return err
					}
					if err := tx.Create(ref.Collection("events").Doc(fmt.Sprintf("%06d", n+1)), map[string]any{"n": n + 1}); err != nil {
						return err
					}
					if useShared {
						time.Sleep(*hold)
						return tx.Set(shared, map[string]any{"n": m + 1})
					}
					return nil
				})
				cancel()
				if useShared && *gate {
					<-sem
				}
				d := time.Since(t0)
				if err != nil {
					failed.Add(1)
					mu.Lock()
					errs[err.Error()]++
					mu.Unlock()
				} else {
					ok.Add(1)
					if useShared {
						sharedOK.Add(1)
					}
					mu.Lock()
					lats = append(lats, d)
					mu.Unlock()
				}
				if s := interval - d; s > 0 {
					time.Sleep(s)
				}
			}
		}(w)
	}
	wg.Wait()
	sort.Slice(lats, func(i, j int) bool { return lats[i] < lats[j] })
	p := func(q float64) time.Duration {
		if len(lats) == 0 {
			return 0
		}
		return lats[int(float64(len(lats)-1)*q)]
	}
	total := ok.Load() + failed.Load()
	fmt.Printf("gate %v, hold %v, workers %d, %.2f tx/s each, shared %.1f tx/s, %s: committed %d (shared %d), failed %d, attempts/tx %.2f, %.1f tx/s, p50 %v p95 %v p99 %v\n",
		*gate, *hold, *workers, *perWorker, *sharedRate, *dur, ok.Load(), sharedOK.Load(), failed.Load(), float64(attempts.Load())/float64(max(total, 1)),
		float64(ok.Load())/dur.Seconds(), p(.5).Round(time.Millisecond/10), p(.95).Round(time.Millisecond/10), p(.99).Round(time.Millisecond/10))
	for e, n := range errs {
		fmt.Printf("  %d x %s\n", n, e)
	}
}
