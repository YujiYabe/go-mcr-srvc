package admin_product

import (
	"github.com/labstack/echo"

	"backend/internal/2_adapter/controller"
)

// func AddRoute(parrent *echo.Group) {
// 	grp := parrent.Group("/product")
// 	grp.GET("", func(c echo.Context) error { return Get(c) })             // 管理画面 html
// 	grp.PATCH("/:number", func(c echo.Context) error { return Patch(c) }) // 管理画面 商品更新
// }

func NewRoute(
	EchoEcho *echo.Echo,
	Controller controller.ToController,
	parrent *echo.Group,
) {
	group := parrent.Group("/product")

	group.GET("", func(c echo.Context) error { return Get(c, Controller) })             // 管理画面 html
	group.PATCH("/:number", func(c echo.Context) error { return Patch(c, Controller) }) // 管理画面 商品更新

}
