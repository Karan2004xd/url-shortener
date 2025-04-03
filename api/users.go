package api

import (
	"net/http"
	"url-shortner/internal"
	"url-shortner/models"

	"github.com/gin-gonic/gin"
)

func createUser(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(
			http.StatusBadRequest, gin.H{
				"message": "Could not parse the request",
				"error": err.Error(),
			})
		return
	}

	err = user.Create()

	if err != nil {
		context.JSON(
			http.StatusInternalServerError, gin.H{
				"message": "Could not save the user",
				"error": err.Error(),
			})
		return
	}

	context.JSON(
		http.StatusCreated, gin.H{
			"message": "User Created",
			"user": user,
		})
}

func validateUser(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(
			http.StatusBadRequest, gin.H{
				"message": "Could not parse the request",
				"error": err.Error(),
			})
		return
	}

	err = user.Validate()

	if err != nil {
		context.JSON(
			http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
		return
	}

	token, err := internal.GenerateToken(user.Email, user.Id)

	if err != nil {
		context.JSON(
			http.StatusUnauthorized, gin.H{
				"message": "Error in token generation",
				"error": err.Error(),
			})
		return
	}

	context.JSON(
		http.StatusOK, gin.H{
			"message": "Login was Successfull",
			"token": token,
		})
}

func getAllUsers(context *gin.Context) {
	users, err := models.GetAllUsers()

	if err != nil {
		context.JSON(
			http.StatusInternalServerError, gin.H{
				"message": "Could not fetch users",
				"error": err.Error(),
			})
		return
	}

	context.JSON(
		http.StatusOK, gin.H{
			"users": users,
		})
}
