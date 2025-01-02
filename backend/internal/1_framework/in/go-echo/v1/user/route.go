package user

import (
	"github.com/labstack/echo"

	"backend/internal/1_framework/in/go-echo/v1/user/callback"
	"backend/internal/1_framework/in/go-echo/v1/user/login"
	"backend/internal/1_framework/in/go-echo/v1/user/logout"
	webUtil "backend/internal/1_framework/in/go-echo/web_util"
	"backend/internal/2_adapter/controller"
)

func NewRoute(
	EchoEcho *echo.Echo,
	toController controller.ToController,
	parent *echo.Group,
) {
	group := parent.Group(
		"/user",
	)

	group.GET(
		"",
		func(c echo.Context) (
			err error,
		) {
			return get(c, toController)
		},
		webUtil.JWTMiddleware(),
	)

	callback.NewRoute(
		EchoEcho,
		toController,
		group,
	)

	login.NewRoute(
		EchoEcho,
		toController,
		group,
	)

	logout.NewRoute(
		EchoEcho,
		toController,
		group,
	)
}
