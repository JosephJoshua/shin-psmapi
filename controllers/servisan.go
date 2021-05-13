package controllers

import (
	"net/http"
	"shin-psmapi/forms"
	"shin-psmapi/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ServisanController struct{}

var servisanModel = new(models.ServisanModel)

func (ServisanController) GetAll(c *gin.Context) {
	form := forms.GetAllServisanForm{}
	if err := c.ShouldBindQuery(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid query parameters",
			"error":   err.Error(),
		})

		return
	}

	servisanList, err := servisanModel.All(form)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal saat mengambil servisan",
			"error":   err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{"data": servisanList})
}

func (ServisanController) GetByNomorNota(c *gin.Context) {
	nomorNota, err := strconv.Atoi(c.Param("nomor_nota"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Nomor nota servisan harus berupa angka"})
		return
	}

	servisan, err := servisanModel.ByNomorNota(nomorNota)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal saat mengambil servisan",
			"error":   err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{"data": servisan})
}
