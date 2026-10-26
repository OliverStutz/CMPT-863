// Package configs loads study config versions with their audio manifests from
// the content source, keeps them in memory, and handles admin uploads.
package configs

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"path"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"

	"study/api/internal/assets"
	"study/api/internal/domain"
	"study/api/internal/store"
	sc "study/api/pkg/studyconfig"
)

// SchemaPath is the schema's path in the content source.
const SchemaPath = "schema/study-config.schema.json"

// Entry is one loaded version.
type Entry struct {
	Content   *domain.Content
	Record    *domain.ConfigRecord
	ScenePack []byte
}

// Registry holds every loaded config version.
type Registry struct {
	Source         assets.Source
	Store          store.Store
	Validator      *sc.Validator
	AllowTest      bool
	AssetURLPrefix string // "/assets/"
	Log            *slog.Logger
	mu             sync.RWMutex
	entries        map[string]*Entry
}

var fileName = regexp.MustCompile(`^study-config\.(.+)\.json$`)

// IsTestConfig reports whether a version is a test config, accepted only with
// ALLOW_TEST_CONFIGS=true (section 15).
// DECISION: the document does not say how a test config is recognised; it is
// one whose version starts with "test".
func IsTestConfig(version string) bool { return strings.HasPrefix(version, "test") }

// Load reads the schema and every config in the source: study-config.<tag>.json
// at the root and configs/<version>.json. For each, the manifest is read from
// audio/<version>/manifest.json, or audio/<tag>/manifest.json; when it is
// missing, durations are estimated and a warning is logged. Configs uploaded
// through the admin API are then loaded from the store.
func (r *Registry) Load(ctx context.Context) error {
	r.mu.Lock()
	r.entries = map[string]*Entry{}
	r.mu.Unlock()
	if r.Log == nil {
		r.Log = slog.Default()
	}
	if obj, err := r.Source.Read(ctx, SchemaPath); err == nil {
		v, err := sc.NewValidator(obj.Data)
		if err != nil {
			return fmt.Errorf("schema: %w", err)
		}
		r.Validator = v
	} else {
		r.Log.Warn("schema not found; configs are checked structurally only", "path", SchemaPath, "source", r.Source.Name())
	}
	var files []string
	root, err := r.Source.List(ctx, "")
	if err != nil {
		return err
	}
	for _, f := range root {
		if fileName.MatchString(path.Base(f)) {
			files = append(files, f)
		}
	}
	more, err := r.Source.List(ctx, "configs")
	if err != nil {
		return err
	}
	for _, f := range more {
		if strings.HasSuffix(f, ".json") {
			files = append(files, f)
		}
	}
	for _, f := range files {
		obj, err := r.Source.Read(ctx, f)
		if err != nil {
			r.Log.Warn("config unreadable", "file", f, "err", err)
			continue
		}
		tag := ""
		if m := fileName.FindStringSubmatch(path.Base(f)); m != nil {
			tag = m[1]
		} else {
			tag = strings.TrimSuffix(path.Base(f), ".json")
		}
		if _, err := r.add(ctx, obj.Data, tag, f, domain.ConfigValidated, false); err != nil {
			r.Log.Warn("config skipped", "file", f, "err", err)
		}
	}
	if r.Store != nil {
		recs, err := r.Store.ListConfigs(ctx)
		if err != nil {
			return err
		}
		for _, rec := range recs {
			if rec.Body == "" {
				continue
			}
			if _, ok := r.Get(rec.Version); ok {
				continue
			}
			if _, err := r.add(ctx, []byte(rec.Body), rec.Version, rec.GCSPath, rec.Status, false); err != nil {
				r.Log.Warn("stored config skipped", "version", rec.Version, "err", err)
			}
		}
	}
	return nil
}

