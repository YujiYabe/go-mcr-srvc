package ws

import (
	"github.com/labstack/echo"

	app "backend/internal/1_framework/instacook/http/v1/ws/app"
	"backend/internal/2_adapter/controller"
)

func NewRoute(
	EchoEcho *echo.Echo,
	Controller controller.ToController,
	parrent *echo.Group,
) {

	group := parrent.Group("/ws")

	app.NewRoute(
		EchoEcho,
		Controller,
		group,
	)

}
