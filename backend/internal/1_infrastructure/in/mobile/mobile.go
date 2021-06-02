package mobile

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"app/internal/2_adapter/controller"
	"app/internal/4_domain/domain"
)

var orderType = "mobile"

type (
	// Mobile ...
	Mobile struct {
		EchoEcho   *echo.Echo
		Controller *controller.Controller
	}
)

// NewMobile ...
func NewMobile(ctrl *controller.Controller) *Mobile {
	mb := &Mobile{}
	mb.EchoEcho = NewEcho()
	mb.Controller = ctrl

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
func (mb *Mobile) Start() {
	mb.EchoEcho.POST("/", mb.IndexPost)
	mb.EchoEcho.GET("/2", mb.Index2)
	mb.EchoEcho.Logger.Fatal(mb.EchoEcho.Start(":1234"))
}

// IndexPost ...
func (mb *Mobile) IndexPost(c echo.Context) error {
	ctx := c.Request().Context()

	order := &domain.Order{}
	if err := c.Bind(order); err != nil {
		return err
	}

	orderNumber := mb.Controller.Reserve(ctx)

	go mb.Controller.Order(ctx, *order)

	c.JSON(200, orderNumber)
	return nil
}

// Index2 ...
func (mb *Mobile) Index2(c echo.Context) error {
	ctx := c.Request().Context()
	res, _ := mb.Controller.Dummy(ctx)

	c.JSON(200, res)
	return nil
}
