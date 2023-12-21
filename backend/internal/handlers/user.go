package handlers

import (
	"errors"
	"fmt"
	"gorest/config"
	"gorest/database"
	"gorest/internal/models"
	utils "gorest/internal/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserHandler struct {
}

func (u *UserHandler) GetUsersAPI(c *fiber.Ctx) error {
	db := database.DB
	var users []models.User

	// Preload the Roles association
	if err := db.Preload("Roles").Find(&users).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error retrieving users",
			"error":   err.Error(),
		})
	}

	if len(users) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No users found",
		})
	}

	// Remove password from response
	for i := range users {
		users[i].Password = ""
	}

	return c.JSON(fiber.Map{"status": "success", "message": "All users", "data": users})
}

func (u *UserHandler) CreateUserAPI(c *fiber.Ctx) error {
	input := new(utils.UserPayload)

	// Parse JSON request body
	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Cannot parse JSON",
		})
	}

	user, err := utils.CreateUser(input)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Could not create user",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created user", "data": user})
}

func (u *UserHandler) GetUserAPI(c *fiber.Ctx) error {
	var user *models.User

	userID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid user ID",
		})
	}

	user, err = utils.GetUserByID(userID)
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

func (u *UserHandler) UpdateUserAPI(c *fiber.Ctx) error {
	userID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid user ID format",
		})
	}

	request := new(utils.UserPayload)

	// Parse JSON request body
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Cannot parse JSON",
		})
	}

	updatedUser, err := utils.UpdateUser(userID, request)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Could not update user",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Updated user", "data": updatedUser})
}

func (u *UserHandler) LoginUserAPI(c *fiber.Ctx) error {
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
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				// "message": "User not found",
				"message": "Could not login",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error retrieving user",
		})
	}

	// Use constant time comparison for password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				// "message": "Invalid password",
				"message": "Could not login",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error comparing passwords",
		})
	}

	fmt.Println(user)
	// Generate a JWT token with user ID and roles
	roleNames, err := utils.GetUserRolesByID(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error retrieving user roles",
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.TokenClaims{
		UserID: user.ID,
		Roles:  roleNames,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
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
