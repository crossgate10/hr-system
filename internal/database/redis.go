package database

import (
	"github.com/go-redis/redis/v8"

	"hr-system/internal/config"
)

func InitRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: config.Get().Redis.Addr,
	})
	return rdb
}
