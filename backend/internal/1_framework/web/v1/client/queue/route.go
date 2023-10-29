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

	group.GET("/:number", func(c echo.Context) error { return Get(c, Controller) })   // クライアント側予約商品jsonの返却
	group.POST("/:number", func(c echo.Context) error { return Post(c, Controller) }) // クライアント側予約商品登録
	// group.DELETE("/:number", func(c echo.Context) error { return Delete(c, Controller) }) // クライアント側予約商品削除
}
