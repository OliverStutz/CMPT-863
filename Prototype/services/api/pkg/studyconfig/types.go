// Package studyconfig mirrors the StudyConfig types of
// "02 - Architecture.md" section 4.1 and content/schema/study-config.schema.json.
//
// Discriminated unions of the TypeScript types (Block, Scene, CheckerPanel,
// Answers, BackgroundItem) are flattened into one Go struct each, with the
// discriminator field (kind, status or mode) deciding which other fields are
// meaningful. Keys that start with "_" are author notes and are ignored when
// decoding. Unknown keys are rejected by the JSON Schema validation in Load.
package studyconfig

import (
	"encoding/json"
	"strings"
)

// Ms is a duration in milliseconds.
type Ms = int64

// Group is the experimental group of a session.
type Group string

const (
	GroupChecker Group = "checker"
	GroupNeutral Group = "neutral"
)

// Groups in the fixed order used by allocation (section 5.5).
var Groups = []Group{GroupChecker, GroupNeutral}

// RichText is plain text or a list of spans. It is kept as raw JSON so that it
// round-trips exactly; PlainText returns the text without span ids.
type RichText json.RawMessage

func (r RichText) MarshalJSON() ([]byte, error) {
	if len(r) == 0 {
		return []byte("null"), nil
	}
	return []byte(r), nil
}

func (r *RichText) UnmarshalJSON(b []byte) error {
	*r = append((*r)[:0], b...)
	return nil
}

// Spans returns the spans of a span list, or nil for plain text.
func (r RichText) Spans() []Span {
	var spans []Span
	if json.Unmarshal(r, &spans) == nil {
		return spans
	}
	return nil
}

// PlainText returns the text of the rich text.
func (r RichText) PlainText() string {
	var s string
	if json.Unmarshal(r, &s) == nil {
		return s
	}
	var parts []string
	for _, sp := range r.Spans() {
		parts = append(parts, sp.Text)
	}
	return strings.Join(parts, "")
}

type Span struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

type Choice struct {
	ID    string `json:"id"`
	Label string `json:"label"`
}

type StudyConfig struct {
	Version string       `json:"version"`
	Locale  string       `json:"locale"`
	Timing  Timing       `json:"timing"`
	Voices  Voices       `json:"voices"`
	Options StudyOptions `json:"options"`
	Labels  Labels       `json:"labels"`
	Colors  Colors       `json:"colors"`
	Consent ConsentSpec  `json:"consent"`
	Setup   SetupSpec    `json:"setup"`
	Rating  RatingSpec   `json:"rating"`
	Actions ActionSpec   `json:"actions"`
	Scenes  []Scene      `json:"scenes"`
}

type Timing struct {
	PostAudioViewMs  Ms `json:"postAudioViewMs"`
	PostAnswerViewMs Ms `json:"postAnswerViewMs"`
	RatingWindowMs   Ms `json:"ratingWindowMs"`
	ActionWindowMs   Ms `json:"actionWindowMs"`
	FormWindowMs     Ms `json:"formWindowMs"`
	// BackgroundWindowMs is the time limit of the background questions screen;
	// 0 means none (Oliver, 2026-10-25: 30 s was too short to fill in the
	// information). nil means formWindowMs, as before.
	BackgroundWindowMs  *Ms `json:"backgroundWindowMs,omitempty"`
	LinePauseMs         Ms  `json:"linePauseMs"`
	AudioEndToleranceMs Ms  `json:"audioEndToleranceMs"`
	LatencyAllowanceMs  Ms  `json:"latencyAllowanceMs"`
	AbandonAfterMs      Ms  `json:"abandonAfterMs"`
	MsPerWordEstimate   *Ms `json:"msPerWordEstimate,omitempty"`
}

// BackgroundWindow returns the background screen's time limit; 0 is untimed.
func (t Timing) BackgroundWindow() Ms {
	if t.BackgroundWindowMs == nil {
		return t.FormWindowMs
	}
	return *t.BackgroundWindowMs
}

type Voices struct {
	CISO   VoiceSpec `json:"ciso"`
	Chair  VoiceSpec `json:"chair"`
	Member VoiceSpec `json:"member"`
}

