package postgres

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	_defaultMaxPoolSize     = 1
	_defaultMaxConnAttempts = 10
	_defaultMaxConnTimeout  = time.Second
)

type Postgres struct {
	maxPoolSize     int
	maxConnAttempts int
	maxConnTimeout  time.Duration
	Squirrel        squirrel.StatementBuilderType
	Pool            *pgxpool.Pool
}

func New(ctx context.Context, pgUrl string, options ...Option) (*Postgres, error) {
	postgres := &Postgres{
		maxPoolSize:     _defaultMaxPoolSize,
		maxConnAttempts: _defaultMaxConnAttempts,
		maxConnTimeout:  _defaultMaxConnTimeout,
	}
	for _, option := range options {
		option(postgres)
	}

	postgres.Squirrel = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	pgxConfig, err := pgxpool.ParseConfig(pgUrl)
	if err != nil {
		return nil, fmt.Errorf("pgxpool.ParseConfig: %w", err)
	}
	pgxConfig.MaxConns = int32(postgres.maxConnAttempts)
	for postgres.maxConnAttempts > 0 {
		postgres.Pool, err = pgxpool.NewWithConfig(ctx, pgxConfig)
		if err == nil {
			break
		}
		postgres.maxConnAttempts--
		slog.Info("postgres is trying to connect", slog.Int("attempts left", postgres.maxConnAttempts))
		time.Sleep(postgres.maxConnTimeout)
	}
	if err != nil {
		return nil, fmt.Errorf("postgres.maxConnAttempts = 0")
	}
	return postgres, nil

}

func (p *Postgres) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}
