package router

import (
	"gorest/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupNoteRoutes(r fiber.Router) {
	handler := handlers.NoteHandler{}

	note := r.Group("/notes")
	note.Post("/", handler.CreateNote)
	note.Get("/", handler.GetNotes)
	note.Get("/:id", handler.GetNote)
	note.Put("/:id", handler.UpdateNote)
	note.Delete("/:id", handler.DeleteNote)
}
