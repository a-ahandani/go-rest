// internal/initializer.go
package initializers

import (
	"fmt"
	"gorest/database"
	userHandlers "gorest/internal/handlers/user"
	"gorest/internal/models"
)

// Initialize creates a superadmin user and basic roles if they don't exist
func CreateAdmin() error {
	// Create basic roles if they don't exist
	if err := createBasicRoles(); err != nil {
		return err
	}

	// Create a superadmin user if it doesn't exist
	if _, err := createSuperAdminUser(); err != nil {
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

func createSuperAdminUser() (*models.User, error) {
	db := database.DB

	// Check if the super admin already exists
	var existingSuperAdmin models.User
	if err := db.Where("email = ?", "a.e.ahandani@gmail.com").First(&existingSuperAdmin).Error; err == nil {
		return &existingSuperAdmin, nil
	}

	// Create a new super admin user request
	superAdminRequest := &userHandlers.UserInputBody{
		Name:     "Super Admin",
		Email:    "a.e.ahandani@gmail.com",
		Password: "123123",          // You should hash the password in a real-world scenario
		Roles:    []string{"admin"}, // Set the roles as an array of role names
	}

	// Create the super admin user and assign roles
	superAdmin, err := userHandlers.CreateUser(superAdminRequest)
	if err != nil {
		return nil, err
	}

	return superAdmin, nil
}
