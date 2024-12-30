package presenter

import (
	"context"
)

// Presenter ...
type Presenter struct {
	ToMonitor ToMonitor
}

type (

	// ToMonitor ...
	ToMonitor interface {
		UpdateOrders(ctx context.Context, orderNumber string, phase string)
	}
)

// NewPresenter ...
func NewPresenter(toMonitor ToMonitor) *Presenter {
	return &Presenter{
		ToMonitor: toMonitor,
	}
}
