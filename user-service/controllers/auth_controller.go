package controllers

import (
	"net/http"
	"os"
	"time"
	"user-service/database"
	"user-service/models"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 14)

	user := models.User{Username: input.Username, Password: string(hashedPassword)}
	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
	var input models.User
	var user models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Where("username = ?", input.Username).First(&user)
	if user.ID == 0 || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error signing token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func GetMe(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":       user.ID,
		"username": user.Username,
	})
}

func GetUsers(c *gin.Context) {
	var users []models.User
	result := database.DB.Find(&users)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	result := database.DB.First(&user, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := database.DB.Delete(&user, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func GetToysFromToyService(c *gin.Context) {
	token := c.GetHeader("Authorization")

	client := resty.New()
	resp, err := client.R().
		SetHeader("Authorization", token).
		SetResult([]map[string]interface{}{}).
		Get("http://toy-service:8082/toys")

	if err != nil || resp.StatusCode() != http.StatusOK {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Failed to fetch toys"})
		return
	}

	c.JSON(http.StatusOK, resp.Result())
}
func GetToys(c *gin.Context) {
	token := c.GetHeader("Authorization")
	client := resty.New()

	queryParams := c.Request.URL.RawQuery
	toyServiceURL := "http://toy-service:8082/toys"
	if queryParams != "" {
		toyServiceURL += "?" + queryParams
	}

	resp, err := client.R().
		SetHeader("Authorization", token).
		SetHeader("Accept", "application/json").
		Get(toyServiceURL)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Ойыншықтарды алу сәтсіз аяқталды"})
		return
	}

	c.Data(resp.StatusCode(), "application/json", resp.Body())
}

func GetToyByID(c *gin.Context) {
	token := c.GetHeader("Authorization")
	id := c.Param("id")
	client := resty.New()
	resp, err := client.R().
		SetHeader("Authorization", token).
		SetResult(map[string]interface{}{}).
		Get("http://toy-service:8082/toys/" + id)

	if err != nil || resp.StatusCode() != http.StatusOK {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Failed to fetch toy"})
		return
	}
	c.JSON(http.StatusOK, resp.Result())
}

func CreateToy(c *gin.Context) {
	token := c.GetHeader("Authorization")
	var toy map[string]interface{}
	if err := c.BindJSON(&toy); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	client := resty.New()
	resp, err := client.R().
		SetHeader("Authorization", token).
		SetBody(toy).
		Post("http://toy-service:8082/toys")

	if err != nil || resp.StatusCode() != http.StatusCreated {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Failed to create toy"})
		return
	}
	c.JSON(http.StatusCreated, resp.Result())
}

func UpdateToy(c *gin.Context) {
	token := c.GetHeader("Authorization")
	id := c.Param("id")
	var toy map[string]interface{}
	if err := c.BindJSON(&toy); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	client := resty.New()
	resp, err := client.R().
		SetHeader("Authorization", token).
		SetBody(toy).
		Put("http://toy-service:8082/toys/" + id)

	if err != nil || resp.StatusCode() != http.StatusOK {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Failed to update toy"})
		return
	}
	c.JSON(http.StatusOK, resp.Result())
}

func DeleteToy(c *gin.Context) {
	token := c.GetHeader("Authorization")
	id := c.Param("id")
	client := resty.New()
	resp, err := client.R().
		SetHeader("Authorization", token).
		Delete("http://toy-service:8082/toys/" + id)

	if err != nil || resp.StatusCode() != http.StatusOK {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Failed to delete toy"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Toy deleted successfully"})
}

func GetCategories(c *gin.Context) {
	token := c.GetHeader("Authorization")
	client := resty.New()
	resp, err := client.R().
		SetHeader("Authorization", token).
		SetResult([]map[string]interface{}{}).
		Get("http://toy-service:8082/categories")

	if err != nil || resp.StatusCode() != http.StatusOK {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Failed to fetch categories"})
		return
	}
	c.JSON(http.StatusOK, resp.Result())
}

func GetCategoryByID(c *gin.Context) {
	token := c.GetHeader("Authorization")
	id := c.Param("id")
	client := resty.New()
	resp, err := client.R().
		SetHeader("Authorization", token).
		SetResult(map[string]interface{}{}).
		Get("http://toy-service:8082/categories/" + id)

	if err != nil || resp.StatusCode() != http.StatusOK {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Failed to fetch category"})
		return
	}
	c.JSON(http.StatusOK, resp.Result())
}

func CreateCategory(c *gin.Context) {
	token := c.GetHeader("Authorization")
	var category map[string]interface{}
	if err := c.BindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	client := resty.New()
	resp, err := client.R().
		SetHeader("Authorization", token).
		SetBody(category).
		Post("http://toy-service:8082/categories")

	if err != nil || resp.StatusCode() != http.StatusCreated {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Failed to create category"})
		return
	}
	c.JSON(http.StatusCreated, resp.Result())
}

func UpdateCategory(c *gin.Context) {
	token := c.GetHeader("Authorization")
	id := c.Param("id")
	var category map[string]interface{}
	if err := c.BindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	client := resty.New()
	resp, err := client.R().
		SetHeader("Authorization", token).
		SetBody(category).
		Put("http://toy-service:8082/categories/" + id)

	if err != nil || resp.StatusCode() != http.StatusOK {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Failed to update category"})
		return
	}
	c.JSON(http.StatusOK, resp.Result())
}

func DeleteCategory(c *gin.Context) {
	token := c.GetHeader("Authorization")
	id := c.Param("id")
	client := resty.New()
	resp, err := client.R().
		SetHeader("Authorization", token).
		Delete("http://toy-service:8082/categories/" + id)

	if err != nil || resp.StatusCode() != http.StatusOK {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Failed to delete category"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}

func SearchToysByName(c *gin.Context) {
	token := c.GetHeader("Authorization")
	query := c.Query("name")
	client := resty.New()
	resp, err := client.R().
		SetHeader("Authorization", token).
		SetResult([]map[string]interface{}{}).
		Get("http://toy-service:8082/toys/search?name=" + query)

	if err != nil || resp.StatusCode() != http.StatusOK {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Search failed"})
		return
	}
	c.JSON(http.StatusOK, resp.Result())
}
