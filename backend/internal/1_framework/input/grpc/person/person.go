package person

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"backend/internal/2_adapter/controller"
	"backend/internal/4_domain/struct_object"
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

	listen, err := net.Listen("tcp", pkg.DeliveryAddress)
	if err != nil {
		pkg.Logging(ctx, err)
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()

	RegisterPersonServer(server, &receiver.Server)
	reflection.Register(server)

	if err := server.Serve(listen); err != nil {
		pkg.Logging(ctx, err)
		log.Fatalf("failed to serve: %v", err)
	}
}

// Implementation of the GetPersonByCondition method
func (receiver *Server) GetPersonByCondition(
	ctx context.Context,
	req *V1PersonParameter,
) (
	v1PersonParameterArray *V1PersonParameterArray,
	err error,
) {
	v1PersonParameterArray = &V1PersonParameterArray{}
	v1PersonParameterList := []*V1PersonParameter{}
	var id int
	if req.Id != nil {
		id = int(req.GetId())
	}

	name := req.Name
	mailAddress := req.MailAddress

	reqPerson := struct_object.NewPerson(
		&struct_object.NewPersonArgs{
			ID:          &id,
			Name:        name,
			MailAddress: mailAddress,
		},
	)
	if reqPerson.Err != nil {
		pkg.Logging(ctx, reqPerson.Err)
		return v1PersonParameterArray, reqPerson.Err
	}

	responseList, err := receiver.Controller.GetPersonByCondition(
		ctx,
		*reqPerson,
	)
	if err != nil {
		pkg.Logging(ctx, err)
		return v1PersonParameterArray, err
	}

	for _, response := range responseList {
		id32 := int32(response.ID.Content.Value)
		v1PersonParameter := &V1PersonParameter{
			Id:          &id32,
			Name:        &response.Name.Content.Value,
			MailAddress: &response.MailAddress.Content.Value,
		}
		v1PersonParameterList = append(v1PersonParameterList, v1PersonParameter)
	}

	v1PersonParameterArray.Persons = v1PersonParameterList
	return v1PersonParameterArray, nil
}
