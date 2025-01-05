package http_middleware

import (
	"context"

	"github.com/google/uuid"
	"github.com/labstack/echo"

	"google.golang.org/grpc/metadata"
)

// contextKey はコンテキストのキー型を定義します
type contextKey string

// traceID は共通リクエストIDを格納するためのコンテキストキーです
const TraceIDKey contextKey = "traceID"

func ContextMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			// traceID設定
			traceID := generateTraceID(
				GetTraceID(c.Request().Context()),
			)
			setTraceIDContext(c, traceID)
			setTraceIDHeader(c, traceID)

			return next(c)
		}
	}
}

func generateTraceID(
	existingID string,
) string {
	if existingID == "" {
		return uuid.New().String()
	}
	return existingID
}

func setTraceIDContext(
	c echo.Context,
	traceID string,
) {
	ctx := context.WithValue(
		c.Request().Context(),
		TraceIDKey,
		traceID,
	)

	c.SetRequest(c.Request().WithContext(ctx))
}

func setTraceIDHeader(
	c echo.Context,
	traceID string,
) {
	c.Response().Header().Set(
		string(TraceIDKey),
		traceID,
	)
}

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
