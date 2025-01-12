package http_middleware

import (
	"context"
	"log"

	"github.com/labstack/echo"

	groupObject "backend/internal/4_domain/group_object"
)

func ContextMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			clientIP := c.RealIP()
			userAgent := c.Request().UserAgent()
			locale := c.Request().Header.Get("Accept-Language")
			timezone := c.Request().Header.Get("Time-Zone")

			newRequestContextArgs := &groupObject.NewRequestContextArgs{
				ClientIP:  &clientIP,
				UserAgent: &userAgent,
				Locale:    &locale,
				Timezone:  &timezone,
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

			c.SetRequest(c.Request().WithContext(ctx))

			return next(c)
		}
	}
}
