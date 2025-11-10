package routes

import (
	"saturn-backend/internal/handlers"
	"saturn-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func LocalRoutes(r *gin.Engine) {
	auth := r.Group("/local", middleware.AuthMiddleware())
	{
		auth.POST("/register", handlers.RegisterLocal)
		auth.POST("/update", handlers.UpdateLocal)
		auth.POST("/delete", handlers.DeleteLocal)
		auth.GET("/local", handlers.GetLocal)
		auth.GET("/local-list", handlers.GetLocalList)
	}
}
