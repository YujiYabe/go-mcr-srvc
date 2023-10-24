package usecase

import (
	"context"

	domain "backend/internal/4_domain"
	"backend/pkg"
)

var (
	myErr        *pkg.MyErr
	orderUseCase = make(chan OrderUseCase)
)

func init() {
	myErr = pkg.NewMyErr("application_business_rule", "usecase")
}

// Start ...
func (receiver *useCase) Start(ctx *context.Context) {
	receiver.SetUpInMemory(ctx)
}

// Reserve ...
func (receiver *useCase) Reserve(ctx context.Context) {}

// Order ...
func (receiver *useCase) Order(ctx *context.Context) error {

	return nil
}

func (receiver *useCase) bulkOrder() {}

// GetProduct ...
func (receiver *useCase) GetProduct(
	ctx context.Context,
	productNumber int,
) *domain.Product {
	return receiver.ToDomain.GetProduct(
		ctx,
		productNumber,
	)

}

// SetUpInMemory ...
func (receiver *useCase) SetUpInMemory(
	ctx *context.Context,
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
