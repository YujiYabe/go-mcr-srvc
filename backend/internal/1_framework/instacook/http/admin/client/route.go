package admin_client

import (
	"github.com/labstack/echo"
)

func AddRoute(parrent *echo.Group) {
	grp := parrent.Group("/client")
	grp.GET("/:number", func(c echo.Context) error { return Get(c) }) // 管理画面 顧客印刷 html
}
