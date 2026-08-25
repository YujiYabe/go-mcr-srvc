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

	return toController.PublishTestTopic(ctx)

}
