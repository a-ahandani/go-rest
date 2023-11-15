package router

import (
	noteRoutes "note-api-fiber/router/note"
	userRoutes "note-api-fiber/router/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())
	noteRoutes.SetupNoteRoutes(api)
	userRoutes.SetupUserRoutes(api)
}