type VoiceSpec struct {
	Provider     string   `json:"provider"`
	Name         string   `json:"name"`
	LanguageCode string   `json:"languageCode"`
	SpeakingRate *float64 `json:"speakingRate,omitempty"`
	Pitch        *float64 `json:"pitch,omitempty"`
}

type StudyOptions struct {
	DeviceGate DeviceGate `json:"deviceGate"`
}

type DeviceGate struct {
	MinWidthPx  int  `json:"minWidthPx"`
	AllowPhones bool `json:"allowPhones"`
}

type Labels struct {
	PanelTitle      string `json:"panelTitle"`
	NoIssues        string `json:"noIssues"`
	PanelIdle       string `json:"panelIdle"`
	RatingSubmit    string `json:"ratingSubmit"`
	TimeLeft        string `json:"timeLeft"`
	Continue        string `json:"continue"`
	PlayTestSound   string `json:"playTestSound"`
	ReplayTestSound string `json:"replayTestSound"`
	StartMeeting    string `json:"startMeeting"`
}

type Colors struct {
	Flag     string `json:"flag"`
	NoIssues string `json:"noIssues"`
}

type ConsentSpec struct {
	Version       string `json:"version"`
	PdfAssetKey   string `json:"pdfAssetKey"`
	CheckboxLabel string `json:"checkboxLabel"`
	AcceptLabel   string `json:"acceptLabel"`
}

// Voice ids of AudioLine.voice.
const (
	VoiceCISO   = "ciso"
	VoiceChair  = "chair"
	VoiceMember = "member"
)

type AudioLine struct {
	ID         string        `json:"id"`
	Voice      string        `json:"voice"`
	Text       string        `json:"text"`
	SSML       string        `json:"ssml,omitempty"`
	AudioKey   string        `json:"audioKey,omitempty"`
	DurationMs *Ms           `json:"durationMs,omitempty"`
	Marks      map[string]Ms `json:"marks,omitempty"`
}

type SetupSpec struct {
	Instruction  string       `json:"instruction"`
	AudioCheck   AudioCheck   `json:"audioCheck"`
	ExampleScene ExampleScene `json:"exampleScene"`
	PreStartText PreStartText `json:"preStartText"`
}

type ExampleScene struct {
	Template string      `json:"template,omitempty"`
	Layout   []Block     `json:"layout"`
	Lines    []AudioLine `json:"lines"`
}

type PreStartText struct {
	Both    []string `json:"both"`
	Checker []string `json:"checker"`
}

type AudioCheck struct {
	Clip            AudioLine `json:"clip"`
	Prompt          string    `json:"prompt"`
	Options         []Choice  `json:"options"`
	CorrectOptionID string    `json:"correctOptionId"`
}

// Block kinds.
const (
	BlockTitle  = "title"
	BlockBullet = "bullet"
	BlockNumber = "number"
	BlockImage  = "image"
	BlockChart  = "chart"
)

// Block is the flattened union of the five block kinds.
type Block struct {
	Kind      string     `json:"kind"`
	ID        string     `json:"id"`
	Text      RichText   `json:"text,omitempty"`
	Value     RichText   `json:"value,omitempty"`
	Caption   RichText   `json:"caption,omitempty"`
	AssetKey  string     `json:"assetKey,omitempty"`
	Alt       string     `json:"alt,omitempty"`
	ChartType string     `json:"chartType,omitempty"`
	Data      *ChartData `json:"data,omitempty"`
}

type ChartData struct {
	Axis    ChartAxis     `json:"axis"`
	Series  []ChartSeries `json:"series"`
	Markers []ChartMarker `json:"markers,omitempty"`
}

type ChartAxis struct {
	Min   float64 `json:"min"`
	Max   float64 `json:"max"`
	Label string  `json:"label,omitempty"`
}

type ChartSeries struct {
	ID     string    `json:"id"`
	Label  string    `json:"label"`
	Value  *float64  `json:"value,omitempty"`
	Start  *float64  `json:"start,omitempty"`
	End    *float64  `json:"end,omitempty"`
	Points []float64 `json:"points,omitempty"`
}

