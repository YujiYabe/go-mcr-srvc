package admin

import (
	"github.com/labstack/echo"

	client "backend/internal/1_framework/instacook/http/admin/client"
	kitchen "backend/internal/1_framework/instacook/http/admin/kitchen"
	language "backend/internal/1_framework/instacook/http/admin/language"
	product "backend/internal/1_framework/instacook/http/admin/product"
)

func AddRoute(parrent *echo.Group) {
	g := parrent.Group("/admin")

	client.AddRoute(g)
	kitchen.AddRoute(g)
	language.AddRoute(g)
	product.AddRoute(g)
}
