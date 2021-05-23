package infrastructure

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"app/interfaces/controllers"
)

// Run ...
func Run() {
	// Echo instance
	e := echo.New()
	//
	userController := controllers.NewUserController(NewSQLHandler())

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/users", func(c echo.Context) error { return userController.IFCNIndex(c) })
	e.GET("/user/:id", func(c echo.Context) error { return userController.IFCNShow(c) })
	e.POST("/create", func(c echo.Context) error { return userController.IFCNCreate(c) })
	e.PUT("/user/:id", func(c echo.Context) error { return userController.IFCNSave(c) })
	e.DELETE("/user/:id", func(c echo.Context) error { return userController.IFCNDelete(c) })

	// Start server
	e.Logger.Fatal(e.Start(":1234"))

}
