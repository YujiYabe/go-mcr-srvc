package gateway

import (
	"context"
)

// GetPatties ...
func (receiver *Gateway) GetPatties(ctx context.Context, requestPatties map[string]int) error {
	err := receiver.ToFreezer.UpdatePatties(ctx, requestPatties)
	if err != nil {
		myErr.Logging(err)
		return err
	}

	return nil
}
