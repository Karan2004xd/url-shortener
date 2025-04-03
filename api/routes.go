package api

import (
	"url-shortner/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST("/users/signup", createUser)
	server.POST("/users/login", validateUser)
	server.GET("/users", getAllUsers)
	server.GET("/:short_url", getLongUrl)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)

	authenticated.POST("/url", generateShortUrl)
}
