package model

type Package struct {
	Id    int64   `json:"package_id" db:"package_id"`
	Price float64 `json:"price" db:"price"`
	Title string  `json:"title" db:"title"`
}
