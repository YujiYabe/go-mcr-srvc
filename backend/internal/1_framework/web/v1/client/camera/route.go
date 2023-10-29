package camera

import (
	"github.com/labstack/echo"

	"backend/internal/2_adapter/controller"
)

func NewRoute(
	EchoEcho *echo.Echo,
	Controller controller.ToController,
	parrent *echo.Group,
) {
	group := parrent.Group("/camera")

	group.POST("/:number", func(c echo.Context) error { return Post(c, Controller) }) // 予約商品の送信
}
