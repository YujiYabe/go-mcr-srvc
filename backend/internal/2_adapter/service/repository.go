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
		ISHandOver(ctx context.Context, order *domain.Order) error
		ISLogging(ctx context.Context, order *domain.Order) error
	}

	// ToRefrigerator ...
	ToRefrigerator interface {
		ISGetVegetables(ctx context.Context, items map[string]int) error
		ISGetIngredients(ctx context.Context, items map[string]int) error
	}

	// ToFreezer ...
	ToFreezer interface {
		ISGetPatties(ctx context.Context, items map[string]int) error
	}

	// ToShelf ...
	ToShelf interface {
		ISGetBans(ctx context.Context, items map[string]int) error
	}

	// ToMonitor ...
	ToMonitor interface {
		ISUpdateOrders(ctx context.Context, orderNumber string, phase string)
	}
)
