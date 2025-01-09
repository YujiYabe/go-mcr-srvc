package person

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"

	httpParameter "backend/internal/1_framework/parameter/http"
	"backend/internal/2_adapter/controller"
	groupObject "backend/internal/4_domain/group_object"
	valueObject "backend/internal/4_domain/value_object"
	"backend/pkg"
)

func get(
	c echo.Context,
	toController controller.ToController,
) (
	err error,
) {
	ctx := c.Request().Context()
	traceID := valueObject.GetTraceID(ctx)
	log.Println("== == == == == == == == == == ")
	pkg.Logging(ctx, traceID)
	timeoutSecond := valueObject.GetTimeoutSecond(ctx)

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(timeoutSecond)*time.Millisecond,
	)
	defer cancel() // コンテキストのキャンセルを必ず呼び出す
	done := make(chan struct{})

	responseList := []httpParameter.V1Person{}
	// ゴルーチンで処理を実行

	go func() {
		person := httpParameter.V1Person{}
		if err := c.Bind(&person); err != nil {
			pkg.Logging(ctx, err)
			err := c.JSON(
				http.StatusBadRequest,
				err,
			)
			if err != nil {
				pkg.Logging(ctx, err)
			}
			return
		}

		reqPerson := groupObject.NewPerson(
			ctx,
			&groupObject.NewPersonArgs{
				ID:          person.ID,
				Name:        person.Name,
				MailAddress: person.MailAddress,
			},
		)

		if reqPerson.GetError() != nil {
			pkg.Logging(ctx, reqPerson.GetError())
			err := c.JSON(
				http.StatusBadRequest,
				reqPerson.GetError(),
			)
			if err != nil {
				pkg.Logging(ctx, err)
			}

			return
		}

		personList := toController.GetPersonByCondition(
			ctx,
			*reqPerson,
		)

		// ダミーのエラーをセット
		// personList.SetError(
		// 	ctx,
		// 	fmt.Errorf("failed to get person: %v", "dummy"),
		// )

		if personList.GetError() != nil {
			pkg.Logging(ctx, personList.GetError())
			err := c.JSON(
				http.StatusBadRequest,
				personList.GetError(),
			)
			if err != nil {
				pkg.Logging(ctx, err)
			}

			return
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

		log.Println("== == == == == == == == == == ")
		pkg.Logging(ctx, traceID)

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
		// タイムアウトした場合
		return c.JSON(
			http.StatusRequestTimeout,
			responseList,
		)
	}

}
