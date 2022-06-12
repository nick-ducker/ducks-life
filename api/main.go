package main

import (
	"fmt"
	"net/http"

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
		fmt.Println("YOU HIT THE MIDDLEWARE!")
		c.Next()
	}
}
