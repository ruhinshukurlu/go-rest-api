package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-rest-api/models"
	"github.com/go-rest-api/utils"
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

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't signup"})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Can't authorize"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login success", "token": token})
}
