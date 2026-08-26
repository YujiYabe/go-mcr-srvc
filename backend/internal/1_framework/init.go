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
	app struct {
		goEcho   *goEcho.GoEcho
		goGRPC   *goGRPC.GoGRPC
		goPubSub *goPubSub.GoPubSub
	}
)

// NewApp ...
func NewApp() (*app, error) {
	ctx := context.Background()
	config, err := env.Load()
	if err != nil {
		return nil, err
	}

	toPostgres, err := postgresClient.NewToPostgres(ctx, config.Database.DSN)
	if err != nil {
		return nil, fmt.Errorf("new postgres client: %w", err)
	}

	toGRPC, err := grpcClient.NewToGRPC(ctx, config.Server.GRPCAddress)
	if err != nil {
		return nil, fmt.Errorf("new grpc client: %w", err)
	}

	toPubSub, err := pubsubPublisher.NewToPubSub(
		ctx,
		config.PubSub.BootstrapServers,
		config.PubSub.TestTopic,
		config.PubSub.FlushTimeoutMS,
		config.PubSub.SampleUserName,
	)
	if err != nil {
		return nil, fmt.Errorf("new pubsub publisher: %w", err)
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

	a := &app{
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

	return a, nil
}

// Start ...
func (receiver *app) Start() error {
	ctx := context.Background()
	if false {
		go func() {
			logger.Logging(ctx, receiver.goPubSub.Start())
		}()
	}
	go func() {
		if err := receiver.goGRPC.Start(); err != nil {
			logger.Logging(ctx, err)
		}
	}()
	return receiver.goEcho.Start()
}
