package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-rest-api/utils"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized request"})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid token request"})
		return
	}

	// set userid to get in the route functions
	context.Set("userId", userId)
	context.Next()
}
