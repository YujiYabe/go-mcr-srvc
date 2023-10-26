package presenter

import (
	"context"

	domain "backend/internal/4_domain"
)

func (receiver *Presenter) DistributeOrder(
	ctx context.Context,
	orderList *domain.OrderList,
) {
	receiver.ToMonitor.DistributeOrder(
		ctx,
		orderList,
	)
}
