// Package firestore is the Firestore implementation of store.Store (7.1).
// With FIRESTORE_EMULATOR_HOST set, the client talks to the emulator.
//
// Every session transaction runs in one Firestore transaction: all reads
// (session, counters, ratings, answers, the consumed event of a retry) happen
// inside the callback, and all writes are buffered and applied at its end, so
// Firestore's reads-before-writes rule holds and a retried attempt starts from
// a clean buffer.
package firestore

import (
	"context"
	"errors"
	"log/slog"
	"sort"
	"sync/atomic"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"study/api/internal/domain"
	"study/api/internal/store"
)

const (
	colRuns     = "runs"
	colSessions = "sessions"
	colCounters = "counters"
	// docAlloc was the single global allocation document until 2026-10-25;
	// counters are now one document per run, counters/run-<runId> (1.19).
	docAlloc   = "allocation"
	colConfigs = "configs"
	subEvents  = "events"
	subRatings = "ratings"
	subAnswers = "answers"
)

// Store is the Firestore implementation.
type Store struct {
	c *firestore.Client
	// alloc admits one allocating transaction (one that reads
	// counters/allocation, that is a consent) at a time per process. See
	// allocGate.
	alloc chan struct{}
}

// New connects to the (default) database of a project.
func New(ctx context.Context, projectID string) (*Store, error) {
	c, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	return &Store{c: c, alloc: make(chan struct{}, 1)}, nil
}

// allocGate serialises, within one process, the transactions that read and
// write the shared counters/allocation document (5.5).
//
// DECISION (reports/concurrency-review.md): without it, concurrent consents
// in one instance all read the counter document and then conflict at commit.
// Under the emulator's pessimistic locking (and a database in pessimistic
// mode) two readers that both want to write deadlock until the lock timeout,
// and every consent arriving meanwhile joins the convoy: at 200 sessions
// started within 10 s, 60 % of consents failed after 5 attempts. Under
// optimistic concurrency (the production database) the same overlap is an
// abort and a retry. With the gate, consents of one instance never contend
// with each other; only consents on different Cloud Run instances can, and
// the transaction still guarantees correctness then. Trade-off: one
// instance allocates at most one session per transaction round trip (about
// 25 to 50 a second in production), far above any expected sign-up rate.
// The gate is held across the client library's retries of the same
// transaction and released when RunTx returns.
type allocGate struct {
	s    *Store
	held bool
}

func (g *allocGate) acquire(ctx context.Context) error {
	if g.held {
		return nil
	}
	select {
	case g.s.alloc <- struct{}{}:
		g.held = true
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (g *allocGate) release() {
	if g.held {
		<-g.s.alloc
		g.held = false
	}
}

func (s *Store) Close() error { return s.c.Close() }

func notFound(err error) bool { return status.Code(err) == codes.NotFound }

func (s *Store) CreateRun(ctx context.Context, r *domain.Run) error {
	return s.c.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		docs, err := tx.Documents(s.c.Collection(colRuns).Where("code", "==", r.Code).Limit(1)).GetAll()
		if err != nil {
			return err
		}
		if len(docs) > 0 {
			return store.ErrConflict
		}
		return tx.Create(s.c.Collection(colRuns).Doc(r.ID), r)
	})
}

func runFrom(d *firestore.DocumentSnapshot) (*domain.Run, error) {
	var r domain.Run
	if err := d.DataTo(&r); err != nil {
		return nil, err
	}
	r.ID = d.Ref.ID
	return &r, nil
}

func (s *Store) GetRun(ctx context.Context, id string) (*domain.Run, error) {
	d, err := s.c.Collection(colRuns).Doc(id).Get(ctx)
	if notFound(err) {
		return nil, store.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return runFrom(d)
}

func (s *Store) GetRunByCode(ctx context.Context, code string) (*domain.Run, error) {
	docs, err := s.c.Collection(colRuns).Where("code", "==", code).Limit(1).Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}
	if len(docs) == 0 {
		return nil, store.ErrNotFound
	}
	return runFrom(docs[0])
}

func (s *Store) ListRuns(ctx context.Context) ([]*domain.Run, error) {
	docs, err := s.c.Collection(colRuns).Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}
	var out []*domain.Run
	for _, d := range docs {
		r, err := runFrom(d)
		if err != nil {
			return nil, err
		}
		out = append(out, r)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].CreatedAt.After(out[j].CreatedAt) })
	return out, nil
}

func (s *Store) UpdateRun(ctx context.Context, r *domain.Run) error {
	_, err := s.c.Collection(colRuns).Doc(r.ID).Set(ctx, r)
	return err
}

