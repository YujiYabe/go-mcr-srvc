package gateway

import (
	"backend/pkg"
	"context"
)

// GetBans ...
func (receiver *Gateway) GetBans(ctx context.Context, requestBans map[string]int) error {
	err := receiver.ToMongo.UpdateBans(ctx, requestBans)
	if err != nil {
		pkg.Logging(ctx, err)
		return err
	}

	return nil
}
