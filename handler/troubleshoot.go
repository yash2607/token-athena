package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/piyushverma013/token-athena/config"
)

type troubleshoot struct {
	appConfig *config.AppConfig
}

type Troubleshoot interface {
	HomePage(ctx *gin.Context)
}

func NewTroubleshoot(appConfig *config.AppConfig) Troubleshoot {
	return &troubleshoot{
		appConfig: appConfig,
	}
}

func (obj *troubleshoot) HomePage(ctx *gin.Context) {
	response := map[string]string{
		"message": "Welcome to TokenAthena!",
	}
	ctx.JSON(http.StatusOK, response)
}
