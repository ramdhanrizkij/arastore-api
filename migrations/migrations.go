package migrations

import (
	"database/sql"
	"embed"

	"github.com/pressly/goose/v3"
)

//go:embed *.sql
var embedMigrations embed.FS

func init() {
	goose.SetBaseFS(embedMigrations)
}

// EmbedMigrations returns the embedded SQL migrations filesystem.
func EmbedMigrations() embed.FS {
	return embedMigrations
}

// Migrate applies all pending migrations.
func Migrate(db *sql.DB) error {
	return goose.Up(db, ".")
}

// RollbackLast rolls back the most recent migration.
func RollbackLast(db *sql.DB) error {
	return goose.Down(db, ".")
}

// Redo rolls back and reapplies the last migration.
func Redo(db *sql.DB) error {
	return goose.Redo(db, ".")
}

// Status shows the migration status.
func Status(db *sql.DB) error {
	return goose.Status(db, ".")
}
