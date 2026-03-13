package dto

import (
	"encoding/json"
	"io"
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