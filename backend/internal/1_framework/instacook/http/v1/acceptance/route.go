package acceptance

import (
	"github.com/labstack/echo"
)

func AddRoute(parrent *echo.Group) {
	g := parrent.Group("/acceptance")

	g.GET("", func(c echo.Context) error { return Get(c) }) // 商品情報
}
