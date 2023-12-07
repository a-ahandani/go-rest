package router

import (
	"gorest/internal/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	userHandler := handlers.UserHandler{}
	api.Post("/auth", userHandler.LoginUserAPI)

	SetupNoteRoutes(api)
	SetupUserRoutes(api)
	SetupResourceRoutes(api)

}
