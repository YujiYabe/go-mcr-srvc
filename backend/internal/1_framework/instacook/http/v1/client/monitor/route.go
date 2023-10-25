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

	// お客様側モニターhtml。エンドポイントの番号を並び順に合わせる
	group.GET("/:number", func(c echo.Context) error { return Get(c, Controller) })
}
