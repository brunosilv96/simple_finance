package dto

import (
	"encoding/json"
	"io"
	"time"
)

type RegisterUserRequest struct {
	Name string `json:"name"`
}

func (dto *RegisterUserRequest) Bind(body io.ReadCloser) error {
	defer body.Close()
	return json.NewDecoder(body).Decode(dto)
}


type RegisterUserResponse struct {
	ID string `json:"id"`
	Name string `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}


func (dto *RegisterUserResponse) Render(w io.Writer) error {
	return json.NewEncoder(w).Encode(dto)
}