package models

import "gorm.io/gorm"

type Local struct {
	gorm.Model
	Name    string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
}
