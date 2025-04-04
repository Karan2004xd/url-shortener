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
	shortUrl := context.Param("short_url")

	url, err := models.GetLongUrl(shortUrl)

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

func updateUrl(context *gin.Context) {
	userId := context.GetInt64("user_id")
	shortUrl := context.Param("short_url")

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
	err = url.UpdateUrl(shortUrl)

	if err != nil {
		context.JSON(
			http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "URL Updated successfully",
	})
}

func getAllUrls(context *gin.Context) {
	userId := context.GetInt64("user_id")

	urls, err := models.GetAllCustomUrls(userId)

	if err != nil {
		context.JSON(
			http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"urls": urls,
	})
}
