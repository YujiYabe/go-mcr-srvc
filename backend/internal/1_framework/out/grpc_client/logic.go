package grpc_client

import (
	"context"

	grpcMiddleware "backend/internal/1_framework/middleware/grpc"
	grpcParameter "backend/internal/1_framework/parameter/grpc"
	groupObject "backend/internal/4_domain/group_object"
	valueObject "backend/internal/4_domain/value_object"
	"backend/pkg"
)

// ...
// ViaGRPC ...
func (receiver *GRPCClient) ViaGRPC(
	ctx context.Context,
	reqPerson groupObject.Person,
) (
	resPersonList groupObject.PersonList,
) {
	traceID := groupObject.GetRequestContext(ctx).TraceID.GetValue()
	pkg.Logging(ctx, traceID)

	var err error
	resPersonList = groupObject.PersonList{}

	// クライアントの作成
	client := grpcParameter.NewPersonServiceClient(receiver.Conn)

	// リクエストの作成
	v1GetPersonByConditionRequest := &grpcParameter.GetPersonListByConditionRequest{
		V1PersonParameter: &grpcParameter.V1PersonParameter{},
	}

	if !reqPerson.Name.GetIsNil() && reqPerson.Name.GetValue() != "" {
		value := reqPerson.Name.GetValue()
		v1GetPersonByConditionRequest.V1PersonParameter.Name = &value
	}

	if !reqPerson.MailAddress.GetIsNil() && reqPerson.MailAddress.GetValue() != "" {
		value := reqPerson.MailAddress.GetValue()
		v1GetPersonByConditionRequest.V1PersonParameter.MailAddress = &value
	}

	ctx = grpcMiddleware.ContextToMetadata(ctx)

	// gRPCリクエストの実行
	grpcPersonList, err := client.GetPersonListByCondition(
		ctx,
		v1GetPersonByConditionRequest,
	)
	if err != nil {
		resPersonList.SetError(ctx, err)
		return
	}
	for _, grpcPerson := range grpcPersonList.V1PersonParameterArray.Persons {
		person := &groupObject.Person{}

		id := int(grpcPerson.GetId())
		person.ID = valueObject.NewID(
			ctx,
			&id,
		)

		name := grpcPerson.GetName()
		person.Name = valueObject.NewName(
			ctx,
			&name,
		)

		mailAddress := grpcPerson.GetMailAddress()
		person.MailAddress = valueObject.NewMailAddress(
			ctx,
			&mailAddress,
		)

		resPersonList.Content = append(resPersonList.Content, *person)
	}

	traceID = groupObject.GetRequestContext(ctx).TraceID.GetValue()
	pkg.Logging(ctx, traceID)

	return

}
