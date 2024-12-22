package gateway

import (
	"context"
)

type (
	Gateway struct {
		ToRedis    ToRedis
		ToPostgres ToPostgres
		ToMySQL    ToMySQL
		ToMongo    ToMongo
	}

	// ToRedis ...
	ToRedis interface {
		ResetPlaceListInRedis(
			ctx context.Context,
		) error
	}

	// ToPostgres ...
	ToPostgres interface {
		UpdateVegetables(ctx context.Context, items map[string]int) error
		UpdateIngredients(ctx context.Context, items map[string]int) error
	}

	// ToMySQL ...
	ToMySQL interface {
		UpdatePatties(ctx context.Context, items map[string]int) error
	}

	// ToMongo ...
	ToMongo interface {
		UpdateBans(ctx context.Context, items map[string]int) error
	}
)

// NewGateway ...
func NewGateway(
	toRedis ToRedis,
	toPostgres ToPostgres,
	toMySQL ToMySQL,
	toMongo ToMongo,
) *Gateway {
	return &Gateway{
		ToRedis:    toRedis,
		ToPostgres: toPostgres,
		ToMySQL:    toMySQL,
		ToMongo:    toMongo,
	}
}
