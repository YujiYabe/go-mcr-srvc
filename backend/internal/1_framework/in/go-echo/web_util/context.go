package web_util

import (
	"github.com/google/uuid"
	"github.com/labstack/echo"
)

func ContextMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			requestID := c.Response().Header().Get(echo.HeaderXRequestID)
			if requestID == "" {
				requestID = uuid.New().String() // 新しいリクエストIDを生成
			}
			// レスポンスヘッダーにリクエストIDを設定
			c.Response().Header().Set(echo.HeaderXRequestID, requestID)

			// 次のハンドラを呼び出す
			return next(c)
		}
	}
}
