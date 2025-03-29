package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// struct untuk representasi jwt claims
type JWTClaims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

// AuthConfig menyimpan konfigurasi untuk auth middleware
type AuthConfig struct {
	JWTSecret string
}

// NewAuthMiddleware membuat middleware autentikasi baru
func NewAuthMiddleware(config AuthConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token, err := extractTokenFromHeader(c)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		userID, err := validateToken(token, config.JWTSecret)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		c.Locals("userID", userID)
		return c.Next()
	}
}

// extractTokenFromHeader untuk mengekstrak token JWT dari header Authorization
func extractTokenFromHeader(c *fiber.Ctx) (string, error) {
	authHeader := c.Get("Authorization")

	if authHeader == "" {
		return "", fiber.NewError(fiber.StatusUnauthorized, "header Authorization not found")
	}

	if len(authHeader) < 8 || authHeader[:7] != "Bearer " {
		return "", fiber.NewError(fiber.StatusUnauthorized, "invalid authorization format")
	}

	return authHeader[7:], nil
}

// validateToken untuk memvalidasi token JWT dan mengembalikan id pengguna
func validateToken(tokenString string, jwtSecret string) (int, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return 0, fiber.NewError(fiber.StatusUnauthorized, "invalid token: "+err.Error())
	}

	if !token.Valid {
		return 0, fiber.NewError(fiber.StatusUnauthorized, "token is invalid")
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return 0, fiber.NewError(fiber.StatusUnauthorized, "invalid token claims")
	}

	return claims.UserID, nil
}
