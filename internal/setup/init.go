package setup

import (
	"context"
	"os"

	"log/slog"

	cfg "github.com/4nd3r5on/ctf1/internal/config"
	db_utils "github.com/4nd3r5on/ctf1/pkg/db"
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

func (a *App) initPostgres(ctx context.Context, mCfg cfg.MigrationConfig) error {
	var err error
	ver, dirty, errx := db_utils.DoMigrate(ctx, db_utils.MigrationConfig{
		MigrationsPath: mCfg.MigrationsPath,
		DBurl:          a.pgURL,
		VersionLimit:   mCfg.VersionLimit,
		Drop:           mCfg.Drop,
		Logger:         a.logger,
	})
	if errx != nil {
		db_utils.LogMigrationErr(errx, a.logger)
		return errx
	}
	a.logger.Debug("Migrations applied successfuly",
		slog.Int("version", int(ver)), slog.Bool("dirty", dirty))
	a.pgPool, err = pgxpool.New(ctx, a.pgURL)
	if err != nil {
		a.logger.Error("Failed to initialize pgx pool")
	}

	return err
}

func (a *App) initRedis(ctx context.Context) {
	a.redisClient = redis.NewClient(a.redisOptions)
}
