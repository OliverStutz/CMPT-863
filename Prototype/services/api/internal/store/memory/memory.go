// Package memory is an in-process Store for tests and quick local runs. One
// mutex serialises every transaction, and writes are applied only when the
// transaction function succeeds, so the semantics of 5.5 and section 6 hold.
package memory

import (
	"context"
	"encoding/json"
	"sort"
	"sync"

	"study/api/internal/domain"
	"study/api/internal/store"
)

type sessionData struct {
	s       *domain.Session
	events  []*domain.Event
	ratings map[string]*domain.Rating
	answers map[string]*domain.AnswerDoc
}

// Store is the memory implementation.
type Store struct {
	mu       sync.Mutex
	runs     map[string]*domain.Run
	sessions map[string]*sessionData
	counters map[string]*domain.Counters // per run id (1.19 revised)
	configs  map[string]*domain.ConfigRecord
	deleted  map[string]*domain.Deleted // tombstones by session id
}

func New() *Store {
	return &Store{runs: map[string]*domain.Run{}, sessions: map[string]*sessionData{}, configs: map[string]*domain.ConfigRecord{},
		counters: map[string]*domain.Counters{}, deleted: map[string]*domain.Deleted{}}
}

// clone deep-copies through JSON so callers never share memory with the store.
func clone[T any](v *T) *T {
	if v == nil {
		return nil
	}
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	var out T
	if err := json.Unmarshal(b, &out); err != nil {
		panic(err)
	}
	return &out
}

// Session and Run hide their id from JSON-free firestore tags but keep it in
// JSON; tokenHash is json:"-", so it is copied by hand.
func cloneSession(s *domain.Session) *domain.Session {
	c := clone(s)
	c.TokenHash = s.TokenHash
	c.LastResult = clone(s.LastResult)
	return c
}

func cloneEvent(e *domain.Event) *domain.Event {
	c := clone(e)
	c.BodyHash, c.Response, c.ResponseStatus = e.BodyHash, e.Response, e.ResponseStatus
	return c
}

func cloneConfig(c *domain.ConfigRecord) *domain.ConfigRecord {
	x := clone(c)
	x.Body = c.Body
	return x
}

func (m *Store) CreateRun(_ context.Context, r *domain.Run) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, x := range m.runs {
		if x.Code == r.Code {
			return store.ErrConflict
		}
	}
	if _, ok := m.runs[r.ID]; ok {
		return store.ErrConflict
	}
	m.runs[r.ID] = clone(r)
	return nil
}

func (m *Store) GetRun(_ context.Context, id string) (*domain.Run, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	r, ok := m.runs[id]
	if !ok {
		return nil, store.ErrNotFound
	}
	return clone(r), nil
}

func (m *Store) GetRunByCode(_ context.Context, code string) (*domain.Run, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, r := range m.runs {
		if r.Code == code {
			return clone(r), nil
		}
	}
	return nil, store.ErrNotFound
}

func (m *Store) ListRuns(_ context.Context) ([]*domain.Run, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	var out []*domain.Run
	for _, r := range m.runs {
		out = append(out, clone(r))
	}
	sort.Slice(out, func(i, j int) bool { return out[i].CreatedAt.After(out[j].CreatedAt) })
	return out, nil
}

func (m *Store) UpdateRun(_ context.Context, r *domain.Run) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.runs[r.ID]; !ok {
		return store.ErrNotFound
	}
	m.runs[r.ID] = clone(r)
	return nil
}

func (m *Store) DeleteRun(_ context.Context, id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.runs[id]; !ok {
		return store.ErrNotFound
	}
	for _, sd := range m.sessions {
		if sd.s.RunID == id {
			return store.ErrConflict
		}
	}
	delete(m.runs, id)
	return nil
}

func (m *Store) CreateSession(_ context.Context, s *domain.Session) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.sessions[s.ID]; ok {
		return store.ErrConflict
	}
	m.sessions[s.ID] = &sessionData{s: cloneSession(s), ratings: map[string]*domain.Rating{}, answers: map[string]*domain.AnswerDoc{}}
	return nil
}

func (m *Store) GetSession(_ context.Context, id string) (*domain.Session, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	sd, ok := m.sessions[id]
	if !ok {
		return nil, store.ErrNotFound
	}
	return cloneSession(sd.s), nil
}

func (m *Store) ListSessions(_ context.Context, f store.SessionFilter) ([]*domain.Session, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	var out []*domain.Session
	for _, sd := range m.sessions {
		if f.Match(sd.s) {
			out = append(out, cloneSession(sd.s))
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].OpenedAt.After(out[j].OpenedAt) })
	return out, nil
}

