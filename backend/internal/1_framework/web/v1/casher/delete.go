package casher

import (
	"net/http"

	"github.com/labstack/echo"

	"backend/internal/2_adapter/controller"
	domain "backend/internal/4_domain"
)

// ---------------------------------------------------------
// 準備中リストの削除
// Delete はキャッシャーからSoldListの情報を削除します。
func Delete(
	c echo.Context,
	Controller controller.ToController,
) error {
	ctx := c.Request().Context()

	number, err := domain.PickOutNumber(c.Param("number"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	Controller.DeleteSold(ctx, number)

	Controller.DistributeOrder(ctx)

	return c.JSON(http.StatusOK, nil)

}
