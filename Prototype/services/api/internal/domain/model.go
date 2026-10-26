// Package domain holds the session plan compiler, the state machine with its
// timing rules, lazy timeout evaluation and group allocation (sections 5 and 6
// of "02 - Architecture.md"), and the records they read and write (7.1).
package domain

import "time"

// Session statuses (7.1, status mapping).
const (
	StatusOpened    = "Opened"
	StatusStarted   = "Started"
	StatusOngoing   = "Ongoing"
	StatusFinished  = "Finished"
	StatusAbandoned = "Abandoned"
	StatusBlocked   = "Blocked"
)

// Flags shown in the admin view (section 6).
const (
	FlagPossibleDuplicate    = "possible_duplicate"
	FlagHoneypot             = "honeypot"
	FlagAudioCheckWrong      = "audio_check_wrong"
	FlagAudioProblem         = "audio_problem"
	FlagManyInvalidEvents    = "many_invalid_events"
	FlagRateLimited          = "rate_limited"
	FlagTabHiddenDuringAudio = "tab_hidden_during_audio"
	FlagReloaded             = "reloaded"
	FlagForcedAllocation     = "forced_allocation" // group fixed by a test link, not by round robin
)

// Run is runs/{runId}.
type Run struct {
	ID            string    `firestore:"-" json:"id"`
	Code          string    `firestore:"code" json:"code"`
	Name          string    `firestore:"name" json:"name"`
	Source        string    `firestore:"source" json:"source"`
	ConfigVersion string    `firestore:"configVersion" json:"configVersion"`
	Status        string    `firestore:"status" json:"status"` // "open" | "closed"
	CreatedAt     time.Time `firestore:"createdAt" json:"createdAt"`
	CreatedFromIP string    `firestore:"createdFromIp" json:"createdFromIp"`
	// TestOverrides lets a session of this run fix its group and denial slide
	// from the link (?group=checker|neutral&denial=<slide id>). For test runs
	// only (Oliver, 2026-10-25): such sessions are flagged forced_allocation
	// and never touch the run's round-robin counters.
	TestOverrides bool `firestore:"testOverrides" json:"testOverrides"`
}

const (
	RunOpen   = "open"
	RunClosed = "closed"
)

// Deadlines of the current step. All times are server times.
type Deadlines struct {
	StepStartedAt  *time.Time `firestore:"stepStartedAt,omitempty" json:"stepStartedAt,omitempty"`
	ClipStartedAt  *time.Time `firestore:"clipStartedAt,omitempty" json:"clipStartedAt,omitempty"`
	RatingOpensAt  *time.Time `firestore:"ratingOpensAt,omitempty" json:"ratingOpensAt,omitempty"`
	RatingClosesAt *time.Time `firestore:"ratingClosesAt,omitempty" json:"ratingClosesAt,omitempty"`
	ActionClosesAt *time.Time `firestore:"actionClosesAt,omitempty" json:"actionClosesAt,omitempty"`
	FormClosesAt   *time.Time `firestore:"formClosesAt,omitempty" json:"formClosesAt,omitempty"`
}

// AudioCheckResult is the session's audio check (2.4).
type AudioCheckResult struct {
	ChoiceID string `firestore:"choiceId" json:"choiceId"`
	Correct  *bool  `firestore:"correct" json:"correct"`
	Plays    int    `firestore:"plays" json:"plays"`
}

// Screen is the screen size sent at session creation.
type Screen struct {
	W int `firestore:"w" json:"w"`
	H int `firestore:"h" json:"h"`
}

// LastResult keeps the response of the last consumed clientSeq so an identical
// retry can be answered without a query.
type LastResult struct {
	ClientSeq int64  `firestore:"clientSeq" json:"clientSeq"`
	BodyHash  string `firestore:"bodyHash" json:"bodyHash"`
	EventNo   int64  `firestore:"eventNo" json:"eventNo"`
	Status    int    `firestore:"status" json:"status"`
	Response  string `firestore:"response" json:"response"`
}

