package go_redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var RedisDB *redis.Client
var Gctx context.Context

func init() {
	RedisDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6380",
		Password: "",
		DB:       0,
	})
	Gctx = context.Background()
}
