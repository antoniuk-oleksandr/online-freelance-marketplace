package dto

type BestFreelance struct {
	Id          int64   `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Image       *string `json:"image"`
}
