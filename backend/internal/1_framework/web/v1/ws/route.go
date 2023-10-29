package ws

import (
	"github.com/labstack/echo"

	app "backend/internal/1_framework/web/v1/ws/app"
	"backend/internal/2_adapter/controller"
)

func NewRoute(
	EchoEcho *echo.Echo,
	Controller controller.ToController,
	parrent *echo.Group,
) {
	go app.SendToAgents()

	group := parrent.Group("/ws")

	app.NewRoute(
		EchoEcho,
		Controller,
		group,
	)

}
