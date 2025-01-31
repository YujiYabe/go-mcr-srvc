package topubsub

import (
	"github.com/labstack/echo/v4"

	"backend/internal/2_adapter/controller"
)

func PublishTestTopic(
	echoContext echo.Context,
	toController controller.ToController,
) error {
	ctx := echoContext.Request().Context()
	toController.PublishTestTopic(ctx)

	return nil

	// toController.ToPubsub()
	// ctx := echoContext.Request().Context()
	// requestContext := groupObject.GetRequestContext(ctx)

	// timeoutMillSecond := requestContext.TimeOutMillSecond.GetValue()

	// ctxWithTimeout, cancel := context.WithTimeout(
	// 	ctx,
	// 	time.Duration(timeoutMillSecond)*time.Millisecond,
	// )
	// defer cancel() // コンテキストのキャンセルを必ず呼び出す
	// done := make(chan struct{})

	// responseList := []httpParameter.V1Person{}
	// var requestErr error

	// // ゴルーチンで処理を実行
	// go func() {
	// 	person := httpParameter.V1Person{
	// 		Name:        getUsersParams.Name,
	// 		MailAddress: getUsersParams.MailAddress,
	// 	}

	// 	responseList, requestErr = handleUsersRequest(
	// 		ctxWithTimeout,
	// 		person,
	// 		toController,
	// 	)
	// 	if requestErr != nil {
	// 		logger.Logging(ctxWithTimeout, requestErr)
	// 		err := echoContext.JSON(http.StatusBadRequest, requestErr)
	// 		if err != nil {
	// 			logger.Logging(ctxWithTimeout, err)
	// 		}
	// 		return
	// 	}

	// 	close(done)
	// }()

	// // タイムアウトまたは処理完了を待つ
	// select {
	// case <-done:
	// 	// 処理が完了した場合
	// 	return echoContext.JSON(
	// 		http.StatusOK,
	// 		responseList,
	// 	)

	// case <-ctxWithTimeout.Done():
	// 	logger.Logging(ctxWithTimeout, ctxWithTimeout.Err())
	// 	// タイムアウトした場合
	// 	return echoContext.JSON(
	// 		http.StatusRequestTimeout,
	// 		responseList,
	// 	)
	// }

}
