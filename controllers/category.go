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
