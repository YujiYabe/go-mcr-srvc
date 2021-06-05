package service

import (
	"context"

	"backend/internal/4_domain/domain"
)

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
		HandOver(ctx context.Context, order *domain.Order) error
		Logging(ctx context.Context, order *domain.Order) error
	}

	// ToRefrigerator ...
	ToRefrigerator interface {
		GetVegetables(ctx context.Context, items map[string]int) error
		GetIngredients(ctx context.Context, items map[string]int) error
		Dummy(ctx context.Context) (string, error)
	}

	// ToFreezer ...
	ToFreezer interface {
		GetPatties(ctx context.Context, items map[string]int) error
		Dummy(ctx context.Context) (string, error)
	}

	// ToShelf ...
	ToShelf interface {
		GetBans(ctx context.Context, items map[string]int) error
		Dummy(ctx context.Context) (string, error)
	}

	// ToClient ...
	ToClient interface {
		HandOver(ctx context.Context) error
	}

	// ToMonitor ...
	ToMonitor interface {
		UpdateOrders(ctx context.Context, order *domain.Order, phase string) error
	}

	// // DatabaseResult ...
	// DatabaseResult interface {
	// 	LastInsertId() (int64, error)
	// 	RowsAffected() (int64, error)
	// }
)
