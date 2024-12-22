package presenter

import (
	"context"

	domain "backend/internal/4_domain"
)

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
