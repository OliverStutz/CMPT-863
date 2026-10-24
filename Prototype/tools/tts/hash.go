package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"

	sc "study/api/pkg/studyconfig"
)

// Audio encoding requested from the API. Part of every key, so a change of
// encoding never reuses an old file.
const (
	audioEncoding   = "MP3"
	sampleRateHertz = 24000
)

// synthKey is the content hash of one synthesis: the voice, the text, the
// SSML and the encoding. The MP3 is stored as <key>.mp3; the name says
// nothing about the line it holds.
func synthKey(v sc.VoiceSpec, text, ssml string) string {
	in := struct {
		Kind         string   `json:"kind"`
		Voice        string   `json:"voice"`
		LanguageCode string   `json:"languageCode"`
		SpeakingRate *float64 `json:"speakingRate,omitempty"`
		Pitch        *float64 `json:"pitch,omitempty"`
		Text         string   `json:"text"`
		SSML         string   `json:"ssml"`
		Encoding     string   `json:"encoding"`
		SampleRate   int      `json:"sampleRateHertz"`
	}{"synth", v.Name, v.LanguageCode, v.SpeakingRate, v.Pitch, text, ssml, audioEncoding, sampleRateHertz}
	return hashJSON(in)
}

// joinKey is the hash of a clip assembled from synthesised parts: the part
// keys in order, the silence between them and the silence appended at the end.
func joinKey(parts []string, gapMs, padFrames int64) string {
	in := struct {
		Kind      string   `json:"kind"`
		Parts     []string `json:"parts"`
		GapMs     int64    `json:"gapMs"`
		PadFrames int64    `json:"padFrames"`
	}{"join", parts, gapMs, padFrames}
	return hashJSON(in)
}

func hashJSON(v any) string {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err) // only plain structs are hashed
	}
	sum := sha256.Sum256(b)
	return hex.EncodeToString(sum[:])
}
