package monitor

import (
	"context"

	wschannel "backend/internal/1_framework/instacook/http/v1/ws/channel"
	domain "backend/internal/4_domain"
)

// DistributeOrder ...
func (mntr *Monitor) DistributeOrder(
	ctx context.Context,
	orderList *domain.OrderList,
) {
	wschannel.Cnnl <- orderList
}
