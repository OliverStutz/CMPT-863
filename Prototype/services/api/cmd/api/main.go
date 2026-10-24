// Command api is the HTTP server: SPAs, assets, participant and admin APIs.
package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/crypto/bcrypt"

	"study/api/internal/app"
	"study/api/internal/configs"
	"study/api/internal/domain"
	apihttp "study/api/internal/http"
	"study/api/internal/secrets"
	"study/api/internal/store"
	"study/api/web"
)

func main() {
	log := app.Logger()
	slog.SetDefault(log)
	if err := run(log); err != nil {
		log.Error("fatal", "err", err)
		os.Exit(1)
	}
}

func run(log *slog.Logger) error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	st, err := app.OpenStore(ctx, log)
	if err != nil {
		return err
	}
	defer st.Close()
	src, err := app.OpenSource(ctx)
	if err != nil {
		return err
	}
	reg, err := app.LoadRegistry(ctx, src, st, log)
	if err != nil {
		return err
	}

	fetch := secrets.SecretManager()
	pwHash := secrets.FromEnv("ADMIN_PASSWORD_HASH", fetch, log)
	if pw := os.Getenv("ADMIN_PASSWORD"); pw != "" && !pwHash.Configured() {
		// Dev convenience (section 14): a plain password, hashed at start.
		h, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		pwHash = secrets.Static("ADMIN_PASSWORD_HASH", h)
		log.Warn("ADMIN_PASSWORD set: dev only")
	}
	sessKey := secrets.FromEnv("ADMIN_SESSION_KEY", fetch, log)
	access := secrets.FromEnv("ADMIN_ACCESS_KEY", fetch, log)
	for _, s := range []*secrets.Secret{pwHash, sessKey, access} {
		if !s.Configured() {
			log.Warn("secret not configured", "secret", s.Name)
		} else if len(s.Get(ctx)) == 0 {
			log.Error("secret configured but empty or unreadable", "secret", s.Name)
		}
	}

	if code := app.Env("SEED_RUN_CODE", ""); code != "" {
		seed(ctx, st, reg.List(), code, log)
	}

	srv := apihttp.New(apihttp.Options{
		Store: st, Registry: reg, Source: src, ParticipantFS: web.Participant(), AdminFS: web.Admin(),
		PasswordHash: pwHash, SessionKey: sessKey, AccessKey: access,
		TrustedProxyHops: app.EnvInt("TRUSTED_PROXY_HOPS", 0), InsecureCookies: app.EnvBool("INSECURE_COOKIES"),
		PublicBaseURL: app.Env("PUBLIC_BASE_URL", ""), Log: log,
	})
	hs := &http.Server{Addr: ":" + app.Env("PORT", "8080"), Handler: srv, ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout: 30 * time.Second, WriteTimeout: 60 * time.Second, IdleTimeout: 120 * time.Second}
	go func() {
		<-ctx.Done()
		sctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_ = hs.Shutdown(sctx)
	}()
	log.Info("listening", "addr", hs.Addr)
	if err := hs.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

// seed creates a dev run with the given code on the newest loaded config,
// unless it exists (make dev; section 14 "make seed").
func seed(ctx context.Context, st store.Store, entries []*configs.Entry, code string, log *slog.Logger) {
	if _, err := st.GetRunByCode(ctx, code); err == nil || len(entries) == 0 {
		return
	}
	v := entries[len(entries)-1].Record.Version
	run := &domain.Run{ID: "dev-" + code, Code: code, Name: "Dev run", Source: "Dev", ConfigVersion: v,
		Status: domain.RunOpen, CreatedAt: time.Now().UTC(), CreatedFromIP: "seed"}
	if err := st.CreateRun(ctx, run); err != nil {
		log.Error("seed run", "err", err)
		return
	}
	log.Info("seeded dev run", "code", code, "configVersion", v, "url", "/s/"+code)
}
