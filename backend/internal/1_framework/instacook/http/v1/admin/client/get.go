package admin_client

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	domain "backend/internal/4_domain"
	// domain "backend/internal/4_domain"
)

// ---------------------------------------------------------
// 管理画面 お客様側商品印刷
func Get(
	c echo.Context,
) error {
	fmt.Println("== == == == == == == == == == ")
	number, err := domain.PickOutNumber(c.Param("number"))
	if err != nil {
		fmt.Println("== err == == == == == == == == == ")
		fmt.Printf("%#v\n", err)
		fmt.Println("== == == == == == == == == == ")

		return c.JSON(http.StatusBadRequest, err)
	}
	fmt.Println("== number == == == == == == == == == ")
	fmt.Printf("%#v\n", number)
	fmt.Println("== == == == == == == == == == ")

	// data := struct {
	// 	Number int
	// }{
	// 	Number: number,
	// }

	// return c.Render(http.StatusOK, "adminPrintClient", data)
	return nil
}
