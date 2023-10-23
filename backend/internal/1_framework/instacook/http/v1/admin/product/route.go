package admin_product

import (
	"github.com/labstack/echo"
)

func AddRoute(parrent *echo.Group) {
	grp := parrent.Group("/product")
	grp.GET("", func(c echo.Context) error { return Get(c) })             // 管理画面 html
	grp.PATCH("/:number", func(c echo.Context) error { return Patch(c) }) // 管理画面 商品更新
}
