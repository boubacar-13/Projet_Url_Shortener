package middleware

import (
    "github.com/AkhilSharma90/Redis-GO-URL-Shortener/utils"
    "github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
    tokenString := c.Get("Authorization")
    if tokenString == "" {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
    }

    userID, err := utils.ValidateJWT(tokenString)
    if err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
    }

    c.Locals("userID", userID)
    return c.Next()
}
