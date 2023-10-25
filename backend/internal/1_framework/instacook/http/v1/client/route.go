package client

import (
	"github.com/labstack/echo"

	monitor "backend/internal/1_framework/instacook/http/v1/client/monitor"

	"backend/internal/2_adapter/controller"
)

func NewRoute(
	EchoEcho *echo.Echo,
	Controller controller.ToController,
	parrent *echo.Group,
) {
	group := parrent.Group("/client")

	monitor.NewRoute(
		EchoEcho,
		Controller,
		group,
	)

	// group.GET("", func(c echo.Context) error { return Get(c, Controller) })     // 管理画面 html
	// group.PATCH("", func(c echo.Context) error { return Patch(c, Controller) }) // 管理画面 商品更新
}
