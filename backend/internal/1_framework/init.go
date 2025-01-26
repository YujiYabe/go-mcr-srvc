package app

import (
	//
	goEcho "backend/internal/1_framework/in/go-echo"
	goGRPC "backend/internal/1_framework/in/go-grpc"

	//
	postgresClient "backend/internal/1_framework/out/db/postgres_client"
	redisClient "backend/internal/1_framework/out/db/redis_client"
	auth0Client "backend/internal/1_framework/out/external/auth0_client"
	grpcClient "backend/internal/1_framework/out/external/grpc_client"

	//
	"backend/internal/2_adapter/controller"
)

type (
	app struct {
		goEcho *goEcho.GoEcho
		goGRPC *goGRPC.GoGRPC
	}
)

// NewApp ...
func NewApp() *app {
	ctrl := controller.NewController(
		postgresClient.NewToPostgres(),
		redisClient.NewToRedis(),
		auth0Client.NewToAuth0(),
		grpcClient.NewToGRPC(),
	)
	ctrl.Start()

	a := &app{
		goGRPC: goGRPC.NewGoGRPC(ctrl),
		goEcho: goEcho.NewGoEcho(ctrl),
	}

	return a
}

// Start ...
func (receiver *app) Start() {
	go receiver.goEcho.Start()
	receiver.goGRPC.Start()
}
