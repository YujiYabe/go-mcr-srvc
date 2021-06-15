package mobile

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"backend/internal/2_interface_adapter/controller"
	"backend/internal/4_enterprise_business_rule/entity"
	"backend/pkg"
)

var (
	orderType = "mobile"
	myErr     *pkg.MyErr
)

func init() {
	myErr = pkg.NewMyErr("framework_driver", "mobile")
}

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
func (mbl *Mobile) Start() {
	mbl.EchoEcho.POST("/", mbl.IndexPost)
	mbl.EchoEcho.Logger.Fatal(mbl.EchoEcho.Start(pkg.MobilePort))
}

// IndexPost ...
func (mbl *Mobile) IndexPost(c echo.Context) error {
	ctx := c.Request().Context()

	product := &entity.Product{}
	if err := c.Bind(product); err != nil {
		myErr.Logging(err)
		return err
	}

	order := &entity.Order{
		Product: *product,
	}

	mbl.Controller.Reserve(ctx, order, orderType)
	c.JSON(200, order.OrderInfo.OrderNumber)

	mbl.Controller.Order(&ctx, order)

	return nil
}
