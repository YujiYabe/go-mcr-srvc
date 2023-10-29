package admin_client

import (
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

	data := struct {
		Number int
	}{
		Number: number,
	}

	return c.Render(http.StatusOK, "adminPrintClient", data)
}
