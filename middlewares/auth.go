package middlewares

import (
	"net/http"
	"url-shortner/internal"

	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(
			http.StatusUnauthorized, gin.H{
				"message": "No authentication token specified.",
			})
		return
	}

	userId, err := internal.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
		return
	}

	context.Set("user_id", userId)
	context.Next()
}
