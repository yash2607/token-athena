package service

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/piyushverma013/token-athena/config"
	"github.com/piyushverma013/token-athena/model"
)

func TestGenerateToken(t *testing.T) {
	appConfig := &config.AppConfig{
		SecretKey:       "my-secret-key",
		Issuer:          "test-issuer",
		TokenExpiryTime: 3600,
	}

	tests := []struct {
		count int
		name  string
		req   model.GenerateTokenRequest
		err   string
		want  model.GenerateTokenResponse
	}{
		{
			count: 1,
			name:  "success - valid token generation",
			req: model.GenerateTokenRequest{
				UserID: "12345",
			},
			want: model.GenerateTokenResponse{
				Status:    http.StatusOK,
				Message:   http.StatusText(http.StatusOK),
				TokenType: "Bearer",
				ExpiresIn: 3600,
				Jti:       "jwt",
			},
		},
	}

	for _, tt := range tests {
		ts := NewTokenService(appConfig)
		t.Run(fmt.Sprintf("case: %v %v", tt.count, tt.name), func(t *testing.T) {
			resp, err := ts.GenerateToken(tt.req)

			if tt.err != "" {
				assert.NotEqual(t, nil, err)
				assert.Equal(t, tt.err, err.Error())
			} else {
				assert.Equal(t, nil, err)
				assert.Equal(t, tt.want.Status, resp.Status)
				assert.Equal(t, tt.want.Message, resp.Message)
				assert.Equal(t, tt.want.TokenType, resp.TokenType)
				assert.Equal(t, tt.want.ExpiresIn, resp.ExpiresIn)
				assert.Equal(t, tt.want.Jti, resp.Jti)
				assert.NotEqual(t, "", resp.Token)
			}
		})
	}
}

func TestValidateToken(t *testing.T) {
	appConfig := &config.AppConfig{
		SecretKey:       "my-secret-key",
		Issuer:          "test-issuer",
		TokenExpiryTime: 3600,
	}

	tests := []struct {
		count int
		name  string
		setup func() string
		req   model.ValidateTokenRequest
		want  *TokenClaims
		err   string
	}{
		{
			count: 1,
			name:  "success - valid token validation",
			setup: func() string {
				ts := NewTokenService(appConfig)
				resp, _ := ts.GenerateToken(model.GenerateTokenRequest{UserID: "12345"})
				return resp.Token
			},
			req: model.ValidateTokenRequest{},
			want: &TokenClaims{
				UserID: "12345",
				RegisteredClaims: jwt.RegisteredClaims{
					Issuer: "test-issuer",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("case: %v %v", tt.count, tt.name), func(t *testing.T) {
			ts := NewTokenService(appConfig)
			if tt.setup != nil {
				tt.req.Token = tt.setup()
			}

			claims, err := ts.ValidateToken(tt.req)

			if tt.err != "" {
				assert.NotEqual(t, nil, err)
				assert.Equal(t, tt.err, err.Error())
			} else {
				assert.Equal(t, tt.want.UserID, claims.UserID)
				assert.Equal(t, tt.want.Issuer, claims.Issuer)
			}
		})
	}
}
