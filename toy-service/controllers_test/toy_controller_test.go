package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"toy-service/controllers"
	"toy-service/database"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func init() {
	database.ConnectDB()
}

func TestGetToys(t *testing.T) {
	r := gin.Default()
	r.GET("/toys", controllers.GetToys)

	w := performRequest(r, "GET", "/toys")

	assert.Equal(t, 200, w.Code)
}

func TestCreateToy(t *testing.T) {
	r := gin.Default()
	r.POST("/toys", controllers.CreateToy)

	toy := map[string]interface{}{
		"name":        "Dino Puzzle",
		"description": "A fun puzzle to improve logical thinking.",
		"price":       25.0,
		"category_id": 2,
	}
	jsonValue, _ := json.Marshal(toy)

	w := performRequestWithJSON(r, "POST", "/toys", jsonValue)

	assert.Equal(t, 201, w.Code)
}

func TestGetToyByID(t *testing.T) {
	r := gin.Default()
	r.GET("/toys/:id", controllers.GetToyByID)

	w := performRequest(r, "GET", "/toys/2")
	assert.Equal(t, 200, w.Code)
}

func TestUpdateToy(t *testing.T) {
	r := gin.Default()
	r.PUT("/toys/:id", controllers.UpdateToy)

	update := map[string]interface{}{
		"name":        "Updated Dino Puzzle",
		"description": "A more challenging puzzle.",
		"price":       40.0,
		"category_id": 2,
	}
	jsonValue, _ := json.Marshal(update)

	w := performRequestWithJSON(r, "PUT", "/toys/2", jsonValue)

	assert.Equal(t, 200, w.Code)
}

func TestDeleteToy(t *testing.T) {
	r := gin.Default()
	r.DELETE("/toys/:id", controllers.DeleteToy)

	w := performRequest(r, "DELETE", "/toys/2")

	assert.Equal(t, 200, w.Code)
}

func TestSearchToysByName(t *testing.T) {
	r := gin.Default()
	r.GET("/toys/search", controllers.SearchToysByName)

	w := performRequest(r, "GET", "/toys/search?name=Dino")

	assert.Equal(t, 200, w.Code)
}

func TestGetToysByCategoryAndPrice(t *testing.T) {
	r := gin.Default()
	r.GET("/toys", controllers.GetToys)

	w := performRequest(r, "GET", "/toys?category_id=2&price=40")

	assert.Equal(t, 200, w.Code)
}

func TestCreateToyInvalidPrice(t *testing.T) {
	r := gin.Default()
	r.POST("/toys", controllers.CreateToy)

	toy := map[string]interface{}{
		"name":        "Faulty Toy",
		"description": "Toy with invalid price.",
		"price":       -15.0,
		"category_id": 3,
	}
	jsonValue, _ := json.Marshal(toy)

	w := performRequestWithJSON(r, "POST", "/toys", jsonValue)

	assert.Equal(t, 400, w.Code)
}

func TestGetToysInvalidPage(t *testing.T) {
	r := gin.Default()
	r.GET("/toys", controllers.GetToys)

	w := performRequest(r, "GET", "/toys?page=xyz")

	assert.Equal(t, 400, w.Code)
}

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func performRequestWithJSON(r http.Handler, method, path string, jsonBody []byte) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
