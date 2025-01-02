package callback

import (
	"github.com/labstack/echo"

	webUtil "backend/internal/1_framework/in/go-echo/web_util"
	// "backend/internal/1_framework/in/go-echo/v1/user/login"
	// "backend/internal/1_framework/in/go-echo/v1/user/logout"
	// "backend/internal/1_framework/in/go-echo/v1/user/callback"

	"backend/internal/2_adapter/controller"
)

func NewRoute(
	EchoEcho *echo.Echo,
	toController controller.ToController,
	parent *echo.Group,
) {
	group := parent.Group(
		"/callback",
		webUtil.JWTMiddleware(),
	)

	group.GET(
		"",
		func(c echo.Context) (
			err error,
		) {
			return get(c, toController)
		},
	)

}
