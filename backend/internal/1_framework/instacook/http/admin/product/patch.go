package admin_product

import (
	"github.com/labstack/echo"
	// "backend/internal/1_framework/out/db"
	// domain "backend/internal/4_domain"
)

// ---------------------------------------------------------
// 管理画面 商品更新
// Patch は管理者用のAPIで、製品情報を更新します。
func Patch(
	c echo.Context,
) error {
	// コンテキストからJANコードを取得
	// number, err := domain.PickOutNumber(c.Param("number"))
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, err)
	// }

	// // リクエストの内容を新製品オブジェクトにバインド
	// newProduct := &domain.Product{}
	// if err := c.Bind(newProduct); err != nil {
	// 	return c.JSON(http.StatusBadRequest, err)
	// }
	// newProduct.JANCode = number

	// // 製品リストの中で該当する製品を見つけて更新
	// domain.UpdateProduct(newProduct)
	// domain.FilterProductValid()

	// product := domain.GetProduct(number)
	// if product == nil {
	// 	return c.JSON(http.StatusOK, nil)
	// }

	// // DB更新
	// err = db.UpdateProduct(product)
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, err)
	// }

	// return c.JSON(http.StatusOK, nil)

	return nil
}
