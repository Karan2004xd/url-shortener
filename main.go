package main

import (
	"url-shortner/api"
	"url-shortner/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDb()
	server := gin.Default()

	api.RegisterRoutes(server)
	server.Run(":8080")
}
