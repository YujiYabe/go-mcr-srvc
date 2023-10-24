package acceptance

import (
	"github.com/labstack/echo"

	"backend/internal/2_adapter/controller"
)

// func AddRoute(parrent *echo.Group) {
// 	g := parrent.Group("/acceptance")

// 	g.GET("", func(c echo.Context) error { return Get(c) }) // 商品情報
// }

type (
	Acceptance struct {
		EchoEcho   *echo.Echo
		Controller controller.ToController
	}
)

func NewRoute(
	EchoEcho *echo.Echo,
	Controller controller.ToController,
	parrent *echo.Group,
) {
	group := parrent.Group("/acceptance")

	group.GET("", func(c echo.Context) error { return Get(c, Controller) }) // 商品情報
}
