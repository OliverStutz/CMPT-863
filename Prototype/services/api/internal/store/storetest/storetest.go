// Package storetest is a conformance test for store.Store implementations.
package storetest

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"

	"study/api/internal/domain"
	"study/api/internal/events"
	"study/api/internal/store"
	"study/api/internal/testutil"
)

// Run checks runs, the transaction rollback, the allocation transaction and
// concurrent events on one session.
func Run(t *testing.T, st store.Store, prefix string) {
	ctx := context.Background()
	ct, err := domain.NewContent(testutil.FastConfig("test-fast", 3, 2), nil, "/a/", "/s/")
	if err != nil {
		t.Fatal(err)
	}
	now := time.Now().UTC().Truncate(time.Millisecond)

	run := &domain.Run{ID: prefix + "run", Code: prefix + "k7m", Name: "n", Source: "s", ConfigVersion: "test-fast", Status: domain.RunOpen, CreatedAt: now}
	if err := st.CreateRun(ctx, run); err != nil {
		t.Fatal(err)
	}
	dup := *run
	dup.ID = prefix + "run2"
	if err := st.CreateRun(ctx, &dup); !errors.Is(err, store.ErrConflict) {
		t.Fatalf("duplicate code: %v", err)
	}
	if got, err := st.GetRunByCode(ctx, run.Code); err != nil || got.ID != run.ID {
		t.Fatalf("by code: %v", err)
	}

	sid := prefix + "sess"
	s := testutil.NewSession(ct, sid, now)
	s.RunID = run.ID
	if err := st.CreateSession(ctx, s); err != nil {
		t.Fatal(err)
	}
	if err := st.DeleteRun(ctx, run.ID); !errors.Is(err, store.ErrConflict) {
		t.Fatalf("delete with sessions: %v", err)
	}

	// A failing transaction writes nothing.
	_ = st.RunTx(ctx, sid, func(tx domain.Tx) error {
		tx.Session().Status = "X"
		tx.PutEvent(&domain.Event{EventNo: 1, Type: "x", Payload: map[string]any{}})
		return fmt.Errorf("boom")
	})
	if got, _ := st.GetSession(ctx, sid); got.Status != domain.StatusOpened {
		t.Fatal("rolled-back transaction changed the session")
	}
	if evs, _ := st.Events(ctx, sid); len(evs) != 0 {
		t.Fatal("rolled-back transaction wrote an event")
	}

	// Concurrent events with the same clientSeq: exactly one is consumed.
	b, _ := json.Marshal(map[string]any{"widthPx": 1440, "heightPx": 900})
	var wg sync.WaitGroup
	results := make([]int, 8)
	for i := range results {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			req := &events.Request{ClientSeq: 1, Type: events.DeviceChecked, ClientTs: i, Payload: b}
			_ = st.RunTx(ctx, sid, func(tx domain.Tx) error {
				r, err := domain.HandleEvent(tx, ct, time.Now().UTC(), req)
				if err == nil {
					results[i] = r.Status
				}
				return err
			})
		}(i)
	}
	wg.Wait()
	ok := 0
	for _, c := range results {
		if c == 200 {
			ok++
		}
	}
	if ok != 1 {
		t.Fatalf("accepted %d of %v", ok, results)
	}
	evs, err := st.Events(ctx, sid)
	if err != nil {
		t.Fatal(err)
	}
	for i, e := range evs {
		if e.EventNo != int64(i+1) {
			t.Fatalf("event numbers not contiguous: %d at %d", e.EventNo, i)
		}
	}
	got, _ := st.GetSession(ctx, sid)
	if got.LastClientSeq != 1 || got.LastEventNo != int64(len(evs)) || got.CurrentState != "consent:consent" {
		t.Fatalf("session after concurrency: seq %d eventNo %d/%d state %s", got.LastClientSeq, got.LastEventNo, len(evs), got.CurrentState)
	}

	// Consent allocates in the same transaction as the counters.
	req := &events.Request{ClientSeq: 2, Type: events.ConsentAccepted, Payload: []byte(`{"consentVersion":"c1","checkbox":true}`)}
	if err := st.RunTx(ctx, sid, func(tx domain.Tx) error {
		_, err := domain.HandleEvent(tx, ct, time.Now().UTC().Add(time.Minute), req)
		return err
	}); err != nil {
		t.Fatal(err)
	}
	got, _ = st.GetSession(ctx, sid)
	c, _ := st.Counters(ctx, run.ID)
	if got.Group == "" || c.Groups[got.Group] < 1 || c.Denial[got.Group][got.DenialSlide] < 1 {
		t.Fatalf("allocation %s %s %+v", got.Group, got.DenialSlide, c)
	}
	list, err := st.ListSessions(ctx, store.SessionFilter{RunID: run.ID})
	if err != nil || len(list) != 1 {
		t.Fatalf("list %d %v", len(list), err)
	}
}
