package api

import (
	"net/http"
	"url-shortner/models"

	"github.com/gin-gonic/gin"
)

func generateShortUrl(context *gin.Context) {
	userId := context.GetInt64("user_id")
	var url models.Url

	err := context.ShouldBindJSON(&url)

	if err != nil {
		context.JSON(
			http.StatusBadRequest, gin.H{
				"message": "Could not parse the request",
				"error": err.Error(),
			})
		return
	}

	url.UserId = userId
	err = url.GenerateShortUrl()

	if err != nil {
		context.JSON(
			http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"url": url,
	})
}

func getLongUrl(context *gin.Context) {
	short_url := context.Param("short_url")

	url, err := models.GetLongUrl(short_url)

	if url == nil {
		context.JSON(
			http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
		return
	}

	if err != nil {
		context.JSON(
			http.StatusInternalServerError, gin.H{
				"message": "Unable to fetch the url.",
			})
		return
	}

	context.Redirect(301, url.LongUrl)
}
