package redis_client

import (
	"github.com/redis/go-redis/v9"

	"backend/internal/2_adapter/gateway"
)

type (
	// RedisClient ...
	RedisClient struct {
		Conn *redis.Client
	}
)

// NewToRedis ...
func NewToRedis() (
	toRedis gateway.ToRedis,
) {
	conn := redis.NewClient(&redis.Options{
		Addr:     "localhost:6739",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	redisClient := new(RedisClient)
	redisClient.Conn = conn

	toRedis = redisClient
	return
}

func init() {
}
