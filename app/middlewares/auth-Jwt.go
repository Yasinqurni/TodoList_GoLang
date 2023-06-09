package middleware

import (
	"strings"
	"todoList_GoLang/pkg"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func AuthorizeJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.ErrUnauthorized
		}

		validate := pkg.NewJWTService()
		tokenString := strings.Split(authHeader, " ")[1]
		token, err := validate.ValidateToken(tokenString)
		if err != nil {
			return echo.ErrUnauthorized
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return echo.ErrUnauthorized
		}

		// Set user ID from JWT claims into context for further use
		c.Set("userId", claims["userId"].(uint))
		return next(c)
	}
}
