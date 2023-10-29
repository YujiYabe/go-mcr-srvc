package admin_language

import (
	"net/http"

	"github.com/labstack/echo"

	"backend/internal/2_adapter/controller"
	domain "backend/internal/4_domain"
)

// ---------------------------------------------------------
// 管理画面 お客様側言語印刷
func Get(
	c echo.Context,
	Controller controller.ToController,
) error {
	number, err := domain.PickOutNumber(c.Param("number"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()

	isVaildlangCodeMap := Controller.GetIsVaildLangCodeMap(ctx)

	data := struct {
		Number int
		Name   string
	}{
		Number: number,
		Name:   isVaildlangCodeMap[number],
	}

	return c.Render(http.StatusOK, "adminPrintLanguage", data)

}
