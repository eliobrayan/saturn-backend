package handlers

import (
	"net/http"
	"saturn-backend/internal/database"
	"saturn-backend/internal/models"

	"github.com/gin-gonic/gin"
)

func GetUserTypes(c *gin.Context) {
	var userTypes []models.UserType

	if err := database.DB.Find(&userTypes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userTypes)
}

// Crear un nuevo usuario
func RegisterUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ⚠️ No guardar contraseñas en texto plano
	// Aquí deberías hashearla (ejemplo con bcrypt)
	// hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	// user.Password = string(hash)

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user created successfully", "user": user})
}

// Actualizar un usuario existente
func UpdateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Model(&models.User{}).Where("id = ?", user.ID).Updates(user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user updated successfully"})
}

// Eliminar un usuario
func DeleteUser(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing user id"})
		return
	}

	if err := database.DB.Delete(&models.User{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
}

// Obtener un usuario por ID
func GetUser(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing user id"})
		return
	}

	var user models.User
	if err := database.DB.Preload("UserType").First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// Listar todos los usuarios
func GetUserList(c *gin.Context) {
	var users []models.User

	if err := database.DB.Preload("UserType").Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users, "count": len(users)})
}

// profile: returns simple profile using token
func Profile(c *gin.Context) {
	u, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found in context"})
		return
	}

	user := u.(models.User)
	var fullUser models.User
	if err := database.DB.Preload("UserType").First(&fullUser, user.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load user type"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         fullUser.ID,
		"username":   fullUser.Username,
		"email":      fullUser.Email,
		"phone":      fullUser.Phone,
		"userTypeId": fullUser.UserTypeID,
		"userType":   fullUser.UserType, // 👈 ahora viene completo
		"createdAt":  fullUser.CreatedAt,
		"updatedAt":  fullUser.UpdatedAt,
	})
}
