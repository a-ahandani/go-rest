package database

import (
	"fmt"
	"gorest/config"
	"gorest/internal/models"
	"log"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		log.Println(err)
		return err
	}

	// Connection URL to connect to PostgreSQL Database
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))

	// Connect to the DB and initialize the DB variable
	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Println("Failed to connect to the database:", err)
		return err
	}

	// Enable the uuid-ossp extension
	err = DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error
	if err != nil {
		log.Println("Failed to enable uuid-ossp extension:", err)
		return err
	}

	// AutoMigrate only if necessary (e.g., during application startup)
	err = migrateDB()
	if err != nil {
		return err
	}

	fmt.Println("Connection Opened to Database")
	return nil
}

func migrateDB() error {
	// AutoMigrate only if necessary (e.g., during application startup)
	err := DB.AutoMigrate(&models.User{}, &models.Role{}, &models.Note{})
	if err != nil {
		log.Println("Failed to auto migrate database:", err)
		return err
	}
	return nil
}
