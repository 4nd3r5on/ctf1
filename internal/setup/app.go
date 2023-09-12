package setup

import (
	"context"
	"errors"
	"log/slog"

	"github.com/4nd3r5on/ctf1/internal/config"
	mail_repo "github.com/4nd3r5on/ctf1/internal/repository/mail_verification"
	user_repo "github.com/4nd3r5on/ctf1/internal/repository/users"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type App struct {
	// HTTP config
	servIP     string
	servDomain string
	httpPort   int
	apiPrefix  string
	// TLS config
	tlsEnabled bool
	pathToCert string
	pathToKey  string

	smtpConfig config.SMTPConfig

	redisOptions *redis.Options
	redisClient  *redis.Client

	pgURL  string
	pgPool *pgxpool.Pool

	userRepo user_repo.UsersRepository
	mailRepo mail_repo.MailVerificationRepository

	logger *slog.Logger

	appStage int
}

func NewApp(ctx context.Context, cfg config.AppConfig) (App, error) {
	if cfg.AppStage == config.APP_STAGE_PROD {
		if !cfg.EnableTLS {
			return App{},
				errors.New("Stage production cannot be used without TLS")
		}
	}
	if cfg.ApiPrefix == "" {
		cfg.ApiPrefix = "/"
	}
	if cfg.HTTPPort == 0 {
		if cfg.EnableTLS {
			cfg.HTTPPort = 433
		}
		cfg.HTTPPort = 80
	}
	if cfg.Address == "" {
		cfg.Address = "127.0.0.1"
	}

	a := App{
		appStage: cfg.AppStage,

		servIP:     cfg.Address,
		servDomain: cfg.Domain,
		httpPort:   cfg.HTTPPort,
		apiPrefix:  cfg.ApiPrefix,

		tlsEnabled: cfg.EnableTLS,
		pathToCert: cfg.PathToCert,
		pathToKey:  cfg.PathToKey,

		pgURL: cfg.PostgresURL,

		redisOptions: cfg.RedisConfig,

		smtpConfig: cfg.SMTPConfig,
	}

	a.initLogger()
	if err := a.initPostgres(ctx, cfg.PgMigrationConfig); err != nil {
		return App{}, err
	}
	if err := a.initRedis(ctx); err != nil {
		return App{}, err
	}

	return a, nil
}
