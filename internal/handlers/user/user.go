package userHandlers

import (
	"errors"
	"gorest/database"
	"gorest/internal/models"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserInputBody struct {
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Roles    []string `json:"roles"`
}

func GetUsersAPI(c *fiber.Ctx) error {

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

func CreateUser(input *UserInputBody) (*models.User, error) {
	db := database.DB

	// Create a new user instance
	user := models.User{
		ID:       uuid.New(),
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}

	// Hash the user's password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)

	// Create the user in the database
	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}

	// Assign roles to the user based on the request body
	roles := input.Roles
	if len(roles) > 0 {
		for _, roleName := range roles {
			roleName = strings.TrimSpace(roleName)
			var role models.Role
			if err := db.Where("name = ?", roleName).First(&role).Error; err == nil {
				db.Model(&user).Association("Roles").Append(&role)
			}
		}
	}

	return &user, nil
}

func CreateUserAPI(c *fiber.Ctx) error {
	input := new(UserInputBody)

	// Parse JSON request body
	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Cannot parse JSON",
		})
	}

	user, err := CreateUser(input)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Could not create user",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created user", "data": user})
}

func GetUserAPI(c *fiber.Ctx) error {
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

func UpdateUser(userID uuid.UUID, request *UserInputBody) (*models.User, error) {
	db := database.DB

	// Find the existing user
	var existingUser models.User
	if err := db.First(&existingUser, userID).Error; err != nil {
		return nil, err
	}

	// Update user fields
	existingUser.Name = request.Name
	existingUser.Email = request.Email

	// Update password if provided
	if request.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		existingUser.Password = string(hashedPassword)
	}

	// Update roles based on the request body
	roles := request.Roles
	if len(roles) > 0 {
		// Clear existing roles
		db.Model(&existingUser).Association("Roles").Clear()

		// Assign new roles
		for _, roleName := range roles {
			roleName = strings.TrimSpace(roleName)
			var role models.Role
			if err := db.Where("name = ?", roleName).First(&role).Error; err == nil {
				db.Model(&existingUser).Association("Roles").Append(&role)
			}
		}
	}

	// Save the updated user to the database
	if err := db.Save(&existingUser).Error; err != nil {
		return nil, err
	}

	return &existingUser, nil
}

func UpdateUserAPI(c *fiber.Ctx) error {
	userID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid user ID format",
		})
	}

	request := new(UserInputBody)

	// Parse JSON request body
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Cannot parse JSON",
		})
	}

	updatedUser, err := UpdateUser(userID, request)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Could not update user",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Updated user", "data": updatedUser})
}

// func LoginUser(c *fiber.Ctx) error {
// 	var loginData models.User
// 	if err := c.BodyParser(&loginData); err != nil {
// 		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
// 			"message": "Cannot parse JSON",
// 		})
// 	}

// 	var user models.User
// 	db := database.DB
// 	result := db.Where("email = ?", loginData.Email).First(&user)
// 	if result.Error != nil {
// 		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 			"message": "Unauthorized",
// 		})
// 	}

// 	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
// 		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 			"message": "Unauthorized",
// 		})
// 	}

// 	// Generate a JWT token with user ID and roles
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.TokenClaims{
// 		UserID: user.ID,
// 		Roles:  user.Roles, // Include user roles in the JWT claims
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
// 		},
// 	})
// 	tokenString, err := token.SignedString([]byte(config.Config("SECRET")))
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"message": "Could not login",
// 		})
// 	}

// 	return c.JSON(fiber.Map{
// 		"status":  "success",
// 		"message": "Success",
// 		"token":   tokenString,
// 	})
// }

// GetUserByID retrieves a user by ID from the database
func GetUserByID(userID uuid.UUID) (*models.User, error) {
	db := database.DB

	var user models.User
	if err := db.Where("id = ?", userID).Preload("Roles").First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
