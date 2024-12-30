package app

import (
	"backend/internal/1_framework/db/postgres"
	"backend/internal/1_framework/db/redis"
	mobile "backend/internal/1_framework/input/go-echo"
	grpcPerson "backend/internal/1_framework/input/grpc/person"
	"backend/internal/2_adapter/controller"
)

type (
	app struct {
		mobile *mobile.Mobile
		person *grpcPerson.Person
	}
)

// NewApp ...
func NewApp() *app {

	ctrl := controller.NewController(
		redis.NewToRedis(),
		postgres.NewToPostgres(),
	)

	a := &app{
		person: grpcPerson.NewPerson(ctrl),
		mobile: mobile.NewMobile(ctrl),
	}
	ctrl.Start()

	return a
}

// Start ...
func (receiver *app) Start() {
	go receiver.mobile.Start()
	receiver.person.Start()
}
