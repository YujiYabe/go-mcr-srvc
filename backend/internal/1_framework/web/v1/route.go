package v1

import (
	"github.com/labstack/echo"

	"backend/internal/1_framework/web/v1/acceptance"
	"backend/internal/1_framework/web/v1/admin"
	"backend/internal/1_framework/web/v1/casher"
	"backend/internal/1_framework/web/v1/client"
	"backend/internal/1_framework/web/v1/delivery"
	"backend/internal/1_framework/web/v1/kitchen"
	"backend/internal/1_framework/web/v1/order"
	"backend/internal/1_framework/web/v1/ws"

	"backend/internal/2_adapter/controller"
	// "backend/internal/1_framework/web/admin"
	// "backend/internal/1_framework/web/v1/casher"
	// "backend/internal/1_framework/web/v1/delivery"
)

func NewRoute(
	EchoEcho *echo.Echo,
	Controller controller.ToController,
	parrent *echo.Group,
) {
	group := parrent.Group("/v1")

	admin.NewRoute(
		EchoEcho,
		Controller,
		group,
	)

	acceptance.NewRoute(
		EchoEcho,
		Controller,
		group,
	)

	delivery.NewRoute(
		EchoEcho,
		Controller,
		group,
	)

	casher.NewRoute(
		EchoEcho,
		Controller,
		group,
	)

	order.NewRoute(
		EchoEcho,
		Controller,
		group,
	)

	kitchen.NewRoute(
		EchoEcho,
		Controller,
		group,
	)

	client.NewRoute(
		EchoEcho,
		Controller,
		group,
	)

	ws.NewRoute(
		EchoEcho,
		Controller,
		group,
	)

}
