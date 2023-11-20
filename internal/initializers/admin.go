// internal/initializer.go
package initializers

import (
	"fmt"
	"gorest/database"
	"gorest/internal/models"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Initialize creates a superadmin user and basic roles if they don't exist
func CreateAdmin() error {
	// Create basic roles if they don't exist
	if err := createBasicRoles(); err != nil {
		return err
	}

	// Create a superadmin user if it doesn't exist
	if err := createSuperAdminUser(); err != nil {
		return err
	}

	return nil
}

// createBasicRoles creates basic roles if they don't exist
func createBasicRoles() error {
	db := database.DB

	roles := []models.Role{
		{Name: "admin", Label: "Admin"},
		{Name: "user", Label: "User"},
	}

	for _, role := range roles {
		// Find or create the role by name
		result := db.Where("name = ?", role.Name).FirstOrCreate(&role)
		if result.Error != nil {
			fmt.Printf("Failed to create role '%s': %v\n", role.Name, result.Error)
			return result.Error
		}
	}

	fmt.Println("Basic roles created successfully")
	return nil
}

func createSuperAdminUser() error {
	db := database.DB

	superadmin := models.User{
		Name:     "Super Admin",
		Email:    "superadmin@example.com",
		Password: "superadminpassword", // You should hash the password in a real-world scenario
		Roles:    []string{"admin"},    // Set the roles as an array of role names
	}

	// Find or create the superadmin user by email
	result := db.Table("users").Where("email = ?", superadmin.Email).FirstOrCreate(&superadmin)
	if result.Error != nil {
		fmt.Printf("Failed to create superadmin user: %v\n", result.Error)
		return result.Error
	}

	fmt.Println("Superadmin user created successfully")
	return nil
}

// CreateUser creates a user using the provided user data
func CreateUser(user *models.User, db *gorm.DB) error {
	user.ID = uuid.New()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	err = db.Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}
