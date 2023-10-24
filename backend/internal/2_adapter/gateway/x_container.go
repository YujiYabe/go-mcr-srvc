package gateway

import (
	"context"

	domain "backend/internal/4_domain"
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
		ToSqlite
	}

	ToSqlite interface {
		GetAllProductList(ctx context.Context) *domain.AllProductList

		UpdateProduct(
			ctx context.Context,
			product domain.Product,
		)
	}
)

// NewGateway ...
func NewGateway(
	toSqlite ToSqlite,
) *Gateway {
	return &Gateway{
		ToSqlite: toSqlite,
	}
}
