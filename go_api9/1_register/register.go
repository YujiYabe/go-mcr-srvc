package register

import (
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// go mod init github.com/YujiYabe/go-docker-template

// Run ...
func Run() {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339}  ${status}  ${method}\t${uri}\n",
	}))

	// Middleware
	e.Use(middleware.Recover())

	// get
	e.GET("/", func(c echo.Context) error { return Index(c) })

	// Start server
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))

}

// Index ...
func Index(c echo.Context) (err error) {
	c.JSON(200, "users")
	return
}
