package goEcho

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	httpMiddleware "backend/internal/1_framework/middleware/http"

	"backend/internal/1_framework/in/go-echo/openapi"
	v1 "backend/internal/1_framework/in/go-echo/v1"
	v1ToPubsub "backend/internal/1_framework/in/go-echo/v1/topubsub"
	v1users "backend/internal/1_framework/in/go-echo/v1/users"
	"backend/internal/2_adapter/controller"
	"backend/internal/env"
)

type (
	// GoEcho ...
	GoEcho struct {
		Controller controller.ToController
		EchoEcho   *echo.Echo
	}

	// ServerInterfaceImpl は生成された ServerInterface を実装する構造体
	ServerInterfaceImpl struct {
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
		Controller: controller,
		EchoEcho:   NewEcho(),
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
	echoEcho.Use(httpMiddleware.ContextMiddleware())

	return echoEcho
}

// Start ...
func (receiver *GoEcho) Start() {
	group := receiver.EchoEcho.Group("")

	server := &ServerInterfaceImpl{
		Controller: receiver.Controller,
	}
	openapi.RegisterHandlers(
		receiver.EchoEcho,
		server,
	)

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
		receiver.EchoEcho.Start(":" + env.ServerConfig.GoEchoPort),
	)
}

// GetUsers は /users GET エンドポイントの実装
func (receiver *ServerInterfaceImpl) GetUsers(
	echoContext echo.Context,
	getUsersParams openapi.GetUsersParams,
) error {
	return v1users.GetUsers(
		echoContext,
		receiver.Controller,
		getUsersParams,
	)
}

// CreateUser は /users POST エンドポイントの実装
func (receiver *ServerInterfaceImpl) CreateUser(ctx echo.Context) error {
	var user openapi.User
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}
	user.Id = 3 // 仮に新しいユーザーIDを割り当て
	return ctx.JSON(http.StatusCreated, user)
}

// GetHealth は /health GET エンドポイントの実装
func (receiver *ServerInterfaceImpl) GetHealth(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "OK")
}

// GetUsers は /users GET エンドポイントの実装
func (receiver *ServerInterfaceImpl) ToPubsub(
	echoContext echo.Context,
) error {

	return v1ToPubsub.PublishTestTopic(
		echoContext,
		receiver.Controller,
	)
}
