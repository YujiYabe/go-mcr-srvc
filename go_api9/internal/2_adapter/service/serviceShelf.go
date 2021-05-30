package service

import (
	"context"
)

// GetVegetables3 ...
func (sv *Service) GetVegetables3(ctx context.Context, requestVegetables map[string]int) error {
	err := sv.ToShelf.GetVegetables(ctx, requestVegetables)
	if err != nil {
		return err
	}

	return nil
}
