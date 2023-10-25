package v1

import (
	"github.com/labstack/echo"

	"backend/internal/1_framework/instacook/http/v1/acceptance"
	"backend/internal/1_framework/instacook/http/v1/admin"
	"backend/internal/1_framework/instacook/http/v1/delivery"

	"backend/internal/2_adapter/controller"
	// "backend/internal/1_framework/instacook/http/admin"
	// "backend/internal/1_framework/instacook/http/v1/casher"
	// "backend/internal/1_framework/instacook/http/v1/client"
	// "backend/internal/1_framework/instacook/http/v1/delivery"
	// "backend/internal/1_framework/instacook/http/v1/kitchen"
	// "backend/internal/1_framework/instacook/http/v1/order"
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

	// client.AddRoute(g)
	// kitchen.AddRoute(g)
	// casher.AddRoute(g)

	// order.AddRoute(g)

}