func (r *Registry) add(ctx context.Context, data []byte, tag, srcPath, status string, keepBody bool) (*Entry, error) {
	cfg, err := sc.Parse(data, r.Validator)
	if err != nil {
		return nil, err
	}
	if IsTestConfig(cfg.Version) && !r.AllowTest {
		return nil, fmt.Errorf("test config %q rejected (ALLOW_TEST_CONFIGS is false)", cfg.Version)
	}
	var man *sc.Manifest
	manPath := ""
	dir := ""
	for _, d := range []string{cfg.Version, tag} {
		if d == "" {
			continue
		}
		p := "audio/" + d + "/manifest.json"
		obj, err := r.Source.Read(ctx, p)
		if err != nil {
			continue
		}
		var m sc.Manifest
		if err := json.Unmarshal(obj.Data, &m); err != nil {
			return nil, fmt.Errorf("%s: %w", p, err)
		}
		man, manPath, dir = &m, p, d
		break
	}
	audioBase := r.AssetURLPrefix + "audio/" + dir + "/"
	ct, err := domain.NewContent(cfg, man, audioBase, r.AssetURLPrefix+"_dev/silence/")
	if err != nil {
		return nil, err
	}
	if man == nil {
		r.Log.Warn("audio manifest missing; using estimated durations and silent dev clips",
			"version", cfg.Version, "tried", []string{"audio/" + cfg.Version + "/manifest.json", "audio/" + tag + "/manifest.json"})
	} else if len(ct.Estimated) > 0 {
		r.Log.Warn("clips missing from the manifest; durations estimated", "version", cfg.Version, "clips", ct.Estimated)
	}
	for _, w := range ct.Warnings {
		r.Log.Warn("config warning", "version", cfg.Version, "warning", w)
	}
	// Optional video manifest (talking heads, 2026-10-25): video/<version>/ or video/<tag>/.
	var vm *sc.VideoManifest
	vdir := ""
	for _, d := range []string{cfg.Version, tag} {
		if d == "" {
			continue
		}
		p := "video/" + d + "/manifest.json"
		obj, err := r.Source.Read(ctx, p)
		if err != nil {
			continue
		}
		var m sc.VideoManifest
		if err := json.Unmarshal(obj.Data, &m); err != nil {
			return nil, fmt.Errorf("%s: %w", p, err)
		}
		vm, vdir = &m, d
		break
	}
	ct.SetVideo(r.AssetURLPrefix+"video/"+vdir+"/", vm)
	pack, err := sc.PublicScenePack(cfg)
	if err != nil {
		return nil, err
	}
	if len(ct.SeatVideos) > 0 {
		// The idle loops of the silent seats travel with the public scene pack.
		var pm map[string]any
		if err := json.Unmarshal(pack, &pm); err != nil {
			return nil, err
		}
		pm["video"] = map[string]any{"seats": ct.SeatVideos}
		if pack, err = json.Marshal(pm); err != nil {
			return nil, err
		}
	}
	sum := sha256.Sum256(data)
	now := time.Now().UTC()
	rec := &domain.ConfigRecord{Version: cfg.Version, SHA256: hex.EncodeToString(sum[:]), GCSPath: srcPath,
		ScenePackPath: "scenes/" + cfg.Version + ".json", ManifestPath: manPath, Status: status, UploadedAt: &now,
		LintReport: ct.Warnings}
	if man == nil {
		rec.LintReport = append(rec.LintReport, "audio manifest missing: durations estimated")
	}
	if vm != nil {
		rec.LintReport = append(rec.LintReport, fmt.Sprintf("video manifest: %d clip(s), %d seat loop(s)", len(ct.Videos), len(ct.SeatVideos)))
	}
	if keepBody {
		rec.Body = string(data)
	}
	e := &Entry{Content: ct, Record: rec, ScenePack: pack}
	r.mu.Lock()
	if prev, ok := r.entries[cfg.Version]; ok {
		r.Log.Warn("config version loaded twice; the later file wins", "version", cfg.Version, "previous", prev.Record.GCSPath, "file", srcPath)
	}
	r.entries[cfg.Version] = e
	r.mu.Unlock()
	return e, nil
}

