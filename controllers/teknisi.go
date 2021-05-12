package controllers

import (
	"net/http"
	"shin-psmapi/forms"
	"shin-psmapi/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TeknisiController struct{}

var teknisiModel = new(models.TeknisiModel)

func (TeknisiController) GetAll(c *gin.Context) {
	form := forms.GetAllTeknisiForm{}
	if err := c.ShouldBindQuery(&form); err != nil {
		form.SearchQuery = ""
	}

	teknisiList, err := teknisiModel.All(form.SearchQuery)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal saat mengambil teknisi",
			"error":   err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{"data": teknisiList})
}

func (TeknisiController) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "ID teknisi harus berupa angka"})
		return
	}

	teknisi, err := teknisiModel.ByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal saat mengambil teknisi",
			"error":   err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{"data": teknisi})
}

func (TeknisiController) Create(c *gin.Context) {
	var form forms.CreateTeknisiForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
			"error":   err.Error(),
		})

		return
	}

	id, err := teknisiModel.Create(form)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal saat membuat teknisi",
			"error":   err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (TeknisiController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "ID teknisi harus berupa angka"})
		return
	}

	if err := teknisiModel.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal saat menghapus teknisi",
			"error":   err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
