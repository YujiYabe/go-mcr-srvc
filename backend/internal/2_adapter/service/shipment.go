package service

import (
	"context"

	"app/internal/4_domain/domain"
)

// Shipment ...
func (sv *Service) Shipment(ctx context.Context, order *domain.Order) error {
	err := sv.ToShipment.HandOver(ctx, order)
	if err != nil {
		return err
	}

	err = sv.ToShipment.Logging(ctx, order)
	if err != nil {
		return err
	}

	return nil
}
