package goEcho

import (
	"net/http"

	"github.com/labstack/echo/v4"

	v1ToPubsub "backend/internal/1_framework/in/go-echo/handlers/v1/topubsub"
	v1users "backend/internal/1_framework/in/go-echo/handlers/v1/users"
	"backend/internal/1_framework/in/go-echo/openapi"
	"backend/internal/2_adapter/controller"
)

// ServerInterfaceImpl は生成された ServerInterface を実装する構造体
type ServerInterfaceImpl struct {
	Controller controller.ToController
}

// V1UsersGet は /v1/users GET エンドポイントの実装
func (receiver *ServerInterfaceImpl) V1UsersGet(
	echoContext echo.Context,
	getUsersParams openapi.V1UsersGetParams,
) error {
	return v1users.Get(
		echoContext,
		receiver.Controller,
		getUsersParams,
	)
}

// V1UsersPost は /v1/users POST エンドポイントの実装
func (receiver *ServerInterfaceImpl) V1UsersPost(ctx echo.Context) error {
	var user openapi.User
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			map[string]string{"error": "Invalid request"},
		)
	}
	user.Id = 3 // 仮に新しいユーザーIDを割り当て
	return ctx.JSON(http.StatusCreated, user)
}

// V1HealthGet は /v1/health GET エンドポイントの実装
func (receiver *ServerInterfaceImpl) V1HealthGet(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "OK")
}

// V1ToPubsubGet は /v1/to-pubsub GET エンドポイントの実装
func (receiver *ServerInterfaceImpl) V1ToPubsubGet(
	echoContext echo.Context,
) error {
	return v1ToPubsub.Get(
		echoContext,
		receiver.Controller,
	)
}
