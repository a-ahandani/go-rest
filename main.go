package main

import (
	"gorest/database"
	"gorest/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	database.ConnectDB()
	router.SetupRoutes(app)

	app.Listen(":3000")
}
