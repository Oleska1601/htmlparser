package redis

import (
	"context"
	"fmt"
)

func (rc *RedisClient) GetValue(ctx context.Context, key string) (string, error) {
	value, err := rc.client.Get(ctx, key).Result()
	if err != nil {
		return "", fmt.Errorf("rc.client.Get: %w", err)
	}
	return value, nil
}

func (rc *RedisClient) SetValue(ctx context.Context, key string, value string) error {
	err := rc.client.Set(ctx, key, value, rc.ttl).Err()
	if err != nil {
		return fmt.Errorf("rc.client.Set: %w", err)
	}
	return nil
}
