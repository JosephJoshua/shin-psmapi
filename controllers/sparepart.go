package controllers

import (
	"net/http"
	"strconv"

	"github.com/JosephJoshua/shin-psmapi/forms"
	"github.com/JosephJoshua/shin-psmapi/models"
	"github.com/gin-gonic/gin"
)

type SparepartController struct{}

var sparepartModel = new(models.SparepartModel)

func (SparepartController) GetAll(c *gin.Context) {
	var form forms.GetAllSparepartForm
	if err := c.ShouldBindQuery(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid query parameters",
			"error":   err.Error(),
		})

		return
	}

	sparepartList, err := sparepartModel.All(form)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal saat mengambil sparepart",
			"error":   err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{"data": sparepartList})
}

func (SparepartController) GetByNomorNota(c *gin.Context) {
	nomorNota, err := strconv.Atoi(c.Param("nomor_nota"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Nomor nota servisan harus berupa angka"})
		return
	}

	sparepartList, err := sparepartModel.ByNomorNota(nomorNota)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal saat mengambil sparepart",
			"error":   err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{"data": sparepartList})
}

func (SparepartController) Create(c *gin.Context) {
	var form forms.CreateSparepartForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
			"error":   err.Error(),
		})

		return
	}

	id, err := sparepartModel.Create(form)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal saat membuat sparepart",
			"error":   err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (SparepartController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "ID sparepart harus berupa angka"})
		return
	}

	if err := sparepartModel.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal saat menghapus sparepart",
			"error":   err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
