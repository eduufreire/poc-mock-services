package cache

import "context"

type CacheService interface {
	Set(ctx context.Context, key string, item any)
	Get(ctx context.Context, key string) *string
}
