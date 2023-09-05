package setup

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"

	"log/slog"

	api "github.com/4nd3r5on/ctf1/internal/api"
	cfg "github.com/4nd3r5on/ctf1/internal/config"
)

func (a *App) Run(ctx context.Context) error {
	var err error
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		err := a.runHttpAPI(ctx)
		if err != nil {
			a.logger.Error("Fatal error while running HTTP API",
				slog.String("error", err.Error()),
				slog.String("app_stage", cfg.StageToString(a.appStage)))
		}
		wg.Done()
	}()

	a.logger.Info("Running HTTP API", slog.String("API_URL", a.apiURL()))
	wg.Wait()
	return err
}

func (a *App) runHttpAPI(ctx context.Context) error {
	var domain string
	if a.servDomain != "" {
		domain = a.servDomain
	} else {
		domain = a.servIP
	}
	r := api.NewAPI(ctx, api.ApiOpts{
		Addr:       domain,
		TlsEnabled: a.tlsEnabled,
		Port:       a.httpPort,
		ApiPrefix:  a.apiPrefix,
		ThisURL:    a.apiURL(),

		SMTPConfig: a.smtpConfig,
		URepo:      a.userRepo,
		MRepo:      a.mailRepo,
		Logger:     a.logger,
	})

	addr := fmt.Sprintf("%s:%d", a.servIP, a.httpPort)

	switch a.appStage {
	case cfg.APP_STAGE_DEV, cfg.APP_STAGE_TEST:
		return http.ListenAndServe(addr, r)
	case cfg.APP_STAGE_PROD:
		return http.ListenAndServeTLS(addr, a.pathToCert, a.pathToKey, r)
	default:
		return errors.New("Cannot run HTTP Server: failed to parse app stage")
	}
}
