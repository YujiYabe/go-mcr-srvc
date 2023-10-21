package presenter

import (
	"context"
)

// Shipment ...
func (prsntr *Presenter) Shipment(ctx context.Context) error {
	// 商品の出荷
	err := prsntr.ToShipment.PutProducts(ctx)
	if err != nil {
		myErr.Logging(err)
		return err
	}

	// 商品の出荷記録
	err = prsntr.ToShipment.WriteLog(ctx)
	if err != nil {
		myErr.Logging(err)
		return err
	}

	return nil
}
