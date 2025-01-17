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

const (
// timeFormat = "06-01-02-15:04:05.000000000"
)

// Server ...
type Server struct {
	grpcParameter.UnimplementedPersonServiceServer
	Controller controller.ToController
}

// NewGoGRPC ...
func NewGoGRPC(
	ctrl controller.ToController,
) *GoGRPC {
	d := &GoGRPC{
		Server: Server{
			Controller: ctrl,
		},
	}
	return d
}

// Start ....
func (receiver *GoGRPC) Start() {
	log.Println("------------------------- start GRPC ------------------------- ")

	listen, err := net.Listen("tcp", env.GRPCAddress)
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
