package person

import (
	"net/http"

	"github.com/labstack/echo"

	"backend/internal/2_adapter/controller"
	"backend/pkg"
)

func get(
	c echo.Context,
	toController controller.ToController,
) (
	err error,
) {
	ctx := pkg.GetNewContext(
		c.Request().Context(),
		c.Response().Header().Get(echo.HeaderXRequestID),
	)

	personList, err := toController.GetPersonList(
		ctx,
	)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			nil,
		)
	}

	pkg.Logging(ctx, personList)

	return c.JSON(
		http.StatusOK,
		nil,
	)
}
