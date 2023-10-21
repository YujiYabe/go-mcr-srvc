package presenter

import (
	"context"
)

// Monitor ...
func (prsntr *Presenter) UpdateOrders(ctx context.Context) {
	// オーダー情報更新
	prsntr.ToMonitor.UpdateOrders(ctx)
}
