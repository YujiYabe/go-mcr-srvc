package goGRPC

import (
	"context"
	"log"
	"time"

	grpcMiddleware "backend/internal/1_framework/middleware/grpc"
	grpcParameter "backend/internal/1_framework/parameter/grpc"
	groupObject "backend/internal/4_domain/group_object"
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
	request *grpcParameter.GetPersonByConditionRequest,
) (
	v1GetPersonByConditionResponse *grpcParameter.GetPersonByConditionResponse,
	err error,
) {
	v1GetPersonByConditionResponse = &grpcParameter.GetPersonByConditionResponse{}

	if request.GetV1CommonParameter() != nil {
		ctx = grpcMiddleware.CommonToContext(
			ctx,
			request.GetV1CommonParameter(),
		)
	}

	traceID := valueObject.GetTraceID(ctx)
	log.Println("== == == == == == == == == == ")
	pkg.Logging(ctx, traceID)

	v1PersonParameterArray := &grpcParameter.V1PersonParameterArray{}
	v1PersonParameterList := []*grpcParameter.V1PersonParameter{}

	var id *int
	if request.V1PersonParameter.GetId() != 0 {
		id = new(int)
		*id = int(request.V1PersonParameter.GetId())
	}

	var name *string
	if request.V1PersonParameter.Name != nil {
		name = request.V1PersonParameter.Name
	}

	var mailAddress *string
	if request.V1PersonParameter.MailAddress != nil {
		mailAddress = request.V1PersonParameter.MailAddress
	}

	reqPerson := groupObject.NewPerson(
		ctx,
		&groupObject.NewPersonArgs{
			ID:          id,
			Name:        name,
			MailAddress: mailAddress,
		},
	)
	if reqPerson.GetError() != nil {
		pkg.Logging(ctx, reqPerson.GetError())
		err = reqPerson.GetError()
		return
	}

	responseList := receiver.Controller.GetPersonByCondition(
		ctx,
		*reqPerson,
	)
	if responseList.GetError() != nil {
		pkg.Logging(ctx, responseList.GetError())
		return
	}

	for _, response := range responseList.Content {
		id32 := uint32(response.ID.GetValue())
		name := response.Name.GetValue()
		mailAddress := response.MailAddress.GetValue()
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
			TraceId: traceID,
		},
		Mutable: &grpcParameter.V1MutableParameter{
			TimeStamp: time.Now().Format(time.RFC3339),
		},
	}

	log.Println("== == == == == == == == == == ")
	pkg.Logging(ctx, traceID)

	return

}
