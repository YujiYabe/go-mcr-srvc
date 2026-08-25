package redis_client

import (
	"github.com/redis/go-redis/v9"

	gatewayDB "backend/internal/2_adapter/gateway/db"
	"backend/internal/env"
)

type (
	// RedisClient ...
	RedisClient struct {
		Conn *redis.Client
	}
)

// NewToRedis ...
func NewToRedis() (
	toRedis gatewayDB.ToRedis,
) {
	conn := redis.NewClient(&redis.Options{
		Addr:     env.RedisConfig.Addr,
		Password: env.RedisConfig.Password,
		DB:       env.RedisConfig.DB,
	})

	redisClient := new(RedisClient)
	redisClient.Conn = conn

	toRedis = redisClient
	return
}

func init() {
}
