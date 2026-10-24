// Package events defines the event types the participant client sends
// (section 5.3), their payloads, and the request envelope.
package events

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

// Event types (5.3), plus the server-written synthetic timeouts (5.4), which
// reuse the client timeout type names with source "server".
const (
	DeviceChecked       = "device_checked"
	ConsentAccepted     = "consent_accepted"
	AudioCheckPlayed    = "audio_check_played"
	AudioCheckAnswered  = "audio_check_answered"
	SceneStarted        = "scene_started"
	MarkingShown        = "marking_shown"
	AudioEnded          = "audio_ended"
	AudioProblem        = "audio_problem"
	SetupDone           = "setup_done"
	RatingSubmitted     = "rating_submitted"
	RatingTimeout       = "rating_timeout"
	ActionChosen        = "action_chosen"
	ActionTimeout       = "action_timeout"
	BudgetSubmitted     = "budget_submitted"
	BudgetTimeout       = "budget_timeout"
	BackgroundSubmitted = "background_submitted"
	BackgroundTimeout   = "background_timeout"
	Finished            = "finished"
	Visibility          = "visibility"
)

// Source of an event.
const (
	SourceClient = "client"
	SourceServer = "server"
)

// Request is the body of POST /api/v1/sessions/{id}/events.
type Request struct {
	ClientSeq int64           `json:"clientSeq"`
	Type      string          `json:"type"`
	ClientTs  any             `json:"clientTs"`
	Payload   json.RawMessage `json:"payload"`
}

// Hash identifies the body for retry detection: the same clientSeq, type,
// clientTs and payload (compacted JSON) give the same hash.
func (r *Request) Hash() string {
	var p bytes.Buffer
	if len(r.Payload) > 0 {
		if err := json.Compact(&p, r.Payload); err != nil {
			p.Write(r.Payload)
		}
	}
	ts, _ := json.Marshal(r.ClientTs)
	h := sha256.New()
	enc := json.NewEncoder(h)
	_ = enc.Encode([]any{r.ClientSeq, r.Type, json.RawMessage(ts), p.String()})
	return hex.EncodeToString(h.Sum(nil))
}

// PayloadMap decodes the payload for storage; a missing or non-object payload
// gives an empty map (non-objects are stored under "_raw").
func (r *Request) PayloadMap() map[string]any {
	m := map[string]any{}
	if len(r.Payload) == 0 || string(r.Payload) == "null" {
		return m
	}
	if err := json.Unmarshal(r.Payload, &m); err != nil {
		return map[string]any{"_raw": string(r.Payload)}
	}
	return m
}

// Decode decodes the payload into v; a missing payload decodes as {}.
func (r *Request) Decode(v any) error {
	if len(r.Payload) == 0 || string(r.Payload) == "null" {
		return json.Unmarshal([]byte("{}"), v)
	}
	return json.Unmarshal(r.Payload, v)
}

type DeviceCheckedPayload struct {
	WidthPx  *float64 `json:"widthPx"`
	HeightPx *float64 `json:"heightPx"`
	Touch    bool     `json:"touch"`
}

type ConsentPayload struct {
	ConsentVersion string `json:"consentVersion"`
	Checkbox       bool   `json:"checkbox"`
}

type ChoicePayload struct {
	ChoiceID string `json:"choiceId"`
}

type ScenePayload struct {
	SceneID      string `json:"sceneId"`
	MarkingID    string `json:"markingId"`
	AfterProblem bool   `json:"afterProblem"`
	Kind         string `json:"kind"`
	AtMs         any    `json:"atMs"`
}

// RatingPayload holds the point and the raw item values; rating_submitted
// sends the items at the top level, rating_timeout under "partial".
type RatingPayload struct {
	Point   string                     `json:"point"`
	Partial map[string]json.RawMessage `json:"partial"`
}

type ActionPayload struct {
	SlideID  string  `json:"slideId"`
	ActionID *string `json:"actionId"`
}

type BudgetPayload struct {
	OptionID   json.RawMessage `json:"optionId"`
	Confidence json.RawMessage `json:"confidence"`
	Partial    *struct {
		OptionID   json.RawMessage `json:"optionId"`
		Confidence json.RawMessage `json:"confidence"`
	} `json:"partial"`
}

type BackgroundPayload struct {
	Items map[string]json.RawMessage `json:"items"`
	// Partial is { items: {...} } or the item map itself.
	Partial map[string]json.RawMessage `json:"partial"`
	HP      any                        `json:"hp"`
}

type VisibilityPayload struct {
	Hidden bool `json:"hidden"`
}
