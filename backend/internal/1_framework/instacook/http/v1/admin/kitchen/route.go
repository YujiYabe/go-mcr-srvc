package admin_kitchen

import (
	"github.com/labstack/echo"

	"backend/internal/2_adapter/controller"
)

type (
	Client struct {
		EchoEcho   *echo.Echo
		Controller controller.ToController
	}
)

func NewRoute(
	EchoEcho *echo.Echo,
	Controller controller.ToController,
	parrent *echo.Group,
) {
	group := parrent.Group("/kitchen")
	group.GET("/:number", func(c echo.Context) error { return Get(c, Controller) }) // 管理画面 厨房印刷 html
}
