package presenter

import (
	"context"

	"backend/pkg"
)

var (
	myErr *pkg.MyErr
)

func init() {
	myErr = pkg.NewMyErr("interface_adapter", "gateway")
}

// Presenter ...
type Presenter struct {
	ToShipment ToShipment
	ToMonitor  ToMonitor
}

type (
	// ToShipment ...
	ToShipment interface {
		PutProducts(ctx context.Context) error
		WriteLog(ctx context.Context) error
	}

	// ToMonitor ...
	ToMonitor interface {
		UpdateOrders(ctx context.Context)
	}
)

// NewPresenter ...
func NewPresenter(
	toShipment ToShipment,
	toMonitor ToMonitor,
) *Presenter {
	return &Presenter{
		ToShipment: toShipment,
		ToMonitor:  toMonitor,
	}
}
