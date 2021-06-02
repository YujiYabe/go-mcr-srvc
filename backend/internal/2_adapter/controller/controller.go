package controller

import (
	"context"
	"fmt"
	"time"

	"app/internal/2_adapter/service"
	"app/internal/3_usecase/usecase"
	"app/internal/4_domain/domain"
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
) *Controller {
	ct := &Controller{
		UseCase: usecase.UseCase{
			ToDomain: domain.NewDomain(),
			ToService: &service.Service{
				// 	ToGrpcOut: toGrpcOut,
				// 	ToWsOrder: toWsOrder,
				ToRefrigerator: toRefrigerator,
				ToFreezer:      toFreezer,
				ToShelf:        toShelf,
				ToShipment:     toShipment,
			},
		},
	}

	return ct
}

// Dummy ...
func (ctrl *Controller) Dummy(ctx context.Context) (string, error) {
	res, _ := ctrl.UseCase.Dummy(ctx)
	return res, nil
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
