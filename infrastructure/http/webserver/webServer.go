package webserver

import "github.com/gin-gonic/gin"

type WebServer struct {
	*gin.Engine
}

func NewWebServer() *WebServer {
	gin.SetMode(gin.ReleaseMode)

	return &WebServer{
		gin.Default(),
	}
}