// Get returns a loaded version.
func (r *Registry) Get(version string) (*Entry, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	e, ok := r.entries[version]
	return e, ok
}

// Content returns the content of a loaded version.
func (r *Registry) Content(version string) (*domain.Content, error) {
	e, ok := r.Get(version)
	if !ok {
		return nil, fmt.Errorf("config version %q is not loaded", version)
	}
	return e.Content, nil
}

// ScenePack implements assets.Generated.
func (r *Registry) ScenePack(version string) ([]byte, bool) {
	e, ok := r.Get(version)
	if !ok {
		return nil, false
	}
	return e.ScenePack, true
}

// List returns all entries sorted by version.
func (r *Registry) List() []*Entry {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]*Entry, 0, len(r.entries))
	for _, e := range r.entries {
		out = append(out, e)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Record.Version < out[j].Record.Version })
	return out
}

// UploadError carries the lint errors of a rejected upload.
type UploadError struct {
	Code   string
	Errors []string
}

func (e *UploadError) Error() string { return e.Code + ": " + strings.Join(e.Errors, "; ") }

// ErrLocked is returned when an upload would replace a locked version.
var ErrLocked = errors.New("config version is locked")

// Upload validates, checks and stores an uploaded config as a draft.
// DECISION: uploaded configs are kept in the store (configs/{version}.body)
// rather than written to the content bucket; the service account only has
// read access to the bucket (12.3).
func (r *Registry) Upload(ctx context.Context, data []byte, locked func(version string) bool) (*Entry, error) {
	cfg, err := sc.Parse(data, r.Validator)
	if err != nil {
		return nil, &UploadError{Code: domain.ErrLintFailed, Errors: splitErr(err)}
	}
	if locked(cfg.Version) {
		return nil, ErrLocked
	}
	// 8.2: 422 if any audioKey in the resolved config is missing from the bucket.
	var missing []string
	for _, k := range audioKeys(cfg) {
		found := false
		for _, d := range []string{cfg.Version} {
			if _, err := r.Source.Read(ctx, "audio/"+d+"/"+k); err == nil {
				found = true
			}
		}
		if !found {
			missing = append(missing, k)
		}
	}
	if len(missing) > 0 {
		return nil, &UploadError{Code: domain.ErrMissingAudio, Errors: missing}
	}
	e, err := r.add(ctx, data, cfg.Version, "store:configs/"+cfg.Version, domain.ConfigDraft, true)
	if err != nil {
		return nil, &UploadError{Code: domain.ErrLintFailed, Errors: splitErr(err)}
	}
	if r.Store != nil {
		if err := r.Store.PutConfig(ctx, e.Record); err != nil {
			return nil, err
		}
	}
	return e, nil
}

func splitErr(err error) []string {
	return strings.Split(err.Error(), "\n")
}

func audioKeys(c *sc.StudyConfig) []string {
	var out []string
	add := func(l *sc.AudioLine) {
		if l != nil && l.AudioKey != "" {
			out = append(out, l.AudioKey)
		}
	}
	add(&c.Setup.AudioCheck.Clip)
	for i := range c.Setup.ExampleScene.Lines {
		add(&c.Setup.ExampleScene.Lines[i])
	}
	for i := range c.Scenes {
		s := &c.Scenes[i]
		for j := range s.Lines {
			add(&s.Lines[j])
		}
		add(s.Narration)
		add(s.RoutineQuestion)
		if q := s.ParticipantQuestions; q != nil {
			add(&q.Explain)
			add(&q.Challenge)
		}
		if a := s.Answers; a != nil {
			add(a.Reason)
			add(a.Denial)
			add(a.Opening)
			add(a.ReasonBody)
			add(a.Closing)
		}
	}
	return out
}
