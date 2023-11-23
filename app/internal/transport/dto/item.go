package dto

import (
	"github.com/vildan-valeev/go-clean-architecture/internal/domain"
	"strconv"
)

type ItemCreateRequest struct {
	ID    string `json:"id,omitempty"`
	Title string `json:"title"`
}

type ItemCreateResponse struct {
	ID string `json:"Code"`
}

func ItemCreateToResponse(c uint64) ItemCreateResponse {
	return ItemCreateResponse{
		ID: strconv.FormatUint(c, 10),
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
