package dto

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type CreateCategoryRequest struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

func (dto *CreateCategoryRequest) Bind(body io.ReadCloser) error {
	defer body.Close()
	return json.NewDecoder(body).Decode(dto)
}

func (dto *CreateCategoryRequest) Render(w io.Writer) error {
	return json.NewEncoder(w).Encode(dto)
}

type CategoryResponse struct {
	ID 			string 		`json:"id"`
	Name 		string 		`json:"name"`
	Description string 		`json:"description,omitempty"`
	CreatedAt 	time.Time 	`json:"created_at"`
}

func (dto *CategoryResponse) Bind(body io.ReadCloser) error {
	defer body.Close()
	return json.NewDecoder(body).Decode(dto)
}

func (dto *CategoryResponse) Render(w http.ResponseWriter,) error {
	return json.NewEncoder(w).Encode(dto)
}
