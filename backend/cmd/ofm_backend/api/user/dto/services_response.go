package dto

import "ofm_backend/cmd/ofm_backend/api/user/model"

type ServicesResponse struct {
	Services        *[]model.UserByIdFreelanceService `json:"services"`
	HasMoreServices bool                              `json:"hasMoreServices"`
	ServicesCursor  *string                           `json:"servicesCursor"`
}
