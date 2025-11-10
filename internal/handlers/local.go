package handlers

import (
	"net/http"
	"saturn-backend/internal/database"
	"saturn-backend/internal/models"

	"github.com/gin-gonic/gin"
)

// Crear un nuevo local
func RegisterLocal(c *gin.Context) {
	var local models.Local

	if err := c.ShouldBindJSON(&local); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&local).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create local"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "local created successfully", "local": local})
}

// Actualizar un local existente
func UpdateLocal(c *gin.Context) {
	var local models.Local

	if err := c.ShouldBindJSON(&local); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Model(&models.Local{}).Where("id = ?", local.ID).Updates(local).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update local"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "local updated successfully"})
}

// Eliminar un local
func DeleteLocal(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing local id"})
		return
	}

	if err := database.DB.Delete(&models.Local{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete local"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "local deleted successfully"})
}

// Obtener un local por ID
func GetLocal(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing local id"})
		return
	}

	var local models.Local
	if err := database.DB.First(&local, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "local not found"})
		return
	}

	c.JSON(http.StatusOK, local)
}

// Listar todos los locales
func GetLocalList(c *gin.Context) {
	var locals []models.Local

	if err := database.DB.Find(&locals).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve locals"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"locals": locals, "count": len(locals)})
}
