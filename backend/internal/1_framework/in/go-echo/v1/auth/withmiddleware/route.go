package withmiddleware

import (
	"github.com/labstack/echo"

	webUtil "backend/internal/1_framework/in/go-echo/web_util"
	"backend/internal/2_adapter/controller"
)

func NewRoute(
	EchoEcho *echo.Echo,
	toController controller.ToController,
	parent *echo.Group,
) {
	group := parent.Group(
		"/withmiddleware",
		webUtil.JWTMiddleware(),
	)

	group.GET(
		"/getAccessToken",
		func(c echo.Context) (err error) { return getAccessToken(c, toController) },
		webUtil.JWTMiddleware(),
	)

	group.GET(
		"/protected",
		func(c echo.Context) (err error) { return protected(c, toController) },
		webUtil.JWTMiddleware(),
	)
}
