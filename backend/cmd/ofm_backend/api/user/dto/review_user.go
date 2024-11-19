package dto

type ReviewUserDTO struct {
	ID        int64   `json:"id"`
	Avatar    *string `json:"avatar"`
	Username  string  `json:"username"`
}
