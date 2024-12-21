package gateway

import (
	"context"

	"backend/pkg"
)

var (
	myErr *pkg.MyErr
)

func init() {
	myErr = pkg.NewMyErr("interface_adapter", "gateway")
}

type (
	Gateway struct {
		ToPostgres ToPostgres
		ToMySQL    ToMySQL
		ToMongo    ToMongo
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
func NewGateway(ToPostgres ToPostgres, toMySQL ToMySQL, toMongo ToMongo) *Gateway {
	return &Gateway{
		ToPostgres: ToPostgres,
		ToMySQL:    toMySQL,
		ToMongo:    toMongo,
	}
}