func (s *Store) DeleteRun(ctx context.Context, id string) error {
	return s.c.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		ref := s.c.Collection(colRuns).Doc(id)
		if _, err := tx.Get(ref); err != nil {
			if notFound(err) {
				return store.ErrNotFound
			}
			return err
		}
		docs, err := tx.Documents(s.c.Collection(colSessions).Where("runId", "==", id).Limit(1)).GetAll()
		if err != nil {
			return err
		}
		if len(docs) > 0 {
			return store.ErrConflict
		}
		return tx.Delete(ref)
	})
}

func (s *Store) CreateSession(ctx context.Context, x *domain.Session) error {
	_, err := s.c.Collection(colSessions).Doc(x.ID).Create(ctx, x)
	if status.Code(err) == codes.AlreadyExists {
		return store.ErrConflict
	}
	return err
}

func sessionFrom(d *firestore.DocumentSnapshot) (*domain.Session, error) {
	var x domain.Session
	if err := d.DataTo(&x); err != nil {
		return nil, err
	}
	x.ID = d.Ref.ID
	return &x, nil
}

func (s *Store) GetSession(ctx context.Context, id string) (*domain.Session, error) {
	d, err := s.c.Collection(colSessions).Doc(id).Get(ctx)
	if notFound(err) {
		return nil, store.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return sessionFrom(d)
}

// ListSessions queries by one indexed field (7.1 indexes) and applies the
// rest of the filter in memory.
// DECISION: with about 40 sessions, filters beyond the indexed field are
// applied in memory rather than through further composite indexes.
func (s *Store) ListSessions(ctx context.Context, f store.SessionFilter) ([]*domain.Session, error) {
	q := s.c.Collection(colSessions).Query
	// DECISION (reports/concurrency-review.md): updatedSince comes first. The
	// admin page polls every 5 s with updatedSince and often a runId; with
	// runId first every poll read every session of the run, with updatedSince
	// first it reads only the sessions changed since the last poll.
	switch {
	case f.UpdatedSince != nil:
		q = q.Where("updatedAt", ">", *f.UpdatedSince)
	case f.RunID != "":
		q = q.Where("runId", "==", f.RunID)
	case f.IP != "":
		q = q.Where("ip", "==", f.IP)
	case len(f.Statuses) > 0:
		q = q.Where("status", "in", f.Statuses)
	}
	it := q.Documents(ctx)
	defer it.Stop()
	var out []*domain.Session
	for {
		d, err := it.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			return nil, err
		}
		x, err := sessionFrom(d)
		if err != nil {
			return nil, err
		}
		if f.Match(x) {
			out = append(out, x)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].OpenedAt.After(out[j].OpenedAt) })
	return out, nil
}

type loader struct {
	s    *Store
	ctx  context.Context
	tx   *firestore.Transaction
	ref  *firestore.DocumentRef
	gate *allocGate
}

// countersDoc is the per-run allocation document id.
func countersDoc(runID string) string { return "run-" + runID }

func (l loader) LoadCounters(runID string) (*domain.Counters, error) {
	if err := l.gate.acquire(l.ctx); err != nil {
		return nil, err
	}
	d, err := l.tx.Get(l.s.c.Collection(colCounters).Doc(countersDoc(runID)))
	if notFound(err) {
		return &domain.Counters{}, nil
	}
	if err != nil {
		return nil, err
	}
	var c domain.Counters
	return &c, d.DataTo(&c)
}

func (l loader) LoadRating(point string) (*domain.Rating, error) {
	d, err := l.tx.Get(l.ref.Collection(subRatings).Doc(point))
	if notFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var r domain.Rating
	return &r, d.DataTo(&r)
}

func (l loader) LoadAnswers(kind string) (*domain.AnswerDoc, error) {
	d, err := l.tx.Get(l.ref.Collection(subAnswers).Doc(kind))
	if notFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var a domain.AnswerDoc
	return &a, d.DataTo(&a)
}

func (l loader) LoadConsumedEvent(seq int64) (*domain.Event, error) {
	docs, err := l.tx.Documents(l.ref.Collection(subEvents).Where("clientSeq", "==", seq).Where("consumed", "==", true).Limit(1)).GetAll()
	if err != nil {
		return nil, err
	}
	if len(docs) == 0 {
		return nil, nil
	}
	var e domain.Event
	return &e, docs[0].DataTo(&e)
}

// TxStats counts session transactions and their attempts since the process
// started. An attempt beyond the first is a retry after contention (Firestore
// answered ABORTED and the client library ran the callback again).
// Added for the concurrency review (reports/concurrency-review.md).
var TxStats struct {
	Transactions, Attempts, Retried, AllocRetried, Failed atomic.Int64
}

