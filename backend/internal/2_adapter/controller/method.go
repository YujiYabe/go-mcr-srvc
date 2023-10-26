package controller

import (
	"context"

	domain "backend/internal/4_domain"
)

// Start ...
func (receiver *controller) Start() {
	ctx := context.Background()

	receiver.ToUseCase.Start(ctx)
}

// product -----------------------
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

// order -----------------------
// GetOrderList ...
func (receiver *controller) GetOrderList(
	ctx context.Context,
) domain.OrderList {
	return receiver.ToUseCase.GetOrderList(
		ctx,
	)
}

// GetReservingList ...
func (receiver *controller) GetReservingList(
	ctx context.Context,
) domain.ReservingList {
	return receiver.ToUseCase.GetReservingList(
		ctx,
	)
}

// GetSoldList ...
func (receiver *controller) GetSoldList(
	ctx context.Context,
) domain.SoldList {
	return receiver.ToUseCase.GetSoldList(
		ctx,
	)
}

// GetPreparingList ...
func (receiver *controller) GetPreparingList(
	ctx context.Context,
) domain.SoldList {
	return receiver.ToUseCase.GetPreparingList(
		ctx,
	)
}

// GetCompletedList ...
func (receiver *controller) GetCompletedList(
	ctx context.Context,
) domain.SoldList {
	return receiver.ToUseCase.GetCompletedList(
		ctx,
	)
}

// GetPassedList ...
func (receiver *controller) GetPassedList(
	ctx context.Context,
) domain.SoldList {
	return receiver.ToUseCase.GetPassedList(
		ctx,
	)
}

// UpdateSold ...
func (receiver *controller) UpdateSoldStatus(
	ctx context.Context,
	newSold domain.Sold,
) {
	receiver.ToUseCase.UpdateSoldStatus(
		ctx,
		newSold,
	)
}

// allergy -----------------------
// GetAllergyDefault ...
func (receiver *controller) GetAllergyDefault(
	ctx context.Context,
) domain.Allergy {
	return receiver.ToUseCase.GetAllergyDefault(
		ctx,
	)
}

// GetAllergyList ...
func (receiver *controller) GetAllergyList(
	ctx context.Context,
) domain.AllergyList {
	return receiver.ToUseCase.GetAllergyList(
		ctx,
	)
}

// language -----------------------
// GetIsVaildLangCodeMap ...
func (receiver *controller) GetIsVaildLangCodeMap(
	ctx context.Context,
) map[int]string {
	return receiver.ToUseCase.GetIsVaildLangCodeMap(
		ctx,
	)
}

// websocket -----------------------
// DistributeOrder ...
func (receiver *controller) DistributeOrder(
	ctx context.Context,
) {
	receiver.ToUseCase.DistributeOrder(
		ctx,
	)
}
