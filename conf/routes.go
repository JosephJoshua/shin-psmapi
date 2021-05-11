package conf

import (
	"shin-psmapi/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.Static("/public", "./public")

	api := r.Group("/api")
	{
		userController := controllers.UserController{}
		api.POST("/register", userController.Register)
	}
}
