package queue

import (
	"net/http"

	"github.com/labstack/echo"

	"backend/internal/2_adapter/controller"
	domain "backend/internal/4_domain"
)

func Post(
	c echo.Context,
	Controller controller.ToController,
) error {
	ctx := c.Request().Context()

	// コンテキストから番号を取得
	number, err := domain.PickOutNumber(c.Param("number"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// 新しい予約の情報をバインド
	newReserving := domain.Reserving{}
	if err := c.Bind(&newReserving); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	newReserving.QueueNo = number

	Controller.SaveReserving(ctx, newReserving)

	// オーダーリストの更新を通知
	Controller.DistributeOrder(ctx)

	return c.JSON(http.StatusOK, nil)
}
