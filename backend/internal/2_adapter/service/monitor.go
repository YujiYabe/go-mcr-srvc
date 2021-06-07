package service

import (
	"context"
)

// Monitor ...
func (sv *Service) UpdateOrders(ctx context.Context, orderNumber string, phase string) {
	sv.ToMonitor.UpdateOrders(ctx, orderNumber, phase)
	return
}
