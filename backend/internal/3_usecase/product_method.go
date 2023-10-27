package usecase

import (
	"context"

	domain "backend/internal/4_domain"
)

// product -----------------------
// GetProduct ...
func (receiver *useCase) GetProduct(
	ctx context.Context,
	productNumber int,
) domain.Product {
	return receiver.ToDomain.GetProduct(
		ctx,
		productNumber,
	)
}

// GetAllProductList ...
func (receiver *useCase) GetAllProductList(
	ctx context.Context,
) domain.AllProductList {
	return receiver.ToDomain.GetAllProductList(
		ctx,
	)
}

// GetProductList ...
func (receiver *useCase) GetProductList(
	ctx context.Context,
) domain.ProductList {
	return receiver.ToDomain.GetProductList(
		ctx,
	)
}

// UpdateProduct ...
func (receiver *useCase) UpdateProduct(
	ctx context.Context,
	product domain.Product,
) {

	// DB更新
	receiver.ToGateway.UpdateProduct(
		ctx,
		product,
	)

	// インメモリの情報を更新
	receiver.SetUpInMemory(ctx)
}

// SetUpInMemory ...
func (receiver *useCase) SetUpInMemory(
	ctx context.Context,
) {
	// localDBから全商品を取得
	allProductList := receiver.ToGateway.GetAllProductList(
		ctx,
	)

	receiver.ToDomain.SaveInMemory(
		ctx,
		allProductList,
	)

}
