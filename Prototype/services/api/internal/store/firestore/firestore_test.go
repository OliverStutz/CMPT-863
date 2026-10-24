package firestore_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"study/api/internal/store/firestore"
	"study/api/internal/store/storetest"
)

// TestConformance runs against the Firestore emulator (make dev starts it);
// it is skipped when FIRESTORE_EMULATOR_HOST is not set.
func TestConformance(t *testing.T) {
	if os.Getenv("FIRESTORE_EMULATOR_HOST") == "" {
		t.Skip("FIRESTORE_EMULATOR_HOST not set")
	}
	st, err := firestore.New(context.Background(), "demo-study-test")
	if err != nil {
		t.Fatal(err)
	}
	defer st.Close()
	storetest.Run(t, st, fmt.Sprintf("fs%d", time.Now().UnixNano()))
}
