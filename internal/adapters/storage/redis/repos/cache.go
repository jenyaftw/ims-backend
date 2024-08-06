package repos

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type CacheRepository struct {
	rdb *redis.Client
}

func NewCacheRepository(rdb *redis.Client) *CacheRepository {
	return &CacheRepository{
		rdb: rdb,
	}
}

func (r CacheRepository) Set(ctx context.Context, key string, value string, expiration time.Duration) error {
	return r.rdb.Set(ctx, key, value, expiration).Err()
}

func (r CacheRepository) Get(ctx context.Context, key string) (string, error) {
	return r.rdb.Get(ctx, key).Result()
}

func (r CacheRepository) Expire(ctx context.Context, key string, duration time.Duration) error {
	return r.rdb.Expire(ctx, key, duration).Err()
}

func (r CacheRepository) Delete(ctx context.Context, key string) error {
	return r.rdb.Del(ctx, key).Err()
}
