package mobile

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"backend/internal/2_adapter/controller"
	"backend/internal/4_domain/domain"
	"backend/pkg"
)

var (
	orderType = "mobile"
	myErr     *pkg.MyErr
)

func init() {
	myErr = pkg.NewMyErr("mobile", "infrastructure")
}

type (
	// Mobile ...
	Mobile struct {
		EchoEcho   *echo.Echo
		Controller *controller.Controller
	}
)

// NewMobile ...
func NewMobile(ctrl *controller.Controller) *Mobile {
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
func (mb *Mobile) Start() {
	mb.EchoEcho.POST("/", mb.IndexPost)
	mb.EchoEcho.Logger.Fatal(mb.EchoEcho.Start(pkg.MobilePort))
}

// IndexPost ...
func (mb *Mobile) IndexPost(c echo.Context) error {
	ctx := c.Request().Context()

	product := &domain.Product{}
	if err := c.Bind(product); err != nil {
		myErr.Logging(err)
		return err
	}

	order := &domain.Order{
		Product: *product,
	}

	mb.Controller.Reserve(ctx, order, orderType)

	go mb.Controller.Order(ctx, order)

	c.JSON(200, order.OrderInfo.OrderNumber)

	return nil
}
