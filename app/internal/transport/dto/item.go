package dto

type ItemUpdateDtoRequest struct {
	ID    int64  `json:"id,omitempty"`
	Key   string `json:"key,omitempty"`
	Value uint8  `json:"value,omitempty"`
}

type ItemUpdateDtoResponse struct {
	Value uint8 `json:"value"`
}

type ItemCreateDtoRequest struct {
	Name string `json:"name"`
	Age  uint8  `json:"age"`
}

type ItemCreateDtoResponse struct {
	ID int64 `json:"id"`
}
