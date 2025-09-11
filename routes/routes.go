package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-rest-api/middlewares"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(server *gin.Engine) {
	// Swagger endpoint
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Create API v1 group
	v1 := server.Group("/api/v1")

	// Event routes
	v1.GET("/events", getEvents)
	v1.GET("/events/:id", getEvent)

	// Protected event routes
	authenticated := v1.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	// Auth routes
	v1.POST("/signup", signup)
	v1.POST("/login", login)
}
