package middleware

import (
	"github.com/piyushverma013/token-athena/constant"

	"github.com/gin-gonic/gin"
)

// Config holds the configuration values.
type Config struct {
	HeaderName   string
	HeaderValue  string
	ResponseCode int
}

func HealthCheck() gin.HandlerFunc {
	return New(Config{
		HeaderName:   constant.DefaultHeaderName,
		HeaderValue:  constant.DefaultHeaderValue,
		ResponseCode: constant.DefaultResponseCode,
	})
}

// New Creates a new middileware with `cfg`.
func New(cfg Config) gin.HandlerFunc {
	if cfg.HeaderName == "" {
		cfg.HeaderName = constant.DefaultHeaderName
	}
	if cfg.HeaderValue == "" {
		cfg.HeaderValue = constant.DefaultHeaderValue
	}
	if cfg.ResponseCode == 0 {
		cfg.ResponseCode = constant.DefaultResponseCode
	}

	return func(ctx *gin.Context) {
		if ctx.GetHeader(cfg.HeaderName) == cfg.HeaderValue {
			ctx.JSON(cfg.ResponseCode, gin.H{
				"status": "healthy",
				"code":   cfg.ResponseCode,
			})
			ctx.Abort()
		}
		ctx.Next()
	}
}
