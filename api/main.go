package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nick-ducker/ducks-life/api/controllers"
	"github.com/nick-ducker/ducks-life/api/models"
)

func init() {
	env := os.Getenv("ENVIRONMENT")
	if env != "production" && env != "docker" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "pong"})
	})

	authed := r.Group("/")
	authed.Use(authHandler())
	{
		authed.GET("/ramblings", controllers.GetRamblings)
		authed.GET("/ramblings/:id", controllers.GetRambling)

		authed.POST("/ramblings", controllers.CreateRambling)
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
