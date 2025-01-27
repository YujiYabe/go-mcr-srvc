package person

import (
	"github.com/labstack/echo/v4"

	"backend/internal/2_adapter/controller"
)

func NewRoute(
	EchoEcho *echo.Echo,
	toController controller.ToController,
	parent *echo.Group,
) {
	group := parent.Group("/person")

	group.GET(
		"",
		func(c echo.Context) (
			err error,
		) {
			return get(c, toController)
		},
	)

	group.GET(
		"/viaGRPC",
		func(c echo.Context) (
			err error,
		) {
			return viaGRPC(c, toController)
		},
	)
}
