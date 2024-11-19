package dto

type FreelanceServiceFreelancer struct {
	ID           int64   `json:"id"`
	Username     string  `json:"username"`
	FirstName    string  `json:"firstName"`
	Surname      string  `json:"surname"`
	Avatar       string  `json:"avatar"`
	Rating       float64 `json:"rating"`
	Level        float64 `json:"level"`
	ReviewsCount int64   `json:"reviewsCount"`
}