package auth

import (
	"github.com/labstack/echo/v4"

	"backend/internal/1_framework/in/go-echo/v1/auth/withmiddleware"
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
		"/auth",
	)

	withmiddleware.NewRoute(
		EchoEcho,
		toController,
		group,
		authConfig,
	)

}
