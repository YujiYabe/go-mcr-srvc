package grpc_util

import (
	"context"
	"log"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	// プロトバフで生成されたパッケージ
)

// UnaryInterceptor の実装
// Define a custom type for the context key
type contextKey string

const requestIDKey contextKey = "requestID"

func RequestIDInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	// メタデータからリクエストIDを取得
	md, ok := metadata.FromIncomingContext(ctx)
	var requestID string
	if ok {
		values := md["x-request-id"]
		if len(values) > 0 {
			requestID = values[0]
		}
	}

	// リクエストIDが無い場合は新規生成
	if requestID == "" {
		requestID = uuid.New().String()
	}

	// リクエストIDをコンテキストに追加
	ctx = context.WithValue(ctx, requestIDKey, requestID)

	// ログ出力
	log.Printf("Handling request: %s", requestID)

	// 次のハンドラーを呼び出す
	return handler(ctx, req)
}
