package grpc_middleware

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// contextKey はコンテキストのキー型を定義します
type contextKey string

// traceID は共通リクエストIDを格納するためのコンテキストキーです
const TraceIDKey contextKey = "traceID"

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
		values := md.Get(string(TraceIDKey))
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
		TraceIDKey,
		traceID,
	)

	// ログ出力

	// 次のハンドラーを呼び出す
	return handler(ctx, req)
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
