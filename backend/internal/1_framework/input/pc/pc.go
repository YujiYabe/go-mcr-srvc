package pc

import (
	"github.com/gin-gonic/gin"

	"backend/internal/2_adapter/controller"
	domain "backend/internal/4_domain"
	"backend/pkg"
)

var (
	orderType = "pc"
)

type (
	// PC ...
	PC struct {
		GinEngine  *gin.Engine
		Controller controller.ToController
	}
)

// NewPC ...
func NewPC(ctrl controller.ToController) *PC {
	pc := &PC{
		GinEngine:  NewGin(),
		Controller: ctrl,
	}

	return pc
}

// NewGin ...
func NewGin() *gin.Engine {
	return gin.Default()
}

// Start ...
func (receiver *PC) Start() {
	receiver.GinEngine.POST("/", receiver.IndexPost)

	receiver.GinEngine.Run(":" + pkg.PCPort)
}

// IndexPost ...
func (receiver *PC) IndexPost(c *gin.Context) {
	// 標準コンテキストを取得
	ctx := c.Request.Context()

	// web_uiのデータ型をControllerに持ち込まないようにproductに変換
	product := &domain.Product{}
	if err := c.Bind(product); err != nil {
		pkg.Logging(ctx, err)
		return
	}
	order := &domain.Order{Product: *product}

	// receiver.Controller.Reserve(ctx, order, orderType) // オーダー番号発行
	// receiver.Controller.Order(&ctx, order)             // オーダー
	c.JSON(200, order.OrderInfo.OrderNumber) // オーダー番号返却
}
