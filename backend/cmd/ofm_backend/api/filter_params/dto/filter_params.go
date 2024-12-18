package dto

import (
	freelance_model "ofm_backend/cmd/ofm_backend/api/freelance/model"
	user_model "ofm_backend/cmd/ofm_backend/api/user/model"
)

type FilterParams struct {
	Language []user_model.Language      `json:"language"`
	Category []freelance_model.Category `json:"category" `
	Skill    []user_model.Skill         `json:"skill"`
}
