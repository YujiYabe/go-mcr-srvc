package service

import (
	"context"
)

// GetVegetables2 ...
func (sv *Service) GetVegetables2(ctx context.Context, requestVegetables map[string]int) error {
	err := sv.ToRefrigerator.GetVegetables(ctx, requestVegetables)
	if err != nil {
		return err
	}

	return nil
}