type loader struct {
	m  *Store
	sd *sessionData
}

func (l loader) LoadCounters(runID string) (*domain.Counters, error) {
	return clone(l.m.counters[runID]), nil
}
func (l loader) LoadRating(p string) (*domain.Rating, error) {
	return clone(l.sd.ratings[p]), nil
}
func (l loader) LoadAnswers(k string) (*domain.AnswerDoc, error) {
	return clone(l.sd.answers[k]), nil
}
func (l loader) LoadConsumedEvent(seq int64) (*domain.Event, error) {
	for _, e := range l.sd.events {
		if e.Consumed && e.ClientSeq != nil && *e.ClientSeq == seq {
			return cloneEvent(e), nil
		}
	}
	return nil, nil
}

func (m *Store) RunTx(_ context.Context, id string, fn func(domain.Tx) error) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	sd, ok := m.sessions[id]
	if !ok {
		return store.ErrNotFound
	}
	b := store.NewBuffer(loader{m, sd}, cloneSession(sd.s))
	if err := fn(b); err != nil {
		return err
	}
	for _, e := range b.Evs {
		for _, x := range sd.events {
			if x.EventNo == e.EventNo {
				return store.ErrConflict // append-only
			}
		}
	}
	sd.s = cloneSession(b.S)
	if b.CountersDirty {
		m.counters[b.S.RunID] = clone(b.CountersValue())
	}
	for p, r := range b.Ratings {
		sd.ratings[p] = clone(r)
	}
	for k, a := range b.Answer {
		sd.answers[k] = clone(a)
	}
	for _, e := range b.Evs {
		sd.events = append(sd.events, cloneEvent(e))
	}
	return nil
}

func (m *Store) Events(_ context.Context, id string) ([]*domain.Event, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	sd, ok := m.sessions[id]
	if !ok {
		return nil, store.ErrNotFound
	}
	out := make([]*domain.Event, 0, len(sd.events))
	for _, e := range sd.events {
		out = append(out, cloneEvent(e))
	}
	sort.Slice(out, func(i, j int) bool { return out[i].EventNo < out[j].EventNo })
	return out, nil
}

func (m *Store) Ratings(_ context.Context, id string) ([]*domain.Rating, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	sd, ok := m.sessions[id]
	if !ok {
		return nil, store.ErrNotFound
	}
	var out []*domain.Rating
	for _, r := range sd.ratings {
		out = append(out, clone(r))
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Point < out[j].Point })
	return out, nil
}

func (m *Store) Answers(_ context.Context, id string) ([]*domain.AnswerDoc, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	sd, ok := m.sessions[id]
	if !ok {
		return nil, store.ErrNotFound
	}
	var out []*domain.AnswerDoc
	for _, a := range sd.answers {
		out = append(out, clone(a))
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Kind < out[j].Kind })
	return out, nil
}

func (m *Store) DeleteSession(_ context.Context, id string, t *domain.Deleted) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.sessions[id]; !ok {
		return store.ErrNotFound
	}
	delete(m.sessions, id) // the session and its events, ratings and answers
	d := clone(t)
	d.SessionID = id
	m.deleted[id] = d
	return nil
}

func (m *Store) ListDeleted(_ context.Context) ([]*domain.Deleted, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	out := []*domain.Deleted{}
	for _, d := range m.deleted {
		out = append(out, clone(d))
	}
	sort.Slice(out, func(i, j int) bool { return out[i].DeletedAt.After(out[j].DeletedAt) })
	return out, nil
}

func (m *Store) Counters(_ context.Context, runID string) (*domain.Counters, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.counters[runID] == nil {
		return &domain.Counters{}, nil
	}
	return clone(m.counters[runID]), nil
}

func (m *Store) PutConfig(_ context.Context, c *domain.ConfigRecord) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.configs[c.Version] = cloneConfig(c)
	return nil
}

func (m *Store) GetConfig(_ context.Context, v string) (*domain.ConfigRecord, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	c, ok := m.configs[v]
	if !ok {
		return nil, store.ErrNotFound
	}
	return cloneConfig(c), nil
}

func (m *Store) ListConfigs(_ context.Context) ([]*domain.ConfigRecord, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	var out []*domain.ConfigRecord
	for _, c := range m.configs {
		out = append(out, cloneConfig(c))
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Version < out[j].Version })
	return out, nil
}

func (m *Store) Close() error { return nil }
