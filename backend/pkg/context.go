// Package pkg はコンテキスト関連のユーティリティを提供します
package pkg

import (
	"context"
)

// contextKey はコンテキストのキー型を定義します
type contextKey string

// traceID は共通リクエストIDを格納するためのコンテキストキーです
const TraceIDKey contextKey = "trace-id"

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
	traceID string,
) (
	newCtx context.Context,
) {
	newCtx = context.WithValue(
		ctx,
		TraceIDKey,
		traceID,
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
