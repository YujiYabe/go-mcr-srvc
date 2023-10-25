package sold

import (
	"github.com/labstack/echo"

	"backend/internal/2_adapter/controller"
)

func NewRoute(
	EchoEcho *echo.Echo,
	Controller controller.ToController,
	parrent *echo.Group,
) {

	group := parrent.Group("/sold")

	group.GET("", func(c echo.Context) error { return Get(c, Controller) })     // 管理画面 html
	// group.PATCH("", func(c echo.Context) error { return Patch(c, Controller) }) // 管理画面 商品更新
}
