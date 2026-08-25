package withmiddleware

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"

	"backend/internal/2_adapter/controller"
)

func protected(
	c echo.Context,
	_ controller.ToController,
) (
	err error,
) {
	claims := c.Get("user").(jwt.MapClaims)
	username := claims["sub"].(string) // Example: "sub" from the token claims

	return c.String(
		http.StatusOK,
		fmt.Sprintf(
			"Welcome to the protected endpoint, %s   !",
			username,
		),
	)

}
