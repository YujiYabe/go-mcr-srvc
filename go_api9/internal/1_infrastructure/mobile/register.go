package mobile

import (
	"app/internal/2_adapter/controller"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type (

	// Mobile ...
	Mobile struct {
		EchoEcho   *echo.Echo
		Controller *controller.Controller
	}
)

// NewMobile ...
func NewMobile(ctrl *controller.Controller) *Mobile {
	mb := &Mobile{}
	mb.EchoEcho = NewEcho()
	mb.Controller = ctrl
	// mb.Agents = make(map[string]*Agent)

	return mb
}

// NewEcho ...
func NewEcho() *echo.Echo {
	e := echo.New()
	e.HideBanner = true

	// e.Renderer = &Template{
	// 	templates: template.Must(template.ParseGlob(shared.IndexFilePath)),
	// }

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339}__${status}__${method}__${uri}\n",
	}))
	e.Use(middleware.Recover())

	// e.Static("/", "public")

	return e
}

// Run ...
// func Run() {
// 	e := echo.New()
// 	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
// 		Format: "${time_rfc3339}  ${status}  ${method}\t${uri}\n",
// 	}))
// 	// Middleware
// 	e.Use(middleware.Recover())
// 	// get
// 	e.GET("/", func(c echo.Context) error { return Index(c) })
// 	// Start server
// 	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
// }

// Start ...
// func (mb *Mobile) Start(address, port string) {
func (mb *Mobile) Start() {
	// log.Println("--------------------------- ")
	// log.Println("http://" + address + ":" + port)
	// log.Println("--------------------------- ")

	// go mb.SendToAgents()

	mb.EchoEcho.GET("/", mb.Index)

	// mb.EchoEcho.Static("/public", shared.PublicPath)
	// mb.EchoEcho.GET("/", mb.Index(address, port))
	// mb.EchoEcho.GET("/ws", mb.WebSocket)
	// mb.EchoEcho.POST("/file_upload", mb.FileUpload)

	mb.EchoEcho.Logger.Fatal(mb.EchoEcho.Start(":1234"))
}

// Index ...
func (mb *Mobile) Index(c echo.Context) (err error) {

	c.JSON(200, "users")

	return
}

// func (mb *Mobile) Index(address, port string) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		// serverInfo := &ServerInfo{
// 		// 	Address: address,
// 		// 	Port:    port,
// 		// }
// 		// data := struct{ *ServerInfo }{ServerInfo: serverInfo}

// 		// return c.Render(http.StatusOK, "index", data)
// 	}
// }
