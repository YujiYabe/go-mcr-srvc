package camera

import (
	"net/http"

	"github.com/labstack/echo"

	"backend/internal/2_adapter/controller"
	domain "backend/internal/4_domain"
)

// ---------------------------------------------------------
// 予約商品画像の受信
// Post はクライアントからの画像を受け取り、それに基づいてJANコードを処理します。
func Post(
	c echo.Context,
	Controller controller.ToController,
) error {
	// コンテキストから番号を取得
	number, err := domain.PickOutNumber(c.Param("number"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// 画像のフォームデータを取得
	file, err := c.FormFile("image")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Image upload failed"})
	}

	ctx := c.Request().Context()

	err = Controller.DetectSaveJANCodes(ctx, number, file)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Image upload failed"})
	}

	Controller.DistributeOrder(ctx)

	return c.JSON(http.StatusOK, nil)
}
