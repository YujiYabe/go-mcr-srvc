package delivery

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	"app/internal/2_adapter/controller"
	"app/internal/4_domain/domain"

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
	lis, err := net.Listen("tcp", "3456")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	RegisterDeliveryServiceServer(s, &dlvr.Server)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// SendContentRPC ...
func (s *Server) SendContentRPC(ctx context.Context, in *DeliveryRequest) (*DeliveryResponse, error) {
	param := &domain.Order{
		// Room:   in.GetRoom(),
		// Object: in.GetObject(),
		// Key:    in.GetKey(),
		// Value:  in.GetValue(),
	}

	s.Controller.Order(ctx, *param)

	return &DeliveryResponse{Message: "ok"}, nil
}
