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
