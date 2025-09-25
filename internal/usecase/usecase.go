package usecase

import "htmlparser/internal/logging"

type Usecase struct {
	cache  CacheInterface
	pgRepo PgRepoInterface
	logger logging.LoggerInterface
}

func New(cache CacheInterface, pgRepo PgRepoInterface, logger logging.LoggerInterface) *Usecase {
	return &Usecase{
		cache:  cache,
		pgRepo: pgRepo,
		logger: logger,
	}
}
