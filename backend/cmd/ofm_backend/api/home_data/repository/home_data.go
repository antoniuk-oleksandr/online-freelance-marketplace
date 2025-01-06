package repository

import "ofm_backend/cmd/ofm_backend/api/home_data/model"

type HomeRepository interface {
	GetHomeData() (*model.HomeData, error)
}