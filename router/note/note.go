package noteRoutes

import (
	handlers "gorest/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupNoteRoutes(r fiber.Router) {
	note := r.Group("/notes")
	note.Post("/", handlers.CreateNote)
	note.Get("/", handlers.GetNotes)
	note.Get("/:id", handlers.GetNote)
	note.Put("/:id", handlers.UpdateNote)
	note.Delete("/:id", handlers.DeleteNote)
}
