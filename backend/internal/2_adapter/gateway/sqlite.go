package gateway

import (
	domain "backend/internal/4_domain"
	"context"
)

// GetPatties ...
func (gw *Gateway) GetAllProductList(ctx context.Context) *domain.AllProductList {
	return gw.ToSqlite.GetAllProductList(ctx)
}

// GetPatties ...
func (gw *Gateway) UpdateProduct(
	ctx context.Context,
	product domain.Product,
) {
	gw.ToSqlite.UpdateProduct(
		ctx,
		product,
	)
}
