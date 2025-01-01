package dto

type ServiceByIdDto struct {
	ID           int64     `json:"id" db:"id"`
	Title        string    `json:"title" db:"title"`
	Image        *string   `json:"image" db:"image"`
	ReviewsCount int64     `json:"reviewsCount"`
	Rating       float64   `json:"rating"`
	MinPrice     float64   `json:"minPrice"`
}
