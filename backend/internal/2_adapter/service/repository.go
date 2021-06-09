package service

import (
	"context"

	"backend/internal/4_domain/domain"
	"backend/pkg"
)

var (
	myErr *pkg.MyErr
)

func init() {
	myErr = pkg.NewMyErr("adapter", "service")
}

// Service ...
type Service struct {
	ToRefrigerator ToRefrigerator
	ToFreezer      ToFreezer
	ToShelf        ToShelf
	ToShipment     ToShipment
	ToMonitor      ToMonitor
}

type (
	// ToShipment ...
	ToShipment interface {
		PutProducts(ctx context.Context, order *domain.Order) error
		WriteLog(ctx context.Context, order *domain.Order) error
	}

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

	// ToMonitor ...
	ToMonitor interface {
		UpdateOrders(ctx context.Context, orderNumber string, phase string)
	}
)
