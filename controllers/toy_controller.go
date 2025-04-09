package controllers

import (
	"WelcomeGo/database"
	"WelcomeGo/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetToys(c *gin.Context) {
	var toys []models.Toy
	categoryID := c.DefaultQuery("category_id", "")
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid page number"})
		return
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid limit"})
		return
	}

	offset := (pageInt - 1) * limitInt

	query := database.DB
	if categoryID != "" {
		query = query.Where("category_id = ?", categoryID)
	}

	if err := query.Offset(offset).Limit(limitInt).Find(&toys).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error retrieving toys"})
		return
	}
	c.JSON(http.StatusOK, toys)
}

func CreateToy(c *gin.Context) {
	var toy models.Toy
	if err := c.ShouldBindJSON(&toy); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input format"})
		return
	}

	if toy.Price <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Price must be greater than 0"})
		return
	}

	if err := database.DB.Create(&toy).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error adding toy"})
		return
	}
	c.JSON(http.StatusCreated, toy)
}
func GetToyByID(c *gin.Context) {
	id := c.Param("id")

	toyID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Қате ID форматы"})
		return
	}

	var toy models.Toy
	if err := database.DB.First(&toy, toyID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Ойыншық табылмады"})
		return
	}

	c.JSON(http.StatusOK, toy)
}
