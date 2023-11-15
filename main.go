package main

import (
	"note-api-fiber/database"
	"note-api-fiber/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	database.ConnectDB()
	router.SetupRoutes(app)

	app.Listen(":3000")
}
