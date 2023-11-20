package userHandlers

import (
	"errors"
	"fmt"
	"gorest/config"
	"gorest/database"
	"gorest/internal/models"
	"gorest/internal/validators"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func GetUsers(c *fiber.Ctx) error {

	db := database.DB
	var users []models.User
	db.Find(&users)
	if len(users) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No users found",
		})
	}
	// remove password from response
	for i := range users {
		users[i].Password = ""
	}
	return c.JSON(fiber.Map{"status": "success", "message": "All users", "data": users})

}

func CreateUser(c *fiber.Ctx) error {
	// Check if the Content-Type is application/json
	if !strings.Contains(c.Get("Content-Type"), "application/json") {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Content-Type. It must be application/json",
		})
	}

	db := database.DB

	// Parse JSON request body into a map to handle the Roles field
	var requestBody map[string]interface{}
	if err := c.BodyParser(&requestBody); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Cannot parse JSON",
		})
	}

	// Extract roles from the request body
	rawRoles, exists := requestBody["roles"]
	var roles []string
	if exists {
		roles, _ = rawRoles.([]string)
	}

	// Create a new user with the provided fields
	user := &models.User{
		ID:       uuid.New(),
		Name:     requestBody["name"].(string),
		Email:    requestBody["email"].(string),
		Password: requestBody["password"].(string),
		Verified: new(bool), // You might want to handle the Verified field differently

		// Roles field now accepts an array of role names
		Roles: roles,
	}

	// Validate user input
	valid, errs := validators.Validate(user)
	if !valid {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":  "error",
			"message": errs,
		})
	}

	// Hash the user's password before storing it in the database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to hash password",
			"error":   err.Error(),
		})
	}
	user.Password = string(hashedPassword)

	// Create the user in the database
	if err := db.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Could not create user",
			"error":   err.Error(),
		})
	}

	// Associate roles with the user
	var userRoles []models.Role
	db.Find(&userRoles, "name IN ?", roles)

	// Associate roles with the user
	db.Model(&user).Association("Roles").Append(userRoles)

	// Remove sensitive information from the response
	user.Password = ""

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "User created",
		"data":    user,
	})
}

func GetUser(c *fiber.Ctx) error {
	var user *models.User

	userID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid user ID",
		})
	}

	user, err = GetUserByID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  "error",
				"message": "No user found with given ID",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to retrieve user",
		})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "User found", "data": user})
}

func UpdateUser(c *fiber.Ctx) error {
	db := database.DB
	var user models.User
	db.Find(&user, "id = ?", c.Params("id"))
	if user.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "No user found with given ID",
		})
	}
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Cannot parse JSON",
		})
	}
	// Hash the user's password before updating it in the database
	if user.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Failed to hash password",
				"error":   err.Error(),
			})
		}
		user.Password = string(hashedPassword)
	}

	valid, errs := validators.Validate(user)
	if !valid {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":  "error",
			"message": errs,
		})
	}
	db.Save(&user)
	return c.JSON(fiber.Map{"status": "success", "message": "User updated", "data": user})
}

func LoginUser(c *fiber.Ctx) error {
	var loginData models.User
	if err := c.BodyParser(&loginData); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Cannot parse JSON",
		})
	}

	var user models.User
	db := database.DB
	result := db.Where("email = ?", loginData.Email).First(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	// Generate a JWT token with user ID and roles
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.TokenClaims{
		UserID: user.ID,
		Roles:  user.Roles, // Include user roles in the JWT claims
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
		},
	})
	tokenString, err := token.SignedString([]byte(config.Config("SECRET")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not login",
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Success",
		"token":   tokenString,
	})
}

// GetUserByID retrieves a user by ID from the database
func GetUserByID(userID uuid.UUID) (*models.User, error) {
	db := database.DB

	var user models.User
	if err := db.Where("id = ?", userID).Preload("Roles").First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
