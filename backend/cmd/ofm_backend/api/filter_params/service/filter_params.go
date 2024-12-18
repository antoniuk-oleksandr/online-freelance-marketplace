package service

import (
	"ofm_backend/cmd/ofm_backend/api/filter_params/dto"
	"ofm_backend/cmd/ofm_backend/api/filter_params/mapper"
	"ofm_backend/cmd/ofm_backend/api/filter_params/repository"
	"ofm_backend/internal/database"
)

func FilterParamsGetAll() (*dto.FilterParams, error){
	db := database.GetDB()

	rawFilterParams, err := repository.FilterParamsGetAll(db)
	if err != nil {
		return nil, err
	}
	
	return mapper.MapFilterParamsModelToDTO(rawFilterParams)
}
