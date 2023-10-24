package admin_client

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
	group := parrent.Group("/client")

	group.GET("/:number", func(c echo.Context) error { return Get(c, Controller) }) // 管理画面 顧客印刷 html
}
