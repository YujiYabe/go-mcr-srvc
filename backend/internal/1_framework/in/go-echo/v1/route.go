package v1

import (
	"github.com/labstack/echo"

	httpMiddleware "backend/internal/1_framework/middleware/http"
	"backend/internal/1_framework/in/go-echo/v1/auth"
	"backend/internal/1_framework/in/go-echo/v1/person"
	"backend/internal/2_adapter/controller"
)

func NewRoute(
	EchoEcho *echo.Echo,
	toController controller.ToController,
	parent *echo.Group,
) {
	group := parent.Group(
		"/v1",
		httpMiddleware.ContextMiddleware(),
	)

	person.NewRoute(
		EchoEcho,
		toController,
		group,
	)

	auth.NewRoute(
		EchoEcho,
		toController,
		group,
	)
}