func (s *Store) RunTx(ctx context.Context, id string, fn func(domain.Tx) error) error {
	attempts, alloc := 0, false
	err := s.runTx(ctx, id, func(b *store.Buffer) { attempts++; alloc = b.CountersDirty }, fn)
	TxStats.Transactions.Add(1)
	TxStats.Attempts.Add(int64(attempts))
	if attempts > 1 {
		TxStats.Retried.Add(1)
		if alloc {
			TxStats.AllocRetried.Add(1)
		}
		slog.Info("firestore transaction retried", "session", id, "attempts", attempts, "allocation", alloc, "err", err)
	}
	if err != nil && !errors.Is(err, store.ErrNotFound) {
		TxStats.Failed.Add(1)
	}
	return err
}

func (s *Store) runTx(ctx context.Context, id string, attempt func(*store.Buffer), fn func(domain.Tx) error) error {
	ref := s.c.Collection(colSessions).Doc(id)
	gate := &allocGate{s: s}
	defer gate.release()
	return s.c.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		d, err := tx.Get(ref)
		if notFound(err) {
			return store.ErrNotFound
		}
		if err != nil {
			return err
		}
		sess, err := sessionFrom(d)
		if err != nil {
			return err
		}
		b := store.NewBuffer(loader{s: s, ctx: ctx, tx: tx, ref: ref, gate: gate}, sess)
		err = fn(b)
		attempt(b)
		if err != nil {
			return err
		}
		if err := tx.Set(ref, b.S); err != nil {
			return err
		}
		if b.CountersDirty {
			if err := tx.Set(s.c.Collection(colCounters).Doc(countersDoc(b.S.RunID)), b.CountersValue()); err != nil {
				return err
			}
		}
		for p, r := range b.Ratings {
			if err := tx.Set(ref.Collection(subRatings).Doc(p), r); err != nil {
				return err
			}
		}
		for k, a := range b.Answer {
			if err := tx.Set(ref.Collection(subAnswers).Doc(k), a); err != nil {
				return err
			}
		}
		for _, e := range b.Evs {
			// Create, not Set: events are append-only.
			if err := tx.Create(ref.Collection(subEvents).Doc(store.EventID(e.EventNo)), e); err != nil {
				return err
			}
		}
		return nil
	})
}

func (s *Store) Events(ctx context.Context, id string) ([]*domain.Event, error) {
	docs, err := s.c.Collection(colSessions).Doc(id).Collection(subEvents).OrderBy("eventNo", firestore.Asc).Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}
	out := make([]*domain.Event, 0, len(docs))
	for _, d := range docs {
		var e domain.Event
		if err := d.DataTo(&e); err != nil {
			return nil, err
		}
		out = append(out, &e)
	}
	return out, nil
}

func (s *Store) Ratings(ctx context.Context, id string) ([]*domain.Rating, error) {
	docs, err := s.c.Collection(colSessions).Doc(id).Collection(subRatings).Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}
	var out []*domain.Rating
	for _, d := range docs {
		var r domain.Rating
		if err := d.DataTo(&r); err != nil {
			return nil, err
		}
		out = append(out, &r)
	}
	return out, nil
}

func (s *Store) Answers(ctx context.Context, id string) ([]*domain.AnswerDoc, error) {
	docs, err := s.c.Collection(colSessions).Doc(id).Collection(subAnswers).Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}
	var out []*domain.AnswerDoc
	for _, d := range docs {
		var a domain.AnswerDoc
		if err := d.DataTo(&a); err != nil {
			return nil, err
		}
		out = append(out, &a)
	}
	return out, nil
}

func (s *Store) Counters(ctx context.Context, runID string) (*domain.Counters, error) {
	d, err := s.c.Collection(colCounters).Doc(countersDoc(runID)).Get(ctx)
	if notFound(err) {
		return &domain.Counters{}, nil
	}
	if err != nil {
		return nil, err
	}
	var c domain.Counters
	return &c, d.DataTo(&c)
}

func (s *Store) PutConfig(ctx context.Context, c *domain.ConfigRecord) error {
	_, err := s.c.Collection(colConfigs).Doc(c.Version).Set(ctx, c)
	return err
}

func (s *Store) GetConfig(ctx context.Context, v string) (*domain.ConfigRecord, error) {
	d, err := s.c.Collection(colConfigs).Doc(v).Get(ctx)
	if notFound(err) {
		return nil, store.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	var c domain.ConfigRecord
	return &c, d.DataTo(&c)
}

func (s *Store) ListConfigs(ctx context.Context) ([]*domain.ConfigRecord, error) {
	docs, err := s.c.Collection(colConfigs).Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}
	var out []*domain.ConfigRecord
	for _, d := range docs {
		var c domain.ConfigRecord
		if err := d.DataTo(&c); err != nil {
			return nil, err
		}
		out = append(out, &c)
	}
	return out, nil
}
