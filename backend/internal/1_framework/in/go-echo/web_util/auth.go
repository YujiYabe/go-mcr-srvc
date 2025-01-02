package web_util

import (
	"crypto/rsa"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/MicahParks/keyfunc"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.NewHTTPError(
					http.StatusUnauthorized,
					"Authorization header is required",
				)
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == authHeader { // No "Bearer " prefix
				return echo.NewHTTPError(
					http.StatusUnauthorized,
					"Authorization header format must be Bearer {token}",
				)
			}

			claims := jwt.MapClaims{}
			_, err := jwt.ParseWithClaims(
				tokenString,
				claims,
				func(token *jwt.Token) (interface{}, error) {
					// Ensure signing method is HMAC
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf(
							"unexpected signing method: %v",
							token.Header["alg"],
						)
					}
					return []byte(os.Getenv("AUTH0_CLIENT_SECRET")), nil
				},
			)

			if err != nil {
				return echo.NewHTTPError(
					http.StatusUnauthorized,
					fmt.Sprintf(
						"Invalid token: %v", err),
				)
			}

			// Attach claims to context
			c.Set("user", claims)
			return next(c)
		}
	}
}

// JWTMiddleware validates the JWT token from the Authorization header
func JWTMiddlewareAuth0(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(
				http.StatusUnauthorized,
				"Authorization header is required",
			)
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader { // No "Bearer " prefix
			return echo.NewHTTPError(
				http.StatusUnauthorized,
				"Authorization header format must be Bearer {token}",
			)
		}

		// Auth0の公開鍵を取得
		jwksURL := fmt.Sprintf(
			"https://%s/.well-known/jwks.json", os.Getenv("AUTH0_DOMAIN"))

		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(
			tokenString,
			claims,
			func(token *jwt.Token) (interface{}, error) {
				return validateAndGetKey(token, jwksURL)
			},
		)
		if err != nil || !token.Valid {
			return echo.NewHTTPError(
				http.StatusUnauthorized,
				"Invalid token",
			)
		}

		c.Set("user", claims)

		return next(c)
	}
}

func validateAndGetKey(
	token *jwt.Token,
	jwksURL string,
) (
	interface{},
	error,
) {
	// Tokenの署名方式を確認
	if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
		return nil, fmt.Errorf(
			"unexpected signing method: %v", token.Header["alg"])
	}

	// Auth0から公開鍵を取得し、検証に使用
	cert, err := getRSAPublicKey(jwksURL, token)
	if err != nil {
		return nil, err
	}
	return cert, nil
}

// getRSAPublicKey fetches the RSA public key from Auth0 JWKS URL
func getRSAPublicKey(
	jwksURL string,
	token *jwt.Token,
) (
	*rsa.PublicKey,
	error,
) {
	// Fetch JWKS from the URL
	jwks, err := keyfunc.Get(
		jwksURL,
		keyfunc.Options{},
	)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to get JWKS: %w", err)
	}

	// Extract the RSA public key for the token
	key, err := jwks.Keyfunc(token)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to extract RSA key: %w", err)
	}

	rsaKey, ok := key.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf(
			"key is not an RSA public key")
	}

	return rsaKey, nil
}
