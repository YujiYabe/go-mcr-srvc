package controller

import (
	"context"

	domain "backend/internal/4_domain"
)

// GetProduct ...
func (receiver *controller) GetProduct(
	ctx context.Context,
	productNumber int,
) domain.Product {
	return receiver.ToUseCase.GetProduct(
		ctx,
		productNumber,
	)
}

// GetAllProductList ...
func (receiver *controller) GetAllProductList(
	ctx context.Context,
) domain.AllProductList {
	return receiver.ToUseCase.GetAllProductList(
		ctx,
	)
}

// GetProductList ...
func (receiver *controller) GetProductList(
	ctx context.Context,
) domain.ProductList {
	return receiver.ToUseCase.GetProductList(
		ctx,
	)
}

// UpdateProduct ...
func (receiver *controller) UpdateProduct(
	ctx context.Context,
	product domain.Product,
) {
	receiver.ToUseCase.UpdateProduct(
		ctx,
		product,
	)
}
