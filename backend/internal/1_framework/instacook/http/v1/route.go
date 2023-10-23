package v1

import (
	"backend/internal/1_framework/instacook/http/v1/admin"
	"backend/internal/2_adapter/controller"

	"github.com/labstack/echo"
	// "backend/internal/1_framework/in/http/acceptance"
	// "backend/internal/1_framework/instacook/http/admin"
	// "backend/internal/1_framework/in/http/casher"
	// "backend/internal/1_framework/in/http/client"
	// "backend/internal/1_framework/in/http/delivery"
	// "backend/internal/1_framework/in/http/kitchen"
	// "backend/internal/1_framework/in/http/order"
)

type (
	V1 struct {
		EchoEcho   *echo.Echo
		Controller controller.ToController
	}
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

	// client.AddRoute(g)
	// kitchen.AddRoute(g)
	// casher.AddRoute(g)

	// delivery.AddRoute(g)
	// admin.AddRoute(g)

	// order.AddRoute(g)
	// acceptance.AddRoute(g)

}
