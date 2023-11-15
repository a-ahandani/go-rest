package userHandlers

import (
	"note-api-fiber/database"
	"note-api-fiber/internal/models"
	"note-api-fiber/internal/validators"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetUsers(c *fiber.Ctx) error {

	db := database.DB
	var users []models.User
	db.Find(&users)
	if len(users) == 0 {
		return c.Status(404).JSON(fiber.Map{
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
		return c.Status(400).JSON(fiber.Map{
			"message": "Cannot parse JSON",
		})
	}
	user.ID = uuid.New()
	// if doesn't have password, return error
	if user.Password == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Password cannot be empty",
		})
	}
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
		return c.Status(404).JSON(fiber.Map{
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
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "No user found with given ID",
		})
	}
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Cannot parse JSON",
		})
	}
	valid, errs := validators.Validate(user)
	if !valid {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": errs,
		})
	}
	db.Save(&user)
	return c.JSON(fiber.Map{"status": "success", "message": "User updated", "data": user})
}
