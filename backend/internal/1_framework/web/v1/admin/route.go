package admin

import (
	"github.com/labstack/echo"

	client "backend/internal/1_framework/web/v1/admin/client"
	kitchen "backend/internal/1_framework/web/v1/admin/kitchen"
	language "backend/internal/1_framework/web/v1/admin/language"
	product "backend/internal/1_framework/web/v1/admin/product"
	"backend/internal/2_adapter/controller"
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
