package controllers

import (
	"WelcomeGo/database"
	"WelcomeGo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetToys - барлық ойыншықтарды алу
func GetToys(c *gin.Context) {
	var toys []models.Toy
	if err := database.DB.Find(&toys).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Ойыншықтарды алу кезінде қате пайда болды"})
		return
	}
	c.JSON(http.StatusOK, toys)
}

// CreateToy - жаңа ойыншық қосу
func CreateToy(c *gin.Context) {
	var toy models.Toy
	if err := c.ShouldBindJSON(&toy); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Қате енгізу форматы"})
		return
	}
	if err := database.DB.Create(&toy).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Ойыншықты қосу кезінде қате болды"})
		return
	}
	c.JSON(http.StatusCreated, toy)
}
