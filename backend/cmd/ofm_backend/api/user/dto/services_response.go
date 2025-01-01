package dto

type ServicesResponse struct {
	Services        *[]ServiceByIdDto `json:"services"`
	HasMoreServices bool              `json:"hasMoreServices"`
	ServicesCursor  *string           `json:"servicesCursor"`
}
