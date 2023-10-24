package controller

import (
	"context"

	domain "backend/internal/4_domain"
)

func (receiver *controller) Start() {
	ctx := context.Background()

	receiver.UseCase.Start(ctx)
}

// GetProduct ...
func (receiver *controller) GetProduct(
	ctx context.Context,
	productNumber int,
) *domain.Product {
	return receiver.UseCase.GetProduct(
		ctx,
		productNumber,
	)
}

// GetAllProductList ...
func (receiver *controller) GetAllProductList(
	ctx context.Context,
) *domain.AllProductList {
	return receiver.UseCase.GetAllProductList(
		ctx,
	)
}

// GetAllergyDefault ...
func (receiver *controller) GetAllergyDefault(
	ctx context.Context,
) domain.Allergy {
	return receiver.UseCase.GetAllergyDefault(
		ctx,
	)
}

// GetIsVaildLangCodeMap ...
func (receiver *controller) GetIsVaildLangCodeMap(
	ctx context.Context,
) map[int]string {
	return receiver.UseCase.GetIsVaildLangCodeMap(
		ctx,
	)
}

// UpdateProduct ...
func (receiver *controller) UpdateProduct(
	ctx context.Context,
	product domain.Product,
) {
	receiver.UseCase.UpdateProduct(
		ctx,
		product,
	)
}
