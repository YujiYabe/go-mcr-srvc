package pc

import (
	"app/internal/2_adapter/controller"
	"app/internal/4_domain/domain"
	"context"

	"github.com/gin-gonic/gin"
)

var orderType domain.OrderType = "pc"

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
	pc.GinEngine.POST("/", pc.IndexPost)
	pc.GinEngine.GET("/2", pc.Index2)

	pc.GinEngine.Run(":2345")
}

// IndexPost ...
func (pc *PC) IndexPost(c *gin.Context) {
	ctx := c.Request.Context()

	order := &domain.Order{}
	if err := c.Bind(order); err != nil {
		return
	}

	orderNumber, ctxValue := pc.Controller.Reserve(ctx, orderType)
	orderCtx := context.WithValue(ctx, orderNumber, ctxValue)
	go pc.Controller.Order(orderCtx, *order)

	c.JSON(200, orderNumber)
	return
}

// Index2 ...
func (pc *PC) Index2(c *gin.Context) {
	ctx := c.Request.Context()
	res, _ := pc.Controller.Dummy(ctx)

	c.JSON(200, res)
	return
}
