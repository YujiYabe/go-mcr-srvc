package admin_product

import (
	"github.com/labstack/echo"
	// domain "backend/internal/4_domain"
)

// ---------------------------------------------------------
// 管理画面 html
func Get(
	c echo.Context,
) error {

	// allProductListJson, err := json.Marshal(domain.GetAllProductList())
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// allergyListJson, err := json.Marshal(domain.GetAllergyListJa())
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// data := struct {
	// 	AllProductList string
	// 	AllergyList    string
	// }{
	// 	AllProductList: string(allProductListJson),
	// 	AllergyList:    string(allergyListJson),
	// }

	// return c.Render(http.StatusOK, "adminMonitor", data)

	return nil
}
