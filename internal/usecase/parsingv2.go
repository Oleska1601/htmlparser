package usecase

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"

	"github.com/redis/go-redis/v9"
)

// tokenizer parser
func (u *Usecase) GetParsingDataV2(ctx context.Context, url string) (string, error) {
	data, err := u.cache.GetValue(ctx, url)
	if err != nil && !errors.Is(err, redis.Nil) {
		u.logger.Error("GetParsingDataV2 u.cache.GetValue", slog.Any("error", err))
		return "", errors.New("error of getting parsing data")
	} else if err == nil {
		return data, nil
	}

	parsing, err := u.pgRepo.GetParsing(ctx, url)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		u.logger.Error("GetParsingDataV2 u.pgRepo.GetParsing", slog.Any("error", err))
		return "", errors.New("error of getting parsing data")
	} else if err == nil {
		return parsing.Data, nil
	}
	data, err = tokenizerParser(url)
	if err != nil {
		u.logger.Error("GetParsingDataV2 nodeParser", slog.Any("error", err))
		return "", errors.New("error of getting parsing data")
	}
	err = u.pgRepo.AddParsing(ctx, url, data)
	if err != nil {
		u.logger.Error("GetParsingDataV2 u.pgRepo.AddParsing", slog.Any("error", err))
		return "", errors.New("error of getting parsing data")
	}
	err = u.cache.SetValue(ctx, url, data)
	if err != nil {
		u.logger.Error("GetParsingDataV2 u.cache.SetValue", slog.Any("error", err))
		return "", errors.New("error of getting parsing data")
	}
	return data, nil
}
