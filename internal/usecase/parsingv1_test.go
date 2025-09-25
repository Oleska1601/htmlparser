package usecase_test

import (
	"context"
	"database/sql"
	"errors"
	"htmlparser/internal/entity"
	loggingMocks "htmlparser/internal/logging/mocks"
	"htmlparser/internal/usecase"
	usecaseMocks "htmlparser/internal/usecase/mocks"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/mock"
)

func TestUsecase_GetParsingDataV1(t *testing.T) {
	type args struct {
		ctx context.Context
		url string
	}
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		args       args
		setupMocks func(cache *usecaseMocks.CacheInterface, pgRepo *usecaseMocks.PgRepoInterface,
			logger *loggingMocks.LoggerInterface, url string)
		// Named input parameters for target function.
		//want    string
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				ctx: context.Background(),
				url: "https://google.com",
			},
			setupMocks: func(cache *usecaseMocks.CacheInterface, pgRepo *usecaseMocks.PgRepoInterface,
				logger *loggingMocks.LoggerInterface, url string) {
				cache.On("GetValue", mock.Anything, url).Return("", nil)
			},

			wantErr: false,
		},
		{
			name: "test2",
			args: args{
				ctx: context.Background(),
				url: "https://google.com",
			},
			setupMocks: func(cache *usecaseMocks.CacheInterface, pgRepo *usecaseMocks.PgRepoInterface,
				logger *loggingMocks.LoggerInterface, url string) {
				cache.On("GetValue", mock.Anything, url).Return("", redis.Nil)
				pgRepo.On("GetParsing", mock.Anything, url).Return(entity.Parsing{}, nil)
			},

			wantErr: false,
		},
		{
			name: "test3",
			args: args{
				ctx: context.Background(),
				url: "https://google.com",
			},
			setupMocks: func(cache *usecaseMocks.CacheInterface, pgRepo *usecaseMocks.PgRepoInterface,
				logger *loggingMocks.LoggerInterface, url string) {
				cache.On("GetValue", mock.Anything, url).Return("", errors.New(""))
				logger.On("Error", "GetParsingDataV1 u.cache.GetValue", mock.Anything).Return()
			},

			wantErr: true,
		},
		{
			name: "test4",
			args: args{
				ctx: context.Background(),
				url: "https://google.com",
			},
			setupMocks: func(cache *usecaseMocks.CacheInterface, pgRepo *usecaseMocks.PgRepoInterface,
				logger *loggingMocks.LoggerInterface, url string) {
				cache.On("GetValue", mock.Anything, url).Return("", redis.Nil)
				pgRepo.On("GetParsing", mock.Anything, url).Return(entity.Parsing{}, sql.ErrNoRows)
				pgRepo.On("AddParsing", mock.Anything, url, mock.Anything).Return(nil)
				cache.On("SetValue", mock.Anything, url, mock.Anything).Return(nil)
			},

			wantErr: false,
		},
		{
			name: "test5",
			args: args{
				ctx: context.Background(),
				url: "https://google.com",
			},
			setupMocks: func(cache *usecaseMocks.CacheInterface, pgRepo *usecaseMocks.PgRepoInterface,
				logger *loggingMocks.LoggerInterface, url string) {
				cache.On("GetValue", mock.Anything, url).Return("", redis.Nil)
				pgRepo.On("GetParsing", mock.Anything, url).Return(entity.Parsing{}, errors.New(""))
				logger.On("Error", "GetParsingDataV1 u.pgRepo.GetParsing", mock.Anything).Return()
			},

			wantErr: true,
		},
		{
			name: "test6",
			args: args{
				ctx: context.Background(),
				url: "https://google.com",
			},
			setupMocks: func(cache *usecaseMocks.CacheInterface, pgRepo *usecaseMocks.PgRepoInterface,
				logger *loggingMocks.LoggerInterface, url string) {
				cache.On("GetValue", mock.Anything, url).Return("", redis.Nil)
				pgRepo.On("GetParsing", mock.Anything, url).Return(entity.Parsing{}, sql.ErrNoRows)
				pgRepo.On("AddParsing", mock.Anything, url, mock.Anything).Return(errors.New(""))
				logger.On("Error", "GetParsingDataV1 u.pgRepo.AddParsing", mock.Anything).Return()
			},

			wantErr: true,
		},
		{
			name: "test7",
			args: args{
				ctx: context.Background(),
				url: "https://google.com",
			},
			setupMocks: func(cache *usecaseMocks.CacheInterface, pgRepo *usecaseMocks.PgRepoInterface,
				logger *loggingMocks.LoggerInterface, url string) {
				cache.On("GetValue", mock.Anything, url).Return("", redis.Nil)
				pgRepo.On("GetParsing", mock.Anything, url).Return(entity.Parsing{}, sql.ErrNoRows)
				pgRepo.On("AddParsing", mock.Anything, url, mock.Anything).Return(nil)
				cache.On("SetValue", mock.Anything, url, mock.Anything).Return(errors.New(""))
				logger.On("Error", "GetParsingDataV1 u.cache.SetValue", mock.Anything).Return()
			},

			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache := usecaseMocks.NewCacheInterface(t)
			pgRepo := usecaseMocks.NewPgRepoInterface(t)
			logger := loggingMocks.NewLoggerInterface(t)

			tt.setupMocks(cache, pgRepo, logger, tt.args.url)

			u := usecase.New(cache, pgRepo, logger)
			_, gotErr := u.GetParsingDataV1(tt.args.ctx, tt.args.url)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("GetParsingDataV1() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("GetParsingDataV1() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want. - не проверяю, что имеенно за значение возвратиться, просто наличие ошибки
			//if true {
			//	t.Errorf("GetParsingDataV1() = %v, want %v", got, tt.want)
			//}
		})
	}
}
