package model

type Freelance struct {
	Id      int64   `json:"service_id" db:"service_id"`
	Title   string  `json:"title" db:"title"`
	Image   string  `json:"image" db:"image"`
	Package Package `json:"package" db:"package"`
}
