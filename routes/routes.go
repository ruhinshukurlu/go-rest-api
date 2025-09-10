package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-rest-api/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	// Event routes
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	// protected event routes
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("events/:id", updateEvent)
	authenticated.DELETE("events/:id", deleteEvent)

	// Auth routes
	server.POST("/signup", signup)
	server.POST("/login", login)
}
