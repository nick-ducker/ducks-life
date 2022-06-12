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
