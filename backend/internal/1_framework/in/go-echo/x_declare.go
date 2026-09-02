package goEcho

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"backend/internal/1_framework/in/go-echo/openapi"
	httpMiddleware "backend/internal/1_framework/middleware/http"
	"backend/internal/2_adapter/controller"
)

type (
	// GoEcho ...
	GoEcho struct {
		Controller controller.ToController
		EchoEcho   *echo.Echo
		port       string
		authConfig httpMiddleware.AuthConfig
	}
)

// NewGoEcho ...
func NewGoEcho(
	controller controller.ToController,
	port string,
	auth0Domain string,
	auth0ClientSecret string,
) (
	goEcho *GoEcho,
) {
	goEcho = &GoEcho{
		Controller: controller,
		EchoEcho:   NewEcho(),
		port:       port,
		authConfig: httpMiddleware.AuthConfig{
			Domain:       auth0Domain,
			ClientSecret: auth0ClientSecret,
		},
	}

	return
}

// NewEcho ...
func NewEcho() (
	echoEcho *echo.Echo,
) {
	echoEcho = echo.New()
	echoEcho.HideBanner = true

	echoEcho.Use(
		middleware.RequestLoggerWithConfig(
			middleware.RequestLoggerConfig{
				LogMethod: true,
				LogStatus: true,
				LogURI:    true,
				LogValuesFunc: func(echoContext echo.Context, values middleware.RequestLoggerValues) error {
					echoContext.Logger().Printf(
						"%s__%d__%s__%s\n",
						values.StartTime.Format("15:04:05"),
						values.Status,
						values.Method,
						values.URI,
					)
					return nil
				},
			},
		),
	)
	echoEcho.Use(middleware.Recover())
	// echoEcho.Use(middleware.RequestID())
	echoEcho.Use(httpMiddleware.ContextMiddleware())

	return
}

// Start ...
func (receiver *GoEcho) Start() (
	err error,
) {
	server := &ServerInterfaceImpl{
		Controller: receiver.Controller,
	}
	openapi.RegisterHandlers(
		receiver.EchoEcho,
		server,
	)

	err = receiver.EchoEcho.Start(":" + receiver.port)
	return
}
