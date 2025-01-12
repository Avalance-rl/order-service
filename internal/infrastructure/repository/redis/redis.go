package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	rdb *redis.Client
}

func New(ctx context.Context, addr, pass string) *Redis {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       0,
	})

	return &Redis{rdb}
}
