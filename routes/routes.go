package routes

import (
	"WelcomeGo/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)

		api.GET("/toys", controllers.GetToys)
		api.GET("/toys/:id", controllers.GetToyByID)
		api.POST("/toys", controllers.CreateToy)

		api.GET("/categories", controllers.GetCategories)
		api.POST("/categories", controllers.CreateCategory)
	}
}
