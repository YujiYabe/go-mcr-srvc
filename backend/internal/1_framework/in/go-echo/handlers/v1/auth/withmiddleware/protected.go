package withmiddleware

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"

	"backend/internal/2_adapter/controller"
)

func protected(
	echoContext echo.Context,
	_ controller.ToController,
) (
	err error,
) {
	claims, ok := echoContext.Get("user").(jwt.MapClaims)
	if !ok {
		err = echoContext.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid token claims"})
		return
	}

	username, ok := claims["sub"].(string)
	if !ok {
		err = echoContext.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid token subject"})
		return
	}

	err = echoContext.String(
		http.StatusOK,
		fmt.Sprintf(
			"Welcome to the protected endpoint, %s   !",
			username,
		),
	)
	return

}
