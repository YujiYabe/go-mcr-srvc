package grpc_client

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	// "backend/internal/1_framework/grpc_parameter"
	"backend/internal/1_framework/grpc_parameter"
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
	client := grpc_parameter.NewPersonClient(conn)

	name := "a"
	// リクエストの作成
	v1GetPersonByConditionRequest := &grpc_parameter.V1GetPersonByConditionRequest{
		V1PersonParameter: &grpc_parameter.V1PersonParameter{
			Name: &name,
		},
		V1CommonParameter: &grpc_parameter.V1CommonParameter{
			TraceID: pkg.GetTraceID(ctx),
		},
	}

	ctx = metadata.AppendToOutgoingContext(
		ctx,
		string(pkg.TraceIDKey),
		pkg.GetTraceID(ctx),
	)
	log.Println("== == == == == == == == == == ")
	pkg.Logging(ctx, pkg.GetTraceID(ctx))
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
