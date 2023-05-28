package webserver

import "github.com/gin-gonic/gin"

func NewWebServer() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	return gin.Default()
}
