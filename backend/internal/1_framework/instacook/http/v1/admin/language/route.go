package admin_language

import (
	"github.com/labstack/echo"
)

func AddRoute(parrent *echo.Group) {
	grp := parrent.Group("/language")
	grp.GET("/:number", func(c echo.Context) error { return Get(c) }) // 管理画面 言語印刷 html
}
