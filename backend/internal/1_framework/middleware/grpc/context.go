package grpc_middleware

import (
	"context"

	valueObject "backend/internal/4_domain/value_object"

	"github.com/google/uuid"
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

	// ログ出力
	if ok {
		ctx = traceIDToContext(ctx, md)
	}

	// 次のハンドラーを呼び出す
	return handler(ctx, req)
}

func traceIDToContext(
	nowCtx context.Context,
	md metadata.MD,
) (
	newCtx context.Context,
) {
	var traceID string

	values := md.Get(string(valueObject.TraceIDMetaName))
	if len(values) > 0 {
		traceID = values[0]
	}

	// リクエストIDが無い場合は新規生成
	if traceID == "" {
		traceID = uuid.New().String()
	}

	// リクエストIDをコンテキストに追加
	newCtx = context.WithValue(
		nowCtx,
		valueObject.TraceIDContextName,
		traceID,
	)

	return
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
