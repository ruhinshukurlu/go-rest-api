package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-rest-api/models"
)

func signup(context *gin.Context) {

	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't signup"})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't save user to db"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "user created successfully!"})
}
