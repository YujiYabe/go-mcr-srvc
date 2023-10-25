package casher

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	"backend/internal/2_adapter/controller"
)

// ---------------------------------------------------------
// レジタブレットhtml
// reserving_listの画面とレジの画面は違うので決済中はcaching_listとする。
// タブレットの操作で注文をpreparing_listに追加する。
func Get(
	c echo.Context,
	Controller controller.ToController,
) error {

	ctx := c.Request().Context()

	productListJson, err := json.Marshal(
		Controller.GetProductList(ctx),
	)
	if err != nil {
		fmt.Println(err)
	}
	preparingList := Controller.GetPreparingList(ctx)

	
	preparingListJson, err := json.Marshal(
		preparingList,
	)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("== == == == == == == == == == ")
	fmt.Printf("%#v\n", preparingListJson)
	fmt.Println("== == == == == == == == == == ")

	data := struct {
		ProductList   string
		orderListJson string
	}{
		ProductList: string(productListJson),
	}

	return c.Render(http.StatusOK, "casherMonitor", data)
}
