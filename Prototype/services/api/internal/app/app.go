// Package app reads the environment and builds the store, the content source
// and the config registry for cmd/api and cmd/sweeper.
package app

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"

	"study/api/internal/assets"
	"study/api/internal/configs"
	"study/api/internal/store"
	"study/api/internal/store/firestore"
	"study/api/internal/store/memory"
)

// Env returns an environment variable or a default.
func Env(key, def string) string {
	if v := strings.TrimSpace(os.Getenv(key)); v != "" {
		return v
	}
	return def
}

// EnvBool reads "true"/"1".
func EnvBool(key string) bool {
	v := strings.ToLower(Env(key, ""))
	return v == "true" || v == "1" || v == "yes"
}

// EnvInt reads an integer with a default.
func EnvInt(key string, def int) int {
	if v, err := strconv.Atoi(Env(key, "")); err == nil {
		return v
	}
	return def
}

// Logger returns a JSON logger (12.4: structured JSON logs).
func Logger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
}

// OpenStore picks the store: STORE=memory|firestore; by default firestore
// when FIRESTORE_EMULATOR_HOST or a project id is set, memory otherwise.
func OpenStore(ctx context.Context, log *slog.Logger) (store.Store, error) {
	project := Env("FIRESTORE_PROJECT_ID", Env("FIRESTORE_PROJECT", Env("GOOGLE_CLOUD_PROJECT", "")))
	kind := Env("STORE", "")
	if kind == "" {
		if os.Getenv("FIRESTORE_EMULATOR_HOST") != "" || project != "" {
			kind = "firestore"
		} else {
			kind = "memory"
		}
	}
	switch kind {
	case "memory":
		log.Warn("using the in-memory store; data is lost when the process ends")
		return memory.New(), nil
	case "firestore":
		if project == "" {
			project = "demo-study"
		}
		log.Info("using firestore", "project", project, "emulator", os.Getenv("FIRESTORE_EMULATOR_HOST"))
		return firestore.New(ctx, project)
	}
	return nil, fmt.Errorf("unknown STORE %q", kind)
}

// OpenSource returns the content source: CONTENT_BUCKET (GCS) or CONTENT_DIR.
func OpenSource(ctx context.Context) (assets.Source, error) {
	if b := Env("CONTENT_BUCKET", ""); b != "" {
		return assets.NewGCS(ctx, b)
	}
	dir := Env("CONTENT_DIR", "")
	if dir == "" {
		return nil, errors.New("set CONTENT_DIR or CONTENT_BUCKET")
	}
	if st, err := os.Stat(dir); err != nil || !st.IsDir() {
		return nil, fmt.Errorf("CONTENT_DIR %q is not a directory", dir)
	}
	return assets.Local{Root: dir}, nil
}

// LoadRegistry loads every config version from the source.
func LoadRegistry(ctx context.Context, src assets.Source, st store.Store, log *slog.Logger) (*configs.Registry, error) {
	reg := &configs.Registry{Source: src, Store: st, AllowTest: EnvBool("ALLOW_TEST_CONFIGS"), AssetURLPrefix: "/assets/", Log: log}
	if err := reg.Load(ctx); err != nil {
		return nil, err
	}
	var versions []string
	for _, e := range reg.List() {
		versions = append(versions, e.Record.Version)
	}
	log.Info("configs loaded", "source", src.Name(), "versions", versions)
	return reg, nil
}
