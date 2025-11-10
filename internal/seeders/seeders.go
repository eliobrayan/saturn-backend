package seeders

import (
	"log"
	"saturn-backend/internal/models"

	"gorm.io/gorm"
)

// SeedUserTypes crea los tipos de usuario si no existen
func SeedUserTypes(db *gorm.DB) error {
	userTypes := []models.UserType{
		{Name: "admin"},
		{Name: "kitchen"},
		{Name: "server"},
		{Name: "cashier"},
	}
	log.Printf("Creando tipos de usuario por defecto")
	for _, ut := range userTypes {
		var existing models.UserType
		if err := db.First(&existing, "name = ?", ut.Name).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				db.Create(&ut)
				log.Printf("✅ Tipo de usuario creado: %s\n", ut.Name)
			}
		}
	}
	return nil
}
