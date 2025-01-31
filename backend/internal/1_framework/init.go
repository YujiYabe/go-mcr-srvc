package app

import (
	//
	goEcho "backend/internal/1_framework/in/go-echo"
	goGRPC "backend/internal/1_framework/in/go-grpc"
	goPubSub "backend/internal/1_framework/in/go-pubsub"

	//
	postgresClient "backend/internal/1_framework/out/db/postgres_client"
	redisClient "backend/internal/1_framework/out/db/redis_client"
	auth0Client "backend/internal/1_framework/out/external/auth0_client"
	grpcClient "backend/internal/1_framework/out/external/grpc_client"
	pubsubPublisher "backend/internal/1_framework/out/external/pubsub_publisher"

	//
	"backend/internal/2_adapter/controller"
)

type (
	app struct {
		goEcho   *goEcho.GoEcho
		goGRPC   *goGRPC.GoGRPC
		goPubSub *goPubSub.GoPubSub
	}
)

// NewApp ...
func NewApp() *app {
	ctrl := controller.NewController(
		postgresClient.NewToPostgres(),
		redisClient.NewToRedis(),
		auth0Client.NewToAuth0(),
		grpcClient.NewToGRPC(),
		pubsubPublisher.NewToPubSub(),
	)
	ctrl.Start()

	a := &app{
		goGRPC:   goGRPC.NewGoGRPC(ctrl),
		goEcho:   goEcho.NewGoEcho(ctrl),
		goPubSub: goPubSub.NewGoPubSub(ctrl),
	}

	return a
}

// Start ...
func (receiver *app) Start() {
	go receiver.goEcho.Start()
	go receiver.goGRPC.Start()
	receiver.goPubSub.Start()
}
