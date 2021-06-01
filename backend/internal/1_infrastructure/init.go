package app

import (
	// 	"app/internal/1_infrastructure/grpc/grpcin"
	// 	"app/internal/1_infrastructure/grpc/grpcout"
	// 	"app/internal/1_infrastructure/network"
	// "app/internal/1_infrastructure/mobile"
	// 	"app/internal/1_infrastructure/ws/wsorder"

	"app/internal/1_infrastructure/in/mobile"
	"app/internal/1_infrastructure/in/pc"
	"app/internal/1_infrastructure/in/delivery"
	"app/internal/1_infrastructure/stock/freezer"
	"app/internal/1_infrastructure/stock/refrigerator"
	"app/internal/1_infrastructure/stock/shelf"
	"app/internal/2_adapter/controller"
)

type (
	app struct {
		mobile *mobile.Mobile
		pc     *pc.PC
		delivery     *delivery.Delivery
	}
)

// NewApp ...
func NewApp() *app {
	a := &app{}

	// grpcOut := grpcout.NewToGrpcOut()
	// wsOrder := wsorder.NewToWsOrder()
	refrigerator := refrigerator.NewToRefrigerator()
	freezer := freezer.NewToFreezer()
	shelf := shelf.NewToShelf()
	// ctrl := controller.NewController(grpcOut, wsOrder)
	ctrl := controller.NewController(refrigerator, freezer, shelf)
	a.GrpcIn = delivery.NewDelivery(ctrl)
	a.mobile = mobile.NewMobile(ctrl)
	a.pc = pc.NewPC(ctrl)

	return a
}

// Start ...
func (a *app) Start() {
	go a.mobile.Start()
	a.pc.Start()
}
