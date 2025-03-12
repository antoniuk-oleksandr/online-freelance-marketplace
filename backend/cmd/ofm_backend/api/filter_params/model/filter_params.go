package model

import (
	freelance_models "ofm_backend/cmd/ofm_backend/api/freelance/model"
	user_models "ofm_backend/cmd/ofm_backend/api/user/model"
)

type FilterParams struct {
	Languages []user_models.Language
	Categories []freelance_models.Category
	Skills []user_models.Skill
}