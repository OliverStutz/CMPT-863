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

	deleteSession(t, st, ct, run, sid, prefix, now)
}

// deleteSession checks DeleteSession: the session and its sub-documents are
// gone, the tombstone is written, other sessions and the counters stay.
func deleteSession(t *testing.T, st store.Store, ct *domain.Content, run *domain.Run, sid, prefix string, now time.Time) {
	ctx := context.Background()
	other := testutil.NewSession(ct, prefix+"other", now)
	other.RunID = run.ID
	if err := st.CreateSession(ctx, other); err != nil {
		t.Fatal(err)
	}
	trust := int64(7)
	for _, id := range []string{sid, other.ID} {
		if err := st.RunTx(ctx, id, func(tx domain.Tx) error {
			tx.PutRating(&domain.Rating{Point: "T0", Phase: domain.PhaseBaseline, Values: map[string]*int64{"trust": &trust}, Status: domain.RecordComplete})
			tx.PutAnswers(&domain.AnswerDoc{Kind: domain.AnswersBudget, Values: map[string]any{"optionId": "approve"}, Status: domain.RecordComplete})
			return nil
		}); err != nil {
			t.Fatal(err)
		}
	}
	before, _ := st.Counters(ctx, run.ID)
	got, _ := st.GetSession(ctx, sid)
	tomb := &domain.Deleted{SessionID: sid, RunID: run.ID, Group: got.Group, DeletedAt: now, Reason: "participant asked"}
	if err := st.DeleteSession(ctx, prefix+"nosuch", tomb); !errors.Is(err, store.ErrNotFound) {
		t.Fatalf("delete unknown: %v", err)
	}
	if err := st.DeleteSession(ctx, sid, tomb); err != nil {
		t.Fatal(err)
	}
	if _, err := st.GetSession(ctx, sid); !errors.Is(err, store.ErrNotFound) {
		t.Fatalf("deleted session still readable: %v", err)
	}
	evs, _ := st.Events(ctx, sid)
	rts, _ := st.Ratings(ctx, sid)
	ans, _ := st.Answers(ctx, sid)
	if len(evs)+len(rts)+len(ans) != 0 {
		t.Fatalf("sub-documents left: %d events, %d ratings, %d answers", len(evs), len(rts), len(ans))
	}
	if err := st.DeleteSession(ctx, sid, tomb); !errors.Is(err, store.ErrNotFound) {
		t.Fatalf("second delete: %v", err)
	}
	list, _ := st.ListSessions(ctx, store.SessionFilter{RunID: run.ID})
	if len(list) != 1 || list[0].ID != other.ID {
		t.Fatalf("other sessions after delete: %d", len(list))
	}
	if rts, _ := st.Ratings(ctx, other.ID); len(rts) != 1 {
		t.Fatalf("other session lost its ratings: %d", len(rts))
	}
	after, _ := st.Counters(ctx, run.ID)
	if fmt.Sprint(before.Groups, before.Denial) != fmt.Sprint(after.Groups, after.Denial) {
		t.Fatalf("counters changed: %+v -> %+v", before, after)
	}
	dels, err := st.ListDeleted(ctx)
	if err != nil {
		t.Fatal(err)
	}
	found := false
	for _, d := range dels {
		if d.SessionID == sid {
			found = d.RunID == run.ID && d.Group == got.Group && d.Reason == "participant asked" && d.DeletedAt.Equal(now)
		}
	}
	if !found {
		t.Fatalf("tombstone missing or wrong: %+v", dels)
	}
}
