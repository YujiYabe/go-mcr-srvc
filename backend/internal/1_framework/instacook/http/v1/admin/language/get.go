package admin_language

import (
	"github.com/labstack/echo"
	// domain "backend/internal/4_domain"
)

// ---------------------------------------------------------
// 管理画面 お客様側言語印刷
func Get(
	c echo.Context,
) error {
	// number, err := domain.PickOutNumber(c.Param("number"))
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, err)
	// }

	// isVaildlangCodeMap := domain.GetIsVaildlangCodeMap()
	// data := struct {
	// 	Number int
	// 	Name   string
	// }{
	// 	Number: number,
	// 	Name:   isVaildlangCodeMap[number],
	// }

	// return c.Render(http.StatusOK, "adminPrintLanguage", data)
	return nil
}
