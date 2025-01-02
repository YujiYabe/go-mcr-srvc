package app

import (
	goEcho "backend/internal/1_framework/in/go-echo"
	grpcPerson "backend/internal/1_framework/in/grpc/person"
	"backend/internal/1_framework/out/auth0"
	"backend/internal/1_framework/out/db/postgres"
	"backend/internal/1_framework/out/db/redis"
	"backend/internal/2_adapter/controller"
)

type (
	app struct {
		goEcho *goEcho.GoEcho
		person *grpcPerson.Person
	}
)

// NewApp ...
func NewApp() *app {

	ctrl := controller.NewController(
		redis.NewToRedis(),
		postgres.NewToPostgres(),
		auth0.NewToAuth0(),
	)

	ctrl.Start()

	a := &app{
		person: grpcPerson.NewPerson(ctrl),
		goEcho: goEcho.NewGoEcho(ctrl),
	}

	return a
}

// Start ...
func (receiver *app) Start() {
	go receiver.goEcho.Start()
	receiver.person.Start()
}
