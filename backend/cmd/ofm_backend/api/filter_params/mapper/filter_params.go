package mapper

import (
	"encoding/json"
	"ofm_backend/cmd/ofm_backend/api/filter_params/dto"
	"ofm_backend/cmd/ofm_backend/api/filter_params/model"
)

func MapFilterParamsModelToDTO(
	rawFilterParams *model.FilterParams,
) (*dto.FilterParams, error) {
	var filterParams dto.FilterParams

	if err := json.Unmarshal(rawFilterParams.Languages, &filterParams.Language); err != nil {
		return nil, err
	}
	if err := json.Unmarshal(rawFilterParams.Skills, &filterParams.Skill); err != nil {
		return nil, err
	}
	if err := json.Unmarshal(rawFilterParams.Categories, &filterParams.Category); err != nil {
		return nil, err
	}

	return &filterParams, nil
}
