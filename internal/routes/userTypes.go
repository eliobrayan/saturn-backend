package routes

import (
	"saturn-backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

func UserTypeRoutes(r *gin.Engine) {
	r.GET("/user-type", handlers.GetUserTypes)
}
