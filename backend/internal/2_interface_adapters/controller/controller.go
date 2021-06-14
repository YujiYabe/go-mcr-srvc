package controller

import (
	"context"
	"fmt"
	"time"

	"backend/internal/2_interface_adapters/gateways"
	"backend/internal/3_application_business_rules/usecase"
	"backend/internal/4_enterprise_business_rules/entities"
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
		order *entities.Order
	}
)

// OrderChannel ...
var orderController = make(chan OrderChannel)

// NewController ...
func NewController(
	toRefrigerator gateways.ToRefrigerator,
	toFreezer gateways.ToFreezer,
	toShelf gateways.ToShelf,
	toShipment gateways.ToShipment,
	toMonitor gateways.ToMonitor,
) *Controller {
	ct := &Controller{
		UseCase: usecase.UseCase{
			ToDomain: entities.NewDomain(),
			ToService: &gateways.Service{
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
func (ctrl *Controller) Reserve(ctx context.Context, order *entities.Order, orderType string) {
	ctrl.OrderNumber++

	order.OrderInfo.OrderNumber = fmt.Sprintf("%03d", ctrl.OrderNumber)
	order.OrderInfo.OrderType = orderType
	order.OrderInfo.OrderTime = time.Now()

	ctrl.UseCase.Reserve(ctx, &order.OrderInfo)
	return
}

// Order ...
func (ctrl *Controller) Order(ctx *context.Context, order *entities.Order) {
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
