package person

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	httpParameter "backend/internal/1_framework/parameter/http"
	"backend/internal/2_adapter/controller"
	groupObject "backend/internal/4_domain/group_object"
	"backend/internal/logger"
)

func get(
	c echo.Context,
	toController controller.ToController,
) (
	err error,
) {
	ctx := c.Request().Context()
	requestContext := groupObject.GetRequestContext(ctx)
	timeoutMillSecond := requestContext.TimeOutMillSecond.GetValue()

	ctx, cancel := context.WithTimeout(
		ctx,
		time.Duration(timeoutMillSecond)*time.Millisecond,
	)
	defer cancel() // コンテキストのキャンセルを必ず呼び出す
	done := make(chan struct{})

	responseList := []httpParameter.V1Person{}
	var requestErr error

	// ゴルーチンで処理を実行
	go func() {
		person := httpParameter.V1Person{}
		if err = c.Bind(&person); err != nil {
			logger.Logging(ctx, err)
			err := c.JSON(http.StatusBadRequest, err)
			if err != nil { // httpレスポンス返却失敗
				logger.Logging(ctx, err)
			}
			return
		}

		responseList, requestErr = handlePersonRequest(
			ctx,
			person,
			toController,
		)
		if requestErr != nil {
			logger.Logging(ctx, err)
			err := c.JSON(http.StatusBadRequest, requestErr)
			if err != nil {
				logger.Logging(ctx, err)
			}
			return
		}

		close(done)
	}()

	// タイムアウトまたは処理完了を待つ
	select {
	case <-done:
		// 処理が完了した場合
		return c.JSON(
			http.StatusOK,
			responseList,
		)

	case <-ctx.Done():
		logger.Logging(ctx, ctx.Err())
		// タイムアウトした場合
		return c.JSON(
			http.StatusRequestTimeout,
			responseList,
		)
	}

}

func handlePersonRequest(
	ctx context.Context,
	person httpParameter.V1Person,
	toController controller.ToController,
) (
	responseList []httpParameter.V1Person,
	err error,
) {
	responseList = []httpParameter.V1Person{}

	reqPerson := groupObject.NewPerson(
		ctx,
		&groupObject.NewPersonArgs{
			ID:          person.ID,
			Name:        person.Name,
			MailAddress: person.MailAddress,
		},
	)

	if reqPerson.GetError() != nil {
		logger.Logging(ctx, reqPerson.GetError())
		return nil, reqPerson.GetError()
	}

	personList := toController.GetPersonListByCondition(
		ctx,
		*reqPerson,
	)

	if personList.GetError() != nil {
		logger.Logging(ctx, personList.GetError())
		return nil, personList.GetError()
	}

	for _, person := range personList.Content {
		id := person.ID.GetValue()
		name := person.Name.GetValue()
		mailAddress := person.MailAddress.GetValue()
		responseList = append(
			responseList,
			httpParameter.V1Person{
				ID:          &id,
				Name:        &name,
				MailAddress: &mailAddress,
			},
		)
	}

	return responseList, nil
}
