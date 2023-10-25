package delivery

import (
	"net/http"

	"github.com/labstack/echo"

	// ws "backend/internal/1_framework/in/http/websocket"
	"backend/internal/2_adapter/controller"
	domain "backend/internal/4_domain"
)

// ---------------------------------------------------------
// Patch はオーダーリスト内の特定の配送のステータスを更新します。
// 受け渡し口 情報更新
func Patch(
	c echo.Context,
	Controller controller.ToController,
) error {
	newSold := &domain.Sold{}
	if err := c.Bind(newSold); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()

	Controller.UpdateSoldStatus(ctx, *newSold)

	// ws.OrderListChan <- true
	return c.JSON(http.StatusOK, nil)
}
