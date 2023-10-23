package instacook

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"backend/internal/2_adapter/controller"
	"backend/internal/1_framework/instacook/http/v1"

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
func (receiver *InstaCook) Start() {
	group := receiver.EchoEcho.Group("")

	v1.NewRoute(
		receiver.EchoEcho,
		receiver.Controller,
		group,
	)

	receiver.EchoEcho.GET("/", receiver.IndexPost)

	routes := receiver.EchoEcho.Routes()
	for _, route := range routes {
		fmt.Printf("%#v\n", route)
	}

	receiver.EchoEcho.Logger.Fatal(
		// receiver.EchoEcho.Start(":" + pkg.InstaCookPort),
		receiver.EchoEcho.StartTLS(":5678", "openssl/server.crt", "openssl/server.key"),
	)

}

// IndexPost ...
func (receiver *InstaCook) IndexPost(c echo.Context) error {
	// 標準コンテキストを取得

	return nil
}
