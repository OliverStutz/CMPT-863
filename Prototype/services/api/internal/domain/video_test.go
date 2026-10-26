package domain_test

import (
	"strings"
	"testing"
	"time"

	"study/api/internal/domain"
	"study/api/internal/testutil"
	sc "study/api/pkg/studyconfig"
)

// A version with a video manifest puts the clip's video on the answer step and
// in the prefetch list, and the seat loops on the content; a clip without a
// video entry, and a version without a manifest, carry nothing (2026-10-25).
func TestVideoManifestOnSteps(t *testing.T) {
	c := testutil.FastConfig("test-fast", 3, 2)
	var slide sc.Scene
	var other string
	for _, sc0 := range c.Scenes {
		if sc0.Kind == sc.SceneSlide && sc0.Answers != nil && slide.ID == "" {
			slide = sc0
		} else if sc0.Kind == sc.SceneSlide && other == "" {
			other = sc0.ID
		}
	}
	reasonID, _ := sc.AnswerClipIDs(&slide)
	if reasonID == "" || other == "" {
		t.Fatalf("fixture has no flagged slide with answers plus another slide (flagged %q, other %q)", slide.ID, other)
	}
	man := &sc.Manifest{ConfigVersion: c.Version, Clips: map[string]sc.ManifestClip{}}
	ct, err := domain.NewContent(c, man, "/assets/audio/x/", "/assets/_dev/silence/")
	if err != nil {
		t.Fatal(err)
	}
	// Give the reason clip a real audio key, as tools/tts would.
	ct.Clips[reasonID].AudioKey = strings.Repeat("a", 64) + ".mp3"
	ct.Clips[reasonID].Estimated = false
	ct.SetVideo("/assets/video/x/", &sc.VideoManifest{
		Clips: map[string]sc.VideoClip{reasonID: {VideoKey: strings.Repeat("b", 64) + ".mp4", DurationMs: 15562}},
		Seats: map[string]sc.VideoClip{"member2": {VideoKey: strings.Repeat("c", 64) + ".mp4", DurationMs: 9937}},
	})
	if got := ct.VideoURL(ct.Clips[reasonID]); got != "/assets/video/x/"+strings.Repeat("b", 64)+".mp4" {
		t.Fatalf("VideoURL = %q", got)
	}
	if got := ct.VideoURL(ct.Clips[slide.Narration.ID]); got != "" {
		t.Fatalf("narration has no video, got %q", got)
	}
	if got := ct.SeatVideos["member2"]; got != "/assets/video/x/"+strings.Repeat("c", 64)+".mp4" {
		t.Fatalf("seat loop = %q", got)
	}
	// Walk a neutral session to the reason answer of that slide and check the step.
	s := &domain.Session{ID: "s1", Group: string(sc.GroupNeutral), DenialSlide: other, ActionChoices: map[string]string{}}
	var answerIdx, prevAudioIdx = -1, -1
	for i, st := range ct.Plan.Steps {
		if st.Kind == domain.KindAnswer && st.SlideID == slide.ID {
			answerIdx = i
			break
		}
	}
	if answerIdx < 0 {
		t.Fatal("no answer step for the flagged slide")
	}
	for i := answerIdx - 1; i >= 0; i-- {
		if ct.Plan.Steps[i].Audio() {
			prevAudioIdx = i
			break
		}
	}
	s.ActionChoices[slide.ID] = slide.ParticipantQuestions.Challenge.ID
	s.CurrentState = ct.Plan.Steps[answerIdx].Key()
	step := ct.Describe(s, time.Now(), false)
	if step.VideoURL == "" || !strings.HasSuffix(step.VideoURL, ".mp4") {
		t.Fatalf("answer step without videoUrl: %+v", step)
	}
	if prevAudioIdx >= 0 && prevAudioIdx >= answerIdx-3 {
		s.CurrentState = ct.Plan.Steps[prevAudioIdx].Key()
		prev := ct.Describe(s, time.Now(), false)
		found := false
		for _, u := range prev.Prefetch {
			if u == step.VideoURL {
				found = true
			}
		}
		if !found {
			t.Fatalf("previous audio step does not prefetch the video: %v", prev.Prefetch)
		}
	}
	// Without a manifest nothing changes.
	ct.SetVideo("", nil)
	if got := ct.Describe(s, time.Now(), false).VideoURL; got != "" {
		t.Fatalf("videoUrl without manifest: %q", got)
	}
}
