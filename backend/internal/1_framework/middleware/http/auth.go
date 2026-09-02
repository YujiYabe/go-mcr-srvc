package http_middleware

import (
	"crypto/rsa"
	"fmt"
	"net/http"
	"strings"

	"github.com/MicahParks/keyfunc"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type AuthConfig struct {
	Domain       string
	ClientSecret string
}

func JWTMiddleware(
	config AuthConfig,
) (
	middlewareFunc echo.MiddlewareFunc,
) {
	middlewareFunc = func(next echo.HandlerFunc) echo.HandlerFunc {
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
					return []byte(config.ClientSecret), nil
				},
			)

			if err != nil {
				return echo.NewHTTPError(
					http.StatusUnauthorized,
					fmt.Sprintf(
						"Invalid token: %v",
						err,
					),
				)
			}

			// Attach claims to context
			c.Set("user", claims)
			return next(c)
		}
	}
	return //nolint:nakedret // Use the project-wide named return convention.
}

// JWTMiddleware validates the JWT token from the Authorization header
func JWTMiddlewareAuth0(
	config AuthConfig,
) (
	middlewareFunc echo.MiddlewareFunc,
) {
	middlewareFunc = func(next echo.HandlerFunc) echo.HandlerFunc {
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
				"https://%s/.well-known/jwks.json",
				config.Domain,
			)

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
	return //nolint:nakedret // Use the project-wide named return convention.
}

func validateAndGetKey(
	token *jwt.Token,
	jwksURL string,
) (
	value interface{},
	err error,
) {
	// Tokenの署名方式を確認
	if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
		value, err = nil, fmt.Errorf(
			"unexpected signing method: %v",
			token.Header["alg"],
		)
		return
	}

	// Auth0から公開鍵を取得し、検証に使用
	cert, err := getRSAPublicKey(
		jwksURL,
		token,
	)
	if err != nil {
		value = nil
		return
	}

	value, err = cert, nil
	return
}

// getRSAPublicKey fetches the RSA public key from Auth0 JWKS URL
func getRSAPublicKey(
	jwksURL string,
	token *jwt.Token,
) (
	publicKey *rsa.PublicKey,
	err error,
) {
	// Fetch JWKS from the URL
	jwks, err := keyfunc.Get(
		jwksURL,
		keyfunc.Options{},
	)
	if err != nil {
		publicKey, err = nil, fmt.Errorf("failed to get JWKS: %w", err)
		return //nolint:nakedret // Use the project-wide named return convention.
	}

	// Extract the RSA public key for the token
	key, err := jwks.Keyfunc(token)
	if err != nil {
		publicKey, err = nil, fmt.Errorf("failed to extract RSA key: %w", err)
		return //nolint:nakedret // Use the project-wide named return convention.
	}

	rsaKey, ok := key.(*rsa.PublicKey)
	if !ok {
		publicKey, err = nil, fmt.Errorf("key is not an RSA public key")
		return //nolint:nakedret // Use the project-wide named return convention.
	}

	publicKey, err = rsaKey, nil
	return //nolint:nakedret // Use the project-wide named return convention.
}