type ChartMarker struct {
	ID    string  `json:"id"`
	At    float64 `json:"at"`
	Label string  `json:"label"`
}

type Marking struct {
	ID           string      `json:"id"`
	TargetIDs    []string    `json:"targetIds"`
	BBox         *[4]float64 `json:"bbox,omitempty"`
	AppearAtMark string      `json:"appearAtMark,omitempty"`
	AppearAtMs   *Ms         `json:"appearAtMs,omitempty"`
	AnchorText   string      `json:"anchorText,omitempty"`
	Color        string      `json:"color,omitempty"`
	Shape        string      `json:"shape"`
	Label        string      `json:"label"`
	Detail       string      `json:"detail"`
	VLMRun       string      `json:"vlmRun,omitempty"`
}

// Checker panel statuses.
const (
	PanelNoIssues = "noIssues"
	PanelFlagged  = "flagged"
)

// CheckerPanel is the flattened union of the two panel statuses.
type CheckerPanel struct {
	Status       string    `json:"status"`
	AppearAtMark string    `json:"appearAtMark,omitempty"`
	AppearAtMs   *Ms       `json:"appearAtMs,omitempty"`
	Text         string    `json:"text,omitempty"`
	Markings     []Marking `json:"markings,omitempty"`
}

// Answer modes.
const (
	AnswersWhole  = "whole"
	AnswersJoined = "joined"
)

type Answers struct {
	Mode              string     `json:"mode"`
	Reason            *AudioLine `json:"reason,omitempty"`
	Denial            *AudioLine `json:"denial,omitempty"`
	Opening           *AudioLine `json:"opening,omitempty"`
	ReasonBody        *AudioLine `json:"reasonBody,omitempty"`
	Closing           *AudioLine `json:"closing,omitempty"`
	PadDenialToReason *bool      `json:"padDenialToReason,omitempty"`
}

type ParticipantQuestions struct {
	Explain   AudioLine `json:"explain"`
	Challenge AudioLine `json:"challenge"`
}

type ProfileCard struct {
	Name           string   `json:"name"`
	Role           string   `json:"role"`
	Lines          []string `json:"lines"`
	PortraitKey    string   `json:"portraitKey,omitempty"`
	ShowFromLineID string   `json:"showFromLineId"`
}

// Scene kinds.
const (
	SceneOpening    = "opening"
	SceneSlide      = "slide"
	SceneRating     = "rating"
	SceneAction     = "action"
	SceneQuestion   = "question"
	SceneAnswer     = "answer"
	SceneBudget     = "budget"
	SceneBackground = "background"
	SceneDebrief    = "debrief"
)

// Scene is the flattened union of the nine scene kinds.
type Scene struct {
	Kind string `json:"kind"`
	ID   string `json:"id"`

	// opening, slide
	Template string  `json:"template,omitempty"`
	Layout   []Block `json:"layout,omitempty"`

	// opening
	ProfileCard *ProfileCard `json:"profileCard,omitempty"`
	Lines       []AudioLine  `json:"lines,omitempty"`

	// slide
	Narration            *AudioLine            `json:"narration,omitempty"`
	Checker              *CheckerPanel         `json:"checker,omitempty"`
	RoutineQuestion      *AudioLine            `json:"routineQuestion,omitempty"`
	ParticipantQuestions *ParticipantQuestions `json:"participantQuestions,omitempty"`
	Answers              *Answers              `json:"answers,omitempty"`

	// rating
	Point string `json:"point,omitempty"`
	After string `json:"after,omitempty"`

	// action, question, answer
	ForSlide string `json:"forSlide,omitempty"`

	// budget
	Prompt     string      `json:"prompt,omitempty"`
	Options    []Choice    `json:"options,omitempty"`
	Confidence *Confidence `json:"confidence,omitempty"`

	// background
	Intro               string           `json:"intro,omitempty"`
	PreferNotToSayLabel string           `json:"preferNotToSayLabel,omitempty"`
	Items               []BackgroundItem `json:"items,omitempty"`

	// debrief
	Paragraphs       *DebriefParagraphs `json:"paragraphs,omitempty"`
	DoNotDiscuss     string             `json:"doNotDiscuss,omitempty"`
	SessionCodeLabel string             `json:"sessionCodeLabel,omitempty"`
}

