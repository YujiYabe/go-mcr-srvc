// Package pkg はコンテキスト関連のユーティリティを提供します
package pkg

import (
	"context"
)

// contextKey はコンテキストのキー型を定義します
type contextKey string

// correlationID は共通リクエストIDを格納するためのコンテキストキーです
const CorrelationIDKey contextKey = "X-Correlation-ID"

// GetNewContext は新しいコンテキストを生成します
// X-Request-IDヘッダーの値をコンテキストに埋め込みます
//
// パラメータ:
//   - ctx: 親コンテキスト
//   - correlationID: X-Request-IDヘッダーの値
//
// 戻り値:
//   - newCtx: 新しく生成されたコンテキスト
func GetNewContext(
	ctx context.Context,
	correlationID string,
) (
	newCtx context.Context,
) {
	newCtx = context.WithValue(
		ctx,
		CorrelationIDKey,
		correlationID,
	)

	return
}

// GetCorrelationID はコンテキストからリクエストIDを取得します
//
// パラメータ:
//   - ctx: リクエストIDを含むコンテキスト
//
// 戻り値:
//   - requestIDString: 取得したリクエストID。取得できない場合は空文字列
func GetCorrelationID(
	ctx context.Context,
) (
	correlationIDString string,
) {
	requestID, ok := ctx.Value(CorrelationIDKey).(string)
	if ok {
		correlationIDString = requestID
	}

	return
}
