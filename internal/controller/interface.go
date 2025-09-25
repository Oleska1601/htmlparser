package controller

import "context"

//go:generate go run github.com/vektra/mockery/v2@latest --name=UsecaseInterface
type UsecaseInterface interface {
	GetParsingDataV1(context.Context, string) (string, error)
	GetParsingDataV2(context.Context, string) (string, error)
}
