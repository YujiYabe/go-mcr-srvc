package person

import (
	"github.com/labstack/echo"

	// "backend/internal/1_framework/web/v1/admin"
	// "backend/internal/1_framework/web/v1/alcohol"
	// "backend/internal/1_framework/web/v1/internet"
	// "backend/internal/1_framework/web/v1/monitor"
	// "backend/internal/1_framework/web/v1/place"
	// "backend/internal/1_framework/web/v1/smaregi"
	// "backend/internal/1_framework/web/v1/ws"

	"backend/internal/2_adapter/controller"
)

func NewRoute(
	EchoEcho *echo.Echo,
	toController controller.ToController,
	parent *echo.Group,
) {
	group := parent.Group("/mail_address")

	group.GET(
		"/:mail_address",
		func(c echo.Context) (err error) { return get(c, toController) },
	)

}
