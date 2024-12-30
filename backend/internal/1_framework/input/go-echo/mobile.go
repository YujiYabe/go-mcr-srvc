package mobile

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	v1 "backend/internal/1_framework/input/go-echo/v1"
	"backend/internal/2_adapter/controller"
	"backend/pkg"
)

type (
	// Mobile ...
	Mobile struct {
		EchoEcho   *echo.Echo
		Controller controller.ToController
	}
)

// NewMobile ...
func NewMobile(
	controller controller.ToController,
) (
	mobile *Mobile,
) {
	mobile = &Mobile{
		EchoEcho:   NewEcho(),
		Controller: controller,
	}

	return mobile
}

// NewEcho ...
func NewEcho() *echo.Echo {
	echoEcho := echo.New()
	echoEcho.HideBanner = true

	echoEcho.Use(
		middleware.LoggerWithConfig(
			middleware.LoggerConfig{
				Format:           "${time_custom}__${status}__${method}__${uri}\n",
				CustomTimeFormat: "06/01/02-15:04:05",
			},
		),
	)
	echoEcho.Use(middleware.Recover())
	echoEcho.Use(middleware.RequestID())

	return echoEcho
}

// Start ...
func (receiver *Mobile) Start() {
	group := receiver.EchoEcho.Group("")

	v1.NewRoute(
		receiver.EchoEcho,
		receiver.Controller,
		group,
	)

	receiver.EchoEcho.Logger.Fatal(receiver.EchoEcho.Start(":" + pkg.MobilePort))
}
