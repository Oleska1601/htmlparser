package repo

import "htmlparser/pkg/postgres"

type PgRepo struct {
	postgres *postgres.Postgres
}

func New(postgres *postgres.Postgres) *PgRepo {
	return &PgRepo{
		postgres: postgres,
	}
}
