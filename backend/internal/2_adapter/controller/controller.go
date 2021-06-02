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
// func NewController(toGrpcOut service.ToGrpcOut, toWsOrder service.ToWsOrder) *Controller {
func NewController(
	toRefrigerator service.ToRefrigerator,
	toFreezer service.ToFreezer,
	toShelf service.ToShelf,
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
func (ctrl *Controller) Reserve(ctx context.Context, orderType domain.OrderType) (string, string) {
	ctrl.OrderNumber++
	orderNumber := fmt.Sprintf("%03d", ctrl.OrderNumber)

	t := time.Now()
	orderTime := fmt.Sprint(t.Format("2006-01-02 15:04:05"))

	ctxValue := fmt.Sprintf("%s,%s", orderTime, orderType)

	return orderNumber, ctxValue
}

// Order ...
func (ctrl *Controller) Order(ctx context.Context, order domain.Order) error {
	err := ctrl.UseCase.Order(ctx, order)
	if err != nil {
		return err
	}

	return nil
}
