package admin_product

import (
	"net/http"

	"github.com/labstack/echo"

	"backend/internal/2_adapter/controller"
	domain "backend/internal/4_domain"
)

// ---------------------------------------------------------
// 管理画面 商品更新
// Patch は管理者用のAPIで、製品情報を更新します。
func Patch(
	c echo.Context,
	Controller controller.ToController,
) error {
	// コンテキストからJANコードを取得
	number, err := domain.PickOutNumber(c.Param("number"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// // リクエストの内容を新製品オブジェクトにバインド
	newProduct := &domain.Product{}
	if err := c.Bind(newProduct); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	newProduct.JANCode = number

	ctx := c.Request().Context()
	Controller.UpdateProduct(ctx, *newProduct)

	return c.JSON(http.StatusOK, nil)
}
