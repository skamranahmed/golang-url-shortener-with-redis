package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/skamranahmed/golang-url-shortner-with-redis/handlers"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/generate", handlers.HandleGenerateShortUrl)
	r.GET("/:urlPath/info", handlers.HandleGetShortUrlInfo)
	return r
}
