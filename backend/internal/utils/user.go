package utils

import (
	"gorest/database"
	"gorest/internal/models"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserPayload struct {
	models.User
	Roles []string `json:"roles"`
}

// GetUserRolesByID fetches role names for a given user ID
func GetUserRolesByID(userID uuid.UUID) ([]string, error) {
	db := database.DB
	var roles []models.Role

	if err := db.Model(&models.User{ID: userID}).Association("Roles").Find(&roles); err != nil {
		return nil, err
	}

	var roleNames []string
	for _, role := range roles {
		roleNames = append(roleNames, role.Name)
	}

	return roleNames, nil
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
func CreateUser(input *UserPayload) (*models.User, error) {
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
func UpdateUser(userID uuid.UUID, request *UserPayload) (*models.User, error) {
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
