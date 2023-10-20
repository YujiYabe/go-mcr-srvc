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
		ToSqlite   ToSqlite
		ToPostgres ToPostgres
		ToMysql    ToMysql
		ToMongo    ToMongo
	}

	ToSqlite interface {
	}

	// ToPostgres ...
	ToPostgres interface {
		UpdateVegetables(ctx context.Context, items map[string]int) error
		UpdateIngredients(ctx context.Context, items map[string]int) error
	}

	// ToMysql ...
	ToMysql interface {
		UpdatePatties(ctx context.Context, items map[string]int) error
	}

	// ToMongo ...
	ToMongo interface {
		UpdateBans(ctx context.Context, items map[string]int) error
	}
)

// NewGateway ...
func NewGateway(
	toSqlite ToSqlite,
	toPostgres ToPostgres,
	toMysql ToMysql,
	toMongo ToMongo,
) *Gateway {
	return &Gateway{
		ToSqlite:   toSqlite,
		ToPostgres: toPostgres,
		ToMysql:    toMysql,
		ToMongo:    toMongo,
	}
}
