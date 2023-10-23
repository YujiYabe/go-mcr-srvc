package admin_client

import (
	"backend/internal/2_adapter/controller"

	"github.com/labstack/echo"
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
	grp := parrent.Group("/client")

	grp.GET("/:number", func(c echo.Context) error { return Get(c) }) // 管理画面 顧客印刷 html
}
