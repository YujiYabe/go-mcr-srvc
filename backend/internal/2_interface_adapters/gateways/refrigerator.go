package gateways

import (
	"context"
)

// GetVegetables ...
func (sv *Service) GetVegetables(ctx context.Context, requestVegetables map[string]int) error {
	err := sv.ToRefrigerator.UpdateVegetables(ctx, requestVegetables)
	if err != nil {
		myErr.Logging(err)
		return err
	}

	return nil
}

// GetIngredients ...
func (sv *Service) GetIngredients(ctx context.Context, requestIngredients map[string]int) error {
	err := sv.ToRefrigerator.UpdateIngredients(ctx, requestIngredients)
	if err != nil {
		myErr.Logging(err)
		return err
	}

	return nil
}
