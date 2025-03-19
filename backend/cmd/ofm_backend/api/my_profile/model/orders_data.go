package model

type OrdersData struct {
	OrderTableData []OrderTableData `json:"order_table_data" db:"order_table_data"`
	TotalPages     int              `json:"total_pages" db:"total_pages"`
}
