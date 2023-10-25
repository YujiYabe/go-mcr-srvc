package app

import (
	"net/http"

	"github.com/labstack/echo"

	"backend/internal/2_adapter/controller"
)

func Get(
	c echo.Context,
	Controller controller.ToController,
) error {

	ctx := c.Request().Context()

	return c.JSON(
		http.StatusOK,
		Controller.GetSoldList(ctx),
	)
}
