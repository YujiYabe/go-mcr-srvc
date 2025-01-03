package person

import (
	"context"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	// grpcUtil "backend/internal/1_framework/in/grpc/grpc_util"
	grpcUtil "backend/internal/1_framework/in/grpc/grpc_util"
	"backend/internal/2_adapter/controller"
	"backend/internal/4_domain/struct_object"
	"backend/pkg"
)

const (
	timeFormat = "06-01-02-15:04:05.000000"
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
	ctx, cancel := context.WithTimeout(
		context.Background(),
		30*time.Second,
	)
	defer cancel()
	log.Println("start GRPC ------------------------- ")

	listen, err := net.Listen("tcp", pkg.DeliveryAddress)
	if err != nil {
		pkg.Logging(ctx, err)
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpcUtil.RequestIDInterceptor), // UnaryInterceptor を設定
	)

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
	v1GetPersonByConditionResponse *V1GetPersonByConditionResponse,
	err error,
) {
	v1GetPersonByConditionResponse = &V1GetPersonByConditionResponse{}
	v1PersonParameterArray := &V1PersonParameterArray{}
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
		err = reqPerson.Err
		return
	}

	responseList, err := receiver.Controller.GetPersonByCondition(
		ctx,
		*reqPerson,
	)
	if err != nil {
		pkg.Logging(ctx, err)
		return
	}

	for _, response := range responseList {
		id32 := uint32(response.ID.Content.GetValue())
		name := response.Name.Content.GetValue()
		mailAddress := response.MailAddress.Content.GetValue()
		v1PersonParameter := &V1PersonParameter{
			Id:          &id32,
			Name:        &name,
			MailAddress: &mailAddress,
		}
		v1PersonParameterList = append(
			v1PersonParameterList,
			v1PersonParameter,
		)
	}

	v1PersonParameterArray.Persons = v1PersonParameterList
	v1GetPersonByConditionResponse.V1PersonParameterArray = v1PersonParameterArray
	v1GetPersonByConditionResponse.V1CommonParameter = &V1CommonParameter{
		XCorrelationID: pkg.GetCorrelationID(ctx),
		Timestamp:      time.Now().Format(timeFormat),
	}

	return
}
