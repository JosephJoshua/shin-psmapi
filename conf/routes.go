package conf

import (
	"shin-psmapi/controllers"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {
	r.Static("/public", "./public")

	api := r.Group("/api")
	{
		userController := controllers.UserController{}

		api.POST("/login", authMiddleware.LoginHandler)

		authRoutes := api.Group("/auth")
		authRoutes.Use(authMiddleware.MiddlewareFunc())
		{
			// Register requires the user to be admin otherwise people can just create
			// admin accounts and login to that account.
			authRoutes.POST("/register", userController.Register)
			authRoutes.GET("/refresh_token", authMiddleware.RefreshHandler)
		}
	}
}
