package topubsub

import (
	"github.com/labstack/echo/v4"

	"backend/internal/2_adapter/controller"
)

func Get(
	echoContext echo.Context,
	toController controller.ToController,
) (
	err error,
) {
	ctx := echoContext.Request().Context()

	err = toController.PublishTestTopic(ctx)

	return
}
