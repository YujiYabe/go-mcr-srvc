package delivery

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	"backend/internal/2_adapter/controller"
)

// ---------------------------------------------------------
// 受け渡し口タブレットhtml
// − 準備中→準備完了
// − 準備完了→お渡し済み
// にステータスを変更する。
func Get(
	c echo.Context,
	Controller controller.ToController,
) error {

	ctx := c.Request().Context()

	soldListJson, err := json.Marshal(
		Controller.GetSoldList(ctx),
	)
	if err != nil {
		fmt.Println(err)
	}

	data := struct {
		Delivery string
	}{
		Delivery: string(soldListJson),
	}
	fmt.Println("== == == == == == == == == == ")
	fmt.Printf("%#v\n", data)
	fmt.Println("== == == == == == == == == == ")

	return c.Render(http.StatusOK, "deliveryMonitor", data)
	// return nil
}
