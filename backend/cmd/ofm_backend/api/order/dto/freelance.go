package dto

type Freelance struct {
	Id      int64   `json:"id"`
	Title   string  `json:"title"`
	Image   string  `json:"image"`
	Package Package `json:"package"`
}
