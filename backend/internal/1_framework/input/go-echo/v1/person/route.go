package person

import (
	"github.com/labstack/echo"

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
		func(c echo.Context) (err error) { return get(c, toController) },
	)
}
