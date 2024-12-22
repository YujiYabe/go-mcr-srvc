package mobile

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"backend/internal/2_adapter/controller"
	domain "backend/internal/4_domain"
	"backend/pkg"
)

var (
	orderType = "mobile"
)

type (
	// Mobile ...
	Mobile struct {
		EchoEcho   *echo.Echo
		Controller controller.ToController
	}
)

// NewMobile ...
func NewMobile(ctrl controller.ToController) *Mobile {
	mb := &Mobile{
		EchoEcho:   NewEcho(),
		Controller: ctrl,
	}

	return mb
}

// NewEcho ...
func NewEcho() *echo.Echo {
	e := echo.New()
	e.HideBanner = true

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339}__${status}__${method}__${uri}\n",
	}))
	e.Use(middleware.Recover())

	return e
}

// Start ...
func (receiver *Mobile) Start() {
	receiver.EchoEcho.POST("/", receiver.IndexPost)
	receiver.EchoEcho.Logger.Fatal(receiver.EchoEcho.Start(":" + pkg.MobilePort))
}

// IndexPost ...
func (receiver *Mobile) IndexPost(c echo.Context) error {
	// 標準コンテキストを取得
	ctx := c.Request().Context()

	// web_uiのデータ型をControllerに持ち込まないようにproductに変換
	product := &domain.Product{}
	if err := c.Bind(product); err != nil {
		pkg.Logging(ctx, err)
		return err
	}
	order := &domain.Order{Product: *product}

	receiver.Controller.Reserve(ctx, order, orderType) // オーダー番号発行
	receiver.Controller.Order(&ctx, order)             // オーダー
	c.JSON(200, order.OrderInfo.OrderNumber)           // オーダー番号返却

	return nil
}
