package goGRPC

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	grpcMiddleware "backend/internal/1_framework/middleware/grpc"
	grpcParameter "backend/internal/1_framework/parameter/grpc"
	"backend/internal/2_adapter/controller"
	"backend/internal/logger"
)

// Server ...
type Server struct {
	Controller controller.ToController
	grpcParameter.UnimplementedUserServiceServer
}

// NewGoGRPC ...
func NewGoGRPC(
	controller controller.ToController,
	address string,
) (
	goGRPC *GoGRPC,
) {
	goGRPC = &GoGRPC{
		Server: Server{
			Controller: controller,
		},
		address: address,
	}
	return
}

// Start ....
func (receiver *GoGRPC) Start(
	ctx context.Context,
) (
	err error,
) {
	logger.Logging(ctx, "start GRPC")
	listenConfig := &net.ListenConfig{}
	listen, err := listenConfig.Listen(ctx, "tcp", receiver.address)
	if err != nil {
		err = fmt.Errorf("listen grpc: %w", err)
		return
	}
	server := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpcMiddleware.UnaryServerInterceptor(),
		),
	)

	grpcParameter.RegisterUserServiceServer(server, &receiver.Server)
	reflection.Register(server)

	if err := server.Serve(listen); err != nil {
		return fmt.Errorf("serve grpc: %w", err)
	}

	err = nil
	return
}
