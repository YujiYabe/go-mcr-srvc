package mobile

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"app/internal/2_adapter/controller"
)

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
	mb.EchoEcho.GET("/", mb.Index)
	mb.EchoEcho.Logger.Fatal(mb.EchoEcho.Start(":1234"))
}

// Index ...
func (mb *Mobile) Index(c echo.Context) (err error) {
	ctx := c.Request().Context()
	res, _ := mb.Controller.Dummy(ctx)
	fmt.Println(" ============================== ")
	fmt.Printf("%+v\n", res)
	fmt.Println(" ============================== ")

	c.JSON(200, "users")

	return
}
