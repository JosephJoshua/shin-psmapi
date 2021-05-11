package conf

import (
	"shin-psmapi/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.Static("/public", "./public")

	userController := controllers.UserController{}
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/register", userController.Register)
	}
}
