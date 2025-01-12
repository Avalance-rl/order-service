package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	rdb *redis.Client
}

func NewCache(addr, pass string) (*Redis, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       0,
	})

	stcmd := rdb.Ping(context.Background())
	if err := stcmd.Err(); err != nil {
		return nil, err
	}

	return &Redis{rdb}, nil
}
