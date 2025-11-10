package models

import "gorm.io/gorm"

// User model
type User struct {
	gorm.Model
	Username   string   `gorm:"uniqueIndex;not null" json:"username"`
	Email      string   `gorm:"uniqueIndex;not null" json:"email"`
	Password   string   `gorm:"not null" json:"-"`
	UserTypeID uint     `gorm:"not null" json:"userTypeId"`
	UserType   UserType `gorm:"not null" json:"userType"`
	Phone      *string  `gorm:"uniqueIndex" json:"phone"`
}
