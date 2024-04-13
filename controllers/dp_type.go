package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/JosephJoshua/shin-psmapi/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var dpTypeModel = new(models.DPTypeModel)

type DPTypeController struct{}

func (DPTypeController) GetAll(c *gin.Context) {
	dpTypeList, err := dpTypeModel.All()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal saat mengambil tipe DP",
			"error": err.Error(),
		})
		
		return
	}
	
	c.JSON(http.StatusOK, dpTypeList)
}

func (DPTypeController) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID harus berupa angka"})
		return
	}
	
	dpType, err := dpTypeModel.ByID(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.Status(http.StatusNotFound)
		return
	}
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal saat mengambil tipe DP",
			"error": 	 err.Error(),
		})
		
		return
	}
	
	c.JSON(http.StatusOK, dpType)
}
