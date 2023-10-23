package admin_kitchen

import (
	"backend/internal/2_adapter/controller"

	"github.com/labstack/echo"
)

type (
	Client struct {
		EchoEcho   *echo.Echo
		Controller controller.ToController
	}
)

func NewRoute(
	EchoEcho *echo.Echo,
	Controller controller.ToController,
	parrent *echo.Group,
) {
	grp := parrent.Group("/kitchen")

	grp.GET("/:number", func(c echo.Context) error { return Get(c, Controller) }) // 管理画面 厨房印刷 html
	// grp.GET("/:number", func(
	// 	c echo.Context,
	// 	Controller controller.ToController,
	// ) error {
	// 	return Get(c, Controller)
	// }) // 管理画面 顧客印刷 html

}

// func AddRoute(parrent *echo.Group) {
// 	grp := parrent.Group("/kitchen")
// 	grp.GET("/:number", func(c echo.Context) error { return Get(c) }) // 管理画面 厨房印刷 html
// }
