package app

import (
	"app/internal/1_infrastructure/in/delivery"
	"app/internal/1_infrastructure/in/mobile"
	"app/internal/1_infrastructure/in/pc"
	"app/internal/1_infrastructure/in/register"
	"app/internal/1_infrastructure/out/shipment"
	"app/internal/1_infrastructure/stock/freezer"
	"app/internal/1_infrastructure/stock/refrigerator"
	"app/internal/1_infrastructure/stock/shelf"
	"app/internal/2_adapter/controller"
)

type (
	app struct {
		mobile   *mobile.Mobile
		pc       *pc.PC
		delivery *delivery.Delivery
		register *register.Register
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

	return a
}

// Start ...
func (a *app) Start() {
	go a.mobile.Start()
	go a.pc.Start()
	go a.register.Start()
	a.delivery.Start()
}
