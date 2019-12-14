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

	var logFormat string
	logFormat += "time:${time_rfc3339} "
	logFormat += "method:${method} "
	logFormat += "status:${status} "
	logFormat += "uri:${uri} "
	// logFormat += "host:${remote_ip} "
	// logFormat += "forwardedfor:${header:x-forwarded-for} "
	// logFormat += "req:- "
	// logFormat += "size:${bytes_out} "
	// logFormat += "referer:${referer} "
	// logFormat += "ua:${user_agent} "
	// logFormat += "reqtime_ns:${latency} "
	// logFormat += "cache:- "
	// logFormat += "runtime:- "
	// logFormat += "apptime:- "
	// logFormat += "vhost:${host} "
	// logFormat += "reqtime_human:${latency_human} "
	// logFormat += "x-request-id:${id} "
	// logFormat += "host:${host} "
	logFormat += "n"

	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: logFormat,
	}))

	// get
	e.GET("/users", func(c echo.Context) error { return userController.IFCNIndex(c) })
	e.GET("/user/:id", func(c echo.Context) error { return userController.IFCNShow(c) })

	// post
	e.POST("/create", func(c echo.Context) error { return userController.IFCNCreate(c) })

	// put
	e.PUT("/user/:id", func(c echo.Context) error { return userController.IFCNSave(c) })

	// delete
	e.DELETE("/user/:id", func(c echo.Context) error { return userController.IFCNDelete(c) })

	// Start server
	e.Logger.Fatal(e.Start(":1234"))

}
