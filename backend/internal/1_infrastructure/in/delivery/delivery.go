package delivery

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"app/internal/2_adapter/controller"
	"app/internal/4_domain/domain"

	"github.com/jinzhu/copier"
	"google.golang.org/grpc/reflection"
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

// Start ....
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
	// fmt.Println("==============================")
	// debugTarget := in.String()
	// fmt.Printf("%#v\n", debugTarget)
	// fmt.Println("==============================")

	param := &domain.Order{
		Hamburgers: []domain.Hamburger{},
	}

	copier.Copy(param, in.Order)
	// copier.Copy(order, in.Order)
	// err := json.Unmarshal([]byte(in.String()), order)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return &DeliveryResponse{Message: "ng"}, err
	// }

	fmt.Printf("%#v\n", param)

	s.Controller.Order(ctx, *param)

	return &DeliveryResponse{Message: "ok"}, nil
}
