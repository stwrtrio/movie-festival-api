package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/stwrtrio/movie-festival-api/pkg/utils"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		claims, err := validateToken(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
		}

		c.Set("user", claims)

		return next(c)
	}
}

func validateToken(c echo.Context) (*utils.Claims, error) {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return nil, errors.New("authorization header is missing")
	}

	// Check if the token has "Bearer" prefix
	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		return nil, errors.New("invalid Authorization format")
	}

	// Extract the token
	token := tokenParts[1]

	// Validate the token
	claims, err := utils.ValidateJWTToken(token)
	if err != nil {
		return nil, errors.New("invalid or expired token")
	}

	return claims, nil
}
