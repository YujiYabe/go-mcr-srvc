package presenter

import (
	"context"

	domain "backend/internal/4_domain"
)

// Shipment ...
func (prsntr *Presenter) Shipment(ctx context.Context, order *domain.Order) error {
	// 商品の出荷
	err := prsntr.ToShipment.PutProducts(ctx, order)
	if err != nil {
		myErr.Logging(err)
		return err
	}

	// 商品の出荷記録
	err = prsntr.ToShipment.WriteLog(ctx, order)
	if err != nil {
		myErr.Logging(err)
		return err
	}

	return nil
}
