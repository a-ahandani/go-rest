package middlewares

import (
	"gorest/config"
	"gorest/internal/models"
	utils "gorest/internal/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func validateToken(tokenString string) (*models.TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config("SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(*models.TokenClaims)

	if !ok {
		return nil, err
	}

	return claims, nil
}

func AuthRequired(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")

	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   "Unauthorized",
			"message": "No authorization token provided",
		})
	}

	parts := strings.Split(tokenString, " ")
	if len(parts) != 2 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   "Malformed Token",
			"message": "Invalid token format. It should be in the form 'Bearer <token>'",
		})
	}

	tokenString = parts[1]

	claims, err := validateToken(tokenString)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   "Malformed Token",
			"message": "Unable to parse the provided token",
		})
	}

	// Fetch user from the database based on the user ID
	user, err := utils.GetUserByID(claims.UserID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Internal Server Error",
			"message": "Error fetching user from the database",
		})
	}

	c.Locals("user", user)

	return c.Next()
}
