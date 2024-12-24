package person

import (
	"net/http"

	"github.com/labstack/echo"

	"backend/internal/1_framework/input/mobile/http_parameter"
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

	dataList := []http_parameter.HTTPParameter{}
	for _, person := range personList {
		dataList = append(
			dataList,
			http_parameter.HTTPParameter{
				ID:          person.ID.Content.Value,
				Name:        person.Name.Content.Value,
				MailAddress: person.MailAddress.Content.Value,
			},
		)
	}

	return c.JSON(
		http.StatusOK,
		dataList,
	)
}
