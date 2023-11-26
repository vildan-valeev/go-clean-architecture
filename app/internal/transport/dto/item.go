package dto

import (
	"github.com/google/uuid"
	"github.com/vildan-valeev/go-clean-architecture/internal/domain"
)

type ItemCreateRequest struct {
	ID    string `json:"id,omitempty"`
	Title string `json:"title"`
}

type ItemCreateResponse struct {
	ID string `json:"Code"`
}

func ItemCreateToResponse(id uuid.UUID) ItemCreateResponse {
	return ItemCreateResponse{
		ID: id.String(),
	}
}

type ItemDeleteRequest struct {
	ID string `json:"id"`
}

type ItemDeleteResponse struct {
}

func ItemDeleteToResponse() ItemDeleteResponse {
	return ItemDeleteResponse{}
}

type ItemUpdateRequest struct {
	ID string `json:"id"`
}

type ItemUpdateResponse struct {
}

func ItemUpdateToResponse() ItemUpdateResponse {
	return ItemUpdateResponse{}
}

type ItemReadRequest struct {
	ID string `json:"id"`
}

type ItemReadResponse struct {
}

func ItemReadToResponse(i domain.Item) ItemReadResponse {
	return ItemReadResponse{}
}
