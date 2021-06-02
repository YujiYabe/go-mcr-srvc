package service

import (
	"context"

	"app/internal/4_domain/domain"
)

// Service ...
type Service struct {
	ToRefrigerator ToRefrigerator
	ToFreezer      ToFreezer
	ToShelf        ToShelf
	ToShipment     ToShipment
}

type (
	// ToRefrigerator ...
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

	// ToGrpcOut ...
	ToGrpcOut interface {
		// IsSendContent(address string, cc *shared.CommonContent) (string, error)
		// IsReceiveContent(address, funcName string) (string, error)
		// IsFileUpload(address, fileName string, fileBody *bytes.Buffer) (string, error)
	}

	// ToWsOrder ...
	ToWsOrder interface {
		// IsSendToAgent(agentID string, cc *shared.CommonContent)
	}

	// ToClient ...
	ToClient interface {
		HandOver(ctx context.Context) error
	}

	// // DatabaseResult ...
	// DatabaseResult interface {
	// 	LastInsertId() (int64, error)
	// 	RowsAffected() (int64, error)
	// }
)
