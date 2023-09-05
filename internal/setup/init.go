package setup

import (
	"context"
	"os"

	"log/slog"

	cfg "github.com/4nd3r5on/ctf1/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func (a *App) initLogger() {
	var logLevel slog.Level

	switch a.appStage {
	case cfg.APP_STAGE_PROD:
		logLevel = slog.LevelInfo
	case cfg.APP_STAGE_DEV, cfg.APP_STAGE_TEST:
		logLevel = slog.LevelDebug
	default:
		logLevel = slog.LevelDebug
	}

	a.logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: logLevel,
	}))
}

func (a *App) initPostgres(ctx context.Context) error {
	var err error
	a.pgPool, err = pgxpool.New(ctx, a.pgURL)
	return err
}

func (a *App) initRedis(ctx context.Context) {
	a.redisClient = redis.NewClient(a.redisOptions)
}
