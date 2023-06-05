package webserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type WebServer struct {
	*gin.Engine
}

func NewWebServer() *WebServer {
	gin.SetMode(gin.ReleaseMode)

	webServer := &WebServer{
		gin.Default(),
	}

	webServer.EnableCors()

	return webServer
}

func (webServer *WebServer) EnableCors() {
	webServer.Use(func(ctx *gin.Context) {
		ctx.Writer.Header().Add("Access-Control-Allow-Methods", "*")
		ctx.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Add("Access-Control-Allow-Headers", "*")
		ctx.Writer.Header().Add("Access-Control-Max-Age", "3600")

		if ctx.Request.Method == http.MethodOptions {
			ctx.JSON(http.StatusOK, "")
		}
	})
}
