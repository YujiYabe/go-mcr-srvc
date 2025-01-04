package withmiddleware

import (
	"github.com/labstack/echo"

	httpMiddleware "backend/internal/1_framework/middleware/http"
	"backend/internal/2_adapter/controller"
)

func NewRoute(
	EchoEcho *echo.Echo,
	toController controller.ToController,
	parent *echo.Group,
) {
	group := parent.Group(
		"/withmiddleware",
		// httpMiddleware.JWTMiddleware(),
	)

	group.POST(
		"/fetchAccessToken",
		func(c echo.Context) (
			err error,
		) {
			return fetchAccessToken(c, toController)
		},
		// httpMiddleware.JWTMiddleware(),
	)

	group.GET(
		"/protected",
		func(c echo.Context) (
			err error,
		) {
			return protected(c, toController)
		},
		httpMiddleware.JWTMiddleware(),
	)
}
