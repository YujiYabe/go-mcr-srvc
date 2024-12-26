package login

import (
	"github.com/labstack/echo"

	webUtil "backend/internal/1_framework/input/mobile/web_util"
	// "backend/internal/1_framework/input/mobile/v1/user/login"
	// "backend/internal/1_framework/input/mobile/v1/user/logout"
	// "backend/internal/1_framework/input/mobile/v1/user/callback"

	"backend/internal/2_adapter/controller"
)

func NewRoute(
	EchoEcho *echo.Echo,
	toController controller.ToController,
	parent *echo.Group,
) {
	group := parent.Group(
		"/login",
		webUtil.JWTMiddleware(),
	)

	group.GET(
		"",
		func(c echo.Context) (err error) { return get(c, toController) },
	)

}
