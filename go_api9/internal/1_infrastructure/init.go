package app

import (
	// 	"app/internal/1_infrastructure/grpc/grpcin"
	// 	"app/internal/1_infrastructure/grpc/grpcout"
	// 	"app/internal/1_infrastructure/network"
	// "app/internal/1_infrastructure/mobile"
	// 	"app/internal/1_infrastructure/ws/wsorder"

	"app/internal/1_infrastructure/mobile"
	"app/internal/1_infrastructure/stocker"
	"app/internal/2_adapter/controller"
	// 	"app/pkg/shared"
)

type (
	app struct {
		mobile *mobile.Mobile
	}
)

// NewApp ...
func NewApp() *app {
	a := &app{}

	// grpcOut := grpcout.NewToGrpcOut()
	// wsOrder := wsorder.NewToWsOrder()
	stocker := stocker.NewToStocker()
	// ctrl := controller.NewController(grpcOut, wsOrder)
	ctrl := controller.NewController(stocker)
	// a.GrpcIn = grpcin.NewGrpcIn(ctrl)
	a.mobile = mobile.NewMobile(ctrl)

	return a
}

// Start ...
func (a *app) Start() {
	a.mobile.Start()
}
