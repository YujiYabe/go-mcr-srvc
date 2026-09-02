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
	gatewayDB "backend/internal/2_adapter/gateway/db"
	gatewayExternal "backend/internal/2_adapter/gateway/external"
	usecase "backend/internal/3_usecase"
	domain "backend/internal/4_domain"
	"backend/internal/env"
	"backend/internal/logger"
)

type (
	App struct {
		goEcho   *goEcho.GoEcho
		goGRPC   *goGRPC.GoGRPC
		goPubSub *goPubSub.GoPubSub
	}
)

// NewApp ...
func NewApp() (
	app *App,
	err error,
) {
	ctx := context.Background()
	config, err := env.Load()
	if err != nil {
		app = nil
		return //nolint:nakedret // Use the project-wide named return convention.
	}

	toPostgres, err := postgresClient.NewToPostgres(ctx, config.Database.DSN)
	if err != nil {
		app, err = nil, fmt.Errorf("new postgres client: %w", err)
		return //nolint:nakedret // Use the project-wide named return convention.
	}

	toGRPC, err := grpcClient.NewToGRPC(ctx, config.Server.GRPCAddress)
	if err != nil {
		app, err = nil, fmt.Errorf("new grpc client: %w", err)
		return //nolint:nakedret // Use the project-wide named return convention.
	}

	toPubSub, err := pubsubPublisher.NewToPubSub(
		ctx,
		config.PubSub.BootstrapServers,
		config.PubSub.TestTopic,
		config.PubSub.FlushTimeoutMS,
		config.PubSub.SampleUserName,
	)
	if err != nil {
		app, err = nil, fmt.Errorf("new pubsub publisher: %w", err)
		return //nolint:nakedret // Use the project-wide named return convention.
	}

	toGatewayDB := gatewayDB.NewGatewayDB(
		toPostgres,
		redisClient.NewToRedis(
			config.Redis.Addr,
			config.Redis.Password,
			config.Redis.DB,
		),
	)

	toGatewayExternal := gatewayExternal.NewGatewayExternal(
		auth0Client.NewToAuth0(
			config.Auth0.TokenURL,
			config.Auth0.Audience,
			config.Auth0.GrantType,
		),
		toGRPC,
		toPubSub,
	)

	useCase := usecase.NewUseCase(
		domain.NewDomain(),
		toGatewayDB,
		toGatewayExternal,
	)

	ctrl := controller.NewController(useCase)

	app = &App{
		goGRPC: goGRPC.NewGoGRPC(
			ctrl,
			config.Server.GRPCAddress,
		),
		goEcho: goEcho.NewGoEcho(
			ctrl,
			config.Server.GoEchoPort,
			config.Auth0.Domain,
			config.Auth0.ClientSecret,
		),
		goPubSub: goPubSub.NewGoPubSub(
			ctrl,
			config.PubSub.BootstrapServers,
			config.PubSub.ConsumerGroupID,
			config.PubSub.TestTopic,
			config.PubSub.OtherTopic,
		),
	}

	err = nil
	return //nolint:nakedret // Use the project-wide named return convention.
}

// Start ...
func (receiver *App) Start() (
	err error,
) {
	ctx := context.Background()
	if false {
		go func() {
			logger.Logging(ctx, receiver.goPubSub.Start(ctx))
		}()
	}
	go func() {
		if err := receiver.goGRPC.Start(ctx); err != nil {
			logger.Logging(ctx, err)
		}
	}()
	err = receiver.goEcho.Start()
	return
}
