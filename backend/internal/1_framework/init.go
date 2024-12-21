package app

import (
	"backend/internal/1_framework/db/freezer"
	"backend/internal/1_framework/db/refrigerator"
	"backend/internal/1_framework/db/shelf"
	"backend/internal/1_framework/external_interface/monitor"
	"backend/internal/1_framework/external_interface/shipment"
	"backend/internal/1_framework/input/delivery"
	"backend/internal/1_framework/input/mobile"
	"backend/internal/1_framework/input/pc"
	"backend/internal/1_framework/input/register"
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

	ctrl := controller.NewController(
		refrigerator.NewToRefrigerator(),
		freezer.NewToFreezer(),
		shelf.NewToShelf(),
		shipment.NewToShipment(),
		monitor.NewToMonitor(),
	)

	a := &app{
		delivery: delivery.NewDelivery(ctrl),
		mobile:   mobile.NewMobile(ctrl),
		pc:       pc.NewPC(ctrl),
		register: register.NewRegister(ctrl),
		monitor:  monitor.NewMonitor(),
	}
	ctrl.Start()

	return a
}

// Start ...
func (receiver *app) Start() {
	go receiver.monitor.Start()
	go receiver.mobile.Start()
	go receiver.pc.Start()
	go receiver.register.Start()
	receiver.delivery.Start()
}
