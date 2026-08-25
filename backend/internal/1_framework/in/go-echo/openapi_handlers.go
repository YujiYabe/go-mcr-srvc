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
		return ctx.JSON(
			http.StatusBadRequest,
			map[string]string{"error": "Invalid request"},
		)
	}
	user.Id = 3 // 仮に新しいユーザーIDを割り当て
	return ctx.JSON(http.StatusCreated, user)
}

// GetHealth は /health GET エンドポイントの実装
func (receiver *ServerInterfaceImpl) GetHealth(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "OK")
}

// ToPubsub は /users GET エンドポイントの実装
func (receiver *ServerInterfaceImpl) ToPubsub(
	echoContext echo.Context,
) error {
	return v1ToPubsub.PublishTestTopic(
		echoContext,
		receiver.Controller,
	)
}
