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

	number, err := domain.PickOutNumber(c.Param("number"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	orderList := Controller.GetOrderList(ctx)

	counter := 0

	// ステータスがPreparingのSoldListを取得
	item := orderList.FindPreparingSoldItem(number, &counter)
	if item != nil {
		return c.JSON(http.StatusOK, item)
	}

	// ReservingListの対応するアイテムを取得
	item = orderList.FindPreparingSoldItem(number, &counter)
	if item != nil {
		return c.JSON(http.StatusOK, item)
	}

	return nil
}
