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
// 管理画面 お客様側商品印刷
func Get(
	c echo.Context,
	Controller controller.ToController,
) error {
	number, err := domain.PickOutNumber(c.Param("number"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	fmt.Println("== == == == == == == == == == ")
	fmt.Printf("%#v\n", number)
	fmt.Println("== == == == == == == == == == ")
	// productJson, err := json.Marshal(domain.GetProduct(number))
	productJson, err := json.Marshal(
		Controller.GetProduct(
			ctx,
			number,
		),
	)
	fmt.Println("== == == == == == == == == == ")
	fmt.Printf("%#v\n", productJson)
	fmt.Println("== == == == == == == == == == ")

	if err != nil {
		fmt.Println(err)
	}

	data := struct {
		Product string
	}{
		Product: string(productJson),
	}
	fmt.Println("== == == == == == == == == == ")
	fmt.Printf("%#v\n", data)
	fmt.Println("== == == == == == == == == == ")

	return c.Render(http.StatusOK, "adminPrintKitchen", data)

	// return nil
}