// Flagged reports whether a slide scene is flagged (checker.status "flagged").
func (s *Scene) Flagged() bool {
	return s.Kind == SceneSlide && s.Checker != nil && s.Checker.Status == PanelFlagged
}

type Confidence struct {
	Label string    `json:"label"`
	Scale ScaleSpec `json:"scale"`
}

type DebriefParagraphs struct {
	Both    []string `json:"both"`
	Checker []string `json:"checker"`
	Neutral []string `json:"neutral"`
	End     []string `json:"end"`
}

// Background item kinds.
const (
	ItemSingle   = "single"
	ItemScale    = "scale"
	ItemCountry  = "country"
	ItemHoneypot = "honeypot"
)

type BackgroundItem struct {
	Kind          string     `json:"kind"`
	ID            string     `json:"id"`
	Label         string     `json:"label"`
	Options       []Choice   `json:"options,omitempty"`
	OtherOptionID string     `json:"otherOptionId,omitempty"`
	Scale         *ScaleSpec `json:"scale,omitempty"`
}

type RatingSpec struct {
	Instruction string       `json:"instruction"`
	Items       []RatingItem `json:"items"`
}

type RatingItem struct {
	ID    string    `json:"id"`
	Label string    `json:"label,omitempty"`
	Scale ScaleSpec `json:"scale"`
}

type ScaleSpec struct {
	Min         float64           `json:"min"`
	Max         float64           `json:"max"`
	MinLabel    string            `json:"minLabel"`
	MaxLabel    string            `json:"maxLabel"`
	PointLabels map[string]string `json:"pointLabels,omitempty"`
}

type ActionSpec struct {
	Checker ActionGroupSpec `json:"checker"`
	Neutral ActionGroupSpec `json:"neutral"`
}

type ActionGroupSpec struct {
	Prompt  string   `json:"prompt"`
	Note    string   `json:"note,omitempty"`
	Options []Choice `json:"options"`
}

// ForGroup returns the action options of a group.
func (a ActionSpec) ForGroup(g Group) ActionGroupSpec {
	if g == GroupChecker {
		return a.Checker
	}
	return a.Neutral
}

// SceneByID returns the scene with the given id, or nil.
func (c *StudyConfig) SceneByID(id string) *Scene {
	for i := range c.Scenes {
		if c.Scenes[i].ID == id {
			return &c.Scenes[i]
		}
	}
	return nil
}

// SceneOfKind returns the first scene of a kind, or nil.
func (c *StudyConfig) SceneOfKind(kind string) *Scene {
	for i := range c.Scenes {
		if c.Scenes[i].Kind == kind {
			return &c.Scenes[i]
		}
	}
	return nil
}

// SceneFor returns the scene of a kind that names forSlide, or nil.
func (c *StudyConfig) SceneFor(kind, forSlide string) *Scene {
	for i := range c.Scenes {
		if c.Scenes[i].Kind == kind && c.Scenes[i].ForSlide == forSlide {
			return &c.Scenes[i]
		}
	}
	return nil
}

// RatingAfter returns the rating scene that follows scene id after, or nil.
func (c *StudyConfig) RatingAfter(after string) *Scene {
	for i := range c.Scenes {
		if c.Scenes[i].Kind == SceneRating && c.Scenes[i].After == after {
			return &c.Scenes[i]
		}
	}
	return nil
}

// Slides returns the slide scenes in config order.
func (c *StudyConfig) Slides() []*Scene {
	var out []*Scene
	for i := range c.Scenes {
		if c.Scenes[i].Kind == SceneSlide {
			out = append(out, &c.Scenes[i])
		}
	}
	return out
}

// FlaggedSlideIDs returns the ids of the flagged slides in slide order.
func (c *StudyConfig) FlaggedSlideIDs() []string {
	var out []string
	for _, s := range c.Slides() {
		if s.Flagged() {
			out = append(out, s.ID)
		}
	}
	return out
}
