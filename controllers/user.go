package controllers

import (
	"net/http"

	"github.com/JosephJoshua/shin-psmapi/forms"
	"github.com/JosephJoshua/shin-psmapi/models"
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

	c.JSON(http.StatusOK, gin.H{"data": userList})
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

	c.JSON(http.StatusCreated, gin.H{"data": user})
}
