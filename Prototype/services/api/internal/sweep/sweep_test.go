package sweep_test

import (
	"context"
	"testing"
	"time"

	"study/api/internal/domain"
	"study/api/internal/store/memory"
	"study/api/internal/sweep"
	"study/api/internal/testutil"
)

func TestSweepMarksAbandoned(t *testing.T) {
	ct, err := domain.NewContent(testutil.FastConfig("test-fast", 3, 2), nil, "/a/", "/s/")
	if err != nil {
		t.Fatal(err)
	}
	st := memory.New()
	now := time.Date(2026, 11, 1, 10, 0, 0, 0, time.UTC)
	ctx := context.Background()
	_ = st.CreateSession(ctx, testutil.NewSession(ct, "old", now))
	_ = st.CreateSession(ctx, testutil.NewSession(ct, "fresh", now.Add(50*time.Second)))
	content := func(string) (*domain.Content, error) { return ct, nil }
	// abandonAfterMs is 60 s in the fast config.
	n, err := sweep.Run(ctx, st, content, now.Add(61*time.Second), 30*time.Minute, nil)
	if err != nil || n != 1 {
		t.Fatalf("n=%d err=%v", n, err)
	}
	old, _ := st.GetSession(ctx, "old")
	fresh, _ := st.GetSession(ctx, "fresh")
	if old.Status != domain.StatusAbandoned || fresh.Status != domain.StatusOpened {
		t.Fatalf("%s %s", old.Status, fresh.Status)
	}
}
