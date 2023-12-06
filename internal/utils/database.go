package utils

import (
	"fmt"
	"gorest/database"
	"gorest/internal/models"
)

func CreateBasicRoles() error {
	roles := []models.Role{
		{Name: "admin", Label: "Admin"},
		{Name: "user", Label: "User"},
	}

	for _, role := range roles {
		// Find or create the role by name
		result := database.DB.Where("name = ?", role.Name).FirstOrCreate(&role)
		if result.Error != nil {
			fmt.Printf("Failed to create role '%s': %v\n", role.Name, result.Error)
			return result.Error
		}
	}

	fmt.Println("Basic roles created successfully")
	return nil
}

func CreateSuperAdminUser() (*models.User, error) {
	// Check if the super admin already exists
	var existingSuperAdmin models.User
	if err := database.DB.Where("email = ?", "a.e.ahandani@gmail.com").First(&existingSuperAdmin).Error; err == nil {
		return &existingSuperAdmin, nil
	}

	// Create a new super admin user request
	superAdminRequest := &UserPayload{
		User: models.User{
			Name:     "Super Admin",
			Email:    "a.e.ahandani@gmail.com",
			Password: "123123",
		},
		Roles: []string{"admin"}, // Set the roles as an array of role names
	}

	// Create the super admin user and assign roles
	superAdmin, err := CreateUser(superAdminRequest)
	if err != nil {
		return nil, err
	}

	return superAdmin, nil
}
