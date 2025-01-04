// Package pkg はコンテキスト関連のユーティリティを提供します
package pkg

import (
	"context"

	"google.golang.org/grpc/metadata"
)

// contextKey はコンテキストのキー型を定義します
type contextKey string

// traceID は共通リクエストIDを格納するためのコンテキストキーです
const TraceIDKey contextKey = "traceID"

// GetNewContext は新しいコンテキストを生成します
// X-Trace-IDヘッダーの値をコンテキストに埋め込みます
//
// パラメータ:
//   - ctx: 親コンテキスト
//   - traceID: X-Trace-IDヘッダーの値
//
// 戻り値:
//   - newCtx: 新しく生成されたコンテキスト
func GetNewContext(
	ctx context.Context,
	keyName contextKey,
	keyValue interface{},
) (
	newCtx context.Context,
) {
	newCtx = context.WithValue(
		ctx,
		keyName,
		keyValue,
	)

	return
}

// GetTraceID はコンテキストからリクエストIDを取得します
//
// パラメータ:
//   - ctx: リクエストIDを含むコンテキスト
//
// 戻り値:
//   - traceIDString: 取得したリクエストID。取得できない場合は空文字列
func GetTraceID(
	ctx context.Context,
) (
	traceIDString string,
) {
	traceID, ok := ctx.Value(TraceIDKey).(string)
	if ok {
		traceIDString = traceID
	}

	return
}

// gRPCメタデータとの変換を行うレイヤーで適切に変換処理を実装
func ConvertToMetadata(
	ctx context.Context,
) metadata.MD {
	traceID := ctx.Value(TraceIDKey).(string)
	return metadata.New(
		map[string]string{
			"trace-id": traceID, // メタデータ用にハイフン区切りに変換
		},
	)
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
