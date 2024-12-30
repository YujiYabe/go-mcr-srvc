package app

import (
	"backend/internal/1_framework/db/postgres"
	"backend/internal/1_framework/db/redis"
	"backend/internal/1_framework/external_interface/monitor"
	person "backend/internal/1_framework/input/grpc/person"
	"backend/internal/1_framework/input/mobile"
	"backend/internal/1_framework/input/pc"
	"backend/internal/2_adapter/controller"
)

type (
	app struct {
		mobile   *mobile.Mobile
		pc       *pc.PC
		person   *person.Person
		monitor  *monitor.Monitor
	}
)

// NewApp ...
func NewApp() *app {

	ctrl := controller.NewController(
		redis.NewToRedis(),
		postgres.NewToPostgres(),
		monitor.NewToMonitor(),
	)

	a := &app{
		person:   person.NewPerson(ctrl),
		mobile:   mobile.NewMobile(ctrl),
		pc:       pc.NewPC(ctrl),
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
	receiver.person.Start()
}
