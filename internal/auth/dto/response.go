package dto

import (
	"time"
)

type UserResponse struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
}
