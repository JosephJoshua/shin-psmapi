package controllers

import (
	"net/http"
	"strconv"

	"github.com/JosephJoshua/shin-psmapi/forms"
	"github.com/JosephJoshua/shin-psmapi/models"
	"github.com/gin-gonic/gin"
)

type TeknisiController struct{}

var teknisiModel = new(models.TeknisiModel)

func (TeknisiController) GetAll(c *gin.Context) {
	if HasBuyerRole(c) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Buyer tidak dapat melihat teknisi"})
		return
	}

	form := forms.GetAllTeknisiForm{}
	if err := c.ShouldBindQuery(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid query parameters",
			"error":   err.Error(),
		})

		return
	}

	teknisiList, err := teknisiModel.All(form.SearchQuery)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal saat mengambil teknisi",
			"error":   err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, teknisiList)
}

func (TeknisiController) GetByID(c *gin.Context) {
	if HasBuyerRole(c) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Buyer tidak dapat melihat teknisi"})
		return
	}

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

	c.JSON(http.StatusOK, teknisi)
}

func (TeknisiController) Create(c *gin.Context) {
	if !HasAdminRole(c) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Hanya admin yang dapat membuat teknisi baru"})
		return
	}

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

	c.JSON(http.StatusCreated, id)
}

func (TeknisiController) Delete(c *gin.Context) {
	if !HasAdminRole(c) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Hanya admin yang dapat menghapus teknisi"})
		return
	}

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
