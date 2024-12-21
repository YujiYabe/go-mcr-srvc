package presenter

import (
	"context"

	domain "backend/internal/4_domain"
)

// Shipment ...
func (receiver *Presenter) Shipment(ctx context.Context, order *domain.Order) error {
	// 商品の出荷
	err := receiver.ToShipment.PutProducts(ctx, order)
	if err != nil {
		myErr.Logging(err)
		return err
	}

	// 商品の出荷記録
	err = receiver.ToShipment.WriteLog(ctx, order)
	if err != nil {
		myErr.Logging(err)
		return err
	}

	return nil
}
