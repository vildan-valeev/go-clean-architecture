package dto

import (
	"github.com/vildan-valeev/go-clean-architecture/internal/domain"
	"strconv"
)

type CategoryCreateRequest struct {
	ID    string `json:"id,omitempty"`
	Title string `json:"title"`
}

type CategoryCreateResponse struct {
	ID string `json:"Code"`
}

func CategoryCreateToResponse(c uint64) CategoryCreateResponse {
	return CategoryCreateResponse{
		ID: strconv.FormatUint(c, 10),
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
	ID string `json:"id"`
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
}

func CategoryReadToResponse(c domain.Category) CategoryReadResponse {
	return CategoryReadResponse{}
}
