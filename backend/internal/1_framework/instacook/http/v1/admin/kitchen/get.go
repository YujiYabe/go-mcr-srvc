package admin_kitchen

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	"backend/internal/2_adapter/controller"
	domain "backend/internal/4_domain"
)

// ---------------------------------------------------------
// 管理画面 厨房側商品印刷
func Get(
	c echo.Context,
	Controller controller.ToController,
) error {
	number, err := domain.PickOutNumber(c.Param("number"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()

	productJson, err := json.Marshal(
		Controller.GetProduct(
			ctx,
			number,
		),
	)

	if err != nil {
		fmt.Println(err)
	}

	data := struct {
		Product string
	}{
		Product: string(productJson),
	}

	return c.Render(http.StatusOK, "adminPrintKitchen", data)

	// return nil
}
