package redis_client

import (
	"context"

	logger "backend/internal/logger"
)

// ResetPlaceListInRedis ...
func (receiver *RedisClient) ResetPlaceListInRedis(
	ctx context.Context,
) (
	err error,
) {
	_, err = receiver.Conn.Del(
		ctx,
		"placeList",
	).Result()
	if err != nil {
		logger.Logging(ctx, err)
	}

	return
}
