package goEcho

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"backend/internal/1_framework/in/go-echo/openapi"
	v1 "backend/internal/1_framework/in/go-echo/v1"
	httpMiddleware "backend/internal/1_framework/middleware/http"
	"backend/internal/2_adapter/controller"
	"backend/internal/env"
	"backend/internal/logger"

	httpParameter "backend/internal/1_framework/parameter/http"
	groupObject "backend/internal/4_domain/group_object"
)

type (
	// GoEcho ...
	GoEcho struct {
		EchoEcho   *echo.Echo
		Controller controller.ToController
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

	ctx := echoContext.Request().Context()
	requestContext := groupObject.GetRequestContext(ctx)

	timeoutMillSecond := requestContext.TimeOutMillSecond.GetValue()

	ctxWithTimeout, cancel := context.WithTimeout(
		ctx,
		time.Duration(timeoutMillSecond)*time.Millisecond,
	)
	defer cancel() // コンテキストのキャンセルを必ず呼び出す
	done := make(chan struct{})

	responseList := []httpParameter.V1Person{}
	var requestErr error

	// ゴルーチンで処理を実行
	go func() {

		person := httpParameter.V1Person{
			Name:        getUsersParams.Name,
			MailAddress: getUsersParams.MailAddress,
		}

		responseList, requestErr = handlePersonRequest(
			ctxWithTimeout,
			person,
			receiver.Controller,
		)
		if requestErr != nil {
			logger.Logging(ctxWithTimeout, requestErr)
			err := echoContext.JSON(http.StatusBadRequest, requestErr)
			if err != nil {
				logger.Logging(ctxWithTimeout, err)
			}
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
