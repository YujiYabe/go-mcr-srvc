package grpc_client

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	// "backend/internal/1_framework/grpc_parameter"
	"backend/internal/1_framework/grpc_parameter"
)

// ...
// ViaGRPC ...
func (receiver *GRPCClient) ViaGRPC(
	ctx context.Context,
) (
	err error,
) {
	// gRPCコネクションの作成
	conn, err := grpc.DialContext(
		ctx,
		// "172.18.0.4:3456",
		"backend:3456",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// クライアントの作成
	client := grpc_parameter.NewPersonClient(conn)

	name := "a"
	// リクエストの作成
	v1PersonParameter := &grpc_parameter.V1PersonParameter{
		Name: &name,
	}
	// v1GetPersonByConditionRequest := &grpc_parameter.V1GetPersonByConditionRequest{
	// 	V1PersonParameter: v1PersonParameter,
	// }

	// gRPCリクエストの実行
	resp, err := client.GetPersonByCondition(ctx, v1PersonParameter)
	log.Println("== == == == == == == == == == ")
	log.Printf("%#v\n", *resp.V1PersonParameterArray.Persons[0].MailAddress)
	log.Println("== == == == == == == == == == ")
	if err != nil {
		return fmt.Errorf("failed to get person: %v", err)
	}

	return nil
}
