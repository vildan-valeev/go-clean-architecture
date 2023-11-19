package dto

import "github.com/vildan-valeev/go-clean-architecture/internal/domain"

type CategoryDtoRequest struct {
	ID    string `json:"id,omitempty"`
	Title string `json:"title"`
}

type CategoryDtoResponse struct {
	ID string `json:"Code"`
}

func ToDTO(sign domain.Category) CategoryDtoResponse {
	return CategoryDtoResponse{
		ID: sign.Tag,
	}
}
