package gateway

import (
	"backend/pkg"
	"context"
)

// GetPatties ...
func (receiver *Gateway) GetPatties(ctx context.Context, requestPatties map[string]int) error {
	err := receiver.ToMySQL.UpdatePatties(ctx, requestPatties)
	if err != nil {
		pkg.Logging(ctx, err)
		return err
	}

	return nil
}
