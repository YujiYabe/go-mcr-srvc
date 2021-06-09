package service

import (
	"context"

	"backend/internal/4_domain/domain"
)

// Shipment ...
func (sv *Service) Shipment(ctx context.Context, order *domain.Order) error {
	err := sv.ToShipment.ISHandOver(ctx, order)
	if err != nil {
		myErr.Logging(err)
		return err
	}

	err = sv.ToShipment.ISLogging(ctx, order)
	if err != nil {
		myErr.Logging(err)
		return err
	}

	return nil
}
