package controller

import (
	"context"

	domain "backend/internal/4_domain"
)

func (receiver *controller) Start() {
	ctx := context.Background()

	receiver.UseCase.Start(&ctx)
}

// Reserve ...
func (receiver *controller) Reserve(
	ctx context.Context,
	order *domain.Order,
	orderType string,
) {
	receiver.UseCase.Reserve(ctx) // オーダー情報更新
}

// Order ...
func (receiver *controller) Order(
	ctx *context.Context,
	order *domain.Order,
) {
	oc := &orderChannel{
		order: order,
	}

	// オーダー番号をweb_uiに即時返却する必要があるため、
	// 後続処理をチャネル経由で処理
	odrChnnl <- *oc
}

func (receiver *controller) bulkOrder() {
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
	ctx *context.Context,
) *domain.Allergy {
	return receiver.UseCase.GetAllergyDefault(
		ctx,
	)
}

// GetProduct ...
func (receiver *controller) GetIsVaildLangCodeMap(
	ctx *context.Context,
) map[int]string {
	return receiver.UseCase.GetIsVaildLangCodeMap(
		ctx,
	)
}
