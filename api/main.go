package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/nick-ducker/ducks-life/api/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	authed := r.Group("/")
	authed.Use(authHandler())
	{
		authed.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"data": "hello world"})
		})
	}

	r.Run()
}

func authHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header.Get("X-API-KEY") != "" {
			if c.Request.Header.Get("X-API-KEY") == os.Getenv("API-KEY") {
				c.Next()
			}
		}
		c.AbortWithStatus(401)
	}
}
