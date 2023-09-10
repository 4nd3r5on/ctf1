package db_utils

import (
	"log/slog"

	"gitlab.com/4nd3rs0n/errorsx"
	"gitlab.com/4nd3rs0n/errorsx/slog_helpers"
)

type MigrationErr = errorsx.CustomErr[map[string]any]

// Those are actually not only errors, also used for logging
const (
	OnDbConn       = "Establishing database connection"
	OnNewInstance  = "Creating new migrate instance"
	OnVersionCheck = "Checking database version"
	OnDrop         = "Dropping database"
	OnUpgrade      = "Upgrading database"
)

func NewMigrationErr(err string, vars map[string]any) MigrationErr {
	return errorsx.NewCustomErr(err, vars)
}

func LogMigrationErr(mErr MigrationErr, logger *slog.Logger) {
	log := slog_helpers.MapToSlog(mErr.Info())
	logger.Error(mErr.Error(), log...)
}
