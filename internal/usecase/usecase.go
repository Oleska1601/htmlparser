package usecase

import "htmlparser/pkg/logger"

type Usecase struct {
	cache  CacheInterface
	pgRepo PgRepoInterface
	logger logger.LoggerInterface
}

func New(cache CacheInterface, pgRepo PgRepoInterface, logger logger.LoggerInterface) *Usecase {
	return &Usecase{
		cache:  cache,
		pgRepo: pgRepo,
		logger: logger,
	}
}
