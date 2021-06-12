package controller

import (
	"context"
	"fmt"
	"time"

	"backend/internal/2_adapter/service"
	"backend/internal/3_usecase/usecase"
	"backend/internal/4_domain/domain"
	"backend/pkg"
)

var (
	myErr *pkg.MyErr
)

func init() {
	myErr = pkg.NewMyErr("adapter", "controller")
}

type (
	// Controller ...
	Controller struct {
		UseCase     usecase.UseCase
		OrderNumber int
	}

	// OrderChannel ...
	OrderChannel struct {
		ctx   *context.Context
		order *domain.Order
	}
)

// OrderChannel ...
var orderController = make(chan OrderChannel)

// NewController ...
func NewController(
	toRefrigerator service.ToRefrigerator,
	toFreezer service.ToFreezer,
	toShelf service.ToShelf,
	toShipment service.ToShipment,
	toMonitor service.ToMonitor,
) *Controller {
	ct := &Controller{
		UseCase: usecase.UseCase{
			ToDomain: domain.NewDomain(),
			ToService: &service.Service{
				ToRefrigerator: toRefrigerator,
				ToFreezer:      toFreezer,
				ToShelf:        toShelf,
				ToShipment:     toShipment,
				ToMonitor:      toMonitor,
			},
		},
	}

	return ct
}

func (ctrl *Controller) Start() {
	go ctrl.bulkReception()
	go ctrl.UseCase.Start()
}

// Reserve ...
func (ctrl *Controller) Reserve(ctx context.Context, order *domain.Order, orderType string) {
	ctrl.OrderNumber++

	order.OrderInfo.OrderNumber = fmt.Sprintf("%03d", ctrl.OrderNumber)
	order.OrderInfo.OrderType = orderType
	order.OrderInfo.OrderTime = time.Now()

	ctrl.UseCase.Reserve(ctx, &order.OrderInfo)
	return
}

// Order ...
func (ctrl *Controller) Order(ctx *context.Context, order *domain.Order) {
	oc := &OrderChannel{
		ctx:   ctx,
		order: order,
	}
	orderController <- *oc

	return
}

func (ctrl *Controller) bulkReception() {
	for {
		oc := <-orderController
		ctxWithTimeout, _ := context.WithTimeout(*oc.ctx, time.Minute*10)

		err := ctrl.UseCase.Order(&ctxWithTimeout, oc.order)
		if err != nil {
			myErr.Logging(err)
		}
	}
}
