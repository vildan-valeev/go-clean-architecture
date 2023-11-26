package dto

import (
	"github.com/google/uuid"
	"github.com/vildan-valeev/go-clean-architecture/internal/domain"
)

type CategoryCreateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Tag         string `json:"tag"`
}

type CategoryCreateResponse struct {
	ID string `json:"id"`
}

func CategoryCreateToResponse(id uuid.UUID) CategoryCreateResponse {
	return CategoryCreateResponse{
		ID: id.String(),
	}
}

type CategoryDeleteRequest struct {
	ID string `json:"id"`
}

type CategoryDeleteResponse struct {
}

func CategoryDeleteToResponse() CategoryDeleteResponse {
	return CategoryDeleteResponse{}
}

type CategoryUpdateRequest struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Tag         string `json:"tag"`
}

type CategoryUpdateResponse struct {
}

func CategoryUpdateToResponse() CategoryUpdateResponse {
	return CategoryUpdateResponse{}
}

type CategoryReadRequest struct {
	ID string `json:"id"`
}

type CategoryReadResponse struct {
	title       string `json:"title"`
	description string `json:"description"`
	tag         string `json:"tag"`
}

func CategoryReadToResponse(c domain.Category) CategoryReadResponse {
	return CategoryReadResponse{
		title:       c.Title,
		description: c.Description,
		tag:         c.Tag,
	}
}
