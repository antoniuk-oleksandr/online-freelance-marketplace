package dto

type OrdersResponse struct {
	Orders     []OrderTableData `json:"orders"`
	TotalPages int              `json:"totalPages"`
}
