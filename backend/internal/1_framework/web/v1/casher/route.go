package casher

import (
	"github.com/labstack/echo"

	"backend/internal/2_adapter/controller"
)

func NewRoute(
	EchoEcho *echo.Echo,
	Controller controller.ToController,
	parrent *echo.Group,
) {

	group := parrent.Group("/casher")

	group.GET("", func(c echo.Context) error { return Get(c, Controller) })               // 管理画面 html
	group.POST("", func(c echo.Context) error { return Post(c, Controller) })             // 管理画面 html
	group.DELETE("/:number", func(c echo.Context) error { return Delete(c, Controller) }) // 管理画面 html
}
