package service

import (
	"context"
)

// GetVegetables ...
func (sv *Service) GetVegetables(ctx context.Context, requestVegetables map[string]int) error {
	err := sv.ToRefrigerator.GetVegetables(ctx, requestVegetables)
	if err != nil {
		return err
	}

	return nil
}

// Dummy ...
func (sv *Service) Dummy(ctx context.Context) (string, error) {
	res, _ := sv.ToRefrigerator.Dummy(ctx)

	return res, nil
}
