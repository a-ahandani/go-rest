package router

import (
	_ "gorest/docs"
	"gorest/internal/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	userHandler := handlers.UserHandler{}
	api.Post("/auth", userHandler.LoginUserAPI)
	api.Get("/docs/*", swagger.HandlerDefault) // default

	SetupNoteRoutes(api)
	SetupUserRoutes(api)
	SetupResourceRoutes(api)

}
