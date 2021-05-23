package infrastructure

import (
	// "net/http"

	"app/interfaces/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() {
	// router := gin.Default()

	// userController := controllers.NewUserController(NewSqlHandler())

	// router.POST("/users", func(c *gin.Context) { userController.Create(c) })
	// router.GET("/users", func(c *gin.Context) { userController.Index(c) })
	// router.GET("/users/:id", func(c *gin.Context) { userController.Show(c) })

	// Router = router

	// Echo instance
	e := echo.New()

	userController := controllers.NewUserController(NewSqlHandler())

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// e.GET("/", hello)
	e.GET("/users", func(c echo.Context) error { return userController.Index(c) })
	e.GET("/users/:id", func(c echo.Context) error { return userController.Show(c) })
	e.POST("/users", func(c echo.Context) error { return userController.Create(c) })

	// e.PUT("/users/:id", func(c echo.Context) error { return userController.Save(c) })
	// e.DELETE("/users/:id", func(c echo.Context) error { return userController.Delete(c) })

	// Start server
	e.Logger.Fatal(e.Start(":1234"))

}
