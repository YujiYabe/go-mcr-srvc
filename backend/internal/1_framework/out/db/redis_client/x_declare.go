package redis_client

import (
	"github.com/redis/go-redis/v9"

	gatewayDB "backend/internal/2_adapter/gateway/db"
)

type (
	// RedisClient ...
	RedisClient struct {
		Conn *redis.Client
	}
)

func NewToRedis(
	addr string,
	password string,
	db int,
) (
	toRedis gatewayDB.ToRedis,
) {
	conn := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	redisClient := new(RedisClient)
	redisClient.Conn = conn

	toRedis = redisClient

	return
}
