package userRoutes

import (
	userHandlers "note-api-fiber/internal/handlers/user"
	"note-api-fiber/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(r fiber.Router) {
	user := r.Group("/users")
	user.Use(middlewares.AuthRequired)

	user.Post("/", userHandlers.CreateUser)
	user.Get("/", userHandlers.GetUsers)
	user.Get("/:id", userHandlers.GetUser)
	user.Put("/:id", userHandlers.UpdateUser)

}
