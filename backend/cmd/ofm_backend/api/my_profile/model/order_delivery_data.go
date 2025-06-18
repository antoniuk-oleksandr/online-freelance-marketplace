package model

type OrderDeliveryData struct {
	Message string   `json:"message" db:"message"`
	Date    string   `json:"date" db:"date"`
	Files   []string `json:"files" db:"files"`
}
