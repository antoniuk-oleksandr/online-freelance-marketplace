package service

import "ofm_backend/cmd/ofm_backend/api/home_data/dto"

type HomeService interface {
	GetHomeData() (*dto.HomeData, error)
}
