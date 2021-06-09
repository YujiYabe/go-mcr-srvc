package service

import (
	"context"
)

// GetVegetables ...
func (sv *Service) GetVegetables(ctx context.Context, requestVegetables map[string]int) error {
	err := sv.ToRefrigerator.ISGetVegetables(ctx, requestVegetables)
	if err != nil {
		myErr.Logging(err)
		return err
	}

	return nil
}

// GetIngredients ...
func (sv *Service) GetIngredients(ctx context.Context, requestIngredients map[string]int) error {
	err := sv.ToRefrigerator.ISGetIngredients(ctx, requestIngredients)
	if err != nil {
		myErr.Logging(err)
		return err
	}

	return nil
}
