package main

import (
	"log"
	"saturn-backend/internal/database"
	"saturn-backend/internal/seeders"
)

func main() {
	database.Connect()
	if err := seeders.SeedUserTypes(database.DB); err != nil {
		log.Fatal("❌ Error seeding user types:", err)
	}
	log.Println("✅ Seeders executed successfully")
}
