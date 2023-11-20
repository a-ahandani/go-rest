package middlewares

import (
	"gorest/config"
	userHandlers "gorest/internal/handlers/user"
	"gorest/internal/models"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func AuthRequired(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")

	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	parts := strings.Split(tokenString, " ")
	if len(parts) != 2 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Malformed Token 1",
		})
	}

	tokenString = parts[1]

	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &models.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config("SECRET")), nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Malformed Token 2",
		})
	}

	if !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Token is not valid",
		})
	}

	claims, ok := token.Claims.(*models.TokenClaims)

	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Malformed Token 3",
		})
	}

	// Fetch user from the database based on the user ID
	user, err := userHandlers.GetUserByID(claims.UserID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error fetching user",
		})
	}

	// Set user and roles in the context locals
	c.Locals("user", user)
	c.Locals("roles", user.Roles) // Adjusted to use user.Roles instead of user.Roles.IDs

	return c.Next()
}
