package grpc_middleware

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	valueObject "backend/internal/4_domain/value_object"
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

// GetTraceID はコンテキストからリクエストIDを取得します
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
	traceID, ok := ctx.Value(valueObject.TraceIDContextName).(string)
	if ok {
		traceIDString = traceID
	}

	return
}
