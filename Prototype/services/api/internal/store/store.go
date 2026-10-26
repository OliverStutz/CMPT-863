// Package store defines the Store interface (7.1) and the transaction buffer
// shared by its implementations (memory and firestore).
package store

import (
	"context"
	"errors"
	"fmt"
	"time"

	"study/api/internal/domain"
)

var (
	ErrNotFound = errors.New("not found")
	ErrConflict = errors.New("conflict")
)

// SessionFilter selects sessions. Zero fields do not filter.
type SessionFilter struct {
	RunID           string
	Statuses        []string
	IP              string
	OpenedSince     *time.Time
	UpdatedSince    *time.Time
	LastEventBefore *time.Time
}

// Match applies the filter in memory. Implementations that query by an
// indexed field use it for the remaining conditions.
func (f SessionFilter) Match(s *domain.Session) bool {
	if f.RunID != "" && s.RunID != f.RunID {
		return false
	}
	if len(f.Statuses) > 0 {
		ok := false
		for _, st := range f.Statuses {
			if s.Status == st {
				ok = true
			}
		}
		if !ok {
			return false
		}
	}
	if f.IP != "" && s.IP != f.IP {
		return false
	}
	if f.OpenedSince != nil && s.OpenedAt.Before(*f.OpenedSince) {
		return false
	}
	if f.UpdatedSince != nil && !s.UpdatedAt.After(*f.UpdatedSince) {
		return false
	}
	if f.LastEventBefore != nil && !s.LastEventAt.Before(*f.LastEventBefore) {
		return false
	}
	return true
}

// Store is the persistence interface the HTTP layer, the sweeper and the
// export use.
type Store interface {
	CreateRun(ctx context.Context, r *domain.Run) error // ErrConflict if the code is taken
	GetRun(ctx context.Context, id string) (*domain.Run, error)
	GetRunByCode(ctx context.Context, code string) (*domain.Run, error)
	ListRuns(ctx context.Context) ([]*domain.Run, error)
	UpdateRun(ctx context.Context, r *domain.Run) error
	// DeleteRun deletes a run without sessions; ErrConflict otherwise.
	DeleteRun(ctx context.Context, id string) error

	CreateSession(ctx context.Context, s *domain.Session) error
	GetSession(ctx context.Context, id string) (*domain.Session, error)
	// ListSessions returns matching sessions, newest openedAt first.
	ListSessions(ctx context.Context, f SessionFilter) ([]*domain.Session, error)
	// RunTx runs fn in one transaction on a session (section 6 and 5.5).
	// Writes made through the Tx are applied only if fn returns nil.
	RunTx(ctx context.Context, sessionID string, fn func(domain.Tx) error) error

	Events(ctx context.Context, sessionID string) ([]*domain.Event, error)
	Ratings(ctx context.Context, sessionID string) ([]*domain.Rating, error)
	Answers(ctx context.Context, sessionID string) ([]*domain.AnswerDoc, error)
	// Counters returns the allocation counters of one run (per-run round robin,
	// decision 1.19 revised 2026-10-25).
	Counters(ctx context.Context, runID string) (*domain.Counters, error)

	PutConfig(ctx context.Context, c *domain.ConfigRecord) error
	GetConfig(ctx context.Context, version string) (*domain.ConfigRecord, error)
	ListConfigs(ctx context.Context) ([]*domain.ConfigRecord, error)

	Close() error
}

// Loader reads documents inside a transaction.
type Loader interface {
	LoadCounters(runID string) (*domain.Counters, error)
	LoadRating(point string) (*domain.Rating, error)
	LoadAnswers(kind string) (*domain.AnswerDoc, error)
	LoadConsumedEvent(clientSeq int64) (*domain.Event, error)
}

// Buffer implements domain.Tx over a Loader: reads see earlier buffered
// writes, and every write is kept until the store commits it.
type Buffer struct {
	L             Loader
	S             *domain.Session
	counter       *domain.Counters
	CountersDirty bool
	Ratings       map[string]*domain.Rating
	Answer        map[string]*domain.AnswerDoc
	Evs           []*domain.Event
	loadedR       map[string]bool
	loadedA       map[string]bool
}

func NewBuffer(l Loader, s *domain.Session) *Buffer {
	return &Buffer{L: l, S: s, Ratings: map[string]*domain.Rating{}, Answer: map[string]*domain.AnswerDoc{},
		loadedR: map[string]bool{}, loadedA: map[string]bool{}}
}

func (b *Buffer) Session() *domain.Session { return b.S }

func (b *Buffer) Counters() (*domain.Counters, error) {
	if b.counter == nil {
		c, err := b.L.LoadCounters(b.S.RunID)
		if err != nil {
			return nil, err
		}
		if c == nil {
			c = &domain.Counters{}
		}
		b.counter = c
	}
	return b.counter, nil
}

func (b *Buffer) PutCounters(c *domain.Counters) { b.counter = c; b.CountersDirty = true }

// CountersValue returns the buffered counters.
func (b *Buffer) CountersValue() *domain.Counters { return b.counter }

func (b *Buffer) Rating(point string) (*domain.Rating, error) {
	if r, ok := b.Ratings[point]; ok {
		return r, nil
	}
	if b.loadedR[point] {
		return nil, nil
	}
	r, err := b.L.LoadRating(point)
	if err != nil {
		return nil, err
	}
	b.loadedR[point] = true
	return r, nil
}

func (b *Buffer) PutRating(r *domain.Rating) { b.Ratings[r.Point] = r }

func (b *Buffer) Answers(kind string) (*domain.AnswerDoc, error) {
	if a, ok := b.Answer[kind]; ok {
		return a, nil
	}
	if b.loadedA[kind] {
		return nil, nil
	}
	a, err := b.L.LoadAnswers(kind)
	if err != nil {
		return nil, err
	}
	b.loadedA[kind] = true
	return a, nil
}

func (b *Buffer) PutAnswers(a *domain.AnswerDoc) { b.Answer[a.Kind] = a }

func (b *Buffer) PutEvent(e *domain.Event) { b.Evs = append(b.Evs, e) }

func (b *Buffer) ConsumedEvent(seq int64) (*domain.Event, error) {
	for _, e := range b.Evs {
		if e.Consumed && e.ClientSeq != nil && *e.ClientSeq == seq {
			return e, nil
		}
	}
	return b.L.LoadConsumedEvent(seq)
}

// EventID is the zero-padded document id of an event, e.g. "000042".
func EventID(no int64) string { return fmt.Sprintf("%06d", no) }
