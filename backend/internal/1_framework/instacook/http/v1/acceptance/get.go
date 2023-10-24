package acceptance

import (
	"backend/internal/2_adapter/controller"
	"net/http"

	"github.com/labstack/echo"
)

// ---------------------------------------------------------
// お客様受取りhtml
// - 準備中
// - 準備完了
// 以上の注文番号を表示する。
func Get(
	c echo.Context,
	Controller controller.ToController,
) error {
	data := struct{}{}
	return c.Render(http.StatusOK, "acceptance", data)
}
