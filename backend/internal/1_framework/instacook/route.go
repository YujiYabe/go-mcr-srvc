package instacook

import (
	"github.com/labstack/echo"

	// "backend/internal/1_framework/in/http/acceptance"
	"backend/internal/1_framework/instacook/http/admin"
	// "backend/internal/1_framework/in/http/casher"
	// "backend/internal/1_framework/in/http/client"
	// "backend/internal/1_framework/in/http/delivery"
	// "backend/internal/1_framework/in/http/kitchen"
	// "backend/internal/1_framework/in/http/order"
)

func AddRoute(parrent *echo.Group) {
	g := parrent.Group("")

	// client.AddRoute(g)
	// kitchen.AddRoute(g)
	// casher.AddRoute(g)

	// delivery.AddRoute(g)
	admin.AddRoute(g)
	// order.AddRoute(g)
	// acceptance.AddRoute(g)
}
