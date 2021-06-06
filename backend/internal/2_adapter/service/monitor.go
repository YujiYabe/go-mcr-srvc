package service

import (
	"context"
)

// Monitor ...
func (sv *Service) UpdateOrders(ctx context.Context, orderNumber string, phase string) error {
	err := sv.ToMonitor.UpdateOrders(ctx, orderNumber, phase)
	if err != nil {
		return err
	}

	return nil
}
