package router

import (
	"gorest/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupResourceRoutes(r fiber.Router) {

	resource := r.Group("/resources")
	// resource.Use(middlewares.AuthRequired)
	// resource.Use(middlewares.CasbinMiddleware)

	resource.Get("/", handlers.ListAPIEndpoints)

}
