package dto

type ReviewUserDTO struct {
	ID        int64   `json:"id"`
	FirstName string  `json:"first_name"`
	Surname   string  `json:"surname"`
	Avatar    *string `json:"avatar"`
}
