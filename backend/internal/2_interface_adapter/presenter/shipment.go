package presenter

import (
	"context"

	"backend/internal/4_enterprise_business_rule/entity"
)

// Shipment ...
func (pr *Presenter) Shipment(ctx context.Context, order *entity.Order) error {
	err := pr.ToShipment.PutProducts(ctx, order)
	if err != nil {
		myErr.Logging(err)
		return err
	}

	err = pr.ToShipment.WriteLog(ctx, order)
	if err != nil {
		myErr.Logging(err)
		return err
	}

	return nil
}
