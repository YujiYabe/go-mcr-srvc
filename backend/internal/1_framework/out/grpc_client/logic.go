package grpc_client

import (
	"context"
)

// ...
// RequestToAuth0 ...
func (receiver *GRPCClient) RequestToAuth0(
	ctx context.Context,
) (
	err error,
) {
	// conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	// if err != nil {
	// 	return nil, err
	// }
	// defer conn.Close()
	// cpbconn := cpb.NewSendContentServiceClient(conn)

	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()

	// res, err := cpbconn.SendContentRPC(ctx, request)
	// if err != nil {
	// 	return nil, err
	// }

	// gRPCコネクションの作成
	// conn, err := grpc.Dial("hogehoge:3456", grpc.WithInsecure())
	// if err != nil {
	// 	return fmt.Errorf("failed to connect: %v", err)
	// }
	// defer conn.Close()

	// // クライアントの作成
	// client := person.NewPersonClient(conn)

	// // リクエストの作成
	// req := &person.GetPersonByConditionRequest{
	// 	// 必要なパラメータをセット
	// }

	// // gRPCリクエストの実行
	// resp, err := client.GetPersonByCondition(ctx, req)
	// if err != nil {
	// 	return fmt.Errorf("failed to get person: %v", err)
	// }

	// レスポンスの処理
	// resp を使用した処理をここに実装

	return nil
}
