package client

import (
	"github.com/labstack/echo"

	"backend/internal/1_framework/instacook/http/v1/client/camera"
	"backend/internal/1_framework/instacook/http/v1/client/monitor"
	"backend/internal/1_framework/instacook/http/v1/client/queue"

	"backend/internal/2_adapter/controller"
)

func NewRoute(
	EchoEcho *echo.Echo,
	Controller controller.ToController,
	parrent *echo.Group,
) {
	group := parrent.Group("/client")

	camera.NewRoute(
		EchoEcho,
		Controller,
		group,
	)
	monitor.NewRoute(
		EchoEcho,
		Controller,
		group,
	)
	queue.NewRoute(
		EchoEcho,
		Controller,
		group,
	)

}
