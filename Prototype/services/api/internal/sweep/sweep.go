// Package sweep marks sessions Abandoned after abandonAfterMs without events
// (5.4, 2.13), after running lazy evaluation so expired windows are recorded.
package sweep

import (
	"context"
	"log/slog"
	"time"

	"study/api/internal/domain"
	"study/api/internal/store"
)

// ContentFunc returns the content of a config version, or an error.
type ContentFunc func(version string) (*domain.Content, error)

// Run sweeps once and returns the number of sessions marked Abandoned.
// DECISION: the threshold is each session's config abandonAfterMs; when the
// config cannot be loaded (12.3 gives the sweeper no bucket access), lazy
// evaluation is skipped and fallback is used.
func Run(ctx context.Context, st store.Store, content ContentFunc, now time.Time, fallback time.Duration, log *slog.Logger) (int, error) {
	if log == nil {
		log = slog.Default()
	}
	// Candidates are read broadly, then checked per session.
	cut := now.Add(-time.Minute)
	sessions, err := st.ListSessions(ctx, store.SessionFilter{
		Statuses:        []string{domain.StatusOpened, domain.StatusStarted, domain.StatusOngoing},
		LastEventBefore: &cut,
	})
	if err != nil {
		return 0, err
	}
	n := 0
	for _, s := range sessions {
		ct, cerr := content(s.ConfigVersion)
		after := fallback
		if cerr == nil {
			after = time.Duration(ct.Config.Timing.AbandonAfterMs) * time.Millisecond
		}
		if now.Sub(s.LastEventAt) <= after {
			continue
		}
		marked := false
		err := st.RunTx(ctx, s.ID, func(tx domain.Tx) error {
			marked = false
			if cerr == nil {
				if err := domain.Advance(tx, ct, now); err != nil {
					return err
				}
			}
			x := tx.Session()
			if x.Final() || x.Status == domain.StatusAbandoned || now.Sub(x.LastEventAt) <= after {
				return nil
			}
			x.Status = domain.StatusAbandoned
			x.UpdatedAt = now
			marked = true
			return nil
		})
		if err != nil {
			log.Error("sweep session", "session", s.ID, "err", err)
			continue
		}
		if marked {
			n++
		}
	}
	return n, nil
}
