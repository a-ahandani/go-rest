package router

import (
	"fmt"
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
	listMethods(&handler.UserHandler{})
	return c.JSON("-->")
}

func listMethods(obj interface{}) {
	t := reflect.TypeOf(obj)
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		fmt.Println(method.Name)
	}
}
