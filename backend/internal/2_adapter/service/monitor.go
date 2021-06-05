package service

import (
	"context"

	"backend/internal/4_domain/domain"
)

// Monitor ...
func (sv *Service) UpdateOrders(ctx context.Context, order *domain.Order) error {
	err := sv.ToMonitor.UpdateOrders(ctx, order)
	if err != nil {
		return err
	}

	// err = sv.ToMonitor.Logging(ctx, order)
	// if err != nil {
	// 	return err
	// }

	return nil
}
