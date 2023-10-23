package controller

import (
	"context"

	domain "backend/internal/4_domain"
)

func (ctrl *controller) Start() {
	go ctrl.bulkOrder()
	go ctrl.UseCase.Start()
}

// Reserve ...
func (ctrl *controller) Reserve(
	ctx context.Context,
	order *domain.Order,
	orderType string,
) {
	ctrl.UseCase.Reserve(ctx) // オーダー情報更新
}

// Order ...
func (ctrl *controller) Order(
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

func (ctrl *controller) bulkOrder() {
}

// GetProduct ...
func (ctrl *controller) GetProduct(
	ctx context.Context,
	productNumber int,
) *domain.Product {
	return ctrl.UseCase.GetProduct(
		ctx,
		productNumber,
	)
}
