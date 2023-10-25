package internal

import (
	"backend/internal/1_framework/db/sqlite"

	"backend/internal/1_framework/external/monitor"
	"backend/internal/1_framework/external/shipment"

	"backend/internal/1_framework/instacook"

	// "backend/internal/1_framework/web_ui/delivery"
	// "backend/internal/1_framework/web_ui/mobile"
	// "backend/internal/1_framework/web_ui/pc"
	// "backend/internal/1_framework/web_ui/register"

	"backend/internal/2_adapter/controller"
)

type (
	app struct {
		instaCook *instacook.InstaCook
		// mobile    *mobile.Mobile
		// pc        *pc.PC
		// delivery  *delivery.Delivery
		// register  *register.Register
		monitor *monitor.Monitor
	}
)

// NewApp ...
func NewApp() *app {

	ctrl := controller.NewController(
		sqlite.NewToSQLite(),
		shipment.NewToShipment(),
		monitor.NewToMonitor(),
	)

	a := &app{
		instaCook: instacook.NewInstaCook(ctrl),
		// delivery:  delivery.NewDelivery(ctrl),
		// mobile:    mobile.NewMobile(ctrl),
		// pc:        pc.NewPC(ctrl),
		// register:  register.NewRegister(ctrl),
		monitor: monitor.NewMonitor(),
	}

	ctrl.Start()

	return a
}

// Start ...
func (receiver *app) Start() {
	go receiver.instaCook.Start()
	receiver.monitor.Start()
	// go a.monitor.Start()
	// go a.mobile.Start()
	// go a.pc.Start()
	// go a.register.Start()
	// a.delivery.Start()
}
