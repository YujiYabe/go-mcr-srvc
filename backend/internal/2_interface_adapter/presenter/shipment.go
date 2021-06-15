package presenter

import (
	"context"

	"backend/internal/4_enterprise_business_rule/entity"
)

// Shipment ...
func (prsntr *Presenter) Shipment(ctx context.Context, order *entity.Order) error {
	err := prsntr.ToShipment.PutProducts(ctx, order)
	if err != nil {
		myErr.Logging(err)
		return err
	}

	err = prsntr.ToShipment.WriteLog(ctx, order)
	if err != nil {
		myErr.Logging(err)
		return err
	}

	return nil
}
