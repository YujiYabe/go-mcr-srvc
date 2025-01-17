package http_middleware

import (
	"context"
	"log"

	"github.com/labstack/echo"

	groupObject "backend/internal/4_domain/group_object"
	valueObject "backend/internal/4_domain/value_object"
)

func ContextMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			clientIP := c.RealIP()
			userAgent := c.Request().UserAgent()
			locale := c.Request().Header.Get("Accept-Language")
			timeZone := c.Request().Header.Get("Time-Zone")

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
				log.Println(requestContext.GetError())
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
				valueObject.TraceIDContextName,
				requestContext.TraceID.GetValue(),
			)

			c.SetRequest(c.Request().WithContext(ctx))

			return next(c)
		}
	}
}
