package controllers

import (
	"WelcomeGo/toy-service/database"
	"WelcomeGo/toy-service/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetToys(c *gin.Context) {
	var toys []models.Toy
	categoryID := c.DefaultQuery("category_id", "")
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")
	price := c.DefaultQuery("price", "")

	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Бет нөмірі дұрыс емес"})
		return
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil || limitInt < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Limit дұрыс емес"})
		return
	}

	var priceFloat float64
	if price != "" {
		priceFloat, err = strconv.ParseFloat(price, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Бағаны дұрыс енгізіңіз"})
			return
		}
	}

	offset := (pageInt - 1) * limitInt

	query := database.DB.Preload("Category")

	if categoryID != "" {
		query = query.Where("category_id = ?", categoryID)
	}

	if priceFloat > 0 {
		query = query.Where("price <= ?", priceFloat)
	}

	if err := query.Offset(offset).Limit(limitInt).Find(&toys).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Ойыншықтарды алу кезінде қате болды"})
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
func UpdateToy(c *gin.Context) {
	id := c.Param("id")
	var toy models.Toy
	if err := database.DB.First(&toy, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Ойыншық табылмады"})
		return
	}

	var input models.Toy
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Қате формат"})
		return
	}

	toy.Name = input.Name
	toy.Description = input.Description
	toy.Price = input.Price
	toy.CategoryID = input.CategoryID

	if err := database.DB.Save(&toy).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Жаңарту кезінде қате"})
		return
	}

	c.JSON(http.StatusOK, toy)
}

func DeleteToy(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.Toy{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Ойыншықты өшіруде қате болды"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Ойыншық өшірілді"})
}

func SearchToysByName(c *gin.Context) {
	query := c.Query("name")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Атауды енгізіңіз"})
		return
	}

	var toys []models.Toy
	if err := database.DB.Where("name ILIKE ?", "%"+query+"%").Find(&toys).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Іздеу кезінде қате пайда болды"})
		return
	}

	c.JSON(http.StatusOK, toys)
}
