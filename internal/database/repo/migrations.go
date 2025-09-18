package repo

import (
	"database/sql"
	"embed"
	"fmt"

	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func (pgRepo *PgRepo) ApplyMigrations() error {

	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("goose.SetDialect: %w", err)
	}
	connector := stdlib.GetPoolConnector(pgRepo.postgres.Pool)
	db := sql.OpenDB(connector)
	defer db.Close()
	if err := goose.Up(db, "migrations"); err != nil {
		return fmt.Errorf("goose.Up: %w", err)
	}
	return nil
}
