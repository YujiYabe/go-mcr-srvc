package gateway

import (
	domain "backend/internal/4_domain"
	"context"
)

// GetPatties ...
func (receiver *Gateway) GetAllProductList(ctx context.Context) domain.AllProductList {
	return receiver.ToSqlite.GetAllProductList(ctx)
}

// GetPatties ...
func (receiver *Gateway) UpdateProduct(
	ctx context.Context,
	product domain.Product,
) {
	receiver.ToSqlite.UpdateProduct(
		ctx,
		product,
	)
}
