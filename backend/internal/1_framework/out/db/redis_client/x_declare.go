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

// NewToRedis ...
func NewToRedis() (
	toRedis gatewayDB.ToRedis,
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
