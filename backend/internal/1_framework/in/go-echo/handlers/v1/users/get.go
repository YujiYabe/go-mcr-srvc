package users

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	requestContextMiddleware "backend/internal/1_framework/middleware/request_context"
	httpParameter "backend/internal/1_framework/parameter/http"
	groupObject "backend/internal/4_domain/group_object"

	"backend/internal/1_framework/in/go-echo/openapi"
	"backend/internal/2_adapter/controller"
	"backend/internal/logger"
)

func Get(
	echoContext echo.Context,
	toController controller.ToController,
	getUsersParams openapi.V1UsersGetParams,
) error {
	ctx := echoContext.Request().Context()
	requestContext := requestContextMiddleware.GetRequestContext(ctx)

	timeoutMillSecond := requestContext.TimeOutMillSecond().GetValue()

	ctxWithTimeout, cancel := context.WithTimeout(
		ctx,
		time.Duration(timeoutMillSecond)*time.Millisecond,
	)
	defer cancel() // コンテキストのキャンセルを必ず呼び出す
	done := make(chan struct{})

	responseList := []httpParameter.V1User{}
	var requestErr error

	// 通常の GET request であれば handler 内で同期的に処理してもよいが、
	// マイクロサービス化すると controller/usecase の先で DB だけでなく
	// gRPC や Pub/Sub など別サービスへの I/O が発生する可能性がある。
	// その場合、下流サービスの遅延・停止・ネットワーク待ちによって
	// HTTP handler が長時間ブロックされると、呼び出し元へ timeout を返す制御や
	// request context のキャンセル伝播が見えづらくなる。
	//
	// ここでは実処理を goroutine 側に逃がし、handler 側は select で
	// 「処理完了」と「request context の timeout/cancel」のどちらが先に来るかを待つ。
	// これにより、下流処理が時間内に終われば通常レスポンスを返し、
	// request context の期限を超えた場合は HTTP request として明示的に
	// StatusRequestTimeout を返せるようにしている。
	go func() {
		user := httpParameter.V1User{
			Name:  getUsersParams.Name,
			Email: getUsersParams.Email,
		}

		responseList, requestErr = handleUsersRequest(
			ctxWithTimeout,
			user,
			toController,
		)
		if requestErr != nil {
			logger.Logging(ctxWithTimeout, requestErr)
			_ = echoContext.JSON(http.StatusBadRequest, requestErr)
			return
		}

		close(done)
	}()

	// タイムアウトまたは処理完了を待つ
	select {
	case <-done:
		// 処理が完了した場合
		return echoContext.JSON(
			http.StatusOK,
			responseList,
		)

	case <-ctxWithTimeout.Done():
		logger.Logging(ctxWithTimeout, ctxWithTimeout.Err())
		// タイムアウトした場合
		return echoContext.JSON(
			http.StatusRequestTimeout,
			responseList,
		)
	}

}

func handleUsersRequest(
	ctx context.Context,
	user httpParameter.V1User,
	toController controller.ToController,
) (
	responseList []httpParameter.V1User,
	err error,
) {
	responseList = []httpParameter.V1User{}

	reqUser, err := groupObject.NewUser(
		&groupObject.NewUserArgs{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	)
	if err != nil {
		logger.Logging(ctx, err)
		return nil, err
	}

	userList, err := toController.GetUserListByCondition(
		ctx,
		*reqUser,
	)
	if err != nil {
		logger.Logging(ctx, err)
		return nil, err
	}

	for _, user := range userList.Content() {
		id := user.ID().GetValue()
		name := user.Name().GetValue()
		email := user.Email().GetValue()
		responseList = append(
			responseList,
			httpParameter.V1User{
				ID:    &id,
				Name:  &name,
				Email: &email,
			},
		)
	}

	return responseList, nil
}
