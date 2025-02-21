package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// Secret Key for JWT
var jwtSecret = []byte("abcd@1234")

// JWTMiddleware validates JWT tokens
func JWTMiddleware(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization") // Get token from headers

	if tokenString == "" {
		return c.Status(401).JSON(fiber.Map{"message": "Missing token"})
	}

	// Parse the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return c.Status(401).JSON(fiber.Map{"message": "Invalid token"})
	}

	// Extract claims (payload) from the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(401).JSON(fiber.Map{"message": "Invalid token format"})
	}

	// Store user email in context
	c.Locals("user_email", claims["email"])

	// Proceed to the next handler
	return c.Next()
}
