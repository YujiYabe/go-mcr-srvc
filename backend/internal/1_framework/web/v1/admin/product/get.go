package admin_product

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	"backend/internal/2_adapter/controller"
)

// ---------------------------------------------------------
// 管理画面 html
func Get(
	c echo.Context,
	Controller controller.ToController,
) error {
	ctx := c.Request().Context()

	allProductListJson, err := json.Marshal(
		Controller.GetAllProductList(ctx),
	)
	if err != nil {
		fmt.Println(err)
	}

	allergyListJson, err := json.Marshal(
		Controller.GetAllergyDefault(ctx),
	)
	if err != nil {
		fmt.Println(err)
	}

	data := struct {
		AllProductList string
		AllergyList    string
	}{
		AllProductList: string(allProductListJson),
		AllergyList:    string(allergyListJson),
	}

	return c.Render(http.StatusOK, "adminProduct", data)
}
