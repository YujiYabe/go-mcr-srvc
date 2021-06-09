package service

import (
	"context"

	"backend/internal/4_domain/domain"
)

// Shipment ...
func (sv *Service) Shipment(ctx context.Context, order *domain.Order) error {
	err := sv.ToShipment.PutProducts(ctx, order)
	if err != nil {
		myErr.Logging(err)
		return err
	}

	err = sv.ToShipment.WriteLog(ctx, order)
	if err != nil {
		myErr.Logging(err)
		return err
	}

	return nil
}
