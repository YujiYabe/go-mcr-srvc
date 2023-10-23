package admin_kitchen

import (
	"github.com/labstack/echo"
)

func AddRoute(parrent *echo.Group) {
	grp := parrent.Group("/kitchen")
	grp.GET("/:number", func(c echo.Context) error { return Get(c) }) // 管理画面 厨房印刷 html
}
