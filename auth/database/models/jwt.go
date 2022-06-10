package models

type Token struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}
