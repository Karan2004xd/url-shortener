package api

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.POST("/users/signup", createUser)
	server.POST("/users/login", validateUser)
	server.GET("/users", getAllUsers)
}
