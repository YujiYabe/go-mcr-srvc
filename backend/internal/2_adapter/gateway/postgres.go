package gateway

import (
	"context"
)

// GetVegetables ...
func (receiver *Gateway) GetVegetables(ctx context.Context, requestVegetables map[string]int) error {
	err := receiver.ToPostgres.UpdateVegetables(ctx, requestVegetables)
	if err != nil {
		myErr.Logging(err)
		return err
	}

	return nil
}

// GetIngredients ...
func (receiver *Gateway) GetIngredients(ctx context.Context, requestIngredients map[string]int) error {
	err := receiver.ToPostgres.UpdateIngredients(ctx, requestIngredients)
	if err != nil {
		myErr.Logging(err)
		return err
	}

	return nil
}
