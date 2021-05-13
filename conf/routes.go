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
		salesController := controllers.SalesController{}
		teknisiController := controllers.TeknisiController{}
		servisanController := controllers.ServisanController{}

		api.POST("/login", authMiddleware.LoginHandler)

		authRoutes := api.Group("/auth")
		authRoutes.Use(authMiddleware.MiddlewareFunc())
		{
			// Register requires the user to be admin otherwise people can just create
			// admin accounts and login to that account.
			authRoutes.POST("/register", userController.Register)
			authRoutes.GET("/refresh_token", authMiddleware.RefreshHandler)
		}

		salesRoutes := api.Group("/sales")
		salesRoutes.Use(authMiddleware.MiddlewareFunc())
		{
			salesRoutes.GET("/", salesController.GetAll)
			salesRoutes.POST("/", salesController.Create)
			salesRoutes.GET("/:id", salesController.GetByID)
			salesRoutes.DELETE("/:id", salesController.Delete)
		}

		teknisiRoutes := api.Group("/teknisi")
		teknisiRoutes.Use(authMiddleware.MiddlewareFunc())
		{
			teknisiRoutes.GET("/", teknisiController.GetAll)
			teknisiRoutes.POST("/", teknisiController.Create)
			teknisiRoutes.GET("/:id", teknisiController.GetByID)
			teknisiRoutes.DELETE("/:id", teknisiController.Delete)
		}

		servisanRoutes := api.Group("/servisan")
		servisanRoutes.Use(authMiddleware.MiddlewareFunc())
		{
			servisanRoutes.GET("/", servisanController.GetAll)
			servisanRoutes.POST("/", servisanController.Create)
			servisanRoutes.GET("/:nomor_nota", servisanController.GetByNomorNota)
			servisanRoutes.DELETE("/:nomor_nota", servisanController.Delete)
		}
	}
}
