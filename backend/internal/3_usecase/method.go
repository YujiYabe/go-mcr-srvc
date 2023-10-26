package usecase

import (
	"context"

	domain "backend/internal/4_domain"
	"backend/pkg"
)

var (
	myErr *pkg.MyErr
)

func init() {
	myErr = pkg.NewMyErr("application_business_rule", "usecase")
}

// Start ...
func (receiver *useCase) Start(ctx context.Context) {
	receiver.SetUpInMemory(ctx)
}

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

// order -----------------------
// GetOrderList ...
func (receiver *useCase) GetOrderList(
	ctx context.Context,
) domain.OrderList {
	return receiver.ToDomain.GetOrderList(
		ctx,
	)
}

// GetReservingList ...
func (receiver *useCase) GetReservingList(
	ctx context.Context,
) domain.ReservingList {
	return receiver.ToDomain.GetReservingList(
		ctx,
	)
}

// GetSoldList ...
func (receiver *useCase) GetSoldList(
	ctx context.Context,
) domain.SoldList {
	return receiver.ToDomain.GetSoldList(
		ctx,
	)
}

// GetPreparingList ...
func (receiver *useCase) GetPreparingList(
	ctx context.Context,
) domain.SoldList {
	return receiver.ToDomain.GetPreparingList(
		ctx,
	)
}

// GetCompletedList ...
func (receiver *useCase) GetCompletedList(
	ctx context.Context,
) domain.SoldList {
	return receiver.ToDomain.GetCompletedList(
		ctx,
	)
}

// GetPassedList ...
func (receiver *useCase) GetPassedList(
	ctx context.Context,
) domain.SoldList {
	return receiver.ToDomain.GetPassedList(
		ctx,
	)
}

// UpdateSold ...
func (receiver *useCase) UpdateSoldStatus(
	ctx context.Context,
	newSold domain.Sold,
) {
	receiver.ToDomain.UpdateSoldStatus(
		ctx,
		newSold,
	)
}

// allergy -----------------------
// GetAllergyDefault ...
func (receiver *useCase) GetAllergyDefault(
	ctx context.Context,
) domain.Allergy {
	return receiver.ToDomain.GetAllergyDefault(
		ctx,
	)
}

// GetAllergyList ...
func (receiver *useCase) GetAllergyList(
	ctx context.Context,
) domain.AllergyList {
	return receiver.ToDomain.GetAllergyList(
		ctx,
	)
}

// language -----------------------
// GetIsVaildLangCodeMap ...
func (receiver *useCase) GetIsVaildLangCodeMap(
	ctx context.Context,
) map[int]string {
	return receiver.ToDomain.GetIsVaildLangCodeMap(
		ctx,
	)
}

// websocket -----------------------
// DistributeOrder ...
func (receiver *useCase) DistributeOrder(
	ctx context.Context,
) {
	orderList := receiver.ToDomain.GetOrderList(
		ctx,
	)

	receiver.ToPresenter.DistributeOrder(
		ctx,
		&orderList,
	)

}
