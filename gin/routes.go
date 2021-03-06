package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/skamranahmed/golang-url-shortener-with-redis/handlers"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/generate", handlers.HandleGenerateShortUrl)
	r.GET("/:urlPath", handlers.HandleRedirectToOriginalUrl)
	r.GET("/:urlPath/info", handlers.HandleGetShortUrlInfo)
	return r
}
