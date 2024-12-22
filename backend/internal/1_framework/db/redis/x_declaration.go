package redis

import (
	"github.com/redis/go-redis/v9"

	"backend/internal/2_adapter/gateway"
)

type (
	// Redis ...
	Redis struct {
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

	redis := new(Redis)
	redis.Conn = conn

	toRedis = redis
	return
}

func init() {
}
