package router

import (
	handler "gorest/internal/handlers"
	"gorest/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(r fiber.Router) {
	handler := handler.UserHandler{}

	user := r.Group("/users")
	user.Use(middlewares.AuthRequired)
	user.Use(middlewares.CasbinMiddleware)

	user.Post("/", handler.CreateUserAPI)
	user.Get("/", handler.GetUsersAPI)
	user.Get("/:id", handler.GetUserAPI)
	user.Put("/:id", handler.UpdateUserAPI)

}
