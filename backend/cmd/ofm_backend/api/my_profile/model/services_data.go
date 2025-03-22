package model

type ServicesData struct {
	ServiceTableData []ServiceTableData `json:"service_table_data" db:"service_table_data"`
	TotalPages       int                `json:"total_pages" db:"total_pages"`
}
