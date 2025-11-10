package routes

import (
	"saturn-backend/internal/handlers"
	"saturn-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine) {
	auth := r.Group("/product", middleware.AuthMiddleware())
	{
		auth.POST("/register", handlers.RegisterProduct)
		auth.POST("/update", handlers.UpdateProduct)
		auth.POST("/delete", handlers.DeleteProduct)
		auth.GET("/product", handlers.GetProduct)
		auth.GET("/product-list", handlers.GetProductList)
	}
}
