package cache

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

type redisService struct {
	client *redis.Client
}

func NewClientRedis() *redisService {
	return &redisService{
		client: redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_URL"),
			Username: "default",
			Password: os.Getenv("REDIS_PASS"),
			DB:       0,
		}),
	}
}

func (r *redisService) Set(ctx context.Context, key string, item any) {
	result := r.client.Set(ctx, key, item, 1*time.Minute)
	if err := result.Err(); err != nil {
		log.Fatal(err)
	}
}

func (r *redisService) Get(ctx context.Context, key string) *string {
	result, err := r.client.Get(ctx, key).Result()
	if err != nil {
		fmt.Println("nao existe cache")
		return nil
	}
	fmt.Println("achou o cache")
	return &result
}
