package queue

import (
	"net/http"

	"github.com/labstack/echo"

	"backend/internal/2_adapter/controller"
	domain "backend/internal/4_domain"
)

func Get(
	c echo.Context,
	Controller controller.ToController,
) error {
	ctx := c.Request().Context()
	// コンテキストから番号を取得

	number, err := domain.PickOutNumber(c.Param("number"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// 指定のキュー番号の予約情報を取得
	reserving := Controller.GetReserving(ctx, number)

	return c.JSON(http.StatusOK, reserving.JANCodeList)
}
