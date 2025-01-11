package model

type FreelanceByIdRestricted struct {
	Id           int64      `json:"id" db:"id"`
	ReviewsCount int64      `json:"reviews_count" db:"reviews_count"`
	Rating       float64    `json:"rating" db:"rating"`
	Title        string     `json:"title" db:"title"`
	Image        *string    `json:"image" db:"image"`
	Packages     *[]Package `json:"packages" db:"packages"`
}
