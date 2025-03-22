package model

type RequestsData struct {
	RequestTableData   []RequestTableData `json:"requests_table_data" db:"requests_table_data"`
	TotalPages int                `json:"total_pages" db:"total_pages"`
}
