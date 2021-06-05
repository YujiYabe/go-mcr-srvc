package service

import (
	"context"

	"backend/internal/4_domain/domain"
)

// Monitor ...
func (sv *Service) UpdateOrders(ctx context.Context, order *domain.Order, phase string) error {
	err := sv.ToMonitor.UpdateOrders(ctx, order, phase)
	if err != nil {
		return err
	}

	return nil
}
