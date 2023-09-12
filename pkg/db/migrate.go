package db_utils

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/golang-migrate/migrate/v4"
	pgx_driver "github.com/golang-migrate/migrate/v4/database/pgx/v5"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type MigrationConfig struct {
	MigrationsPath string
	DBurl          string
	VersionLimit   int
	Drop           bool
	Logger         *slog.Logger
}

func NewMigrateInstance(db *sql.DB, sourceURL string, cfg *pgx_driver.Config) (*migrate.Migrate, error) {
	driver, err := pgx_driver.WithInstance(db, &pgx_driver.Config{})
	if err != nil {
		return nil, err
	}
	m, err := migrate.NewWithDatabaseInstance(sourceURL, "postgres", driver)
	return m, err
}

func DoMigrate(ctx context.Context, cfg MigrationConfig) (uint, bool, MigrationErr) {
	var ver uint
	var dirty bool
	var sourceURL string = fmt.Sprintf("file://%s", cfg.MigrationsPath)
	var driverCfg *pgx_driver.Config

	cfg.Logger.Debug(OnDbConn)
	db, err := sql.Open("pgx/v5", cfg.DBurl)
	if err != nil {
		return 0, false, NewMigrationErr(err.Error(), map[string]any{
			"on": OnDbConn, "db_url": cfg.DBurl})
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		return 0, false, NewMigrationErr(err.Error(), map[string]any{
			"on": OnDbConn, "db_url": cfg.DBurl})
	}

	cfg.Logger.Debug(OnNewInstance)

	m, err := NewMigrateInstance(db, sourceURL, driverCfg)
	if err != nil {
		return 0, false,
			NewMigrationErr(err.Error(), map[string]any{"on": OnNewInstance})
	}

	// Drop db if needed, get current db
	if cfg.Drop {
		cfg.Logger.Warn(OnDrop)

		if err = m.Drop(); err != nil {
			return 0, false,
				NewMigrationErr(err.Error(), map[string]any{"on": OnDrop})
		}
		// After drop we have to create new migrate instance
		cfg.Logger.Debug(OnNewInstance)
		m, err = NewMigrateInstance(db, sourceURL, driverCfg)
		if err != nil {
			return 0, false,
				NewMigrationErr(err.Error(), map[string]any{"on": OnNewInstance})
		}

	} else {
		// It's strange to check DB version after we drop it
		// So I put version checking into else statement
		cfg.Logger.Debug(OnVersionCheck)
		ver, dirty, err = m.Version()
		if err != nil && err != migrate.ErrNilVersion {
			return 0, false,
				NewMigrationErr(err.Error(), map[string]any{"on": OnVersionCheck})
		}
		if dirty {
			cfg.Logger.Info(OnDrop,
				slog.Bool("dirty", dirty))
			if err = m.Drop(); err != nil {
				return ver, dirty,
					NewMigrationErr(err.Error(), map[string]any{"on": OnDrop})
			}
			// After drop we have to create new migrate instance
			m, err = NewMigrateInstance(db, sourceURL, driverCfg)
			if err != nil {
				return ver, dirty,
					NewMigrationErr(err.Error(), map[string]any{"on": OnNewInstance})
			}
			// As we dropped DB
			ver = 0
			dirty = false
		}
	}

	if cfg.VersionLimit == int(ver) {
		cfg.Logger.Debug("DB is already up to date")
		return ver, dirty, nil
	}

	// Upgrade DB
	if cfg.VersionLimit > 0 {
		log := "Upgrading"
		if ver > uint(cfg.VersionLimit) {
			log = "Downgrading"
		}
		cfg.Logger.Debug(log+" DB to a limit",
			slog.Int("version", int(ver)),
			slog.Int("limit", int(cfg.VersionLimit)))
		err = m.Migrate(uint(cfg.VersionLimit))
	} else if cfg.VersionLimit == 0 {
		cfg.Logger.Debug("Downgrading DB to 0")
		err = m.Down()
	} else {
		cfg.Logger.Debug("Upgrading DB to the latest version")
		err = m.Up()
	}
	if err != nil {
		if err != migrate.ErrNoChange {
			return 0, false, NewMigrationErr(err.Error(), map[string]any{
				"on": OnNewInstance, "version_limit": cfg.VersionLimit})
		}
		cfg.Logger.Debug("DB is already up to date")
		return ver, dirty, nil
	}

	ver, dirty, err = m.Version()
	if err != nil {
		return ver, dirty,
			NewMigrationErr(err.Error(), map[string]any{"on": OnVersionCheck})
	}
	return ver, dirty, nil
}
