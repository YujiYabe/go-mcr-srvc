package monitor

import (
	"github.com/labstack/echo"

	"backend/internal/2_adapter/controller"
)

func NewRoute(
	EchoEcho *echo.Echo,
	Controller controller.ToController,
	parrent *echo.Group,
) {

	group := parrent.Group("/monitor")

	group.GET("/:number", func(c echo.Context) error { return Get(c, Controller) }) // 厨房htmlの返却
	// group.PATCH("", func(c echo.Context) error { return Patch(c, Controller) }) // 管理画面 商品更新

}
