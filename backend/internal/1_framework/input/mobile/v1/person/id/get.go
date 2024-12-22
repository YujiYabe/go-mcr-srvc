package person

import (
	"net/http"

	"github.com/labstack/echo"

	"backend/internal/2_adapter/controller"
)

func get(
	c echo.Context,
	toController controller.ToController,
) (
	err error,
) {
	// id := c.Param("id")

	// ctx := pkg.GetNewContext(
	// 	c.Request().Context(),
	// 	c.Response().Header().Get(echo.HeaderXRequestID),
	// )

	// person := toController.GetPersonByID(
	// 	ctx,
	// 	id,
	// )
	// pkg.Logging(ctx, person)

	return c.JSON(
		http.StatusOK,
		nil,
	)
}
