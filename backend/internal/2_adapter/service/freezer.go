package service

import (
	"context"
)

// GetPatties ...
func (sv *Service) GetPatties(ctx context.Context, requestPatties map[string]int) error {
	err := sv.ToFreezer.UpdatePatties(ctx, requestPatties)
	if err != nil {
		myErr.Logging(err)
		return err
	}

	return nil
}
