package main

import (
	"WelcomeGo/user-service/controllers"
	"WelcomeGo/user-service/database"
	"WelcomeGo/user-service/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()

	r := gin.Default()
	r.Use(middleware.LoggerMiddleware())

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	protected := r.Group("/")
	protected.Use(middleware.JWTMiddleware())
	protected.GET("/get-toys", controllers.GetToysFromToyService)

	r.Run(":8081")
}
