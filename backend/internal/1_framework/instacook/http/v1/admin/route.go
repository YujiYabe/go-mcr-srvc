package admin

import (
	"github.com/labstack/echo"

	client "backend/internal/1_framework/instacook/http/v1/admin/client"
	kitchen "backend/internal/1_framework/instacook/http/v1/admin/kitchen"
	language "backend/internal/1_framework/instacook/http/v1/admin/language"
	product "backend/internal/1_framework/instacook/http/v1/admin/product"
	"backend/internal/2_adapter/controller"
)

type (
	Admin struct {
		EchoEcho   *echo.Echo
		Controller controller.ToController
	}
)

func NewRoute(
	EchoEcho *echo.Echo,
	Controller controller.ToController,
	parrent *echo.Group,
) {

	group := parrent.Group("/admin")

	client.NewRoute(
		EchoEcho,
		Controller,
		group,
	)

	kitchen.NewRoute(
		EchoEcho,
		Controller,
		group,
	)

	language.NewRoute(
		EchoEcho,
		Controller,
		group,
	)

	product.NewRoute(
		EchoEcho,
		Controller,
		group,
	)

	// client.AddRoute(g)
	// kitchen.AddRoute(g)
	// language.AddRoute(g)
	// product.AddRoute(g)

}
