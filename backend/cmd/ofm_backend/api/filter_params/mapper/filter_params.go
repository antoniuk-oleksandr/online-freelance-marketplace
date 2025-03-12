package mapper

import (
	"encoding/json"
	"ofm_backend/cmd/ofm_backend/utils"
	"ofm_backend/cmd/ofm_backend/api/filter_params/dto"
	"ofm_backend/cmd/ofm_backend/api/filter_params/model"
)

func MapFilterParamsModelToDTO(
	rawFilterParams *model.FilterParamsJSON,
) (*dto.FilterParams, error) {
	var filterParamsModel model.FilterParams

	if err := json.Unmarshal(rawFilterParams.Languages, &filterParamsModel.Languages); err != nil {
		return nil, err
	}
	if err := json.Unmarshal(rawFilterParams.Skills, &filterParamsModel.Skills); err != nil {
		return nil, err
	}
	if err := json.Unmarshal(rawFilterParams.Categories, &filterParamsModel.Categories); err != nil {
		return nil, err
	}

	var filterParamsDTO dto.FilterParams

	filterParamsDTO.Languages = *utils.MapFilterParamModelsToDTOs(&filterParamsModel.Languages)
	filterParamsDTO.Categories = *utils.MapFilterParamModelsToDTOs(&filterParamsModel.Categories)
	filterParamsDTO.Skills = *utils.MapFilterParamModelsToDTOs(&filterParamsModel.Skills)

	return &filterParamsDTO, nil
}


