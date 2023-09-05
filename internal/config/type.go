package config

import (
	"log/slog"

	"github.com/redis/go-redis/v9"
)

const (
	APP_STAGE_PROD = iota
	APP_STAGE_TEST = iota
	APP_STAGE_DEV  = iota
)

// Used for logging
func StageToString(stage int) string {
	switch stage {
	case APP_STAGE_PROD:
		return "production"
	case APP_STAGE_TEST:
		return "testing"
	case APP_STAGE_DEV:
		return "development"
	default:
		return "unknown"
	}
}

type AppConfig struct {
	AppStage int // Default production

	Address string // Default "127.0.0.1"
	Domain  string // Default: ""
	// Default 80 if TLS disabled and 433 if enabled
	HTTPPort  int
	ApiPrefix string // Default "/"

	// TLS config
	PathToCert string // .crt || .pem file
	PathToKey  string // .key file
	// necessary to be true if stage is production
	EnableTLS bool

	PostgresURL       string
	PgMigrationConfig MigrationConfig

	RedisConfig *redis.Options

	SMTPConfig SMTPConfig
}

type SMTPConfig struct {
	// From is the address from which server sends e-mails
	From         string
	SMTPServer   string // smtp.google.com:587
	SMTPPassword string
}

type MigrationConfig struct {
	MigrationsPath string
	VersionLimit   uint
	DropDev        bool
	Logger         *slog.Logger
}
