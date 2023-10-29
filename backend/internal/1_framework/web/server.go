package web

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	v1 "backend/internal/1_framework/web/v1"
	"backend/internal/2_adapter/controller"
)

var instaCookPort string

func init() {
	currentPath, _ := os.Getwd()

	err := godotenv.Load(filepath.Join(currentPath, ".env"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	instaCookPort = os.Getenv("INSTA_COOK_PORT")
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
func (receiver *InstaCook) Start(isShowRoute bool) {
	group := receiver.EchoEcho.Group("")

	v1.NewRoute(
		receiver.EchoEcho,
		receiver.Controller,
		group,
	)

	if isShowRoute {
		routes := receiver.EchoEcho.Routes()
		for _, route := range routes {
			fmt.Printf("%#v\n", route)
		}
	}

	receiver.EchoEcho.Logger.Fatal(
		receiver.EchoEcho.StartTLS(":"+instaCookPort, "openssl/server.crt", "openssl/server.key"),
	)

}
