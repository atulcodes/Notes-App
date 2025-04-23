package middleware

import (
    "github.com/gofiber/fiber/v2"
)

func JWTProtected() fiber.Handler {
    return func(c *fiber.Ctx) error {
        // Add JWT validation logic here
        return c.Next()
    }
}
