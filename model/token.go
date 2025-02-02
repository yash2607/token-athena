package model

type GenerateTokenRequest struct {
	UserID string `json:"user_id" binding:"required"`
}

type GenerateTokenResponse struct {
	Status    int    `json:"status"`
	Message   string `json:"message"`
	Token     string `json:"token"`
	ExpiresIn int    `json:"expires_in"`
	TokenType string `json:"token_type"`
	Jti       string `json:"jti"`
}

type ValidateTokenRequest struct {
	Token string `json:"token" binding:"required"`
}
