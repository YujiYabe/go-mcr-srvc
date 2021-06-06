package controller

import (
	"context"
	"fmt"
	"time"

	"backend/internal/2_adapter/service"
	"backend/internal/3_usecase/usecase"
	"backend/internal/4_domain/domain"
)

func init() {}

type (
	// Controller ...
	Controller struct {
		UseCase     usecase.UseCase
		OrderNumber int
	}
)

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

// Reserve ...
func (ctrl *Controller) Reserve(ctx context.Context, order *domain.Order, orderType string) {
	ctrl.OrderNumber++

	order.OrderInfo.OrderNumber = fmt.Sprintf("%03d", ctrl.OrderNumber)
	order.OrderInfo.OrderType = orderType
	order.OrderInfo.OrderTime = time.Now()

	return
}

// Order ...
func (ctrl *Controller) Order(ctx context.Context, order *domain.Order) error {
	err := ctrl.UseCase.Order(ctx, order)
	if err != nil {
		return err
	}

	return nil
}
