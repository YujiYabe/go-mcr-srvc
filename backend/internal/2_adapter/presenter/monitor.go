package presenter

import (
	"context"
)

// Monitor ...
func (receiver *Presenter) UpdateOrders(ctx context.Context, orderNumber string, phase string) {
	// オーダー情報更新
	receiver.ToMonitor.UpdateOrders(ctx, orderNumber, phase)

}
