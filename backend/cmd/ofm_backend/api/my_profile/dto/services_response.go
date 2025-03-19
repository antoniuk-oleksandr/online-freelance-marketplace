package dto

type ServicesResponse struct {
	TotalPages int                `json:"totalPages"`
	Services   []ServiceTableData `json:"services"`
}
