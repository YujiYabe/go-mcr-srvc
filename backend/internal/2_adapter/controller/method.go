package controller

import (
	"context"

	domain "backend/internal/4_domain"
)

func (receiver *controller) Start() {
	ctx := context.Background()

	receiver.ToUseCase.Start(ctx)
}

// GetProduct ...
func (receiver *controller) GetProduct(
	ctx context.Context,
	productNumber int,
) *domain.Product {
	return receiver.ToUseCase.GetProduct(
		ctx,
		productNumber,
	)
}

// GetAllProductList ...
func (receiver *controller) GetAllProductList(
	ctx context.Context,
) *domain.AllProductList {
	return receiver.ToUseCase.GetAllProductList(
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

// GetAllergyDefault ...
func (receiver *controller) GetAllergyDefault(
	ctx context.Context,
) domain.Allergy {
	return receiver.ToUseCase.GetAllergyDefault(
		ctx,
	)
}

// GetIsVaildLangCodeMap ...
func (receiver *controller) GetIsVaildLangCodeMap(
	ctx context.Context,
) map[int]string {
	return receiver.ToUseCase.GetIsVaildLangCodeMap(
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
