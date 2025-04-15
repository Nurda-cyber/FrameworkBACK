package routes

import (
	"WelcomeGo/controllers"
	"WelcomeGo/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{

		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)

		protected := api.Group("/")
		protected.Use(middleware.JWTMiddleware())
		{
			protected.GET("/toys", controllers.GetToys)
			protected.GET("/toys/:id", controllers.GetToyByID)
			protected.POST("/toys", controllers.CreateToy)
			protected.DELETE("/toys/:id", controllers.DeleteToy)
			protected.PUT("/toys/:id", controllers.UpdateToy)

			protected.GET("/categories", controllers.GetCategories)
			protected.POST("/categories", controllers.CreateCategory)
			protected.PUT("/categories/:id", controllers.UpdateCategory)
			protected.DELETE("/categories/:id", controllers.DeleteCategory)
			protected.GET("/categories/:id", controllers.GetCategoryByID)

		}
	}
}
