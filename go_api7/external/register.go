package external

import (
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"app/interfaces/controllers"
)

// Register ...
func Register() {
	// Echo instance
	e := echo.New()

	drinkController := controllers.NewDrinkController(NewDrinkStocker())

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339}  ${status}  ${method}\t${uri}\n",
	}))

	// Middleware
	e.Use(middleware.Recover())

	// get
	e.GET("/drinks", func(c echo.Context) error { return drinkController.ShowAllDrinks(c) })
	// e.GET("/drinks", func(c echo.Context) error { return drinkController.IFCNIndex(c) })
	// e.GET("/drink/:name", func(c echo.Context) error { return drinkController.IFCNShow(c) })
	// e.GET("/drink/:name", func(c echo.Context) error { return drinkController.ShowDetailDrink(c) })

	// // post
	// e.POST("/create", func(c echo.Context) error { return drinkController.IFCNCreate(c) })

	// // put
	// e.PUT("/user/:id", func(c echo.Context) error { return drinkController.IFCNSave(c) })

	// // delete
	// e.DELETE("/user/:id", func(c echo.Context) error { return drinkController.IFCNDelete(c) })

	// Start server
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))

}
