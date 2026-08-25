package grpc_client

import (
	"context"

	grpcMiddleware "backend/internal/1_framework/middleware/grpc"
	grpcParameter "backend/internal/1_framework/parameter/grpc"
	groupObject "backend/internal/4_domain/group_object"
)

// ...
// ViaGRPC ...
func (receiver *GRPCClient) ViaGRPC(
	ctx context.Context,
	reqPerson groupObject.Person,
) (
	resPersonList groupObject.PersonList,
	err error,
) {
	// traceID := requestContextMiddleware.GetRequestContext(ctx).TraceID.GetValue()
	// logger.Logging(ctx, traceID)

	// クライアントの作成
	client := grpcParameter.NewPersonServiceClient(receiver.Conn)

	// リクエストの作成
	v1GetPersonByConditionRequest := &grpcParameter.GetPersonListByConditionRequest{
		V1PersonParameter: &grpcParameter.V1PersonParameter{},
	}

	if !reqPerson.Name().GetIsNil() && reqPerson.Name().GetValue() != "" {
		value := reqPerson.Name().GetValue()
		v1GetPersonByConditionRequest.V1PersonParameter.Name = &value
	}

	if !reqPerson.MailAddress().GetIsNil() && reqPerson.MailAddress().GetValue() != "" {
		value := reqPerson.MailAddress().GetValue()
		v1GetPersonByConditionRequest.V1PersonParameter.MailAddress = &value
	}

	ctx = grpcMiddleware.ContextToMetadata(ctx)

	// gRPCリクエストの実行
	grpcPersonList, err := client.GetPersonListByCondition(
		ctx,
		v1GetPersonByConditionRequest,
	)
	if err != nil {
		return
	}
	personArgs := make([]groupObject.NewPersonArgs, 0, len(grpcPersonList.V1PersonParameterArray.Persons))
	for _, grpcPerson := range grpcPersonList.V1PersonParameterArray.Persons {
		id := int(grpcPerson.GetId())
		name := grpcPerson.GetName()
		mailAddress := grpcPerson.GetMailAddress()
		personArgs = append(personArgs, groupObject.NewPersonArgs{
			ID:          &id,
			Name:        &name,
			MailAddress: &mailAddress,
		})
	}

	// traceID = requestContextMiddleware.GetRequestContext(ctx).TraceID.GetValue()
	// logger.Logging(ctx, traceID)

	return groupObject.NewPersonList(&groupObject.NewPersonListArgs{
		Content: personArgs,
	})

}
