package gateway

import (
	domain "backend/internal/4_domain"
	"context"
)

// GetPatties ...
func (gw *Gateway) GetAllProductList(ctx *context.Context) *domain.AllProductList {
	return gw.ToSqlite.GetAllProductList(ctx)
}
