package dto

import (
	"time"
)

type RegisterUserRequest struct {
	Name string `json:"name" binding:"required"`
}

type UserResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
