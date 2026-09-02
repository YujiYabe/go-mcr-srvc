package users

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	middlewareRequestContext "backend/internal/1_framework/middleware/request_context"
	groupObject "backend/internal/4_domain/group_object"

	"backend/internal/1_framework/in/go-echo/openapi"
	"backend/internal/2_adapter/controller"
	"backend/internal/logger"
)

func Get(
	echoContext echo.Context,
	toController controller.ToController,
	getUsersParams openapi.V1UsersGetParams,
) (
	err error,
) {
	err = nil //nolint:wastedassign // Explicitly initialize the named return value.
	ctx := echoContext.Request().Context()
	requestContext := middlewareRequestContext.GetRequestContext(ctx)

	timeoutMillSecond := requestContext.TimeOutMillSecond().GetValue()

	ctxWithTimeout, cancel := context.WithTimeout(
		ctx,
		time.Duration(timeoutMillSecond)*time.Millisecond,
	)
	defer cancel() // コンテキストのキャンセルを必ず呼び出す

	responseList, requestErr := handleUsersRequest(
		ctxWithTimeout,
		getUsersParams,
		toController,
	)
	if ctxWithTimeout.Err() != nil {
		logger.Logging(ctxWithTimeout, ctxWithTimeout.Err())
		err = echoContext.JSON(
			http.StatusRequestTimeout,
			[]openapi.User{},
		)
		return //nolint:nakedret // Use the project-wide named return convention.
	}
	if requestErr != nil {
		logger.Logging(ctxWithTimeout, requestErr)
		err = echoContext.JSON(http.StatusBadRequest, requestErr)
		return //nolint:nakedret // Use the project-wide named return convention.
	}

	err = echoContext.JSON(
		http.StatusOK,
		responseList,
	)
	return //nolint:nakedret // Use the project-wide named return convention.
}

func handleUsersRequest(
	ctx context.Context,
	getUsersParams openapi.V1UsersGetParams,
	toController controller.ToController,
) (
	responseList []openapi.User,
	err error,
) {
	responseList = []openapi.User{}

	reqUser, err := groupObject.NewUser(
		&groupObject.NewUserArgs{
			Name:  getUsersParams.Name,
			Email: getUsersParams.Email,
		},
	)
	if err != nil {
		logger.Logging(ctx, err)
		responseList = nil
		return //nolint:nakedret // Use the project-wide named return convention.
	}

	userList, err := toController.GetUserListByCondition(
		ctx,
		*reqUser,
	)
	if err != nil {
		logger.Logging(ctx, err)
		return //nolint:nakedret // Use the project-wide named return convention.
	}

	for _, user := range userList.Content() {
		userID := user.ID().GetValue()
		name := user.Name().GetValue()
		email := user.Email().GetValue()
		responseList = append(
			responseList,
			openapi.User{
				Id:    userID,
				Name:  name,
				Email: email,
			},
		)
	}

	return //nolint:nakedret // Use the project-wide named return convention.
}
