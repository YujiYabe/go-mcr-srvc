package app

import (
	"backend/internal/1_infrastructure/in/delivery"
	"backend/internal/1_infrastructure/in/mobile"
	"backend/internal/1_infrastructure/in/pc"
	"backend/internal/1_infrastructure/in/register"
	"backend/internal/1_infrastructure/out/monitor"
	"backend/internal/1_infrastructure/out/shipment"
	"backend/internal/1_infrastructure/stock/freezer"
	"backend/internal/1_infrastructure/stock/refrigerator"
	"backend/internal/1_infrastructure/stock/shelf"
	"backend/internal/2_adapter/controller"
)

type (
	app struct {
		mobile   *mobile.Mobile
		pc       *pc.PC
		delivery *delivery.Delivery
		register *register.Register
		monitor  *monitor.Monitor
	}
)

// NewApp ...
func NewApp() *app {
	a := &app{}

	refrigerator := refrigerator.NewToRefrigerator()
	freezer := freezer.NewToFreezer()
	shelf := shelf.NewToShelf()
	shipment := shipment.NewToShipment()

	ctrl := controller.NewController(refrigerator, freezer, shelf, shipment)

	a.delivery = delivery.NewDelivery(ctrl)
	a.mobile = mobile.NewMobile(ctrl)
	a.pc = pc.NewPC(ctrl)
	a.register = register.NewRegister(ctrl)

	a.monitor = monitor.NewMonitor()

	return a
}

// Start ...
func (a *app) Start() {
	go a.monitor.Start()
	go a.mobile.Start()
	go a.pc.Start()
	go a.register.Start()
	a.delivery.Start()
}
