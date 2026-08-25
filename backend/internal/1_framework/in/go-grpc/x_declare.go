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
	"backend/internal/env"
	"backend/internal/logger"
)

// Server ...
type Server struct {
	Controller controller.ToController
	grpcParameter.UnimplementedPersonServiceServer
}

// NewGoGRPC ...
func NewGoGRPC(
	controller controller.ToController,
) *GoGRPC {
	goGRPC := &GoGRPC{
		Server: Server{
			Controller: controller,
		},
	}
	return goGRPC
}

// Start ....
func (receiver *GoGRPC) Start() error {
	logger.Logging(context.Background(), "start GRPC")
	listen, err := net.Listen(
		"tcp",
		env.ServerConfig.GRPCAddress,
	)
	if err != nil {
		return fmt.Errorf("listen grpc: %w", err)
	}
	server := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpcMiddleware.UnaryServerInterceptor(),
		),
	)

	grpcParameter.RegisterPersonServiceServer(server, &receiver.Server)
	reflection.Register(server)

	if err := server.Serve(listen); err != nil {
		return fmt.Errorf("serve grpc: %w", err)
	}

	return nil
}
