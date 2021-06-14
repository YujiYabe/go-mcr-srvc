package presenter

import (
	"context"
)

// Monitor ...
func (pr *Presenter) UpdateOrders(ctx context.Context, orderNumber string, phase string) {
	pr.ToMonitor.UpdateOrders(ctx, orderNumber, phase)
	return
}
