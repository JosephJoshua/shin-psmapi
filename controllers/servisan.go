package controllers

import (
	"net/http"
	"strconv"

	"github.com/JosephJoshua/shin-psmapi/forms"
	"github.com/JosephJoshua/shin-psmapi/models"
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

	c.JSON(http.StatusOK, servisanList)
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

	c.JSON(http.StatusOK, servisan)
}

func (ServisanController) Create(c *gin.Context) {
	var form forms.CreateServisanForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
			"error":   err.Error(),
		})

		return
	}

	nomorNota, err := servisanModel.Create(form)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal saat membuat servisan",
			"error":   err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, nomorNota)
}

func (ServisanController) Update(c *gin.Context) {
	nomorNota, err := strconv.Atoi(c.Param("nomor_nota"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Nomor nota servisan harus berupa angka"})
		return
	}

	var form forms.UpdateServisanForm
	if err = c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
			"error":   err.Error(),
		})

		return
	}

	if err = servisanModel.Update(nomorNota, form); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal saat meng-update servisan",
			"error":   err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (ServisanController) Delete(c *gin.Context) {
	nomorNota, err := strconv.Atoi(c.Param("nomor_nota"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Nomor nota servisan harus berupa angka"})
		return
	}

	if err := servisanModel.Delete(nomorNota); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal saat menghapus servisan",
			"error":   err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
