package gateway

import (
	"context"
)

// GetBans ...
func (receiver *Gateway) GetBans(ctx context.Context, requestBans map[string]int) error {
	err := receiver.ToMongo.UpdateBans(ctx, requestBans)
	if err != nil {
		myErr.Logging(err)
		return err
	}

	return nil
}
