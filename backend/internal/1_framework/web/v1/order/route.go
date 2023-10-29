package order

import (
	"github.com/labstack/echo"

	sold "backend/internal/1_framework/web/v1/order/sold"
	"backend/internal/2_adapter/controller"
)

func NewRoute(
	EchoEcho *echo.Echo,
	Controller controller.ToController,
	parrent *echo.Group,
) {

	group := parrent.Group("/order")

	sold.NewRoute(
		EchoEcho,
		Controller,
		group,
	)

	// group.GET("", func(c echo.Context) error { return Get(c, Controller) })     // 管理画面 html
	// group.PATCH("", func(c echo.Context) error { return Patch(c, Controller) }) // 管理画面 商品更新
}
