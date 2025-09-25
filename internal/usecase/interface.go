package usecase

import (
	"context"
	"htmlparser/internal/entity"
)

//go:generate go run github.com/vektra/mockery/v2@latest --name=CacheInterface
type CacheInterface interface {
	GetValue(context.Context, string) (string, error)
	SetValue(context.Context, string, string) error
}

//go:generate go run github.com/vektra/mockery/v2@latest --name=PgRepoInterface
type PgRepoInterface interface {
	GetParsing(context.Context, string) (entity.Parsing, error)
	AddParsing(context.Context, string, string) error
}
