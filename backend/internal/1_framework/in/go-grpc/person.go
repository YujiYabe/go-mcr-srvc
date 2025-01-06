package goGRPC

import (
	"context"
	"log"
	"time"

	grpcParameter "backend/internal/1_framework/parameter/grpc"
	structObject "backend/internal/4_domain/struct_object"
	valueObject "backend/internal/4_domain/value_object"

	"backend/pkg"
)

// GoGRPC ...
type GoGRPC struct {
	Server
}

// Implementation of the GetPersonByCondition method
func (receiver *Server) GetPersonByCondition(
	ctx context.Context,
	req *grpcParameter.V1GetPersonByConditionRequest,
) (
	v1GetPersonByConditionResponse *grpcParameter.V1GetPersonByConditionResponse,
	err error,
) {
	traceID := valueObject.GetTraceID(ctx)

	log.Println("== == == == == == == == == == ")
	pkg.Logging(ctx, traceID)
	log.Println("== == == == == == == == == == ")

	v1GetPersonByConditionResponse = &grpcParameter.V1GetPersonByConditionResponse{}
	v1PersonParameterArray := &grpcParameter.V1PersonParameterArray{}
	v1PersonParameterList := []*grpcParameter.V1PersonParameter{}

	var id *int
	if req.V1PersonParameter.GetId() != 0 {
		id = new(int)
		*id = int(req.V1PersonParameter.GetId())
	}

	var name *string
	if req.V1PersonParameter.Name != nil {
		name = req.V1PersonParameter.Name
	}

	var mailAddress *string
	if req.V1PersonParameter.MailAddress != nil {
		mailAddress = req.V1PersonParameter.MailAddress
	}

	reqPerson := structObject.NewPerson(
		&structObject.NewPersonArgs{
			ID:          id,
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
		v1PersonParameter := &grpcParameter.V1PersonParameter{
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
	v1GetPersonByConditionResponse.V1CommonParameter = &grpcParameter.V1CommonParameter{
		Immutable: &grpcParameter.V1ImmutableParameter{
			TraceID: traceID,
		},
		Mutable: &grpcParameter.V1MutableParameter{
			TimeStamp: time.Now().Format(timeFormat),
		},
	}

	return
}
