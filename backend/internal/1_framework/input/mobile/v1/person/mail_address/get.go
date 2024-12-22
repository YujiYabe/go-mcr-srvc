package person

import (
	"log"
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
	mailAddress := c.Param("mail_address")
	log.Println("== == == == == == == == == == ")
	log.Printf("%#v\n", mailAddress)
	log.Println("== == == == == == == == == == ")

	// ctx := pkg.GetNewContext(
	// 	c.Request().Context(),
	// 	c.Response().Header().Get(echo.HeaderXRequestID),
	// )

	// placeList := toController.GetPlaceList(
	// 	ctx,
	// )
	// pkg.Logging(ctx, placeList)

	return c.JSON(
		http.StatusOK,
		nil,
	)
}