// Session is sessions/{sessionId}. Fields beyond section 7.1 are marked.
type Session struct {
	ID            string `firestore:"-" json:"sessionId"`
	RunID         string `firestore:"runId" json:"runId"`
	RunCode       string `firestore:"runCode" json:"runCode"`
	Source        string `firestore:"source" json:"source"`
	ConfigVersion string `firestore:"configVersion" json:"configVersion"`
	Group         string `firestore:"group" json:"group"`
	DenialSlide   string `firestore:"denialSlide" json:"denialSlide"`
	// ForcedGroup and ForcedDenialSlide come from a test link on a run with
	// TestOverrides; applied at consent instead of the round robin.
	ForcedGroup       string    `firestore:"forcedGroup,omitempty" json:"forcedGroup,omitempty"`
	ForcedDenialSlide string    `firestore:"forcedDenialSlide,omitempty" json:"forcedDenialSlide,omitempty"`
	Status            string    `firestore:"status" json:"status"`
	CurrentState      string    `firestore:"currentState" json:"currentState"`
	Plan              []string  `firestore:"plan" json:"plan"`
	ProgressIndex     int       `firestore:"progressIndex" json:"progressIndex"`
	ProgressTotal     int       `firestore:"progressTotal" json:"progressTotal"`
	LastClientSeq     int64     `firestore:"lastClientSeq" json:"lastClientSeq"`
	LastEventNo       int64     `firestore:"lastEventNo" json:"lastEventNo"`
	Deadlines         Deadlines `firestore:"deadlines" json:"deadlines"`

	OpenedAt    time.Time  `firestore:"openedAt" json:"openedAt"`
	ConsentAt   *time.Time `firestore:"consentAt" json:"consentAt"`
	StartedAt   *time.Time `firestore:"startedAt" json:"startedAt"`
	OngoingAt   *time.Time `firestore:"ongoingAt" json:"ongoingAt"`
	FinishedAt  *time.Time `firestore:"finishedAt" json:"finishedAt"`
	LastEventAt time.Time  `firestore:"lastEventAt" json:"lastEventAt"`
	UpdatedAt   time.Time  `firestore:"updatedAt" json:"updatedAt"`

	AudioCheck  *AudioCheckResult `firestore:"audioCheck" json:"audioCheck"`
	ReloadCount int               `firestore:"reloadCount" json:"reloadCount"`
	TokenHash   string            `firestore:"tokenHash" json:"-"`
	IP          string            `firestore:"ip" json:"ip"`
	UserAgent   string            `firestore:"userAgent" json:"userAgent"`
	UAFamily    string            `firestore:"uaFamily" json:"uaFamily"`
	Flags       []string          `firestore:"flags" json:"flags"`

	// Beyond 7.1:
	Screen          Screen            `firestore:"screen" json:"screen"`                   // from POST sessions
	TZ              string            `firestore:"tz" json:"tz"`                           // from POST sessions
	RestartAllowed  bool              `firestore:"restartAllowed" json:"restartAllowed"`   // one scene_started repeat after a reload (2.5)
	ActionChoices   map[string]string `firestore:"actionChoices" json:"actionChoices"`     // flagged slide id -> actionId ("" on timeout); copy of answers/actions for the question step
	InvalidCount    int               `firestore:"invalidCount" json:"invalidCount"`       // for many_invalid_events
	VisibilityCount int               `firestore:"visibilityCount" json:"visibilityCount"` // max 30 (5.3)
	LastResult      *LastResult       `firestore:"lastResult" json:"-"`
}

// HasFlag reports whether the session carries a flag.
func (s *Session) HasFlag(f string) bool {
	for _, x := range s.Flags {
		if x == f {
			return true
		}
	}
	return false
}

// AddFlag adds a flag once.
func (s *Session) AddFlag(f string) {
	if !s.HasFlag(f) {
		s.Flags = append(s.Flags, f)
	}
}

// Final reports whether the session can take no more events.
func (s *Session) Final() bool {
	return s.Status == StatusFinished || s.Status == StatusBlocked
}

// Event is sessions/{id}/events/{eventNo}: append-only.
type Event struct {
	EventNo     int64          `firestore:"eventNo" json:"eventNo"`
	ClientSeq   *int64         `firestore:"clientSeq" json:"clientSeq"`
	Type        string         `firestore:"type" json:"type"`
	Source      string         `firestore:"source" json:"source"` // "client" | "server"
	ServerTs    time.Time      `firestore:"serverTs" json:"serverTs"`
	ClientTs    any            `firestore:"clientTs" json:"clientTs"`
	Payload     map[string]any `firestore:"payload" json:"payload"`
	StateBefore string         `firestore:"stateBefore" json:"stateBefore"`
	StateAfter  string         `firestore:"stateAfter" json:"stateAfter"`
	Valid       bool           `firestore:"valid" json:"valid"`
	Reason      string         `firestore:"reason" json:"reason"`
	// Beyond 7.1, for retries: whether this event consumed its clientSeq, the
	// hash of the request body, and the response that was returned.
	Consumed       bool   `firestore:"consumed" json:"consumed"`
	BodyHash       string `firestore:"bodyHash" json:"-"`
	ResponseStatus int    `firestore:"responseStatus" json:"-"`
	Response       string `firestore:"response" json:"-"`
}

