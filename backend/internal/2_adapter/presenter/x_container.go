package presenter

import (
	"context"

	domain "backend/internal/4_domain"
)

// Presenter ...
type Presenter struct {
	ToMonitor ToMonitor
}

type (

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
	toMonitor ToMonitor,
) *Presenter {
	return &Presenter{
		ToMonitor: toMonitor,
	}
}
