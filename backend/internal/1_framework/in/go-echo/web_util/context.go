package web_util

import (
	"backend/pkg"
	"context"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

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
		pkg.TraceIDKey,
		traceID,
	)

	c.SetRequest(c.Request().WithContext(ctx))
}

func setTraceIDHeader(
	c echo.Context,
	traceID string,
) {
	c.Response().Header().Set(
		string(pkg.TraceIDKey),
		traceID,
	)
}

func ContextMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			traceID := generateTraceID(
				pkg.GetTraceID(c.Request().Context()),
			)

			setTraceIDContext(
				c,
				traceID,
			)

			setTraceIDHeader(
				c,
				traceID,
			)

			return next(c)
		}
	}
}
