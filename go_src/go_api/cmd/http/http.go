package http

import (
    "fmt"
    "log"
    "net/http"
    "os"

    "go_api/cmd/http/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
    // "github.com/labstack/echo/v4"
    // "github.com/labstack/echo/v4/middleware"
)

var e = createMux()

func Run() {
    http.Handle("/", e)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
        log.Printf("Defaulting to port %s", port)
    }

    log.Printf("Listening on port %s", port)
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func init() {
    ctrl := &controller.Controller{}

    e.GET("/:message", ctrl.HandleMessage)
}

func createMux() *echo.Echo {
    e := echo.New()

    e.Use(middleware.Recover())
    e.Use(middleware.Logger())
    e.Use(middleware.Gzip())

    return e
}
