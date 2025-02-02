package server

import (
	"github.com/gin-gonic/gin"
	"github.com/piyushverma013/token-athena/config"
	"github.com/piyushverma013/token-athena/handler"
)

func (*HTTPServer) setUpRoutes(router *gin.Engine, appConfig *config.AppConfig) {
	troubleshootHandler := handler.NewTroubleshoot(appConfig)
	troubleshootRoutes(router, troubleshootHandler)

	tokenHandler := handler.NewTokenHandler(appConfig)
	tokenRoutes(router, tokenHandler)
}

func troubleshootRoutes(router *gin.Engine, troubleshootHandler handler.Troubleshoot) {
	router.GET("/", troubleshootHandler.HomePage)
}

func tokenRoutes(router *gin.Engine, tokenHandler handler.TokenHandler) {
	tokenGroup := router.Group("/token")
	{
		tokenGroup.POST("/generate", tokenHandler.GenerateToken)
		tokenGroup.POST("/validate", tokenHandler.ValidateToken)
	}
}
