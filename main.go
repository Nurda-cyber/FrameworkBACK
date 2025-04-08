package main

import (
	"WelcomeGo/database"
	"WelcomeGo/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.ConnectDB()
	routes.RegisterRoutes(r)
	r.Run(":8080")
}
