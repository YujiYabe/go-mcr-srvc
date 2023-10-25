package instacook

import (
	"fmt"
	"html/template"
	"io"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	v1 "backend/internal/1_framework/instacook/http/v1"
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

type Template struct {
	templates *template.Template
}

func (receiver *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return receiver.templates.ExecuteTemplate(w, name, data)
}

// NewEcho ...
func NewEcho() *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	t := &Template{
		templates: template.Must(template.ParseGlob("web/*.html")),
	}
	e.Renderer = t
	e.Static("/js", "web/js")
	e.Static("/css", "web/css")
	e.Static("/image", "web/image")

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

	if false {
		routes := receiver.EchoEcho.Routes()
		for _, route := range routes {
			fmt.Printf("%#v\n", route)
		}
	}

	receiver.EchoEcho.Logger.Fatal(
		// receiver.EchoEcho.Start(":" + pkg.InstaCookPort),
		receiver.EchoEcho.StartTLS(":"+pkg.InstaCookPort, "openssl/server.crt", "openssl/server.key"),
	)

}

// IndexPost ...
func (receiver *InstaCook) IndexPost(c echo.Context) error {
	// 標準コンテキストを取得

	return nil
}
