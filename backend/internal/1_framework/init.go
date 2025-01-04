package app

import (
	goEcho "backend/internal/1_framework/in/go-echo"
	grpcPerson "backend/internal/1_framework/in/grpc/person"
	auth0Client "backend/internal/1_framework/out/auth0_client"
	postgresClient "backend/internal/1_framework/out/db/postgres_client"
	redisClient "backend/internal/1_framework/out/db/redis_client"
	grpcClient "backend/internal/1_framework/out/grpc_client"
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
		redisClient.NewToRedis(),
		postgresClient.NewToPostgres(),
		auth0Client.NewToAuth0(),
		grpcClient.NewToGRPC(),
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
