package service

import (
	"context"
	"fmt"
)

// Service ...
type Service struct {
	ToStocker ToStocker
}

// Order ...
func (sv *Service) Order(ctx context.Context) (string, error) {
	res, _ := sv.ToStocker.StockFind(ctx)

	return res, nil
}

// GetVegetables ...
func (sv *Service) GetVegetables(ctx context.Context, requestVegetables map[string]int) error {
	err := sv.ToStocker.GetVegetables(ctx, requestVegetables)
	if err != nil {
		return err
	}

	res, _ := sv.ToStocker.StockFind(ctx)
	fmt.Println("==============================")
	fmt.Printf("%#v\n", res)
	fmt.Println("==============================")

	return nil
}

// Dummy ...
func (sv *Service) Dummy(ctx context.Context) (string, error) {
	res, _ := sv.ToStocker.Dummy(ctx)

	return res, nil
}
