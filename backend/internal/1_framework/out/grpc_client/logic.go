package grpc_client

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	grpcMiddleware "backend/internal/1_framework/middleware/grpc"
	grpcParameter "backend/internal/1_framework/parameter/grpc"
	valueObject "backend/internal/4_domain/value_object"

	"backend/pkg"
)

// ...
// ViaGRPC ...
func (receiver *GRPCClient) ViaGRPC(
	ctx context.Context,
) (
	err error,
) {
	// gRPCコネクションの作成
	conn, err := grpc.NewClient(
		"backend:3456",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// クライアントの作成
	client := grpcParameter.NewPersonClient(conn)

	name := "a"
	// リクエストの作成
	v1GetPersonByConditionRequest := &grpcParameter.V1GetPersonByConditionRequest{
		V1PersonParameter: &grpcParameter.V1PersonParameter{
			Name: &name,
		},
		V1CommonParameter: &grpcParameter.V1CommonParameter{
			Immutable: &grpcParameter.V1ImmutableParameter{
				TraceID: grpcMiddleware.GetTraceID(ctx),
			},
		},
	}

	ctx = metadata.AppendToOutgoingContext(
		ctx,
		string(valueObject.TraceIDMetaName),
		grpcMiddleware.GetTraceID(ctx),
	)
	log.Println("== == == == == == == == == == ")
	pkg.Logging(ctx, grpcMiddleware.GetTraceID(ctx))
	log.Println("== == == == == == == == == == ")

	// gRPCリクエストの実行
	resp, err := client.GetPersonByCondition(
		ctx,
		v1GetPersonByConditionRequest,
	)
	log.Println("== == == == == == == == == == ")
	log.Printf("%#v\n", *resp.V1PersonParameterArray.Persons[0].MailAddress)
	log.Println("== == == == == == == == == == ")

	if err != nil {
		return fmt.Errorf("failed to get person: %v", err)
	}

	return nil
}
