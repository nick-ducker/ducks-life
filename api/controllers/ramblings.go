package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nick-ducker/ducks-life/api/models"
)

type CreateRamblingInput struct {
	Title    string `json:"title" binding:"required"`
	Markdown string `json:"markdown" binding:"required"`
}

type UpdateRamblingInput struct {
	Title    string `json:"title"`
	Markdown string `json:"markdown"`
}

// Public

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

func CreateRambling(c *gin.Context) {
	var input CreateRamblingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rambling := models.Rambling{Title: input.Title, Markdown: input.Markdown}
	models.DB.Create(&rambling)

	c.JSON(http.StatusOK, gin.H{"data": rambling})
}

func UpdateRambling(c *gin.Context) {
	var rambling models.Rambling
	if err := models.DB.Where("id = ?", c.Param("id")).First(&rambling).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdateRamblingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&rambling).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": rambling})
}

func DeleteRambling(c *gin.Context) {
	var rambling models.Rambling
	if err := models.DB.Where("id = ?", c.Param("id")).First(&rambling).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&rambling)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
