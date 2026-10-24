// Command sweeper marks sessions Abandoned after abandonAfterMs without
// events (5.4). It runs once per invocation (Cloud Run job, Cloud Scheduler
// every 5 minutes), or every SWEEP_INTERVAL when set (local use).
package main

import (
	"context"
	"log/slog"
	"os"
	"time"

	"study/api/internal/app"
	"study/api/internal/configs"
	"study/api/internal/domain"
	"study/api/internal/sweep"
)

func main() {
	log := app.Logger()
	slog.SetDefault(log)
	ctx := context.Background()
	st, err := app.OpenStore(ctx, log)
	if err != nil {
		log.Error("store", "err", err)
		os.Exit(1)
	}
	defer st.Close()
	var reg *configs.Registry
	if src, err := app.OpenSource(ctx); err == nil {
		if reg, err = app.LoadRegistry(ctx, src, st, log); err != nil {
			log.Warn("configs not loaded; lazy evaluation skipped", "err", err)
		}
	} else {
		log.Warn("no content source; lazy evaluation skipped, fallback threshold used", "err", err)
	}
	content := func(v string) (*domain.Content, error) {
		if reg == nil {
			return nil, os.ErrNotExist
		}
		return reg.Content(v)
	}
	fallback := time.Duration(app.EnvInt("ABANDON_AFTER_MS", 1800000)) * time.Millisecond
	interval, _ := time.ParseDuration(app.Env("SWEEP_INTERVAL", ""))
	for {
		n, err := sweep.Run(ctx, st, content, time.Now().UTC(), fallback, log)
		if err != nil {
			log.Error("sweep", "err", err)
			if interval == 0 {
				os.Exit(1)
			}
		} else {
			log.Info("sweep done", "abandoned", n)
		}
		if interval == 0 {
			return
		}
		time.Sleep(interval)
	}
}
