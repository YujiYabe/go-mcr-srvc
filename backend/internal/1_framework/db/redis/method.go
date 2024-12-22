package redis

import (
	"context"
	// "backend/internal/env"
	// "backend/internal/pkg"
)

// ResetPlaceListInRedis ...
func (receiver *Redis) ResetPlaceListInRedis(
	ctx context.Context,
) (
	err error,
) {
	_, err = receiver.Conn.Del(
		ctx,
		"placeList",
	).Result()
	if err != nil {
		// pkg.Logging(ctx, err)
	}

	return
}
