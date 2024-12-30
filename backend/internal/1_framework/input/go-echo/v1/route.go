package v1

import (
	"github.com/labstack/echo"

	"backend/internal/1_framework/input/go-echo/v1/person"
	"backend/internal/2_adapter/controller"
)

func NewRoute(
	EchoEcho *echo.Echo,
	toController controller.ToController,
	parent *echo.Group,
) {
	group := parent.Group("/v1")

	person.NewRoute(
		EchoEcho,
		toController,
		group,
	)
}
