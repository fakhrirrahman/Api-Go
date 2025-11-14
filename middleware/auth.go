package middleware

import (
	"goApi/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// JWTMiddlewareFiber - Fiber version
func JWTMiddlewareFiber(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": "Authorization header diperlukan",
			"data":    nil,
		})
	}

	// Extract token from "Bearer <token>"
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": "Format authorization tidak valid",
			"data":    nil,
		})
	}

	claims, err := utils.ValidateToken(tokenString)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": "Token tidak valid: " + err.Error(),
			"data":    nil,
		})
	}

	// Store claims in context
	c.Locals("user_id", claims.UserID)
	c.Locals("user_email", claims.Email)
	c.Locals("user_name", claims.Name)

	return c.Next()
}
