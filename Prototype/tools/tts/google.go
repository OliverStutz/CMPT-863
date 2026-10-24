package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	sc "study/api/pkg/studyconfig"
)

const synthesizeURL = "https://texttospeech.googleapis.com/v1beta1/text:synthesize"

// synthesizer turns one part into MP3 bytes and its SSML mark offsets in ms.
type synthesizer interface {
	Synthesize(ctx context.Context, v sc.VoiceSpec, input string, marks bool) ([]byte, map[string]int64, error)
}

// googleTTS calls the Cloud Text-to-Speech REST API (v1beta1, which returns
// SSML mark timepoints) with application-default credentials.
type googleTTS struct {
	client       *http.Client
	quotaProject string
}

func newGoogleTTS(ctx context.Context, quotaProject string) (*googleTTS, error) {
	creds, err := google.FindDefaultCredentials(ctx, "https://www.googleapis.com/auth/cloud-platform")
	if err != nil {
		return nil, fmt.Errorf("application-default credentials: %w", err)
	}
	if quotaProject == "" {
		var f struct {
			QuotaProjectID string `json:"quota_project_id"`
		}
		_ = json.Unmarshal(creds.JSON, &f)
		quotaProject = f.QuotaProjectID
	}
	if quotaProject == "" {
		quotaProject = creds.ProjectID
	}
	return &googleTTS{client: oauth2.NewClient(ctx, creds.TokenSource), quotaProject: quotaProject}, nil
}

type synthRequest struct {
	Input struct {
		SSML string `json:"ssml"`
	} `json:"input"`
	Voice struct {
		LanguageCode string `json:"languageCode"`
		Name         string `json:"name"`
	} `json:"voice"`
	AudioConfig struct {
		AudioEncoding   string   `json:"audioEncoding"`
		SampleRateHertz int      `json:"sampleRateHertz"`
		SpeakingRate    *float64 `json:"speakingRate,omitempty"`
		Pitch           *float64 `json:"pitch,omitempty"`
	} `json:"audioConfig"`
	EnableTimePointing []string `json:"enableTimePointing,omitempty"`
}

type synthResponse struct {
	AudioContent string `json:"audioContent"`
	Timepoints   []struct {
		MarkName    string  `json:"markName"`
		TimeSeconds float64 `json:"timeSeconds"`
	} `json:"timepoints"`
}

func (g *googleTTS) Synthesize(ctx context.Context, v sc.VoiceSpec, input string, marks bool) ([]byte, map[string]int64, error) {
	var req synthRequest
	req.Input.SSML = input
	req.Voice.LanguageCode = v.LanguageCode
	req.Voice.Name = v.Name
	req.AudioConfig.AudioEncoding = audioEncoding
	req.AudioConfig.SampleRateHertz = sampleRateHertz
	req.AudioConfig.SpeakingRate = v.SpeakingRate
	req.AudioConfig.Pitch = v.Pitch
	if marks {
		req.EnableTimePointing = []string{"SSML_MARK"}
	}
	body, _ := json.Marshal(req)
	var lastErr error
	for attempt := 0; attempt < 5; attempt++ {
		if attempt > 0 {
			time.Sleep(time.Duration(1<<attempt) * 500 * time.Millisecond)
		}
		hr, _ := http.NewRequestWithContext(ctx, http.MethodPost, synthesizeURL, bytes.NewReader(body))
		hr.Header.Set("Content-Type", "application/json")
		if g.quotaProject != "" {
			hr.Header.Set("x-goog-user-project", g.quotaProject)
		}
		resp, err := g.client.Do(hr)
		if err != nil {
			lastErr = err
			continue
		}
		rb, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		if resp.StatusCode == 429 || resp.StatusCode >= 500 {
			lastErr = fmt.Errorf("HTTP %d: %s", resp.StatusCode, rb)
			continue
		}
		if resp.StatusCode != 200 {
			return nil, nil, fmt.Errorf("voice %s: HTTP %d: %s", v.Name, resp.StatusCode, rb)
		}
		var sr synthResponse
		if err := json.Unmarshal(rb, &sr); err != nil {
			return nil, nil, err
		}
		audio, err := base64.StdEncoding.DecodeString(sr.AudioContent)
		if err != nil {
			return nil, nil, err
		}
		var m map[string]int64
		if len(sr.Timepoints) > 0 {
			m = map[string]int64{}
			for _, tp := range sr.Timepoints {
				m[tp.MarkName] = int64(math.Round(tp.TimeSeconds * 1000))
			}
		}
		return audio, m, nil
	}
	return nil, nil, lastErr
}

// cachedPart is the sidecar stored next to <key>.mp3 in the cache.
type cachedPart struct {
	Voice      string           `json:"voice"`
	Input      string           `json:"input"`
	Marks      map[string]int64 `json:"marks,omitempty"`
	DurationMs int64            `json:"durationMs"`
}

// cache stores synthesised parts by key in a local directory.
type cache struct{ dir string }

func (c cache) get(key string) ([]byte, *cachedPart, bool) {
	b, err := os.ReadFile(c.dir + "/" + key + ".mp3")
	if err != nil {
		return nil, nil, false
	}
	mb, err := os.ReadFile(c.dir + "/" + key + ".json")
	if err != nil {
		return nil, nil, false
	}
	var cp cachedPart
	if json.Unmarshal(mb, &cp) != nil {
		return nil, nil, false
	}
	return b, &cp, true
}

func (c cache) put(key string, audio []byte, meta *cachedPart) error {
	if err := os.MkdirAll(c.dir, 0o755); err != nil {
		return err
	}
	mb, _ := json.MarshalIndent(meta, "", "  ")
	if err := writeAtomic(c.dir+"/"+key+".mp3", audio); err != nil {
		return err
	}
	return writeAtomic(c.dir+"/"+key+".json", mb)
}

func writeAtomic(path string, b []byte) error {
	tmp := path + ".tmp"
	if err := os.WriteFile(tmp, b, 0o644); err != nil {
		return err
	}
	return os.Rename(tmp, path)
}
