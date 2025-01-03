package grpc_util

import (
	"backend/pkg"
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func CorrelationIDInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	// メタデータからリクエストIDを取得
	md, ok := metadata.FromIncomingContext(ctx)
	var correlationID string
	if ok {
		values := md["XCorrelationID"]
		if len(values) > 0 {
			correlationID = values[0]
		}
	}

	// リクエストIDが無い場合は新規生成
	if correlationID == "" {
		correlationID = uuid.New().String()
	}

	// リクエストIDをコンテキストに追加
	ctx = context.WithValue(
		ctx,
		pkg.CorrelationIDKey,
		correlationID,
	)

	// ログ出力

	// 次のハンドラーを呼び出す
	return handler(ctx, req)
}
