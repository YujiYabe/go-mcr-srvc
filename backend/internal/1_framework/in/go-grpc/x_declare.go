package goGRPC

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	grpcMiddleware "backend/internal/1_framework/middleware/grpc"
	grpcParameter "backend/internal/1_framework/parameter/grpc"
	"backend/internal/2_adapter/controller"
	"backend/internal/env"
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
func (receiver *GoGRPC) Start() {
	log.Println("------------------------- start GRPC ------------------------- ")

	listen, err := net.Listen(
		"tcp",
		env.ServerConfig.GRPCAddress,
	)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpcMiddleware.UnaryServerInterceptor(),
		),
	)

	grpcParameter.RegisterPersonServiceServer(server, &receiver.Server)
	reflection.Register(server)

	if err := server.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
