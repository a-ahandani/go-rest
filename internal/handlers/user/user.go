package userHandlers

import (
	"gorest/config"
	"gorest/database"
	"gorest/internal/models"
	"gorest/internal/validators"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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
	db := database.DB
	user := new(models.User)
	err := c.BodyParser(user)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Cannot parse JSON",
		})
	}
	user.ID = uuid.New()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to hash password",
			"error":   err.Error(),
		})
	}
	user.Password = string(hashedPassword)

	err = db.Create(&user).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Could not create user",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created user", "data": user})
}

func GetUser(c *fiber.Ctx) error {
	db := database.DB
	var user models.User
	db.Find(&user, "id = ?", c.Params("id"))
	if user.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "No user found with given ID",
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

	// Generate a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.TokenClaims{
		UserID: user.ID,
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
