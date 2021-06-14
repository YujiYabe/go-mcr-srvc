package app

import (
	"backend/internal/1_frameworks_driver/in/delivery"
	"backend/internal/1_frameworks_driver/in/mobile"
	"backend/internal/1_frameworks_driver/in/pc"
	"backend/internal/1_frameworks_driver/in/register"
	"backend/internal/1_frameworks_driver/out/monitor"
	"backend/internal/1_frameworks_driver/out/shipment"
	"backend/internal/1_frameworks_driver/stock/freezer"
	"backend/internal/1_frameworks_driver/stock/refrigerator"
	"backend/internal/1_frameworks_driver/stock/shelf"
	"backend/internal/2_interface_adapter/controller"
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
func (a *app) Start() {
	go a.monitor.Start()
	go a.mobile.Start()
	go a.pc.Start()
	go a.register.Start()
	a.delivery.Start()
}
