package internal

import (
	"backend/internal/1_framework/db/mongo"
	"backend/internal/1_framework/db/mysql"
	"backend/internal/1_framework/db/postgres"
	"backend/internal/1_framework/external_interface/monitor"
	"backend/internal/1_framework/external_interface/shipment"
	"backend/internal/1_framework/web_ui/delivery"
	"backend/internal/1_framework/web_ui/mobile"
	"backend/internal/1_framework/web_ui/pc"
	"backend/internal/1_framework/web_ui/register"
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
		postgres.NewToPostgres(),
		mysql.NewToMysql(),
		mongo.NewToMongo(),
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
