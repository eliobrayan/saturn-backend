package handlers

import (
	"net/http"
	"saturn-backend/internal/database"
	"saturn-backend/internal/models"
	"saturn-backend/internal/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// register: creates a user and returns a token
func Register(c *gin.Context) {
	var input struct {
		Username   string `json:"username" binding:"required,alphanum"`
		Email      string `json:"email" binding:"required,email"`
		Password   string `json:"password" binding:"required,min=6"`
		UserTypeId uint   `json:"userTypeId" binding:"required"`
		Phone      string `json:"phone"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Username:   input.Username,
		Email:      input.Email,
		UserTypeID: input.UserTypeId,
		Phone:      &input.Phone,
	}
	hashed, _ := utils.HashPassword(input.Password)
	user.Password = hashed

	if err := database.DB.Create(&user).Error; err != nil {
		// detect unique constraint
		c.JSON(http.StatusConflict, gin.H{"error": "user with that username or email already exists"})
		return
	}

	token, _ := utils.GenerateToken(user.ID)
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// login: returns token
func Login(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := database.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}

	if !utils.CheckPassword(user.Password, input.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	token, _ := utils.GenerateToken(user.ID)
	c.JSON(http.StatusOK, gin.H{"token": token})
}
