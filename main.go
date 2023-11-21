package main

import (
	"gorest/database"
	"gorest/internal/initializers"
	"gorest/middlewares"
	"gorest/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	database.ConnectDB()
	middlewares.InitCasbin()
	router.SetupRoutes(app)

	if err := initializers.CreateAdmin(); err != nil {
		panic("Failed to initialize the application: " + err.Error())
	}

	app.Listen(":3000")
}
