package goGRPC

import (
	"context"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	grpcParameter "backend/internal/1_framework/parameter/grpc"
	"backend/internal/2_adapter/controller"
	"backend/pkg"
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
	ctx, cancel := context.WithTimeout(
		context.Background(),
		30*time.Second,
	)
	defer cancel()
	log.Println("------------------------- start GRPC ------------------------- ")

	listen, err := net.Listen("tcp", pkg.GRPCAddress)
	if err != nil {
		pkg.Logging(ctx, err)
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer(
	// grpc.UnaryInterceptor(
	// 	grpcMiddleware.MetadataToContext,
	// ),
	)

	grpcParameter.RegisterPersonServiceServer(server, &receiver.Server)
	reflection.Register(server)

	if err := server.Serve(listen); err != nil {
		pkg.Logging(ctx, err)
		log.Fatalf("failed to serve: %v", err)
	}
}
