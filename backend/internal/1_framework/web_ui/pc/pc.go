package pc

import (
	"github.com/gin-gonic/gin"

	"backend/internal/2_adapter/controller"
	"backend/pkg"
)

var (
	orderType = "pc"
	myErr     *pkg.MyErr
)

func init() {
	myErr = pkg.NewMyErr("framework_driver", "pc")
}

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
func (pc *PC) Start() {
	pc.GinEngine.POST("/", pc.IndexPost)

	pc.GinEngine.Run(":" + pkg.PCPort)
}

// IndexPost ...
func (pc *PC) IndexPost(c *gin.Context) {
	// 標準コンテキストを取得
	c.JSON(200, 0) // オーダー番号返却
}
