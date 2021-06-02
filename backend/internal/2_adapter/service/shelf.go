package service

import (
	"context"
)

// GetBans ...
func (sv *Service) GetBans(ctx context.Context, requestBans map[string]int) error {
	err := sv.ToShelf.GetBans(ctx, requestBans)
	if err != nil {
		return err
	}

	return nil
}
