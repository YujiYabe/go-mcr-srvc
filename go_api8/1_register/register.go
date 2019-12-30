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

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339}  ${status}  ${method}\t${uri}\n",
	}))

	// Middleware
	e.Use(middleware.Recover())

	// get
	e.GET("/users", func(c echo.Context) error { return userController.IFCNIndex(c) })
	e.GET("/user/:id", func(c echo.Context) error { return userController.IFCNShow(c) })

	// // post
	// e.POST("/create", func(c echo.Context) error { return userController.IFCNCreate(c) })

	// // put
	// e.PUT("/user/:id", func(c echo.Context) error { return userController.IFCNSave(c) })

	// // delete
	// e.DELETE("/user/:id", func(c echo.Context) error { return userController.IFCNDelete(c) })

	// Start server
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))

}
