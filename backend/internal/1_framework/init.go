package app

import (
	"context"
	"fmt"

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
	"backend/internal/env"
	"backend/internal/logger"
)

type (
	app struct {
		goEcho   *goEcho.GoEcho
		goGRPC   *goGRPC.GoGRPC
		goPubSub *goPubSub.GoPubSub
	}
)

// NewApp ...
func NewApp() (*app, error) {
	ctx := context.Background()
	if err := env.Err(); err != nil {
		return nil, err
	}

	toPostgres, err := postgresClient.NewToPostgres(ctx)
	if err != nil {
		return nil, fmt.Errorf("new postgres client: %w", err)
	}

	toGRPC, err := grpcClient.NewToGRPC(ctx)
	if err != nil {
		return nil, fmt.Errorf("new grpc client: %w", err)
	}

	toPubSub, err := pubsubPublisher.NewToPubSub(ctx)
	if err != nil {
		return nil, fmt.Errorf("new pubsub publisher: %w", err)
	}

	ctrl := controller.NewController(
		toPostgres,
		redisClient.NewToRedis(),
		auth0Client.NewToAuth0(),
		toGRPC,
		toPubSub,
	)
	ctrl.Start()

	a := &app{
		goGRPC:   goGRPC.NewGoGRPC(ctrl),
		goEcho:   goEcho.NewGoEcho(ctrl),
		goPubSub: goPubSub.NewGoPubSub(ctrl),
	}

	return a, nil
}

// Start ...
func (receiver *app) Start() error {
	ctx := context.Background()
	if false {
		go func() {
			if err := receiver.goPubSub.Start(); err != nil {
				logger.Logging(ctx, err)
			}
		}()
	}
	go func() {
		if err := receiver.goGRPC.Start(); err != nil {
			logger.Logging(ctx, err)
		}
	}()
	return receiver.goEcho.Start()
}
