package main

import (
	"toy-service/controllers"
	"toy-service/database"
	"toy-service/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()

	r := gin.Default()
	r.Use(middleware.LoggerMiddleware())

	protected := r.Group("/")
	protected.Use(middleware.JWTMiddleware())

	protected.GET("/toys", controllers.GetToys)
	protected.GET("/toys/:id", controllers.GetToyByID)
	protected.POST("/toys", controllers.CreateToy)
	protected.PUT("/toys/:id", controllers.UpdateToy)
	protected.DELETE("/toys/:id", controllers.DeleteToy)

	protected.GET("/categories", controllers.GetCategories)
	protected.GET("/categories/:id", controllers.GetCategoryByID)
	protected.POST("/categories", controllers.CreateCategory)
	protected.PUT("/categories/:id", controllers.UpdateCategory)
	protected.DELETE("/categories/:id", controllers.DeleteCategory)

	protected.GET("/toys/search", controllers.SearchToysByName)

	r.Run(":8082")
}
