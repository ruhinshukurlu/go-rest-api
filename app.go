package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/", handler)

	server.Run(":8080")
}

func handler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Hello World"})
}
