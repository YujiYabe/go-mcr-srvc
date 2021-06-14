package gateways

import (
	"context"

	"backend/internal/4_enterprise_business_rules/entities"
)

// Shipment ...
func (sv *Service) Shipment(ctx context.Context, order *entities.Order) error {
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
