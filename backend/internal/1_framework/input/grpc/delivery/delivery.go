package delivery

import (
	"context"
	"log"
	"net"

	"github.com/jinzhu/copier"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"backend/internal/2_adapter/controller"
	domain "backend/internal/4_domain"
	"backend/pkg"
)

var (
	orderType = "delivery"
)

// Delivery ...
type Delivery struct {
	Server
}

// Server ...
type Server struct {
	UnimplementedDeliveryServiceServer
	Controller controller.ToController
}

// NewDelivery ...
func NewDelivery(ctrl controller.ToController) *Delivery {
	d := &Delivery{
		Server: Server{
			Controller: ctrl,
		},
	}

	return d
}

// Start ....
func (receiver *Delivery) Start() {
	ctx := context.Background()
	log.Println("start GRPC ------------------------- ")

	lis, err := net.Listen("tcp", pkg.DeliveryAddress)
	if err != nil {
		pkg.Logging(ctx, err)
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	RegisterDeliveryServiceServer(s, &receiver.Server)
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		pkg.Logging(ctx, err)
		log.Fatalf("failed to serve: %v", err)
	}
}

// DeliveryRPC ...
func (receiver *Server) DeliveryRPC(ctx context.Context, in *DeliveryRequest) (*DeliveryResponse, error) {

	// web_uiのデータ型をControllerに持ち込まないようにproductに変換
	product := &domain.Product{}
	err := copier.Copy(product, in.Order)
	if err != nil {
		pkg.Logging(ctx, err)
		return nil, err
	}
	order := &domain.Order{Product: *product}

	// receiver.Controller.Reserve(ctx, order, orderType)                      // オーダー番号発行
	// receiver.Controller.Order(&ctx, order)                                  // オーダー
	return &DeliveryResponse{OrderNumber: order.OrderInfo.OrderNumber}, nil // オーダー番号返却
}
