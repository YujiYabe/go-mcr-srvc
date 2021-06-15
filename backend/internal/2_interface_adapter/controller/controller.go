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
	myErr = pkg.NewMyErr("interface_adapter", "controller")
}

type (
	// controller ...
	controller struct {
		UseCase     usecase.ToUseCase
		OrderNumber int
	}

	// orderChannel ...
	orderChannel struct {
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

// orderChannel ...
var odrChnnl = make(chan orderChannel)

// NewController ...
func NewController(
	toRefrigerator gateway.ToRefrigerator,
	toFreezer gateway.ToFreezer,
	toShelf gateway.ToShelf,
	toShipment presenter.ToShipment,
	toMonitor presenter.ToMonitor,
) ToController {
	toEntity := entity.NewEntity()
	toGateway := gateway.NewGateway(
		toRefrigerator,
		toFreezer,
		toShelf,
	)
	toPresenter := presenter.NewPresenter(
		toShipment,
		toMonitor,
	)
	uscs := usecase.NewUseCase(
		toEntity,
		toGateway,
		toPresenter,
	)
	ct := &controller{
		UseCase: uscs,
	}

	return ct
}

func (ctrl *controller) Start() {
	go ctrl.bulkReception()
	go ctrl.UseCase.Start()
}

// Reserve ...
func (ctrl *controller) Reserve(ctx context.Context, order *entity.Order, orderType string) {
	ctrl.OrderNumber++

	order.OrderInfo.OrderNumber = fmt.Sprintf("%03d", ctrl.OrderNumber)
	order.OrderInfo.OrderType = orderType
	order.OrderInfo.OrderTime = time.Now()

	ctrl.UseCase.Reserve(ctx, &order.OrderInfo)
}

// Order ...
func (ctrl *controller) Order(ctx *context.Context, order *entity.Order) {
	oc := &orderChannel{
		ctx:   ctx,
		order: order,
	}

	// the reason of use channel
	// for return order number immediatamente
	// and continue process to assemble order.
	odrChnnl <- *oc
}

func (ctrl *controller) bulkReception() {
	for {
		oc := <-odrChnnl
		ctxWithTimeout, _ := context.WithTimeout(*oc.ctx, time.Minute*10)

		err := ctrl.UseCase.Order(&ctxWithTimeout, oc.order)
		if err != nil {
			myErr.Logging(err)
		}
	}

}
