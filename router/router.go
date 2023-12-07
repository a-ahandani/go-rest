package router

import (
	"reflect"

	handler "gorest/internal/handlers"
	noteRoutes "gorest/router/note"
	userRoutes "gorest/router/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	handler := handler.UserHandler{}
	api.Post("/auth", handler.LoginUserAPI)

	noteRoutes.SetupNoteRoutes(api)
	userRoutes.SetupUserRoutes(api)

	// Additional endpoint to list and categorize APIs
	api.Get("/list-apis", listAPIEndpoints)
}
func listAPIEndpoints(c *fiber.Ctx) error {
	// Pass the type 'handler.UserHandler' to listMethods
	a := listMethods(&handler.UserHandler{})
	return c.JSON(a)
}

func listMethods(obj interface{}) []string {
	t := reflect.TypeOf(obj)
	m := []string{}
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		m = append(m, method.Name)
	}
	return m
}
