package delivery

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"

	"google.golang.org/grpc/reflection"

	"app/internal/2_adapter/controller"
)

// Delivery ...
type Delivery struct {
	Server
}

// Server ...
type Server struct {
	UnimplementedDeliveryServiceServer
	Controller *controller.Controller
}

// NewDelivery ...
func NewDelivery(ctrl *controller.Controller) *Delivery {
	dlvr := &Delivery{}
	srvr := &Server{}
	srvr.Controller = ctrl
	dlvr.Server = *srvr

	return dlvr
}

// Start ...
func (dlvr *Delivery) Start() {
	log.Println("start GRPC ------------------------- ")
	lis, err := net.Listen("tcp", ":3456")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	RegisterDeliveryServiceServer(s, &dlvr.Server)
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// DeliveryRPC ...
func (s *Server) DeliveryRPC(ctx context.Context, in *DeliveryRequest) (*DeliveryResponse, error) {
	// param := &domain.Order{
	// 	// Room:   in.GetRoom(),
	// 	// Object: in.GetObject(),
	// 	// Key:    in.GetKey(),
	// 	// Value:  in.GetValue(),
	// }xdx

	param := proto.Clone(in)
	fmt.Println("==============================")
	debugTarget := param
	fmt.Printf("%#v\n", debugTarget)
	// fmt.Printf("%v\n", debugTarget)
	// fmt.Printf("%+v\n", debugTarget)
	// fmt.Printf("%T\n", debugTarget)
	fmt.Println("==============================")

	// s.Controller.Order(ctx, *param)

	return &DeliveryResponse{Message: "ok"}, nil
}
