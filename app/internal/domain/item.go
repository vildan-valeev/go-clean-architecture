package domain

type Item struct {
	ID     int64  `json:"id" redis:"id"`
	Title  string `json:"name" redis:"name"`
	Amount uint8  `json:"amount" redis:"amount"`

	Category Category `json:"category"`
}
