package controllers

import (
	"github.com/JosephJoshua/shin-psmapi/db"
	"github.com/JosephJoshua/shin-psmapi/models"
	"github.com/JosephJoshua/shin-psmapi/utils"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func GetUserRole(c *gin.Context) utils.UserRole {
	claims := jwt.ExtractClaims(c)

	var user models.User
	db.GetDB().Select("role").Where("id = ?", claims[utils.JWTIdentityKey]).First(&user)

	return user.Role
}

func HasAdminRole(c *gin.Context) bool {
	return GetUserRole(c) == utils.AdminUserRole
}

func HasBuyerRole(c *gin.Context) bool {
	return GetUserRole(c) == utils.BuyerUserRole
}

func HasCustomerServiceRole(c *gin.Context) bool {
	return GetUserRole(c) == utils.CustomerServiceUserRole
}
