package monitor

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	"backend/internal/2_adapter/controller"
	domain "backend/internal/4_domain"
)

func Get(
	c echo.Context,
	Controller controller.ToController,
) error {

	ctx := c.Request().Context()

	number, err := domain.PickOutNumber(c.Param("number"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	productListJson, err := json.Marshal(
		Controller.GetProductList(ctx),
	)
	if err != nil {
		fmt.Println(err)
	}

	data := struct {
		Number      int
		ProductList string
	}{
		Number:      number,
		ProductList: string(productListJson),
	}

	Controller.DistributeOrder(ctx)

	return c.Render(http.StatusOK, "kitchenMonitor", data)
}
