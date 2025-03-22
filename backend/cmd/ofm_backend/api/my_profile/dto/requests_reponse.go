package dto

type RequestsResponse struct {
	Requests   []RequestTableData `json:"requests"`
	TotalPages int                `json:"totalPages"`
}
