package grpc_util

import (
	"backend/pkg"
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func TraceIDInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	// メタデータからリクエストIDを取得
	md, ok := metadata.FromIncomingContext(ctx)
	var traceID string
	if ok {
		values := md["XTraceID"]
		if len(values) > 0 {
			traceID = values[0]
		}
	}

	// リクエストIDが無い場合は新規生成
	if traceID == "" {
		traceID = uuid.New().String()
	}

	// リクエストIDをコンテキストに追加
	ctx = context.WithValue(
		ctx,
		pkg.TraceIDKey,
		traceID,
	)

	// ログ出力

	// 次のハンドラーを呼び出す
	return handler(ctx, req)
}
