package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Stock       int     `json:"stock" binding:"required,gte=0"`
	Price       float64 `json:"price" binding:"required,gte=0"`
}
