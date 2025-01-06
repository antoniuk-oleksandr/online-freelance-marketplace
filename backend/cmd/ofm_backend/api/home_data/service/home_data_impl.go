package service

import (
	"ofm_backend/cmd/ofm_backend/api/home_data/dto"
	"ofm_backend/cmd/ofm_backend/api/home_data/mapper"
	"ofm_backend/cmd/ofm_backend/api/home_data/repository"
	"ofm_backend/cmd/ofm_backend/utils"
)

type homeService struct {
	homeRepository repository.HomeRepository
}

func NewHomeService(
	homeRepository repository.HomeRepository,
) HomeService {
	return &homeService{
		homeRepository: homeRepository,
	}
}

func (hs *homeService) GetHomeData() (*dto.HomeData, error) {
	homeDataModel, err := hs.homeRepository.GetHomeData()
	if err != nil {
		return nil, err
	} else if homeDataModel == nil {
		return nil, utils.ErrUnexpectedError
	}

	keyMetrictsDto := mapper.MapKeyMetricsModelToDto(&homeDataModel.KeyMetrics)
	bestFreelancerDtos := mapper.MapBestFreelancerModelsToDtos(homeDataModel.BestFreelancers)
	bestFreelanceDtos := mapper.MapBestFreelanceModelsToDto(homeDataModel.BestFreelances)
	bestReviewDtos := mapper.MapBestReviewModelsToDtos(homeDataModel.BestReviews)

	homeDataDto := mapper.MapHomeDataModelToDto(
		keyMetrictsDto,
		bestFreelancerDtos,
		bestFreelanceDtos,
		bestReviewDtos,
	)
	
	homeDataDtoWithImages := utils.AddServerURLToFiles(&homeDataDto)
	
	return homeDataDtoWithImages, nil
}
