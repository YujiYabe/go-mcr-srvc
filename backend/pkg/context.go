// Package pkg はコンテキスト関連のユーティリティを提供します
package pkg

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
