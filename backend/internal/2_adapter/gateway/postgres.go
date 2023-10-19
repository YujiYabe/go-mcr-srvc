package gateway

import (
	"context"
)

// GetVegetables ...
func (gw *Gateway) GetVegetables(ctx context.Context, requestVegetables map[string]int) error {
	err := gw.ToPostgres.UpdateVegetables(ctx, requestVegetables)
	if err != nil {
		myErr.Logging(err)
		return err
	}

	return nil
}

// GetIngredients ...
func (gw *Gateway) GetIngredients(ctx context.Context, requestIngredients map[string]int) error {
	err := gw.ToPostgres.UpdateIngredients(ctx, requestIngredients)
	if err != nil {
		myErr.Logging(err)
		return err
	}

	return nil
}
