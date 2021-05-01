package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skamranahmed/golang-url-shortner-with-redis/models"
)

// HandleGenerateShortUrl
func HandleGenerateShortUrl(c *gin.Context) {

	c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")

	var userInput models.LongUrlInput
	err := c.ShouldBindJSON(&userInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success":  false,
			"response": "Unable to decode JSON request body",
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"response": userInput.LongUrl,
	})
	return

}
