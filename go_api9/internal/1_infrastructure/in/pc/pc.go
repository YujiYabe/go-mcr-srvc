package pc

import (
	"app/internal/2_adapter/controller"

	"github.com/gin-gonic/gin"
)

type (
	// PC ...
	PC struct {
		GinEngine  *gin.Engine
		Controller *controller.Controller
	}
)

// NewPC ...
func NewPC(ctrl *controller.Controller) *PC {
	pc := &PC{}

	pc.GinEngine = NewGin()
	pc.Controller = ctrl

	return pc
}

// NewGin ...
func NewGin() *gin.Engine {
	r := gin.Default()

	return r
}

// Start ...
func (pc *PC) Start() {
	// mb.EchoEcho.POST("/", mb.IndexPost)
	// mb.EchoEcho.GET("/2", mb.Index2)
	// mb.EchoEcho.Logger.Fatal(mb.EchoEcho.Start(":1234"))

	pc.GinEngine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	pc.GinEngine.Run(":2345")
}

// Index2 ...
func (pc *PC) Index2(c *gin.Context) error {
	c.JSON(200, gin.H{
		"message": "pong",
	})

	return nil

}
