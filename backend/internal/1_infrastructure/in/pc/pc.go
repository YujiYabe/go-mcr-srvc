package pc

import (
	"github.com/gin-gonic/gin"

	"backend/internal/2_adapter/controller"
	"backend/internal/4_domain/domain"
	"backend/pkg"
)

var (
	orderType = "pc"
	myErr     *pkg.MyErr
)

func init() {
	myErr = pkg.NewMyErr("infrastructure", "pc")
}

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

	pc.GinEngine.Run(":2345")
}

// IndexPost ...
func (pc *PC) IndexPost(c *gin.Context) {
	ctx := c.Request.Context()

	order := &domain.Order{}
	product := &domain.Product{}
	if err := c.Bind(product); err != nil {
		myErr.Logging(err)
		return
	}

	order.Product = *product

	pc.Controller.Reserve(ctx, order, orderType)

	go pc.Controller.Order(ctx, order)

	c.JSON(200, order.OrderInfo.OrderNumber)
	return
}
