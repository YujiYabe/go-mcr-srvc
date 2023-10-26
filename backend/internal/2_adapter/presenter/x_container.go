package presenter

import (
	"context"

	domain "backend/internal/4_domain"
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
		DistributeOrder(
			ctx context.Context,
			orderList *domain.OrderList,
		)
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
