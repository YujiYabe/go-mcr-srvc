package queue

import (
	"github.com/labstack/echo"

	"backend/internal/2_adapter/controller"
)

func NewRoute(
	EchoEcho *echo.Echo,
	Controller controller.ToController,
	parrent *echo.Group,
) {
	group := parrent.Group("/queue")

	group.GET("/:number", func(c echo.Context) error { return Get(c, Controller) })
	// 厨房待ち合わせjsonの返却。エンドポイントの番号を並び順に合わせる。preparing_listのリスト→reserving_listのリストの順番に表示する。
}
