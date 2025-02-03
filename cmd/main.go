package main

import (
	"log"
	"mygram/config"
	"mygram/database"
	"mygram/models"
	"mygram/routes"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := config.ConnectDB()

	// Run migrations
	database.RunMigrations(db)

	// Auto Migrate models
	db.AutoMigrate(&models.User{}, &models.Photo{}, &models.Comment{}, &models.SocialMedia{})

	router := routes.SetupRouter(db)
	router.Run(":8080")
}
