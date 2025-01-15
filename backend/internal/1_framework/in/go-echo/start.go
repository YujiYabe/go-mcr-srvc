package goEcho

import (
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	v1 "backend/internal/1_framework/in/go-echo/v1"
	"backend/internal/2_adapter/controller"
	"backend/pkg"
)

type (
	// GoEcho ...
	GoEcho struct {
		EchoEcho   *echo.Echo
		Controller controller.ToController
	}
)

// NewGoEcho ...
func NewGoEcho(
	controller controller.ToController,
) (
	goEcho *GoEcho,
) {
	goEcho = &GoEcho{
		EchoEcho:   NewEcho(),
		Controller: controller,
	}

	return goEcho
}

// NewEcho ...
func NewEcho() *echo.Echo {
	echoEcho := echo.New()
	echoEcho.HideBanner = true

	echoEcho.Use(
		middleware.LoggerWithConfig(
			middleware.LoggerConfig{
				Format:           "${time_custom}__${status}__${method}__${uri}\n",
				CustomTimeFormat: "15:04:05",
			},
		),
	)
	echoEcho.Use(middleware.Recover())
	// echoEcho.Use(middleware.RequestID())

	return echoEcho
}

// Start ...
func (receiver *GoEcho) Start() {
	group := receiver.EchoEcho.Group("")

	v1.NewRoute(
		receiver.EchoEcho,
		receiver.Controller,
		group,
	)

	isShowRoute := false
	if isShowRoute {
		routes := receiver.EchoEcho.Routes()
		for _, route := range routes {
			log.Printf("%#v\n", route)
		}
	}

	receiver.EchoEcho.Logger.Fatal(
		receiver.EchoEcho.Start(":" + pkg.GoEchoPort),
	)
}
