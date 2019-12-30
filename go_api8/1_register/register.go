package register

import (
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	controller "app/2_controller"
)

// Run ...
func Run() {
	// Echo instance
	e := echo.New()

	controllerUser := controller.NewControllerUser()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339}  ${status}  ${method}\t${uri}\n",
	}))

	// Middleware
	e.Use(middleware.Recover())

	// get
	e.GET("/users", func(c echo.Context) error { return controllerUser.Index(c) })
	e.GET("/user/:id", func(c echo.Context) error { return controllerUser.Show(c) })

	// // post
	// e.POST("/create", func(c echo.Context) error { return controllerUser.IFCNCreate(c) })

	// // put
	// e.PUT("/user/:id", func(c echo.Context) error { return controllerUser.IFCNSave(c) })

	// // delete
	// e.DELETE("/user/:id", func(c echo.Context) error { return controllerUser.IFCNDelete(c) })

	// Start server
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))

}
