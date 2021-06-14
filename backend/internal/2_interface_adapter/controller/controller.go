package controller

import (
	"context"
	"fmt"
	"time"

	"backend/internal/2_interface_adapter/gateway"
	"backend/internal/2_interface_adapter/presenter"
	"backend/internal/3_application_business_rule/usecase"
	"backend/internal/4_enterprise_business_rule/entity"
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
		order *entity.Order
	}

	// ToController ...
	ToController interface {
		Start()
		Reserve(ctx context.Context, order *entity.Order, orderType string)
		Order(ctx *context.Context, order *entity.Order)
	}
)

// OrderChannel ...
var orderController = make(chan OrderChannel)

// NewController ...
func NewController(
	toRefrigerator gateway.ToRefrigerator,
	toFreezer gateway.ToFreezer,
	toShelf gateway.ToShelf,
	toShipment presenter.ToShipment,
	toMonitor presenter.ToMonitor,
) ToController {
	ct := &Controller{
		UseCase: usecase.UseCase{
			ToDomain: entity.NewDomain(),
			ToGateway: &gateway.Gateway{
				ToRefrigerator: toRefrigerator,
				ToFreezer:      toFreezer,
				ToShelf:        toShelf,
			},
			ToPresenter: &presenter.Presenter{
				ToShipment: toShipment,
				ToMonitor:  toMonitor,
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
func (ctrl *Controller) Reserve(ctx context.Context, order *entity.Order, orderType string) {
	ctrl.OrderNumber++

	order.OrderInfo.OrderNumber = fmt.Sprintf("%03d", ctrl.OrderNumber)
	order.OrderInfo.OrderType = orderType
	order.OrderInfo.OrderTime = time.Now()

	ctrl.UseCase.Reserve(ctx, &order.OrderInfo)
	return
}

// Order ...
func (ctrl *Controller) Order(ctx *context.Context, order *entity.Order) {
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
