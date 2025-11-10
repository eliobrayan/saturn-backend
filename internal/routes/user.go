package routes

import (
	"saturn-backend/internal/handlers"
	"saturn-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	user := r.Group("/user", middleware.AuthMiddleware())
	{
		user.GET("/profile", handlers.Profile)
		user.POST("/register", handlers.RegisterUser)
		user.POST("/update", handlers.UpdateUser)
		user.POST("/delete", handlers.DeleteUser)
		user.GET("/user", handlers.GetUser)
		user.GET("/user-list", handlers.GetUserList)
	}
}
