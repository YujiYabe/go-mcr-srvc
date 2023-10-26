package presenter

import (
	"context"

	domain "backend/internal/4_domain"
)

// UpdateOrders ...
func (receiver *Presenter) UpdateOrders(ctx context.Context) {
	// オーダー情報更新
	receiver.ToMonitor.UpdateOrders(ctx)
}

func (receiver *Presenter) DistributeOrder(
	ctx context.Context,
	orderList *domain.OrderList,
) {
	receiver.ToMonitor.DistributeOrder(
		ctx,
		orderList,
	)
}
