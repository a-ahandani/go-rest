package userRoutes

import (
	userHandlers "gorest/internal/handlers/user"
	"gorest/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(r fiber.Router) {
	user := r.Group("/users")
	user.Use(middlewares.AuthRequired)
	user.Use(middlewares.CasbinMiddleware)

	user.Post("/", userHandlers.CreateUserAPI)
	user.Get("/", userHandlers.GetUsersAPI)
	user.Get("/:id", userHandlers.GetUserAPI)
	user.Put("/:id", userHandlers.UpdateUserAPI)

}
