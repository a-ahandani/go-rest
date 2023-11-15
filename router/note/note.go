package noteRoutes

import (
	noteHandlers "note-api-fiber/internal/handlers/note"

	"github.com/gofiber/fiber/v2"
)

func SetupNoteRoutes(r fiber.Router) {
	note := r.Group("/notes")
	note.Post("/", noteHandlers.CreateNote)
	note.Get("/", noteHandlers.GetNotes)
	note.Get("/:id", noteHandlers.GetNote)
	note.Put("/:id", noteHandlers.UpdateNote)
	note.Delete("/:id", noteHandlers.DeleteNote)
}
