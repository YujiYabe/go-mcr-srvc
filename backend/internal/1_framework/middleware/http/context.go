package http_middleware

import (
	"context"

	"github.com/labstack/echo/v4"

	domainObject "backend/internal/4_domain/domain_object"
	groupObject "backend/internal/4_domain/group_object"
	"backend/internal/logger"
)

func ContextMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			clientIP := c.RealIP()
			userAgent := c.Request().UserAgent()
			locale := c.Request().Header.Get("Accept-Language")
			timeZone := c.Request().Header.Get("Time-Zone")

			// 他に取れそうな情報
			// 認証関連:
			// 		Authorization - 現在も取得している認証トークン
			// 		Cookie - セッション情報やその他のクッキー
			// キャッシュ制御:
			// 		If-Modified-Since
			// 		If-None-Match
			// 		Cache-Control
			// コンテンツネゴシエーション:
			// 		Accept - クライアントが受け入れ可能なコンテンツタイプ
			// 		Accept-Encoding - 圧縮形式
			// 		Accept-Charset - 文字エンコーディング
			// プロキシ/転送情報:
			// 		X-Forwarded-For - プロキシを経由した場合の元のIPアドレス
			// 		X-Forwarded-Proto - 元のプロトコル
			// 		X-Request-ID - リクエスト追跡用ID
			// セキュリティ:
			// 		Origin - CORSリクエストの送信元
			// 		Referer - リクエスト元のURL

			newRequestContextArgs := &groupObject.NewRequestContextArgs{
				ClientIP:  &clientIP,
				UserAgent: &userAgent,
				Locale:    &locale,
				TimeZone:  &timeZone,
			}

			requestContext := groupObject.NewRequestContext(
				c.Request().Context(),
				newRequestContextArgs,
			)

			if requestContext.GetError() != nil {
				logger.Logging(c.Request().Context(), requestContext.GetError())
				return requestContext.GetError()
			}
			ctx := context.WithValue(
				c.Request().Context(),
				groupObject.RequestContextContextName,
				*requestContext,
			)

			// ________________________________
			// logで追跡するために、contextにTraceIDを設定する
			ctx = context.WithValue(
				ctx,
				domainObject.TraceIDContextName,
				requestContext.TraceID.GetValue(),
			)

			c.SetRequest(c.Request().WithContext(ctx))

			return next(c)
		}
	}
}
