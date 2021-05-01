package handlers

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/skamranahmed/golang-url-shortner-with-redis/models"
	"github.com/skamranahmed/golang-url-shortner-with-redis/utils"
)

// HandleGenerateShortUrl
func HandleGenerateShortUrl(c *gin.Context) {
	// write header
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

	if !(strings.Contains(userInput.LongUrl, "https://") || strings.Contains(userInput.LongUrl, "http://")) {
		userInput.LongUrl = fmt.Sprintf("https://%s", userInput.LongUrl)
	}

	_, err = url.ParseRequestURI(userInput.LongUrl)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success":  false,
			"response": err.Error(),
		})
		return
	}

	// generate a unique id each time for the redis record
	var uniqueKey uint64
	uniqueKey = rand.Uint64()
	isAlreadyUsed := models.IsUniqueKeyAlreadyUsed(uniqueKey)
	for isAlreadyUsed {
		uniqueKey = rand.Uint64()
		isAlreadyUsed = models.IsUniqueKeyAlreadyUsed(uniqueKey)
	}

	// convert that unique key to url path
	urlPath := utils.ConvertUniqueKeyToUrlPath(uniqueKey)
	u := url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("%s:%s", os.Getenv("APP_SERVER"), os.Getenv("APP_PORT")),
		Path:   urlPath,
	}

	output := models.SaveUrlRecord(fmt.Sprint(uniqueKey), userInput.LongUrl)
	if output.Err() != nil {
		fmt.Println(output.Err().Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"response": u.String(),
	})
	return

}

// HandleGetShortUrlInfo
func HandleGetShortUrlInfo(c *gin.Context) {
	urlPath := c.Param("urlPath")
	uniqueKey, err := utils.ConvertUrlPathToUniqueKey(urlPath)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success":  false,
			"response": "The short url that you are trying to access does not exist.",
		})
		return
	}

	value := models.DB.Get(fmt.Sprint(uniqueKey))
	if value.Val() == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"success":  false,
			"response": "The short url that you are trying to access has either expired or it does not exist.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":    true,
		"target_url": value.Val(),
	})
	return

}
