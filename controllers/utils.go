package controllers

import (
	"fmt"

	"github.com/JosephJoshua/shin-psmapi/db"
	"github.com/JosephJoshua/shin-psmapi/models"
	"github.com/JosephJoshua/shin-psmapi/utils"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func HasAdminRole(c *gin.Context) bool {
	claims := jwt.ExtractClaims(c)

	var user models.User
	db.GetDB().Select("role").Where("id = ?", claims[utils.JWTIdentityKey]).First(&user)

	fmt.Println(user)

	return user.Role == utils.AdminUserRole
}
