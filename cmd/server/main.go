package main

import (
	"context"
	"log"
	"math/rand"

	cfg "github.com/4nd3r5on/ctf1/internal/config"
	"github.com/4nd3r5on/ctf1/internal/setup"
)

func main() {
	rand.Seed(0)
	ctx := context.Background()

	pgURL, envErr := cfg.GetPgUrlFromEnv()
	if envErr != nil {
		log.Fatalf("Failed to parse %s envinronment variable: %s",
			envErr.Info().VarName, envErr.Error())
	}
	redisOpts, envErr := cfg.GetRedisOptionsFromEnv()
	if envErr != nil {
		log.Fatalf("Failed to parse %s envinronment variable: %s",
			envErr.Info().VarName, envErr.Error())
	}
	smtpCfg, envErr := cfg.GetSmtpCfgFromEnv()
	if envErr != nil {
		log.Fatalf("Failed to parse %s envinronment variable: %s",
			envErr.Info().VarName, envErr.Error())
	}

	app, err := setup.NewApp(ctx, cfg.AppConfig{
		AppStage: cfg.APP_STAGE_DEV,

		Address:   "0.0.0.0",
		Domain:    "localhost",
		HTTPPort:  9000,
		ApiPrefix: "/api",

		EnableTLS: false,

		PostgresURL: pgURL,
		PgMigrationConfig: cfg.MigrationConfig{
			MigrationsPath: "migrations",
			VersionLimit:   -1,
			Drop:           true,
		},

		RedisConfig: redisOpts,

		SMTPConfig: smtpCfg,
	})
	if err != nil {
		panic(err)
	}

	app.Run(ctx)
}
