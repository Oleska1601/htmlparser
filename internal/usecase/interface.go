package usecase

import (
	"context"
	"htmlparser/internal/entity"
)

type CacheInterface interface {
	GetValue(context.Context, string) (string, error)
	SetValue(context.Context, string, string) error
}

type PgRepoInterface interface {
	GetParsing(context.Context, string) (entity.Parsing, error)
	AddParsing(context.Context, string, string) error
}
