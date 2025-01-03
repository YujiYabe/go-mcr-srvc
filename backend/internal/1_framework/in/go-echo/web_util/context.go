package web_util

import (
	"backend/pkg"
	"context"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

func ContextMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			correlationID := pkg.GetCorrelationID(c.Request().Context())

			if correlationID == "" {
				correlationID = uuid.New().String() // 新しいリクエストIDを生成
			}

			// リクエストIDをコンテキストに追加
			ctx := context.WithValue(c.Request().Context(), pkg.CorrelationIDKey, correlationID)
			c.SetRequest(c.Request().WithContext(ctx))

			// レスポンスヘッダーに相関IDを設定 TODO: レスポンスに含めてよいか要調査
			c.Response().Header().Set(
				string(pkg.CorrelationIDKey),
				correlationID,
			)

			// 次のハンドラを呼び出す
			return next(c)
		}
	}
}
