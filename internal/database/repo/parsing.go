package repo

import (
	"context"
	"fmt"
	"htmlparser/internal/entity"

	"github.com/Masterminds/squirrel"
)

func (r *PgRepo) GetParsing(ctx context.Context, url string) (entity.Parsing, error) {
	sql, args, err := r.postgres.Squirrel.Select("url", "data").From("parsings").Where(squirrel.Eq{"url": url}).ToSql()
	if err != nil {
		return entity.Parsing{}, fmt.Errorf("r.postgres.Squirrel.Select: %w", err)
	}
	row := r.postgres.Pool.QueryRow(ctx, sql, args...)
	var parsing entity.Parsing
	if err := row.Scan(&parsing.URL, &parsing.Data); err != nil {
		return entity.Parsing{}, fmt.Errorf("row.Scan: %w", err)
	}
	return parsing, nil
}

func (r *PgRepo) AddParsing(ctx context.Context, url string, data string) error {
	sql, args, err := r.postgres.Squirrel.Insert("parsings").
		Columns("url", "data").Values(url, data).ToSql()
	if err != nil {
		return fmt.Errorf("r.postgres.Squirrel.Insert: %w", err)
	}
	_, err = r.postgres.Pool.Exec(ctx, sql, args...)

	if err != nil {
		return fmt.Errorf("r.postgres.Pool.Exec: %w", err)
	}
	return nil
}
