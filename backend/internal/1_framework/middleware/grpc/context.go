package grpc_middleware

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func MetadataToContext(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	// メタデータからリクエストIDを取得
	md, ok := metadata.FromIncomingContext(ctx)

	if ok { // メタデータが存在する場合、以下各パラメータをコンテキストに追加
		ctx = traceIDToContext(ctx, md)
		ctx = requestStartTimeToContext(ctx, md)
		ctx = timestampToContext(ctx, md)
	}

	// 次のハンドラーを呼び出す
	return handler(ctx, req)
}

// func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
// 	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
// 		// コンテキストから値を取得してメタデータに変換
// 		md := convertToMetadata(ctx)

// 		// 新しいコンテキストにメタデータを設定
// 		newCtx := metadata.NewOutgoingContext(ctx, md)

// 		return handler(newCtx, req)
// 	}
// }

// // クライアント側での使用例
// func UnaryClientInterceptor() grpc.UnaryClientInterceptor {
// 	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
// 		// コンテキストから値を取得してメタデータに変換
// 		md := convertToMetadata(ctx)

// 		// 新しいコンテキストにメタデータを設定
// 		newCtx := metadata.NewOutgoingContext(ctx, md)

// 		return invoker(newCtx, method, req, reply, cc, opts...)
// 	}
// }
