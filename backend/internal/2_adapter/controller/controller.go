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

	OrderChannel struct {
		ctx   *context.Context
		order *domain.Order
	}
)

// OrderChannel ...
var orderChannel = make(chan OrderChannel, 10)

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
}

// Reserve ...
func (ctrl *Controller) Reserve(ctx context.Context, order *domain.Order, orderType string) {
	ctrl.OrderNumber++

	order.OrderInfo.OrderNumber = fmt.Sprintf("%03d", ctrl.OrderNumber)
	order.OrderInfo.OrderType = orderType
	order.OrderInfo.OrderTime = time.Now()

	return
}

// Order ...
func (ctrl *Controller) Order(ctx context.Context, order *domain.Order) {
	fmt.Println("1==============================")
	fmt.Println("==============================")

	oc := &OrderChannel{
		ctx:   &ctx,
		order: order,
	}

	fmt.Println("2==============================")
	fmt.Println("==============================")
	orderChannel <- *oc
	fmt.Println("3==============================")
	fmt.Println("==============================")

	return
}

func (ctrl *Controller) bulkReception() {
	for {
		oc := <-orderChannel
		ctxWithTimeout, _ := context.WithTimeout(*oc.ctx, time.Minute*10)

		err := ctrl.UseCase.Order(ctxWithTimeout, oc.order)
		if err != nil {
			myErr.Logging(err)
		}
	}
}
