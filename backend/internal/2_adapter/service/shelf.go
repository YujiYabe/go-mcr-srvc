package service

import (
	"context"
)

// GetBans ...
func (sv *Service) GetBans(ctx context.Context, requestBans map[string]int) error {
	err := sv.ToShelf.ISGetBans(ctx, requestBans)
	if err != nil {
		myErr.Logging(err)
		return err
	}

	return nil
}
