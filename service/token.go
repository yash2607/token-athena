package service

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/piyushverma013/token-athena/config"
	"github.com/piyushverma013/token-athena/model"
)

type TokenService interface {
	GenerateToken(req model.GenerateTokenRequest) (resp model.GenerateTokenResponse, err error)
	ValidateToken(req model.ValidateTokenRequest) (resp *TokenClaims, err error)
}

type tokenService struct {
	secretKey  string
	issuer     string
	expiryTime time.Duration
}

// TokenClaims represents the custom claims in our JWT
type TokenClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func NewTokenService(appConfig *config.AppConfig) TokenService {
	return &tokenService{
		secretKey:  appConfig.SecretKey,
		issuer:     appConfig.Issuer,
		expiryTime: time.Duration(appConfig.TokenExpiryTime) * time.Second,
	}
}

func (ts *tokenService) GenerateToken(req model.GenerateTokenRequest) (resp model.GenerateTokenResponse, err error) {
	claims := TokenClaims{
		UserID: req.UserID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    ts.issuer,
			Subject:   req.UserID,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ts.expiryTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte(ts.secretKey))

	if err != nil {
		return resp, fmt.Errorf("[GenerateToken] [jwtToken.SignedString], %w:", err)
	}

	resp.Status = http.StatusOK
	resp.Message = http.StatusText(http.StatusOK)
	resp.Token = token
	resp.ExpiresIn = int(ts.expiryTime.Seconds())
	resp.TokenType = "Bearer"
	resp.Jti = "jwt"

	return resp, nil
}

func (ts *tokenService) ValidateToken(request model.ValidateTokenRequest) (resp *TokenClaims, err error) {
	token, err := jwt.ParseWithClaims(
		request.Token,
		&TokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("[ValidateToken] [jwt.SigningMethodHMAC], invalid token signing method")

			}
			return []byte(ts.secretKey), nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("[ValidateToken] [jwt.ParseWithClaims], %w:", err)
	}

	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
