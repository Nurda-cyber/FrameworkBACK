package controllers

import (
	"WelcomeGo/database"
	"WelcomeGo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	var categories []models.Category
	if err := database.DB.Preload("Toys").Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error retrieving categories"})
		return
	}
	c.JSON(http.StatusOK, categories)
}

func CreateCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input format"})
		return
	}
	if err := database.DB.Create(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error adding category"})
		return
	}
	c.JSON(http.StatusCreated, category)
}
func UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	var category models.Category
	if err := database.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Категория табылмады"})
		return
	}

	var input models.Category
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Қате формат"})
		return
	}

	category.Name = input.Name
	if err := database.DB.Save(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Жаңарту кезінде қате"})
		return
	}

	c.JSON(http.StatusOK, category)
}

func DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.Category{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Категорияны өшіруде қате болды"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Категория өшірілді"})
}
func GetCategoryByID(c *gin.Context) {
	id := c.Param("id")
	var category models.Category

	if err := database.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	c.JSON(http.StatusOK, category)
}
