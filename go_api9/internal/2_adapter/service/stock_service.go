package service

import (
	"context"
)

// Service ...
type Service struct {
	ToStocker ToStocker
}

// GetVegetables ...
func (sv *Service) GetVegetables(ctx context.Context, requestVegetables map[string]int) error {
	err := sv.ToStocker.GetVegetables(ctx, requestVegetables)
	if err != nil {
		return err
	}

	return nil
}

// Dummy ...
func (sv *Service) Dummy(ctx context.Context) (string, error) {
	res, _ := sv.ToStocker.Dummy(ctx)

	return res, nil
}
