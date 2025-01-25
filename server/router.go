package server

import (
	"github.com/gin-gonic/gin"
	"github.com/piyushverma013/token-athena/config"
	"github.com/piyushverma013/token-athena/handler"
)

func (*HTTPServer) setUpRoutes(router *gin.Engine, appConfig *config.AppConfig) {
	troubleshootHandler := handler.NewTroubleshoot(appConfig)
	troubleshootRoutes(router, troubleshootHandler)
}

func troubleshootRoutes(router *gin.Engine, troubleshootHandler handler.Troubleshoot) {
	router.GET("/", troubleshootHandler.HomePage)
}
