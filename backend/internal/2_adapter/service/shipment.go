package service

import (
	"context"
)

// Shipment ...
func (sv *Service) Shipment(ctx context.Context) error {
	err := sv.ToShipment.HandOver(ctx)
	if err != nil {
		return err
	}

	err = sv.ToShipment.Logging(ctx)
	if err != nil {
		return err
	}

	return nil
}
