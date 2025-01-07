package grpc_client

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

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
	traceID := valueObject.GetTraceID(ctx)
	log.Println("== == == == == == == == == == ")
	pkg.Logging(ctx, traceID)

	var err error
	resPersonList = groupObject.PersonList{}
	// gRPCコネクションの作成
	conn, err := grpc.NewClient(
		"backend:3456",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		resPersonList.SetError(ctx, err)
		return
	}
	defer conn.Close()

	// クライアントの作成
	client := grpcParameter.NewPersonClient(conn)

	// リクエストの作成
	v1GetPersonByConditionRequest := &grpcParameter.V1GetPersonByConditionRequest{
		V1PersonParameter: &grpcParameter.V1PersonParameter{},
		V1CommonParameter: &grpcParameter.V1CommonParameter{
			Immutable: &grpcParameter.V1ImmutableParameter{
				TraceID: traceID,
			},
		},
	}

	if !reqPerson.Name.GetIsNil() && reqPerson.Name.GetValue() != "" {
		value := reqPerson.Name.GetValue()
		v1GetPersonByConditionRequest.V1PersonParameter.Name = &value
	}

	if !reqPerson.MailAddress.GetIsNil() && reqPerson.MailAddress.GetValue() != "" {
		value := reqPerson.MailAddress.GetValue()
		v1GetPersonByConditionRequest.V1PersonParameter.MailAddress = &value
	}

	ctx = metadata.AppendToOutgoingContext(
		ctx,
		string(valueObject.TraceIDMetaName),
		traceID,
	)

	// gRPCリクエストの実行
	grpcPersonList, err := client.GetPersonByCondition(
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

	log.Println("== == == == == == == == == == ")
	pkg.Logging(ctx, traceID)

	return

}
