package gateway

import (
	"context"

	domain "backend/internal/4_domain"
)

type Gateway struct {
	ToSqlite
}

// NewGateway ...
func NewGateway(
	toSqlite ToSqlite,
) *Gateway {
	return &Gateway{
		ToSqlite: toSqlite,
	}
}

type ToSqlite interface {
	GetAllProductList(ctx context.Context) domain.AllProductList

	UpdateProduct(
		ctx context.Context,
		product domain.Product,
	)
}
