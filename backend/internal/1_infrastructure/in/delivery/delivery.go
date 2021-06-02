package delivery

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	"backend/internal/2_adapter/controller"
	"backend/internal/4_domain/domain"

	"github.com/jinzhu/copier"
	"google.golang.org/grpc/reflection"
)

var orderType = "delivery"

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

// Start ....
func (dlvr *Delivery) Start() {
	log.Println("start GRPC ------------------------- ")
	lis, err := net.Listen("tcp", "localhost:3456")
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
	order := &domain.Order{}
	product := &domain.Product{}

	copier.Copy(product, in.Order)

	order.Product = *product

	s.Controller.Reserve(ctx, order, orderType)

	go s.Controller.Order(ctx, order)

	return &DeliveryResponse{OrderNumber: order.OrderInfo.OrderNumber}, nil
}
