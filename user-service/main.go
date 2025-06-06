package main

import (
	"time"
	"user-service/controllers"
	"user-service/database"
	"user-service/middleware"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Use(middleware.LoggerMiddleware())

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	protected := r.Group("/user")
	protected.Use(middleware.JWTMiddleware())

	protected.GET("/me", controllers.GetMe)

	protected.GET("/users", controllers.GetUsers)
	protected.GET("/users/:id", controllers.GetUserByID)
	protected.PUT("/users/:id", controllers.UpdateUser)
	protected.DELETE("/users/:id", controllers.DeleteUser)

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

	r.Run(":8081")
}
