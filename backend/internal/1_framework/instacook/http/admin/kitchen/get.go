package admin_kitchen

import (
	"github.com/labstack/echo"
	// domain "backend/internal/4_domain"
)

// ---------------------------------------------------------
// 管理画面 お客様側商品印刷
func Get(
	c echo.Context,
) error {
	// number, err := domain.PickOutNumber(c.Param("number"))
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, err)
	// }

	// productJson, err := json.Marshal(domain.GetProduct(number))
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// data := struct {
	// 	Product string
	// }{
	// 	Product: string(productJson),
	// }

	// return c.Render(http.StatusOK, "adminPrintKitchen", data)
	return nil
}
