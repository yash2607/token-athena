package model

type GenerateTokenRequest struct {
	UserID string `json:"user_id" binding:"required"`
}

type GenerateTokenResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

type ValidateTokenRequest struct {
	Token string `json:"token" binding:"required"`
}
