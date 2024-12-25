package __

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"backend/internal/2_adapter/controller"
	"backend/pkg"
)

// Person ...
type Person struct {
	Server
}

// Server ...
type Server struct {
	UnimplementedPersonServer
	Controller controller.ToController
}

// NewPerson ...
func NewPerson(
	ctrl controller.ToController,
) *Person {
	d := &Person{
		Server: Server{
			Controller: ctrl,
		},
	}

	return d
}

// Start ....
func (receiver *Person) Start() {
	ctx := context.Background()
	log.Println("start GRPC ------------------------- ")

	lis, err := net.Listen("tcp", pkg.DeliveryAddress)
	if err != nil {
		pkg.Logging(ctx, err)
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	RegisterPersonServer(s, &receiver.Server)
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		pkg.Logging(ctx, err)
		log.Fatalf("failed to serve: %v", err)
	}
}

// Implementation of the GetPersonByCondition method
func (receiver *Server) GetPersonByCondition(
	ctx context.Context,
	req *V1PersonParameter,
) (
	*V1PersonParameter,
	error,
) {
	id := int32(req.GetId())
	name := req.GetName()
	mailAddress := req.GetMailAddress()

	resp := &V1PersonParameter{
		Id:          &id,
		Name:        &name,
		MailAddress: &mailAddress,
	}
	return resp, nil
}
