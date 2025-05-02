package main

import (
	"WelcomeGo/toy-service/controllers"
	"WelcomeGo/toy-service/database"
	"WelcomeGo/toy-service/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()

	r := gin.Default()
	r.Use(middleware.LoggerMiddleware())

	protected := r.Group("/")
	protected.Use(middleware.JWTMiddleware())
	protected.GET("/toys", controllers.GetToys)

	r.Run(":8082")
}
