package presenter

import (
	"context"

	"backend/pkg"
	domain "backend/internal/4_domain"
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
		PutProducts(ctx context.Context, order *domain.Order) error
		WriteLog(ctx context.Context, order *domain.Order) error
	}

	// ToMonitor ...
	ToMonitor interface {
		UpdateOrders(ctx context.Context, orderNumber string, phase string)
	}
)

// NewPresenter ...
func NewPresenter(toShipment ToShipment, toMonitor ToMonitor) *Presenter {
	return &Presenter{
		ToShipment: toShipment,
		ToMonitor:  toMonitor,
	}
}
