package controller

import (
	"context"

	"app/internal/2_adapter/service"
	"app/internal/3_usecase/usecase"
	"app/internal/4_domain/domain"
)

func init() {}

type (
	// Controller ...
	Controller struct {
		UseCase usecase.UseCase
	}
)

// NewController ...
// func NewController(toGrpcOut service.ToGrpcOut, toWsOrder service.ToWsOrder) *Controller {
func NewController(toStocker service.ToStocker) *Controller {
	ct := &Controller{
		UseCase: usecase.UseCase{
			ToDomain: domain.NewDomain(),
			ToService: &service.Service{
				// 	ToGrpcOut: toGrpcOut,
				// 	ToWsOrder: toWsOrder,
				ToStocker: toStocker,
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

// Order ...
func (ctrl *Controller) Order(ctx context.Context) (string, error) {
	res, _ := ctrl.UseCase.Order(ctx)
	return res, nil
}
