package main

import (
	"gorest/database"

	"gorest/internal/utils"
	"gorest/middlewares"
	"gorest/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	database.ConnectDB()

	utils.CreateBasicRoles()
	utils.CreateSuperAdminUser()

	middlewares.InitCasbin()
	router.SetupRoutes(app)
	app.Listen(":3000")
}
