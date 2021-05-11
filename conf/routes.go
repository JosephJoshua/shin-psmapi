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

		api.POST("/register", userController.Register)
		api.POST("/login", authMiddleware.LoginHandler)

		authRoutes := api.Group("/auth")
		authRoutes.Use(authMiddleware.MiddlewareFunc())
		{
			authRoutes.GET("/refresh_token", authMiddleware.RefreshHandler)
		}
	}
}
