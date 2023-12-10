package middlewares

import (
	"fmt"
	"gorest/database"

	"gorest/internal/models"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"

	"github.com/gofiber/fiber/v2"
)

var Enforcer *casbin.Enforcer

func InitCasbin() {
	// Initialize Casbin enforcer
	adapter, err := gormadapter.NewAdapterByDB(database.DB)
	if err != nil {
		fmt.Println("failed to initialize Casbin adapter", err)
	}
	modelPath := "casbin_model.conf"

	Enforcer, err = casbin.NewEnforcer(modelPath, adapter)
	if err != nil {
		panic("failed to initialize Casbin enforcer: " + err.Error())
	}

	// Ensure Enforcer is not nil before calling LoadPolicy
	if Enforcer == nil {
		panic("Casbin Enforcer is nil")
	}
	// Enforcer.AddPolicy("admin", "*", "*")

	Enforcer.AddPolicy("admin", "/api/users", "GET")
	Enforcer.AddPolicy("admin", "/api/users", "POST")
	Enforcer.AddPolicy("admin", "/api/users", "PUT")
	Enforcer.AddPolicy("admin", "/api/users", "PATCH")

	// Load policy from database
	err = Enforcer.LoadPolicy()
	if err != nil {
		panic("failed to load Casbin policy: " + err.Error())
	}

	fmt.Println("Casbin Enforcer Initialized")
}

func CasbinMiddleware(c *fiber.Ctx) error {
	// Check permissions using Casbin
	obj := c.Path()   // Use the path as the object in this example
	act := c.Method() // Use the HTTP method as the action in this example

	// Get the user from the context
	user, ok := c.Locals("user").(*models.User)
	if !ok || user == nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}
	// Check permission for each role
	for _, role := range user.Roles {
		subject := role.Name

		// Check permission
		ok, err := Enforcer.Enforce(subject, obj, act)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}

		if ok {
			// User has permission, continue to the next middleware or route handler
			return c.Next()
		}
	}

	// User does not have permission for any role, return a forbidden response
	return c.Status(fiber.StatusForbidden).SendString("Permission Denied")
}
