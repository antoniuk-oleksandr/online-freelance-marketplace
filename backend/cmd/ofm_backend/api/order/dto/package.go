package dto

type Package struct {
	Id    int64   `json:"id"`
	Price float64 `json:"price"`
	Title string  `json:"title"`
}
