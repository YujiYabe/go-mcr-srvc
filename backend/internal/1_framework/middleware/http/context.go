package http_middleware

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo"

	valueObject "backend/internal/4_domain/value_object"
)

func ContextMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			// traceID 設定 ---------
			traceID := generateTraceID(
				valueObject.GetTraceID(c.Request().Context()),
			)
			setTraceIDContext(c, traceID)
			setTraceIDHeader(c, traceID)

			// requestStartTime 設定 ---------
			requestStartTime := generateRequestStartTime(
				valueObject.GetRequestStartTime(c.Request().Context()),
			)
			setRequestStartTimeContext(c, requestStartTime)

			// timeoutSecond 設定 ---------
			timeoutSecond := generateTimeoutSecond(
				valueObject.GetRequestStartTime(c.Request().Context()),
			)
			setTimeoutSecondContext(c, timeoutSecond)

			return next(c)
		}
	}
}

func setTimeoutSecondContext(
	c echo.Context,
	timeoutSecond int64,
) {
	ctx := context.WithValue(
		c.Request().Context(),
		valueObject.TimeOutSecondContextName,
		timeoutSecond,
	)

	c.SetRequest(c.Request().WithContext(ctx))
}

func generateTimeoutSecond(
	requestStartTime int64,
) int64 {
	currentTimestamp := time.Now().UnixMilli()
	requestEndTime := time.UnixMilli(requestStartTime).Add(5 * time.Second).UnixMilli()
	timeoutSecond := requestEndTime - currentTimestamp

	return timeoutSecond
}

func generateRequestStartTime(
	existingValue int64,
) int64 {
	if existingValue == 0 {
		return time.Now().UnixMilli()
	}
	return existingValue
}

func setRequestStartTimeContext(
	c echo.Context,
	requestStartTime int64,
) {
	ctx := context.WithValue(
		c.Request().Context(),
		valueObject.RequestStartTimeContextName,
		requestStartTime,
	)

	c.SetRequest(c.Request().WithContext(ctx))
}

func generateTraceID(
	existingValue string,
) string {
	if existingValue == "" {
		return uuid.New().String()
	}
	return existingValue
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
