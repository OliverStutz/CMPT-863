// Package testutil builds fast test configs and sessions for the tests.
package testutil

import (
	"encoding/json"
	"fmt"
	"time"

	"study/api/internal/domain"
	sc "study/api/pkg/studyconfig"
)

func rt(s string) sc.RichText {
	b, _ := json.Marshal(s)
	return sc.RichText(b)
}

func line(id, voice, text string) sc.AudioLine { return sc.AudioLine{ID: id, Voice: voice, Text: text} }

func ms(v int64) *int64 { return &v }

// FastConfig builds a config with n slides, of which the 1-based positions in
// flagged carry a flag, and short timings (section 15's fast test config).
// Flagged slides alternate between "whole" and "joined" answers.
func FastConfig(version string, n int, flagged ...int) *sc.StudyConfig {
	isFlagged := map[int]bool{}
	for _, f := range flagged {
		isFlagged[f] = true
	}
	scale := func(min, max float64) sc.ScaleSpec {
		return sc.ScaleSpec{Min: min, Max: max, MinLabel: "low", MaxLabel: "high"}
	}
	c := &sc.StudyConfig{
		Version: version, Locale: "en",
		Timing: sc.Timing{PostAudioViewMs: 500, PostAnswerViewMs: 500, RatingWindowMs: 2000, ActionWindowMs: 1500,
			FormWindowMs: 3000, LinePauseMs: 100, AudioEndToleranceMs: 100, LatencyAllowanceMs: 300,
			AbandonAfterMs: 60000, MsPerWordEstimate: ms(100)},
		Voices: sc.Voices{
			CISO:   sc.VoiceSpec{Provider: "google-tts", Name: "a", LanguageCode: "en-US"},
			Chair:  sc.VoiceSpec{Provider: "google-tts", Name: "b", LanguageCode: "en-US"},
			Member: sc.VoiceSpec{Provider: "google-tts", Name: "c", LanguageCode: "en-US"},
		},
		Options: sc.StudyOptions{DeviceGate: sc.DeviceGate{MinWidthPx: 1024}},
		Labels: sc.Labels{PanelTitle: "Slide checker", NoIssues: "No issues found", RatingSubmit: "Next",
			TimeLeft: "Time left", Continue: "Continue", PlayTestSound: "Play", ReplayTestSound: "Play again", StartMeeting: "Start"},
		Colors:  sc.Colors{Flag: "#E8A33D", NoIssues: "#5B8C6A"},
		Consent: sc.ConsentSpec{Version: "c1", PdfAssetKey: "consent-test.pdf", CheckboxLabel: "I agree", AcceptLabel: "Accept"},
		Setup: sc.SetupSpec{
			Instruction: "Turn on your sound.",
			AudioCheck: sc.AudioCheck{Clip: line("setup.test", "chair", "lantern"), Prompt: "Which word?",
				Options: []sc.Choice{{ID: "river", Label: "river"}, {ID: "lantern", Label: "lantern"}}, CorrectOptionID: "lantern"},
			ExampleScene: sc.ExampleScene{Layout: []sc.Block{{Kind: "title", ID: "ex.title", Text: rt("Example")}},
				Lines: []sc.AudioLine{line("setup.example", "chair", "This is an example.")}},
			PreStartText: sc.PreStartText{Both: []string{"The meeting is about to start."}, Checker: []string{"FASTCHECKERSENTENCE"}},
		},
		Rating: sc.RatingSpec{Instruction: "Rate the speaker.", Items: []sc.RatingItem{
			{ID: "trust", Label: "Trust", Scale: scale(0, 10)},
			{ID: "competence", Scale: scale(1, 7)},
			{ID: "benevolence", Scale: scale(1, 7)},
			{ID: "integrity", Scale: scale(1, 7)},
		}},
		Actions: sc.ActionSpec{
			Checker: sc.ActionGroupSpec{Prompt: "What do you do now?", Options: []sc.Choice{{ID: "nothing", Label: "Say nothing"},
				{ID: "explain", Label: "Ask to explain"}, {ID: "challenge", Label: "Challenge"}}},
			Neutral: sc.ActionGroupSpec{Prompt: "FASTNEUTRALPROMPT", Note: "The chair will ask.", Options: []sc.Choice{{ID: "hasQuestion", Label: "Yes"},
				{ID: "noQuestion", Label: "No"}}},
		},
	}
	c.Scenes = append(c.Scenes, sc.Scene{Kind: "opening", ID: "opening",
		ProfileCard: &sc.ProfileCard{Name: "Dana", Role: "CISO", Lines: []string{"Ten years"}, ShowFromLineID: "opening.2"},
		Lines:       []sc.AudioLine{line("opening.1", "chair", "Welcome everyone."), line("opening.2", "ciso", "Thank you.")}},
		sc.Scene{Kind: "rating", ID: "rating.T0", Point: "T0", After: "opening"})
	flaggedCount := 0
	for i := 1; i <= n; i++ {
		id := fmt.Sprintf("s%d", i)
		s := sc.Scene{Kind: "slide", ID: id,
			Layout:    []sc.Block{{Kind: "title", ID: id + ".title", Text: rt("Slide " + id)}, {Kind: "bullet", ID: id + ".b1", Text: rt("A point")}},
			Narration: &sc.AudioLine{ID: id + ".narration", Voice: "ciso", Text: "Here is the slide with a few words."}}
		if !isFlagged[i] {
			s.Checker = &sc.CheckerPanel{Status: "noIssues", AppearAtMs: ms(100)}
			c.Scenes = append(c.Scenes, s, sc.Scene{Kind: "rating", ID: "rating.T" + fmt.Sprint(i), Point: "T" + fmt.Sprint(i), After: id})
			continue
		}
		flaggedCount++
		s.Checker = &sc.CheckerPanel{Status: "flagged", Markings: []sc.Marking{{ID: id + ".m1", TargetIDs: []string{id + ".b1"},
			AppearAtMs: ms(200), Shape: "box", Label: "FLAGLABEL" + id, Detail: "Detail " + id}}}
		s.RoutineQuestion = &sc.AudioLine{ID: id + ".q.routine", Voice: "chair", Text: "Why this choice?"}
		s.ParticipantQuestions = &sc.ParticipantQuestions{Explain: line(id+".q.explain", "member", "Could you explain?"),
			Challenge: line(id+".q.challenge", "member", "The checker flags this. Why?")}
		if flaggedCount%2 == 1 {
			s.Answers = &sc.Answers{Mode: "whole", Reason: &sc.AudioLine{ID: id + ".a.reason", Voice: "ciso", Text: "REASONTEXT" + id + " because of cost."},
				Denial: &sc.AudioLine{ID: id + ".a.denial", Voice: "ciso", Text: "DENIALTEXT" + id + " it is right."}}
		} else {
			pad := true
			s.Answers = &sc.Answers{Mode: "joined", Opening: &sc.AudioLine{ID: id + ".a.open", Voice: "ciso", Text: "It is right."},
				ReasonBody: &sc.AudioLine{ID: id + ".a.body", Voice: "ciso", Text: "REASONTEXT" + id + " because of risk."},
				Closing:    &sc.AudioLine{ID: id + ".a.close", Voice: "ciso", Text: "That is my plan."}, PadDenialToReason: &pad}
		}
		c.Scenes = append(c.Scenes, s,
			sc.Scene{Kind: "rating", ID: "rating.T" + fmt.Sprint(i) + "a", Point: "T" + fmt.Sprint(i) + "a", After: id},
			sc.Scene{Kind: "action", ID: id + ".action", ForSlide: id},
			sc.Scene{Kind: "question", ID: id + ".question", ForSlide: id},
			sc.Scene{Kind: "answer", ID: id + ".answer", ForSlide: id},
			sc.Scene{Kind: "rating", ID: "rating.T" + fmt.Sprint(i) + "b", Point: "T" + fmt.Sprint(i) + "b", After: id + ".answer"})
	}
	c.Scenes = append(c.Scenes,
		sc.Scene{Kind: "budget", ID: "budget", Prompt: "Approve?", Options: []sc.Choice{{ID: "approve", Label: "Approve"}, {ID: "reject", Label: "Reject"}},
			Confidence: &sc.Confidence{Label: "How sure?", Scale: scale(0, 10)}},
		sc.Scene{Kind: "background", ID: "background", Intro: "About you.", PreferNotToSayLabel: "Prefer not to say",
			Items: []sc.BackgroundItem{
				{Kind: "single", ID: "field", Label: "Field", Options: []sc.Choice{{ID: "it", Label: "IT"}, {ID: "other", Label: "Other"}}, OtherOptionID: "other"},
				{Kind: "scale", ID: "cyber", Label: "Knowledge", Scale: &sc.ScaleSpec{Min: 1, Max: 5, MinLabel: "none", MaxLabel: "expert"}},
				{Kind: "country", ID: "country", Label: "Country"},
				{Kind: "honeypot", ID: "hp", Label: "Leave empty"},
			}},
		sc.Scene{Kind: "debrief", ID: "debrief", Paragraphs: &sc.DebriefParagraphs{Both: []string{"Thank you."},
			Checker: []string{"FASTCHECKERDEBRIEF"}, Neutral: []string{"FASTNEUTRALDEBRIEF"}, End: []string{"Bye."}},
			DoNotDiscuss: "Please do not discuss the study.", SessionCodeLabel: "Your code:"},
	)
	return c
}

// DesktopUA is a desktop Chrome user agent.
const DesktopUA = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0 Safari/537.36"

// NewSession builds a freshly opened session as POST sessions would.
func NewSession(ct *domain.Content, id string, now time.Time) *domain.Session {
	s := &domain.Session{ID: id, RunID: "run1", RunCode: "abc234", Source: "Test", ConfigVersion: ct.Config.Version,
		Status: domain.StatusOpened, CurrentState: ct.Plan.Steps[0].Key(), ProgressIndex: 1,
		Deadlines: domain.Deadlines{StepStartedAt: &now}, OpenedAt: now, LastEventAt: now, UpdatedAt: now,
		TokenHash: "x", IP: "203.0.113.5", UserAgent: DesktopUA, Flags: []string{}, ActionChoices: map[string]string{}}
	for _, st := range domain.PreAllocationSteps {
		s.Plan = append(s.Plan, st.Key())
	}
	return s
}
