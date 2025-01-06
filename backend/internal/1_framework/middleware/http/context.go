package http_middleware

import (
	"context"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	"google.golang.org/grpc/metadata"

	valueObject "backend/internal/4_domain/value_object"
)

func ContextMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			// traceID設定
			traceID := generateTraceID(
				valueObject.GetTraceID(c.Request().Context()),
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
		valueObject.TraceIDContextName,
		traceID,
	)

	c.SetRequest(c.Request().WithContext(ctx))
}

func setTraceIDHeader(
	c echo.Context,
	traceID string,
) {
	c.Response().Header().Set(
		string(valueObject.TraceIDContextName),
		traceID,
	)
}

// gRPCメタデータとの変換を行うレイヤーで適切に変換処理を実装
func ConvertToMetadata(
	ctx context.Context,
) metadata.MD {
	return metadata.New(
		map[string]string{
			string(valueObject.TraceIDMetaName): ctx.Value(valueObject.TraceIDContextName).(string), // メタデータ用にハイフン区切りに変換
		},
	)
}
