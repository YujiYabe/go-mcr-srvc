package sold

import (
	"net/http"

	"github.com/labstack/echo"

	"backend/internal/2_adapter/controller"
)

// ---------------------------------------------------------
// レジタブレットhtml
// reserving_listの画面とレジの画面は違うので決済中はcaching_listとする。
// タブレットの操作で注文をpreparing_listに追加する。
func Get(
	c echo.Context,
	Controller controller.ToController,
) error {

	ctx := c.Request().Context()

	return c.JSON(
		http.StatusOK,
		Controller.GetProductList(ctx),
	)
}
