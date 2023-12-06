package main

import (
	"gorest/database"
	"gorest/internal/utils"
	"gorest/middlewares"
	"gorest/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	database.ConnectDB()

	utils.CreateBasicRoles()
	utils.CreateSuperAdminUser()

	middlewares.InitCasbin()
	router.SetupRoutes(app)
	app.Listen(":3000")
}
