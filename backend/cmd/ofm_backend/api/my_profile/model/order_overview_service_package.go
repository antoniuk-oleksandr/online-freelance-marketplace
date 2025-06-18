package model

type OrderOverviewServicePackage struct {
	Description  string `json:"description" db:"description"`
	Name         string `json:"name" db:"name"`
	DeliveryTime int    `json:"delivery_time" db:"delivery_time"`
}
