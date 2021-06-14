package gateway

import (
	"context"

	"backend/pkg"
)

var (
	myErr *pkg.MyErr
)

func init() {
	myErr = pkg.NewMyErr("adapter", "gateway")
}

// Gateway ...
type Gateway struct {
	ToRefrigerator ToRefrigerator
	ToFreezer      ToFreezer
	ToShelf        ToShelf
}

type (

	// ToRefrigerator ...
	ToRefrigerator interface {
		UpdateVegetables(ctx context.Context, items map[string]int) error
		UpdateIngredients(ctx context.Context, items map[string]int) error
	}

	// ToFreezer ...
	ToFreezer interface {
		UpdatePatties(ctx context.Context, items map[string]int) error
	}

	// ToShelf ...
	ToShelf interface {
		UpdateBans(ctx context.Context, items map[string]int) error
	}
)
