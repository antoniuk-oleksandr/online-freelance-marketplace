package model

type DeliveryFile struct {
	DeliveryId int `json:"delivery_id" db:"delivery_id"`
	FileId     int `json:"file_id" db:"file_id"`
}
