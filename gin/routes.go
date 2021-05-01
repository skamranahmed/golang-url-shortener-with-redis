package gin

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/generate", handlers.HandleCreateShortUrl)
	return r
}
