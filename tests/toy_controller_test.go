package tests

import (
	"WelcomeGo/mocks"
	"WelcomeGo/models"
	"WelcomeGo/routes"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func SetupRouter(mockRepo *mocks.MockToyRepository) *gin.Engine {
	
	r := gin.Default()
	routes.RegisterRoutes(r)
	return r
}

var token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDc5MTU0NDEsInVzZXJfaWQiOjF9.-L3eNBkZZPtuWoP0Q5nCR5uoiuypu1Y-0Zd0JJm3a4k"

func TestGetAllToys(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := new(mocks.MockToyRepository)
	mockToys := []models.Toy{{ID: 4, Name: "Car"}}
	mockRepo.On("GetAll").Return(mockToys, nil)

	router := SetupRouter(mockRepo)

	req, _ := http.NewRequest(http.MethodGet, "/toys", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestGetToyByID(t *testing.T) {
	mockRepo := new(mocks.MockToyRepository)
	mockToy := models.Toy{ID: 4, Name: "Car"}
	
	mockRepo.On("GetByID", uint(4)).Return(mockToy, nil) 

	router := SetupRouter(mockRepo)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/toys/4", nil) 
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestCreateToy(t *testing.T) {
	mockRepo := new(mocks.MockToyRepository)
	inputToy := models.Toy{Name: "Test"}
	mockRepo.On("Create", inputToy).Return(models.Toy{ID: 2, Name: "Train"}, nil)

	router := SetupRouter(mockRepo)

	body, _ := json.Marshal(inputToy)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/toys", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestUpdateToy(t *testing.T) {
	mockRepo := new(mocks.MockToyRepository)
	inputToy := models.Toy{Name: "Test"}
	mockRepo.On("Update", uint(1), inputToy).Return(models.Toy{ID: 1, Name: "Plane"}, nil)

	router := SetupRouter(mockRepo)

	body, _ := json.Marshal(inputToy)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/toys/4", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestDeleteToy(t *testing.T) {
	mockRepo := new(mocks.MockToyRepository)
	mockRepo.On("Delete", uint(1)).Return(nil)

	router := SetupRouter(mockRepo)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/toys/4", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestGetToyByID_NotFound(t *testing.T) {
	mockRepo := new(mocks.MockToyRepository)
	mockRepo.On("GetByID", uint(999)).Return(models.Toy{}, gorm.ErrRecordNotFound)

	router := SetupRouter(mockRepo)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/toys/999", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestCreateToy_InvalidJSON(t *testing.T) {
	mockRepo := new(mocks.MockToyRepository)
	router := SetupRouter(mockRepo)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/toys", bytes.NewBuffer([]byte(`{invalid}`)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

func TestUpdateToy_NotFound(t *testing.T) {
	mockRepo := new(mocks.MockToyRepository)
	inputToy := models.Toy{Name: "Test"}
	mockRepo.On("Update", uint(404), inputToy).Return(models.Toy{}, gorm.ErrRecordNotFound)

	router := SetupRouter(mockRepo)

	body, _ := json.Marshal(inputToy)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/toys/404", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestDeleteToy_NotFound(t *testing.T) {
	mockRepo := new(mocks.MockToyRepository)
	mockRepo.On("Delete", uint(404)).Return(gorm.ErrRecordNotFound)

	router := SetupRouter(mockRepo)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/toys/404", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestCreateToy_EmptyName(t *testing.T) {
	mockRepo := new(mocks.MockToyRepository)
	router := SetupRouter(mockRepo)

	body, _ := json.Marshal(models.Toy{Name: ""})
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/toys", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

func TestUnauthorizedAccess(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := new(mocks.MockToyRepository)
	mockRepo.On("GetAll").Return([]models.Toy{{ID: 1, Name: "Car"}}, nil)

	router := SetupRouter(mockRepo)

	req, _ := http.NewRequest(http.MethodGet, "/toys", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestMissingAuthorizationHeader(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := new(mocks.MockToyRepository)
	mockRepo.On("GetAll").Return([]models.Toy{{ID: 1, Name: "Toy"}}, nil)

	router := SetupRouter(mockRepo)

	req, _ := http.NewRequest(http.MethodGet, "/toys", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestInvalidRoute(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := new(mocks.MockToyRepository)
	router := SetupRouter(mockRepo)

	req, _ := http.NewRequest(http.MethodGet, "/invalid-route", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
