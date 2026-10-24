package domain

import "net/http"

// Tx is one store transaction on one session (section 6: read session, check
// clientSeq and expected step, write the event under the next eventNo, update
// session), plus the allocation counter document (5.5). Reads go to the store
// or to writes buffered earlier in the same transaction; all writes are
// applied together when the callback returns nil, and dropped otherwise.
// Implementations retry the callback on contention, so it must not have side
// effects outside the Tx.
type Tx interface {
	// Session is the session loaded in this transaction. Changes to it are
	// written on commit.
	Session() *Session
	Counters() (*Counters, error)
	PutCounters(*Counters)
	// Rating returns nil, nil when the rating does not exist.
	Rating(point string) (*Rating, error)
	PutRating(*Rating)
	// Answers returns nil, nil when the document does not exist.
	Answers(kind string) (*AnswerDoc, error)
	PutAnswers(*AnswerDoc)
	PutEvent(*Event)
	// ConsumedEvent finds the event that consumed a clientSeq, or nil.
	ConsumedEvent(clientSeq int64) (*Event, error)
}

// Error codes (8.3).
const (
	ErrBadRequest           = "bad_request"
	ErrBadPayload           = "bad_payload"
	ErrNoSession            = "no_session"
	ErrBadToken             = "bad_token"
	ErrAdminLoginRequired   = "admin_login_required"
	ErrBadPassword          = "bad_password"
	ErrUnknownRun           = "unknown_run"
	ErrRunClosed            = "run_closed"
	ErrWrongState           = "wrong_state"
	ErrSeqGap               = "seq_gap"
	ErrDuplicateSeqMismatch = "duplicate_seq_mismatch"
	ErrLate                 = "late"
	ErrTooEarly             = "too_early"
	ErrOutOfRange           = "out_of_range"
	ErrLintFailed           = "lint_failed"
	ErrMissingAudio         = "missing_audio"
	ErrRateLimited          = "rate_limited"
	ErrInternal             = "internal"
	ErrUnavailable          = "unavailable"
	// Not in 8.3; used by the admin API for a missing record and a conflict.
	ErrNotFound = "not_found"
	ErrConflict = "conflict"
)

// HTTPStatus maps an error code to its HTTP status (8.3).
func HTTPStatus(code string) int {
	switch code {
	case ErrBadRequest, ErrBadPayload:
		return http.StatusBadRequest
	case ErrNoSession, ErrBadToken, ErrAdminLoginRequired, ErrBadPassword:
		return http.StatusUnauthorized
	case ErrUnknownRun, ErrNotFound:
		return http.StatusNotFound
	case ErrRunClosed:
		return http.StatusGone
	case ErrWrongState, ErrSeqGap, ErrDuplicateSeqMismatch, ErrLate, ErrConflict:
		return http.StatusConflict
	case ErrTooEarly, ErrOutOfRange, ErrLintFailed, ErrMissingAudio:
		return http.StatusUnprocessableEntity
	case ErrRateLimited:
		return http.StatusTooManyRequests
	case ErrUnavailable:
		return http.StatusServiceUnavailable
	}
	return http.StatusInternalServerError
}

// APIError is the body of the error format (8.3).
type APIError struct {
	Code              string `json:"code"`
	Message           string `json:"message"`
	Retryable         bool   `json:"retryable"`
	RetryInMs         *int64 `json:"retryInMs,omitempty"`
	ExpectedClientSeq *int64 `json:"expectedClientSeq,omitempty"`
	Next              *Step  `json:"next,omitempty"`
}

// ErrorBody wraps an APIError as { "error": ... }.
type ErrorBody struct {
	Error *APIError `json:"error"`
}
