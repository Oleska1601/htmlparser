package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	_defaultTTL = time.Hour
)

type RedisClient struct {
	client *redis.Client
	ttl    time.Duration
}

func New(ctx context.Context, addr, password string, ttl time.Duration) (*RedisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})

	if _, err := client.Ping(ctx).Result(); err != nil {
		return nil, fmt.Errorf("client.Ping(ctx).Result: %w", err)
	}
	rc := &RedisClient{
		client: client,
		ttl:    ttl,
	}
	return rc, nil
}
