package routes

import (
	"WelcomeGo/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		// Auth routes
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)

		// Toy routes
		api.GET("/toys", controllers.GetToys)
		api.POST("/toys", controllers.CreateToy)
	}
}
