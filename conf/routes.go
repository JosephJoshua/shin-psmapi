package conf

import (
	"github.com/JosephJoshua/shin-psmapi/controllers"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {
	userController := controllers.UserController{}
	salesController := controllers.SalesController{}
	teknisiController := controllers.TeknisiController{}
	servisanController := controllers.ServisanController{}
	sparepartController := controllers.SparepartController{}

	r.POST("/login", authMiddleware.LoginHandler)

	authRoutes := r.Group("/auth")
	authRoutes.Use(authMiddleware.MiddlewareFunc())
	{
		// Register requires the user to be admin otherwise people can just create
		// admin accounts and login to that account.
		authRoutes.POST("/register", userController.Register)
		authRoutes.GET("/refresh_token", authMiddleware.RefreshHandler)
	}

	userRoutes := r.Group("/users")
	userRoutes.Use(authMiddleware.MiddlewareFunc())
	{
		userRoutes.GET("/", userController.GetAll)
		userRoutes.GET("/current", userController.GetLoggedInUser)
	}

	salesRoutes := r.Group("/sales")
	salesRoutes.Use(authMiddleware.MiddlewareFunc())
	{
		salesRoutes.GET("/", salesController.GetAll)
		salesRoutes.POST("/", salesController.Create)
		salesRoutes.GET("/:id", salesController.GetByID)
		salesRoutes.DELETE("/:id", salesController.Delete)
	}

	teknisiRoutes := r.Group("/teknisi")
	teknisiRoutes.Use(authMiddleware.MiddlewareFunc())
	{
		teknisiRoutes.GET("/", teknisiController.GetAll)
		teknisiRoutes.POST("/", teknisiController.Create)
		teknisiRoutes.GET("/:id", teknisiController.GetByID)
		teknisiRoutes.DELETE("/:id", teknisiController.Delete)
	}

	servisanRoutes := r.Group("/servisan")
	servisanRoutes.Use(authMiddleware.MiddlewareFunc())
	{
		servisanRoutes.GET("/", servisanController.GetAll)
		servisanRoutes.POST("/", servisanController.Create)
		servisanRoutes.GET("/:nomor_nota", servisanController.GetByNomorNota)
		servisanRoutes.PUT("/:nomor_nota", servisanController.Update)
		servisanRoutes.DELETE("/:nomor_nota", servisanController.Delete)
	}

	servisanReportRoutes := r.Group("/servisan-reports")
	servisanReportRoutes.Use(authMiddleware.MiddlewareFunc())
	{
		servisanReportRoutes.GET("/laba-rugi", servisanController.GetLabaRugiReport)
		servisanReportRoutes.GET("/teknisi", servisanController.GetTeknisiReport)
		servisanReportRoutes.GET("/sisa", servisanController.GetSisaReport)
	}

	sparepartRoutes := r.Group("/sparepart")
	sparepartRoutes.Use(authMiddleware.MiddlewareFunc())
	{
		sparepartRoutes.GET("/", sparepartController.GetAll)
		sparepartRoutes.POST("/", sparepartController.Create)
		sparepartRoutes.GET("/:nomor_nota", sparepartController.GetByNomorNota)
		sparepartRoutes.DELETE("/:id", sparepartController.Delete)
	}
}
