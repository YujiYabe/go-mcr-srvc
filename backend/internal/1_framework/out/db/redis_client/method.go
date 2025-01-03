package redis_client

import (
	"backend/pkg"
	"context"
	// "backend/internal/env"
	// "backend/internal/pkg"
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
		pkg.Logging(ctx, err)
	}

	return
}
