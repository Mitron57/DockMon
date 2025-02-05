package postgres

import (
    "database/sql"
    "embed"
    "github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var migration embed.FS

func PerformMigration(db *sql.DB) error {
    goose.SetBaseFS(migration)
    if err := goose.Up(db, "migrations"); err != nil {
        return err
    }
    return nil
}
