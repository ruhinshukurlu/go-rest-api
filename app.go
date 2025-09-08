package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-rest-api/database"
	"github.com/go-rest-api/routes"
)

func main() {
	server := gin.Default()
	routes.RegisterRoutes(server)
	database.InitDB()
	server.Run(":8080")
}
