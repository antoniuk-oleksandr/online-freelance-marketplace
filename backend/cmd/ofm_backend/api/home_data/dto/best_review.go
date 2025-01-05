package dto

type BestReview struct {
	FirstName string `json:"firstName"`
	Surname   string `json:"surname"`
	Content   string `json:"content"`
	Rating    int    `json:"rating"`
}
