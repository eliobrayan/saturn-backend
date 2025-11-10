package main

import (
	"log"
	"saturn-backend/internal/database"
	"saturn-backend/internal/models"
)

func main() {
	database.Connect()

	err := database.DB.AutoMigrate(
		&models.User{},
		&models.UserType{},
	)
	if err != nil {
		log.Fatal("❌ Error running migrations:", err)
	}

	log.Println("✅ Migrations completed successfully")
}
