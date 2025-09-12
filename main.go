package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-rest-api/database"
	_ "github.com/go-rest-api/docs"
	"github.com/go-rest-api/routes"
)

// @title Your API
// @version 1.0
// @description This is your API description
// @host localhost:8080
// @BasePath /api/v1
// @schemes http
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token. Example: "Bearer {token}"
func main() {
	server := gin.Default()
	routes.RegisterRoutes(server)
	database.InitDB()
	server.Run(":8080")
}
