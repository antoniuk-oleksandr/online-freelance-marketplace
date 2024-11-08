package dto

type ReviewService struct {
	ID    int64   `json:"id"`
	Price float64 `json:"price"`
	Image *string `json:"image"`
	Title string  `json:"title"`
}
