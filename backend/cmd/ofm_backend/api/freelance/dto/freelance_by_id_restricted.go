package dto

type FreelanceByIdRestricted struct {
	Id           int64      `json:"id"`
	ReviewsCount int64      `json:"reviewsCount"`
	Rating       float64    `json:"rating"`
	Title        string     `json:"title"`
	Image        *string    `json:"image"`
	Packages     *[]Package `json:"packages"`
}
