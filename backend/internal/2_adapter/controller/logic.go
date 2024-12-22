package controller

import (
	"context"
	"fmt"
	"time"

	domain "backend/internal/4_domain"
	"backend/pkg"
)

// orderChannel ...
var orderChannel = make(chan OrderChannel)

func (receiver *controller) Start() {
	go receiver.bulkOrder()
	go receiver.UseCase.Start()
}

// Reserve ...
func (receiver *controller) Reserve(ctx context.Context, order *domain.Order, orderType string) {
	receiver.OrderNumber++ // オーダー番号発行
	if receiver.OrderNumber >= 1000 {
		receiver.OrderNumber = 1 // オーダー番号を3桁以内にする
	}

	order.OrderInfo.OrderNumber = fmt.Sprintf("%03d", receiver.OrderNumber) // オーダー番号を3桁で表示
	order.OrderInfo.OrderType = orderType                                   // オーダーの種類(mobile/pc/delivery/register)
	order.OrderInfo.OrderTime = time.Now()                                  // オーダー受け付け時間

	receiver.UseCase.Reserve(ctx, &order.OrderInfo) // オーダー情報更新
}

// Order ...
func (receiver *controller) Order(ctx *context.Context, order *domain.Order) {
	oc := &OrderChannel{
		ctx:   ctx,
		order: order,
	}

	// オーダー番号をweb_uiに即時返却する必要があるため、
	// 後続処理をチャネル経由で処理
	orderChannel <- *oc
}

func (receiver *controller) bulkOrder() {
	// 無限ループでチャネルを待ち受け
	for {
		oc := <-orderChannel // Orderメソッドのチャネルの受け取り
		err := receiver.UseCase.Order(oc.ctx, oc.order)
		if err != nil {
			pkg.Logging(*oc.ctx, err)
		}
	}
}
