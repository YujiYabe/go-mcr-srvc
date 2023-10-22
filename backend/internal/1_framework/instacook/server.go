package instacook

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"backend/internal/2_adapter/controller"
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
	// InstaCook ...
	InstaCook struct {
		EchoEcho   *echo.Echo
		Controller controller.ToController
	}
)

// NewInstaCook ...
func NewInstaCook(ctrl controller.ToController) *InstaCook {
	mb := &InstaCook{
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
func (mbl *InstaCook) Start() {
	mbl.EchoEcho.POST("/", mbl.IndexPost)

	// g := mbl.EchoEcho.Group("")
	// AddRoute(g)

	mbl.EchoEcho.Logger.Fatal(mbl.EchoEcho.Start(":" + pkg.InstaCookPort))
}

// IndexPost ...
func (mbl *InstaCook) IndexPost(c echo.Context) error {
	// 標準コンテキストを取得

	return nil
}
