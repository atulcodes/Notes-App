package main

import (
    "github.com/gofiber/fiber/v2"
    "notes-app/internal/auth"
    "notes-app/internal/notes"
    "notes-app/internal/middleware"
)

func main() {
    app := fiber.New()

    app.Static("/static", "./static")
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Welcome to Notes App")
    })

    auth.RegisterRoutes(app)
    notes.RegisterRoutes(app)

    app.Listen(":3000")
}
