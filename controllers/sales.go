package controllers

import (
	"net/http"
	"strconv"

	"github.com/JosephJoshua/shin-psmapi/forms"
	"github.com/JosephJoshua/shin-psmapi/models"
	"github.com/gin-gonic/gin"
)

type SalesController struct{}

var salesModel = new(models.SalesModel)

func (SalesController) GetAll(c *gin.Context) {
	form := forms.GetAllSalesForm{}
	if err := c.ShouldBindQuery(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request parameters",
			"error":   err.Error(),
		})

		return
	}

	salesList, err := salesModel.All(form.SearchQuery)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal saat mengambil sales",
			"error":   err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{"data": salesList})
}

func (SalesController) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "ID sales harus berupa angka"})
		return
	}

	sales, err := salesModel.ByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal saat mengambil sales",
			"error":   err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{"data": sales})
}

func (SalesController) Create(c *gin.Context) {
	if !HasAdminRole(c) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Hanya admin yang dapat membuat sales baru"})
		return
	}

	var form forms.CreateSalesForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
			"error":   err.Error(),
		})

		return
	}

	id, err := salesModel.Create(form)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal saat membuat sales",
			"error":   err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (SalesController) Delete(c *gin.Context) {
	if !HasAdminRole(c) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Hanya admin yang dapat menghapus sales"})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "ID sales harus berupa angka"})
		return
	}

	if err := salesModel.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal saat menghapus sales",
			"error":   err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