// Rating phases.
const (
	PhaseBaseline    = "baseline"
	PhaseAfterSlide  = "afterSlide"
	PhaseAfterAnswer = "afterAnswer"
)

// Rating item statuses. "missing" is used for an item left unclicked in an
// on-time submission.
const (
	ItemAnswered = "answered"
	ItemTimedOut = "timedOut"
	ItemMissing  = "missing"
)

// Rating and form statuses.
const (
	RecordComplete = "complete"
	RecordPartial  = "partial"
	RecordTimeout  = "timeout"
)

// Rating is sessions/{id}/ratings/{point}. Item values are keyed by rating
// item id (trust, competence, benevolence, integrity); nil when not clicked.
type Rating struct {
	Point       string            `firestore:"point" json:"point"`
	SlideID     string            `firestore:"slideId" json:"slideId"`
	Phase       string            `firestore:"phase" json:"phase"`
	Values      map[string]*int64 `firestore:"values" json:"values"`
	ItemStatus  map[string]string `firestore:"itemStatus" json:"itemStatus"`
	Status      string            `firestore:"status" json:"status"`
	LateEventNo *int64            `firestore:"lateEventNo" json:"lateEventNo"`
	OpenedAt    *time.Time        `firestore:"openedAt" json:"openedAt"`
	SubmittedAt *time.Time        `firestore:"submittedAt" json:"submittedAt"`
	LatencyMs   *int64            `firestore:"latencyMs" json:"latencyMs"`
}

// Answer doc kinds.
const (
	AnswersActions    = "actions"
	AnswersBudget     = "budget"
	AnswersBackground = "background"
)

// ActionRecord is one flagged slide's entry in answers/actions.
type ActionRecord struct {
	ActionID    *string    `firestore:"actionId" json:"actionId"`
	TimedOut    bool       `firestore:"timedOut" json:"timedOut"`
	QuestionBy  string     `firestore:"questionBy" json:"questionBy"` // "chair" | "member"
	AnswerType  string     `firestore:"answerType" json:"answerType"` // "reason" | "denial"
	OpenedAt    *time.Time `firestore:"openedAt" json:"openedAt"`
	ChosenAt    *time.Time `firestore:"chosenAt" json:"chosenAt"`
	LateEventNo *int64     `firestore:"lateEventNo" json:"lateEventNo"`
}

// AnswerDoc is sessions/{id}/answers/{kind}. Actions is used by kind
// "actions"; the other fields by "budget" and "background".
type AnswerDoc struct {
	Kind        string                  `firestore:"kind" json:"kind"`
	Actions     map[string]ActionRecord `firestore:"actions,omitempty" json:"actions,omitempty"`
	Values      map[string]any          `firestore:"values,omitempty" json:"values,omitempty"`
	ItemStatus  map[string]string       `firestore:"itemStatus,omitempty" json:"itemStatus,omitempty"`
	Status      string                  `firestore:"status,omitempty" json:"status,omitempty"`
	HP          string                  `firestore:"hp,omitempty" json:"hp,omitempty"`
	OpenedAt    *time.Time              `firestore:"openedAt,omitempty" json:"openedAt,omitempty"`
	SubmittedAt *time.Time              `firestore:"submittedAt,omitempty" json:"submittedAt,omitempty"`
	LateEventNo *int64                  `firestore:"lateEventNo,omitempty" json:"lateEventNo,omitempty"`
}

// Counters is counters/allocation (5.5).
type Counters struct {
	Groups     map[string]int64            `firestore:"groups" json:"groups"`
	Denial     map[string]map[string]int64 `firestore:"denial" json:"denial"`
	LastDenial map[string]string           `firestore:"lastDenial" json:"lastDenial"`
}

// ConfigRecord is configs/{version}.
type ConfigRecord struct {
	Version       string     `firestore:"version" json:"version"`
	SHA256        string     `firestore:"sha256" json:"sha256"`
	GCSPath       string     `firestore:"gcsPath" json:"gcsPath"`
	ScenePackPath string     `firestore:"scenePackPath" json:"scenePackPath"`
	ManifestPath  string     `firestore:"manifestPath" json:"manifestPath"`
	Status        string     `firestore:"status" json:"status"` // draft | validated | locked
	UploadedAt    *time.Time `firestore:"uploadedAt" json:"uploadedAt"`
	LintReport    []string   `firestore:"lintReport" json:"lintReport"`
	// Beyond 7.1: the uploaded config JSON (DECISION in internal/configs).
	Body string `firestore:"body" json:"-"`
}

const (
	ConfigDraft     = "draft"
	ConfigValidated = "validated"
	ConfigLocked    = "locked"
)
