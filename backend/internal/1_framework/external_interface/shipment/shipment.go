package shipment

import (
	"context"

	"backend/internal/2_adapter/presenter"
	"backend/pkg"
)

var (
	myErr *pkg.MyErr
)

func init() {
	myErr = pkg.NewMyErr("framework_driver", "shipment")
}

type Shipment struct{}

// NewToShipment ...
func NewToShipment() presenter.ToShipment {
	s := new(Shipment)
	return s
}

// PutProducts ...
func (shpmnt *Shipment) PutProducts(
	ctx context.Context,
) error {

	return nil
}

// WriteLog ...
func (shpmnt *Shipment) WriteLog(
	ctx context.Context,
) error {
	return nil
}
