package notes

import (
    "github.com/gofiber/fiber/v2"
    "notes-app/internal/middleware"
)

func RegisterRoutes(app *fiber.App) {
    notes := app.Group("/notes", middleware.JWTProtected())
    notes.Get("/", listNotes)
    notes.Post("/", createNote)
    notes.Get("/:id", getNote)
    notes.Delete("/:id", deleteNote)
}

func listNotes(c *fiber.Ctx) error {
    return c.SendString("List notes")
}

func createNote(c *fiber.Ctx) error {
    return c.SendString("Create note")
}

func getNote(c *fiber.Ctx) error {
    return c.SendString("Get single note")
}

func deleteNote(c *fiber.Ctx) error {
    return c.SendString("Delete note")
}
