package service

import (
	"context"
)

// Monitor ...
func (sv *Service) UpdateOrders(ctx context.Context, orderNumber string, phase string) {
	sv.ToMonitor.ISUpdateOrders(ctx, orderNumber, phase)
	return
}
