package ports

import (
	"context"
	"time"
)

type CacheRepository interface {
	Set(ctx context.Context, string, value string, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Expire(ctx context.Context, key string, duration time.Duration) error
	Delete(ctx context.Context, key string) error
}
