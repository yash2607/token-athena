package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/piyushverma013/token-athena/config"
	"github.com/piyushverma013/token-athena/model"
	"github.com/piyushverma013/token-athena/service"
)

type TokenHandler interface {
	GenerateToken(ctx *gin.Context)
	ValidateToken(ctx *gin.Context)
}

type tokenHandler struct {
	tokenService service.TokenService
}

func NewTokenHandler(appConfig *config.AppConfig) TokenHandler {
	return &tokenHandler{
		tokenService: service.NewTokenService(appConfig),
	}
}

func (obj *tokenHandler) GenerateToken(ctx *gin.Context) {
	var req model.GenerateTokenRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := obj.tokenService.GenerateToken(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": resp})
}

func (obj *tokenHandler) ValidateToken(ctx *gin.Context) {
	var req model.ValidateTokenRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := obj.tokenService.ValidateToken(req)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user_id": resp.UserID})
}
