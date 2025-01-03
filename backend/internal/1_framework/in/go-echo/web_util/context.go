package web_util

import (
	"backend/pkg"
	"context"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

func generateCorrelationID(
	existingID string,
) string {
	if existingID == "" {
		return uuid.New().String()
	}
	return existingID
}

func setCorrelationIDContext(
	c echo.Context,
	correlationID string,
) {
	ctx := context.WithValue(
		c.Request().Context(),
		pkg.CorrelationIDKey,
		correlationID,
	)

	c.SetRequest(c.Request().WithContext(ctx))
}

func setCorrelationIDHeader(
	c echo.Context,
	correlationID string,
) {
	c.Response().Header().Set(
		string(pkg.CorrelationIDKey),
		correlationID,
	)
}

func ContextMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			correlationID := generateCorrelationID(
				pkg.GetCorrelationID(c.Request().Context()),
			)

			setCorrelationIDContext(
				c,
				correlationID,
			)

			setCorrelationIDHeader(
				c,
				correlationID,
			)

			return next(c)
		}
	}
}
