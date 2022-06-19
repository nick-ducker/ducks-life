package main

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/nick-ducker/ducks-life/api/controllers"
	"github.com/nick-ducker/ducks-life/api/models"
	"github.com/stretchr/testify/assert"
)

var test_db *gorm.DB

func SetupRouter() *gin.Engine {
	router := gin.Default()
	return router
}

var tests = []struct {
	init           func(*http.Request)
	url            string
	method         string
	bodyData       string
	expectedCode   int
	responseRegexg string
	msg            string
}{
	// Ping
	{
		func(req *http.Request) {},
		"/ping",
		"GET",
		`{}`,
		200,
		`{"data":"pong"}`,
		"Should return pong and 200",
	},
	// Create one rambling
	{
		func(req *http.Request) {},
		"/ramblings",
		"POST",
		`{"title": "test1", "markdown": "test markdown 1"}`,
		200,
		`{"data":{"id":1,"title":"test1","markdown":"test markdown 1"}}`,
		"Should create and return one rambling",
	},
	// Update one rambling
	{
		func(req *http.Request) {},
		"/ramblings/1",
		"POST",
		`{"title": "updated test1", "markdown": "updated test markdown 1"}`,
		200,
		`{"data":{"id":1,"title":"updated test1","markdown":"updated test markdown 1"}}`,
		"Should update and return one rambling",
	},
	// Get all ramblings
	{
		func(req *http.Request) {},
		"/ramblings",
		"GET",
		`{}`,
		200,
		`{"data":\[{"id":1,"title":"updated test1","markdown":"updated test markdown 1"}\]}`,
		"Should return all ramblings",
	},
	// Get one rambling
	{
		func(req *http.Request) {},
		"/ramblings/1",
		"GET",
		`{}`,
		200,
		`{"data":{"id":1,"title":"updated test1","markdown":"updated test markdown 1"}}`,
		"Should return one rambling",
	},
	// Delete one rambling
	{
		func(req *http.Request) {},
		"/ramblings/1",
		"DELETE",
		`{}`,
		200,
		`{"data":true}`,
		"Should delete the rambling and return no data",
	},
}

func TestWithoutAuth(t *testing.T) {
	asserts := assert.New(t)

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.New()
	test_db := models.ConnectDatabase(true)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "pong"})
	})

	authed := r.Group("/")
	authed.Use(authHandler())
	{
		authed.GET("/ramblings", controllers.GetRamblings)
		authed.GET("/ramblings/:id", controllers.GetRambling)

		authed.POST("/ramblings", controllers.CreateRambling)
		authed.POST("/ramblings/:id", controllers.UpdateRambling)

		authed.DELETE("/ramblings/:id", controllers.DeleteRambling)
	}

	for _, testData := range tests {
		bodyData := testData.bodyData
		req, err := http.NewRequest(testData.method, testData.url, bytes.NewBufferString(bodyData))
		req.Header.Set("X-API-KEY", os.Getenv("API_KEY"))
		asserts.NoError(err)

		testData.init(req)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		asserts.Equal(testData.expectedCode, w.Code, "Response Status - "+testData.msg)
		asserts.Regexp(testData.responseRegexg, w.Body.String(), "Response Content - "+testData.msg)

	}

	test_db.Close()
	err = os.Remove("./test.db")
	if err != nil {
		log.Fatal("Error clearing test database file")
	}
}
