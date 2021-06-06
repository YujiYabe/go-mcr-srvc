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

// GetIngredients ...
func (sv *Service) GetIngredients(ctx context.Context, requestIngredients map[string]int) error {
	err := sv.ToRefrigerator.GetIngredients(ctx, requestIngredients)
	if err != nil {
		return err
	}

	return nil
}
