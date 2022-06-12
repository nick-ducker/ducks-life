package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nick-ducker/ducks-life/api/models"
)

func GetRamblings(c *gin.Context) {
	var ramblings []models.Rambling
	models.DB.Find(&ramblings)

	c.JSON(http.StatusOK, gin.H{"data": ramblings})
}

func GetRambling(c *gin.Context) {
	var rambling models.Rambling

	if err := models.DB.Where("id = ?", c.Param("id")).First(&rambling).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": rambling})
}
