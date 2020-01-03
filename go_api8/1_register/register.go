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

	userController := controller.NewUserController()
	hambargarController := controller.NewHambargarController()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339}  ${status}  ${method}\t${uri}\n",
	}))

	// Middleware
	e.Use(middleware.Recover())

	// get
	e.GET("/users", func(c echo.Context) error { return userController.Index(c) })
	e.GET("/user/:id", func(c echo.Context) error { return userController.Show(c) })

	// e.GET("/hambargar/:name", func(c echo.Context) error { return hambargarController.Show(c) })

	e.GET("/hambargar_request/:hambargarName", func(c echo.Context) error { return hambargarController.Request(c) })

	// // post
	// e.POST("/create", func(c echo.Context) error { return userController.IFCNCreate(c) })

	// // put
	// e.PUT("/user/:id", func(c echo.Context) error { return userController.IFCNSave(c) })

	// // delete
	// e.DELETE("/user/:id", func(c echo.Context) error { return userController.IFCNDelete(c) })

	// Start server
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))

}
