package v1

import (
	"github.com/labstack/echo/v4"

	"backend/internal/1_framework/in/go-echo/handlers/v1/auth"
	httpMiddleware "backend/internal/1_framework/middleware/http"
	"backend/internal/2_adapter/controller"
)

func NewRoute(
	EchoEcho *echo.Echo,
	toController controller.ToController,
	parent *echo.Group,
	authConfig httpMiddleware.AuthConfig,
) {
	group := parent.Group(
		"/v1",
	)

	auth.NewRoute(
		EchoEcho,
		toController,
		group,
		authConfig,
	)
}
