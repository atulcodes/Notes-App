package auth

import (
    "github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
    app.Post("/login", loginHandler)
    app.Post("/register", registerHandler)
    app.Get("/logout", logoutHandler)
}

func loginHandler(c *fiber.Ctx) error {
    return c.SendString("Login endpoint")
}

func registerHandler(c *fiber.Ctx) error {
    return c.SendString("Register endpoint")
}

func logoutHandler(c *fiber.Ctx) error {
    return c.SendString("Logout endpoint")
}
