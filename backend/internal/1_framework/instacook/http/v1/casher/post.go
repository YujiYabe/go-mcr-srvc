package casher

import (
	"net/http"

	"github.com/labstack/echo"

	"backend/internal/2_adapter/controller"
	domain "backend/internal/4_domain"
)

// ---------------------------------------------------------
// 準備中リストへの登録
// Post はキャッシャーからの新しい販売情報を処理します。
func Post(
	c echo.Context,
	Controller controller.ToController,
) error {
	// リクエストから注文情報を取得
	newSold := &domain.Sold{}
	if err := c.Bind(newSold); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	ctx := c.Request().Context()
	Controller.SaveSold(ctx, *newSold)

	Controller.DistributeOrder(ctx)

	return c.JSON(http.StatusOK, nil)
}
