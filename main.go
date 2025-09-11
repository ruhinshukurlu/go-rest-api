package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-rest-api/database"
	_ "github.com/go-rest-api/docs"
	"github.com/go-rest-api/routes"
)

// @title Go Rest API
// @version 1.0
// @description This is Go Rest API for practice
// @host localhost:8080
// @BasePath /api/v1
// @schemes http
// @name Authorization
func main() {
	server := gin.Default()
	routes.RegisterRoutes(server)
	database.InitDB()
	server.Run(":8080")
}
