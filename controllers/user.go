package controllers

import (
	"net/http"

	"github.com/JosephJoshua/shin-psmapi/forms"
	"github.com/JosephJoshua/shin-psmapi/models"
	"github.com/JosephJoshua/shin-psmapi/utils"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

var userModel = new(models.UserModel)

func (UserController) GetAll(c *gin.Context) {
	if !HasAdminRole(c) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Hanya admin yang dapat mengambil list user"})
		return
	}

	userList, err := userModel.All()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal mengambil user",
			"error":   err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, userList)
}

func (UserController) GetLoggedInUser(c *gin.Context) {
	id, ok := jwt.ExtractClaims(c)[utils.JWTIdentityKey].(float64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "ID user yang di dalam claim JWT tidak valid",
		})

		return
	}

	user, err := userModel.One(int(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal mengambil user",
			"error":   err.Error(),
		})

		return
	}

	// We don't want to return the password
	c.JSON(http.StatusOK, gin.H{
		"id":       user.ID,
		"email":    user.Email,
		"username": user.Username,
		"role":     user.Role,
	})
}

func (UserController) Register(c *gin.Context) {
	if !HasAdminRole(c) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Hanya admin yang dapat register user"})
		return
	}

	var form forms.RegisterForm

	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request body", "error": err.Error()})
		return
	}

	user, err := userModel.Register(form)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal register", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}
