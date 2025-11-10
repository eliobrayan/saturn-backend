package main

import (
	"log"
	"os"
	"saturn-backend/internal/database"
	"saturn-backend/internal/models"
	"saturn-backend/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// load .env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	database.Connect()
	database.DB.AutoMigrate(&models.User{})
	r := gin.Default()
	routes.AuthRoutes(r)
	routes.UserRoutes(r)
	routes.UserTypeRoutes(r)
	routes.LocalRoutes(r)
	routes.ProductRoutes(r)
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	log.Println("Server running on port", port)
	r.Run(":" + port)
}
